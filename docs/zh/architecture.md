# 架构设计

本页解释 **Maven SDK Go** 各部分如何协同工作，以及一次请求如何从你的代码
一路流向 `mvn` 进程再返回。下面每张图都由本页内嵌的 Mermaid 源码渲染而成。

## 系统上下文

在最上层，你的应用（或 AI 智能体）与一组职责单一的小模块交互。只有
`command` 与 `installer` 会触碰外部世界（`mvn` 进程、网络、文件系统）；其余
都是纯 Go 代码，只对路径、XML 和字符串做处理。

```mermaid
flowchart TB
    subgraph Caller["调用方"]
        App["Go 应用 / AI 智能体"]
    end

    subgraph SDK["mvn-skills 模块"]
        Finder["finder<br/>定位 mvn / mvnw"]
        Command["command<br/>构建并执行 mvn"]
        Pom["pom<br/>解析 pom.xml"]
        Settings["settings<br/>解析 settings.xml"]
        Repo["local_repository<br/>解析 ~/.m2 布局"]
        Installer["installer<br/>下载并安装 mvn"]
    end

    subgraph External["外部世界"]
        Mvn["mvn 命令行进程"]
        FS["文件系统<br/>pom.xml / settings.xml / ~/.m2"]
        Net["网络<br/>Apache 镜像"]
    end

    App --> Finder
    App --> Command
    App --> Pom
    App --> Settings
    App --> Repo
    App --> Installer

    Finder -->|发现路径| Command
    Command -->|执行| Mvn
    Pom -->|读取| FS
    Settings -->|读取| FS
    Repo -->|遍历| FS
    Installer -->|下载| Net
    Installer -->|写入| FS

    classDef pure fill:#e8f0fe,stroke:#2563eb,color:#1e3a8a;
    classDef io fill:#fef3c7,stroke:#d97706,color:#7c2d12;
    class Finder,Pom,Settings,Repo pure;
    class Command,Installer io;
```

<div class="tip custom-block" style="padding-top:8px">

蓝色模块是**纯函数 / 无副作用**的，易于单元测试。琥珀色模块
（`command`、`installer`）会执行 **I/O**，测试时用 mock 进程和 HTTP 测试
服务器覆盖。

</div>

## 模块依赖图

内部依赖刻意保持很浅——没有循环依赖，叶子模块只依赖标准库。

```mermaid
flowchart LR
    installer --> finder
    installer --> command
    finder --> command
    command --> stdlib["os/exec · context"]
    pom --> stdlibxml["encoding/xml"]
    settings --> stdlibxml
    local_repository --> command

    classDef leaf fill:#dcfce7,stroke:#16a34a,color:#14532d;
    class pom,settings leaf;
```

| 模块 | 依赖 | 副作用 |
|------|------|--------|
| `finder` | `command`（版本检查） | 读取 `PATH`、`M2_HOME`、文件系统 |
| `command` | 仅标准库 | 启动 `mvn` 进程 |
| `pom` | `encoding/xml` | 无（纯函数） |
| `settings` | `encoding/xml` | 无（纯函数） |
| `local_repository` | `command` | 读取 `~/.m2` |
| `installer` | `finder`、`command` | 网络、文件系统、环境变量 |

## Maven 定位（finder）

`FindBestMaven` 实现了一条**优先级级联**：项目自带的 Maven Wrapper 永远
优先于系统安装，因为 Wrapper 锁定了项目构建时所用的确切 Maven 版本。

```mermaid
flowchart TD
    Start(["FindBestMaven(projectDir)"]) --> HasWrapper{"projectDir 下有<br/>mvnw / mvnw.cmd?"}
    HasWrapper -->|有| UseWrapper["返回 wrapper 路径"]
    HasWrapper -->|无| CheckPath{"PATH 中有 mvn?"}
    CheckPath -->|有| UsePath["返回 PATH 中的 mvn"]
    CheckPath -->|无| CheckM2{"设置了 M2_HOME /<br/>MAVEN_HOME?"}
    CheckM2 -->|有| ValidateHome{"是有效的<br/>Maven home?"}
    ValidateHome -->|是| UseHome["返回 $M2_HOME/bin/mvn"]
    ValidateHome -->|否| Fail
    CheckM2 -->|无| Fail(["ErrNotFoundMaven"])

    classDef ok fill:#dcfce7,stroke:#16a34a,color:#14532d;
    classDef bad fill:#fee2e2,stroke:#dc2626,color:#7f1d1d;
    class UseWrapper,UsePath,UseHome ok;
    class Fail bad;
```

## 命令执行流水线

`command` 模块把流式构建器转换为 `*exec.Cmd`，执行它，并把非零退出码映射为
结构化的 `*MavenError`（而非裸字符串）。

```mermaid
sequenceDiagram
    autonumber
    participant U as 你的代码
    participant B as CommandBuilder
    participant E as 执行层
    participant M as mvn 进程

    U->>B: WithBatchMode().WithSkipTests()
    Note over B: 每个 With* 返回一个副本<br/>（构建器不可变）
    U->>B: CleanInstall()
    B->>E: 构建 argv ["-B","-DskipTests","clean","install"]
    E->>M: exec.CommandContext(ctx, mvn, argv...)
    activate M
    M-->>E: stdout / stderr / 退出码
    deactivate M
    alt 退出码 == 0
        E-->>U: stdout 字符串, nil
    else 退出码 != 0
        E-->>U: "", *MavenError{ExitCode, Stderr, Args}
    end
```

### 构建器的不可变性

便捷方法（`Clean()`、`Install()`……）**不会**修改接收者本身。每次调用都
返回一个追加了目标的全新构建器，因此一个已配置好的构建器可以被复用来执行
多条命令而互不干扰。

```mermaid
flowchart LR
    base["base = NewCommandBuilder()<br/>.WithBatchMode()"]
    base -->|".Clean()"| c1["副本 + clean"]
    base -->|".Install()"| c2["副本 + install"]
    base -->|".Test()"| c3["副本 + test"]
    base -.->|保持不变| base

    classDef immut fill:#e8f0fe,stroke:#2563eb,color:#1e3a8a;
    class base immut;
```

## Maven 构建生命周期

便捷方法映射到 Maven 内置的三条生命周期。执行某个阶段会连带执行同一条
生命周期中它之前的所有阶段。

```mermaid
flowchart LR
    subgraph cleanLc["clean 生命周期"]
        direction LR
        preclean[pre-clean] --> clean0[clean] --> postclean[post-clean]
    end

    subgraph defaultLc["default 生命周期"]
        direction LR
        validate --> compile --> test --> package --> verify --> install --> deploy
    end

    subgraph siteLc["site 生命周期"]
        direction LR
        presite[pre-site] --> site0[site] --> postsite[post-site] --> sitedeploy[site-deploy]
    end

    clean0 -.->|"CleanInstall()"| validate
```

| SDK 方法 | 实际命令 | 典型用途 |
|----------|----------|----------|
| `Install(mvn)` | `mvn clean install` | 本地构建并安装到 `~/.m2` |
| `CleanPackage()` | `mvn clean package` | 产出构件，不安装 |
| `CleanDeploy()` | `mvn clean deploy` | 发布到远程仓库 |
| `Verify(mvn)` | `mvn verify` | 运行集成测试与检查 |

## 安装器：端到端流程

`InstallWithOptions` 具备幂等性且感知平台。它优先尝试成本最低的方案（已经
安装好的 Maven），再尝试原生包管理器，只有在万不得已时才下载二进制归档。

```mermaid
flowchart TD
    Start(["InstallWithOptions(opts)"]) --> Force{"opts.Force?"}
    Force -->|否| Idem{"已安装且版本<br/>足够新的可用 mvn?"}
    Idem -->|是| Done(["返回现有 MAVEN_HOME"])
    Force -->|是| OS
    Idem -->|否| OS{"runtime.GOOS"}

    OS -->|linux| PkgLinux{"apt / dnf / yum /<br/>apk / pacman / zypper?"}
    OS -->|darwin| Brew{"是否有 brew?"}
    OS -->|windows| Binary

    PkgLinux -->|找到| PkgRun["运行包管理器"] --> Verify
    PkgLinux -->|没有| Binary
    Brew -->|有| BrewRun["brew install maven"] --> Verify
    Brew -->|无| Binary

    Binary["下载二进制归档"] --> Mirrors["按序尝试镜像<br/>+ 指数退避重试"]
    Mirrors --> Checksum{"SHA512 匹配?"}
    Checksum -->|否| NextMirror["下一个镜像"] --> Mirrors
    Checksum -->|是| Extract["解压 (tar.gz / zip)<br/>带路径穿越防护"]
    Extract --> Env["配置 PATH / MAVEN_HOME"]
    Env --> Verify{"mvn -v 可用?"}
    Verify -->|是| Done
    Verify -->|否| Err(["安装错误"])

    classDef ok fill:#dcfce7,stroke:#16a34a,color:#14532d;
    classDef bad fill:#fee2e2,stroke:#dc2626,color:#7f1d1d;
    class Done ok;
    class Err bad;
```

### 带镜像回退的下载

每个镜像会以指数退避重试若干次后再切换；校验和不匹配被视作一次失败下载，
同样触发切换到下一个镜像。

```mermaid
sequenceDiagram
    autonumber
    participant I as installer
    participant M1 as archive.apache.org
    participant M2 as 阿里云镜像
    participant FS as 临时文件

    I->>M1: GET apache-maven-3.9.11-bin.tar.gz
    M1--xI: 503（瞬时故障）
    Note over I: 退避后重试（MaxRetries）
    I->>M1: GET（重试）
    M1--xI: 503
    Note over I: 该镜像耗尽 → 下一个
    I->>M2: GET apache-maven-3.9.11-bin.tar.gz
    M2-->>FS: 200 + 字节流
    I->>M2: GET ...tar.gz.sha512
    M2-->>I: 校验和
    I->>FS: 计算 SHA512
    alt 匹配
        I-->>I: 进入解压环节
    else 不匹配
        I-->>I: 丢弃，尝试下一个镜像
    end
```

### 各平台环境变量配置

三种操作系统持久化 `MAVEN_HOME` 和扩展 `PATH` 的方式差异很大。Windows 最
棘手：`setx PATH` 会在 1024 字符处截断，因此安装器从注册表读取现有的用户
`PATH` 再安全地追加，若超长则降级为打印提示而不是破坏 `PATH`。

```mermaid
flowchart TD
    Cfg(["configureEnvironment(home, opts)"]) --> Skip{"opts.SkipEnvSetup?"}
    Skip -->|是| NoOp(["空操作"])
    Skip -->|否| Which{"runtime.GOOS"}

    Which -->|windows| Win["setx MAVEN_HOME home<br/>从注册表读取用户 PATH<br/>不超 1024 字符则追加 home\\bin"]
    Which -->|darwin| Mac["把 export 语句写入<br/>~/.zshrc（默认 shell）"]
    Which -->|linux| Lin["把 export 语句写入<br/>~/.bashrc / fish 配置"]

    Win --> Len{"新 PATH &gt; 1024?"}
    Len -->|是| Warn(["返回警告，<br/>打印手动配置提示"])
    Len -->|否| OkW(["PATH 已更新"])

    classDef ok fill:#dcfce7,stroke:#16a34a,color:#14532d;
    class OkW,Mac,Lin ok;
```

## POM 对象模型

`pom.ParseFile` 把 `pom.xml` 反序列化为 `Project` 树。访问器方法
（`GetGAV`、`GetDependencies`……）提供空安全读取并带有合理默认值
（例如 packaging 默认为 `jar`）。

```mermaid
classDiagram
    class Project {
        +string GroupId
        +string ArtifactId
        +string Version
        +string Packaging
        +GetGAV()
        +GetDependencies()
        +IsMultiModule() bool
        +HasParent() bool
    }
    class Parent {
        +string GroupId
        +string ArtifactId
        +string Version
    }
    class Dependency {
        +string GroupId
        +string ArtifactId
        +string Version
        +string Scope
        +bool Optional
    }
    class Build {
        +string FinalName
        +Plugins
    }
    class Plugin {
        +string GroupId
        +string ArtifactId
        +string Version
    }
    Project "1" o-- "0..1" Parent
    Project "1" o-- "*" Dependency
    Project "1" o-- "0..1" Build
    Build "1" o-- "*" Plugin
```

## 本地仓库布局

`local_repository` 把一个 Maven 坐标（GAV）映射到它在 `~/.m2/repository`
下的磁盘路径，与 Maven 自身的目录约定一致。

```mermaid
flowchart TD
    GAV["org.springframework : spring-core : 5.3.21"]
    GAV --> Path["~/.m2/repository/<br/>org/springframework/spring-core/5.3.21/"]
    Path --> Jar["spring-core-5.3.21.jar"]
    Path --> Pom["spring-core-5.3.21.pom"]
    Path --> Sources["spring-core-5.3.21-sources.jar"]
    Path --> Sha["*.sha1 / *.md5 校验和"]

    classDef coord fill:#e8f0fe,stroke:#2563eb,color:#1e3a8a;
    class GAV coord;
```

group id `org.springframework` 会变成嵌套目录 `org/springframework`——每个
点号都是一个路径分隔符。

## 下一步

- [API 参考](/zh/api) —— 完整的函数接口，含各模块示意图
- [快速开始](/zh/) —— 安装与首次构建
