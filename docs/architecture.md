# Architecture

This page explains how the pieces of **Maven SDK Go** fit together, and how a
request flows from your code down to the `mvn` process and back. Every diagram
below is rendered from Mermaid source embedded in this page.

## System Context

At the top level, your application (or an AI agent) talks to a handful of small,
focused packages. Only `command` and `installer` ever touch the outside world
(the `mvn` process, the network, the filesystem); the rest are pure Go that
operate on paths, XML, and strings.

```mermaid
flowchart TB
    subgraph Caller["Caller"]
        App["Go application / AI agent"]
    end

    subgraph SDK["mvn-skills packages"]
        Finder["finder<br/>locate mvn / mvnw"]
        Command["command<br/>build &amp; run mvn"]
        Pom["pom<br/>parse pom.xml"]
        Settings["settings<br/>parse settings.xml"]
        Repo["local_repository<br/>resolve ~/.m2 layout"]
        Installer["installer<br/>download &amp; install mvn"]
    end

    subgraph External["External world"]
        Mvn["mvn CLI process"]
        FS["Filesystem<br/>pom.xml / settings.xml / ~/.m2"]
        Net["Network<br/>Apache mirrors"]
    end

    App --> Finder
    App --> Command
    App --> Pom
    App --> Settings
    App --> Repo
    App --> Installer

    Finder -->|discovers path| Command
    Command -->|exec| Mvn
    Pom -->|reads| FS
    Settings -->|reads| FS
    Repo -->|walks| FS
    Installer -->|downloads| Net
    Installer -->|writes| FS

    classDef pure fill:#e8f0fe,stroke:#2563eb,color:#1e3a8a;
    classDef io fill:#fef3c7,stroke:#d97706,color:#7c2d12;
    class Finder,Pom,Settings,Repo pure;
    class Command,Installer io;
```

<div class="tip custom-block" style="padding-top:8px">

Blue packages are **pure / side-effect free** — trivially unit-testable.
Amber packages (`command`, `installer`) perform **I/O** and are exercised with
mocked processes and HTTP test servers.

</div>

## Package Dependency Graph

Internal dependencies are intentionally shallow — there are no cycles, and the
leaf packages depend on nothing but the standard library.

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

| Package | Depends on | Side effects |
|---------|-----------|--------------|
| `finder` | `command` (version check) | reads `PATH`, `M2_HOME`, filesystem |
| `command` | stdlib only | spawns `mvn` |
| `pom` | `encoding/xml` | none (pure) |
| `settings` | `encoding/xml` | none (pure) |
| `local_repository` | `command` | reads `~/.m2` |
| `installer` | `finder`, `command` | network, filesystem, env |

## Maven Resolution (finder)

`FindBestMaven` implements a **preference cascade**: a project-local Maven
Wrapper always wins over a system install, because the wrapper pins the exact
Maven version the project was built with.

```mermaid
flowchart TD
    Start(["FindBestMaven(projectDir)"]) --> HasWrapper{"mvnw / mvnw.cmd<br/>in projectDir?"}
    HasWrapper -->|yes| UseWrapper["return wrapper path"]
    HasWrapper -->|no| CheckPath{"mvn on PATH?"}
    CheckPath -->|yes| UsePath["return PATH mvn"]
    CheckPath -->|no| CheckM2{"M2_HOME /<br/>MAVEN_HOME set?"}
    CheckM2 -->|yes| ValidateHome{"valid Maven<br/>home?"}
    ValidateHome -->|yes| UseHome["return $M2_HOME/bin/mvn"]
    ValidateHome -->|no| Fail
    CheckM2 -->|no| Fail(["ErrNotFoundMaven"])

    classDef ok fill:#dcfce7,stroke:#16a34a,color:#14532d;
    classDef bad fill:#fee2e2,stroke:#dc2626,color:#7f1d1d;
    class UseWrapper,UsePath,UseHome ok;
    class Fail bad;
```

## Command Execution Pipeline

The `command` package turns a fluent builder into an `*exec.Cmd`, runs it, and
maps a non-zero exit into a structured `*MavenError` (never a bare string).

```mermaid
sequenceDiagram
    autonumber
    participant U as Your code
    participant B as CommandBuilder
    participant E as Exec layer
    participant M as mvn process

    U->>B: WithBatchMode().WithSkipTests()
    Note over B: each With* returns a copy<br/>(builder is immutable)
    U->>B: CleanInstall()
    B->>E: build argv ["-B","-DskipTests","clean","install"]
    E->>M: exec.CommandContext(ctx, mvn, argv...)
    activate M
    M-->>E: stdout / stderr / exit code
    deactivate M
    alt exit code == 0
        E-->>U: stdout string, nil
    else exit code != 0
        E-->>U: "", *MavenError{ExitCode, Stderr, Args}
    end
```

### Builder Immutability

Convenience methods (`Clean()`, `Install()`, …) do **not** mutate the receiver.
Each returns a fresh builder with the extra goal appended, so a single
configured builder can be reused for many commands without cross-talk.

```mermaid
flowchart LR
    base["base = NewCommandBuilder()<br/>.WithBatchMode()"]
    base -->|".Clean()"| c1["copy + clean"]
    base -->|".Install()"| c2["copy + install"]
    base -->|".Test()"| c3["copy + test"]
    base -.->|unchanged| base

    classDef immut fill:#e8f0fe,stroke:#2563eb,color:#1e3a8a;
    class base immut;
```

## Maven Build Lifecycle

The convenience methods map onto Maven's three built-in lifecycles. Running a
phase runs every phase before it in the same lifecycle.

```mermaid
flowchart LR
    subgraph cleanLc["clean lifecycle"]
        direction LR
        preclean[pre-clean] --> clean0[clean] --> postclean[post-clean]
    end

    subgraph defaultLc["default lifecycle"]
        direction LR
        validate --> compile --> test --> package --> verify --> install --> deploy
    end

    subgraph siteLc["site lifecycle"]
        direction LR
        presite[pre-site] --> site0[site] --> postsite[post-site] --> sitedeploy[site-deploy]
    end

    clean0 -.->|"CleanInstall()"| validate
```

| SDK method | Effective command | Typical use |
|------------|-------------------|-------------|
| `Install(mvn)` | `mvn clean install` | local build + install to `~/.m2` |
| `CleanPackage()` | `mvn clean package` | produce artifact, skip install |
| `CleanDeploy()` | `mvn clean deploy` | publish to a remote repo |
| `Verify(mvn)` | `mvn verify` | run integration tests + checks |

## Installer: End-to-End Flow

`InstallWithOptions` is idempotent and platform-aware. It tries the cheapest
option first (an already-installed Maven), then a native package manager, and
only downloads a binary archive as a last resort.

```mermaid
flowchart TD
    Start(["InstallWithOptions(opts)"]) --> Force{"opts.Force?"}
    Force -->|no| Idem{"usable mvn already<br/>installed &amp; new enough?"}
    Idem -->|yes| Done(["return existing MAVEN_HOME"])
    Force -->|yes| OS
    Idem -->|no| OS{"runtime.GOOS"}

    OS -->|linux| PkgLinux{"apt / dnf / yum /<br/>apk / pacman / zypper?"}
    OS -->|darwin| Brew{"brew available?"}
    OS -->|windows| Binary

    PkgLinux -->|found| PkgRun["run package manager"] --> Verify
    PkgLinux -->|none| Binary
    Brew -->|yes| BrewRun["brew install maven"] --> Verify
    Brew -->|no| Binary

    Binary["download binary archive"] --> Mirrors["try mirrors in order<br/>+ retry with backoff"]
    Mirrors --> Checksum{"SHA512 matches?"}
    Checksum -->|no| NextMirror["next mirror"] --> Mirrors
    Checksum -->|yes| Extract["extract (tar.gz / zip)<br/>with path-traversal guard"]
    Extract --> Env["configure PATH / MAVEN_HOME"]
    Env --> Verify{"mvn -v works?"}
    Verify -->|yes| Done
    Verify -->|no| Err(["installation error"])

    classDef ok fill:#dcfce7,stroke:#16a34a,color:#14532d;
    classDef bad fill:#fee2e2,stroke:#dc2626,color:#7f1d1d;
    class Done ok;
    class Err bad;
```

### Download with Mirror Fallback

Each mirror is retried with exponential backoff before moving on; a checksum
mismatch is treated like a failed download and advances to the next mirror.

```mermaid
sequenceDiagram
    autonumber
    participant I as installer
    participant M1 as archive.apache.org
    participant M2 as Aliyun mirror
    participant FS as temp file

    I->>M1: GET apache-maven-3.9.11-bin.tar.gz
    M1--xI: 503 (transient)
    Note over I: backoff, retry (MaxRetries)
    I->>M1: GET (retry)
    M1--xI: 503
    Note over I: mirror exhausted → next
    I->>M2: GET apache-maven-3.9.11-bin.tar.gz
    M2-->>FS: 200 + bytes
    I->>M2: GET ...tar.gz.sha512
    M2-->>I: checksum
    I->>FS: compute SHA512
    alt matches
        I-->>I: proceed to extract
    else mismatch
        I-->>I: discard, try next mirror
    end
```

### Platform Environment Configuration

The three operating systems persist `MAVEN_HOME` and extend `PATH` very
differently. Windows is the tricky one: `setx PATH` truncates at 1024 chars, so
the installer reads the existing user `PATH` from the registry and appends
safely, degrading to a printed hint rather than corrupting `PATH`.

```mermaid
flowchart TD
    Cfg(["configureEnvironment(home, opts)"]) --> Skip{"opts.SkipEnvSetup?"}
    Skip -->|yes| NoOp(["no-op"])
    Skip -->|no| Which{"runtime.GOOS"}

    Which -->|windows| Win["setx MAVEN_HOME home<br/>read user PATH from registry<br/>append home\\bin if &lt; 1024 chars"]
    Which -->|darwin| Mac["write export lines to<br/>~/.zshrc (default shell)"]
    Which -->|linux| Lin["write export lines to<br/>~/.bashrc / fish config"]

    Win --> Len{"new PATH &gt; 1024?"}
    Len -->|yes| Warn(["return warning,<br/>print manual hint"])
    Len -->|no| OkW(["PATH updated"])

    classDef ok fill:#dcfce7,stroke:#16a34a,color:#14532d;
    class OkW,Mac,Lin ok;
```

## POM Object Model

`pom.ParseFile` deserializes `pom.xml` into a `Project` tree. The accessor
methods (`GetGAV`, `GetDependencies`, …) provide null-safe reads with sensible
defaults (e.g. packaging defaults to `jar`).

```mermaid
classDiagram
    class Project {
        +string GroupId
        +string ArtifactId
        +string Version
        +string Packaging
        +GetGAV() (string, string, string)
        +GetDependencies() []Dependency
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
        +[]Plugin Plugins
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

## Local Repository Layout

`local_repository` maps a Maven coordinate (GAV) onto its on-disk path inside
`~/.m2/repository`, mirroring Maven's own directory convention.

```mermaid
flowchart TD
    GAV["org.springframework : spring-core : 5.3.21"]
    GAV --> Path["~/.m2/repository/<br/>org/springframework/spring-core/5.3.21/"]
    Path --> Jar["spring-core-5.3.21.jar"]
    Path --> Pom["spring-core-5.3.21.pom"]
    Path --> Sources["spring-core-5.3.21-sources.jar"]
    Path --> Sha["*.sha1 / *.md5 checksums"]

    classDef coord fill:#e8f0fe,stroke:#2563eb,color:#1e3a8a;
    class GAV coord;
```

The group id `org.springframework` becomes the nested directory
`org/springframework` — each dot is a path separator.

## Where to Next

- [API Reference](/api) — the full function surface, with per-package diagrams
- [Getting Started](/) — install and first build
