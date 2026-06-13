package settings

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

const testSettingsXML = `<?xml version="1.0" encoding="UTF-8"?>
<settings xmlns="http://maven.apache.org/SETTINGS/1.2.0"
          xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
          xsi:schemaLocation="http://maven.apache.org/SETTINGS/1.2.0 https://maven.apache.org/xsd/settings-1.2.0.xsd">
    <localRepository>/home/user/.m2/repository</localRepository>
    <offline>false</offline>

    <servers>
        <server>
            <id>internal-repo</id>
            <username>admin</username>
            <password>secret123</password>
        </server>
        <server>
            <id>github</id>
            <username>user</username>
            <password>ghp_token</password>
        </server>
    </servers>

    <mirrors>
        <mirror>
            <id>aliyun</id>
            <name>Aliyun Maven Mirror</name>
            <url>https://maven.aliyun.com/repository/public</url>
            <mirrorOf>central</mirrorOf>
        </mirror>
        <mirror>
            <id>company-mirror</id>
            <name>Company Internal Mirror</name>
            <url>https://repo.company.com/maven</url>
            <mirrorOf>*</mirrorOf>
        </mirror>
    </mirrors>

    <proxies>
        <proxy>
            <id>my-proxy</id>
            <active>true</active>
            <protocol>http</protocol>
            <host>proxy.company.com</host>
            <port>8080</port>
            <username>proxyuser</username>
            <password>proxypass</password>
        </proxy>
    </proxies>

    <profiles>
        <profile>
            <id>dev</id>
            <repositories>
                <repository>
                    <id>dev-repo</id>
                    <url>https://dev.repo.company.com/maven</url>
                    <releases>
                        <enabled>true</enabled>
                    </releases>
                    <snapshots>
                        <enabled>true</enabled>
                    </snapshots>
                </repository>
            </repositories>
        </profile>
    </profiles>

    <activeProfiles>
        <activeProfile>dev</activeProfile>
    </activeProfiles>
</settings>
`

func TestParseBytes(t *testing.T) {
	settings, err := ParseBytes([]byte(testSettingsXML))
	assert.Nil(t, err)
	assert.NotNil(t, settings)

	assert.Equal(t, "/home/user/.m2/repository", settings.LocalRepository)
	assert.False(t, settings.Offline)
}

func TestGetServers(t *testing.T) {
	settings, err := ParseBytes([]byte(testSettingsXML))
	assert.Nil(t, err)

	servers := settings.GetServers()
	assert.Len(t, servers, 2)

	assert.Equal(t, "internal-repo", servers[0].Id)
	assert.Equal(t, "admin", servers[0].Username)
	assert.Equal(t, "secret123", servers[0].Password)

	assert.Equal(t, "github", servers[1].Id)
}

func TestFindServer(t *testing.T) {
	settings, err := ParseBytes([]byte(testSettingsXML))
	assert.Nil(t, err)

	server := settings.FindServer("github")
	assert.NotNil(t, server)
	assert.Equal(t, "user", server.Username)

	server = settings.FindServer("nonexistent")
	assert.Nil(t, server)
}

func TestGetMirrors(t *testing.T) {
	settings, err := ParseBytes([]byte(testSettingsXML))
	assert.Nil(t, err)

	mirrors := settings.GetMirrors()
	assert.Len(t, mirrors, 2)

	assert.Equal(t, "aliyun", mirrors[0].Id)
	assert.Equal(t, "https://maven.aliyun.com/repository/public", mirrors[0].URL)
	assert.Equal(t, "central", mirrors[0].MirrorOf)
}

func TestFindMirror(t *testing.T) {
	settings, err := ParseBytes([]byte(testSettingsXML))
	assert.Nil(t, err)

	mirror := settings.FindMirror("aliyun")
	assert.NotNil(t, mirror)
	assert.Equal(t, "Aliyun Maven Mirror", mirror.Name)
}

func TestFindMirrorOf(t *testing.T) {
	settings, err := ParseBytes([]byte(testSettingsXML))
	assert.Nil(t, err)

	// 查找 central 的镜像
	mirror := settings.FindMirrorOf("central")
	assert.NotNil(t, mirror)
	assert.Equal(t, "aliyun", mirror.Id)

	// * 通配符镜像
	mirror = settings.FindMirrorOf("some-other-repo")
	assert.NotNil(t, mirror)
	assert.Equal(t, "company-mirror", mirror.Id)
}

func TestGetProxies(t *testing.T) {
	settings, err := ParseBytes([]byte(testSettingsXML))
	assert.Nil(t, err)

	proxies := settings.GetProxies()
	assert.Len(t, proxies, 1)

	assert.Equal(t, "my-proxy", proxies[0].Id)
	assert.True(t, proxies[0].Active)
	assert.Equal(t, "http", proxies[0].Protocol)
	assert.Equal(t, "proxy.company.com", proxies[0].Host)
	assert.Equal(t, 8080, proxies[0].Port)
}

func TestFindActiveProxy(t *testing.T) {
	settings, err := ParseBytes([]byte(testSettingsXML))
	assert.Nil(t, err)

	proxy := settings.FindActiveProxy()
	assert.NotNil(t, proxy)
	assert.Equal(t, "proxy.company.com", proxy.Host)
}

func TestGetProfiles(t *testing.T) {
	settings, err := ParseBytes([]byte(testSettingsXML))
	assert.Nil(t, err)

	profiles := settings.GetProfiles()
	assert.Len(t, profiles, 1)
	assert.Equal(t, "dev", profiles[0].Id)
}

func TestGetActiveProfileIds(t *testing.T) {
	settings, err := ParseBytes([]byte(testSettingsXML))
	assert.Nil(t, err)

	activeIds := settings.GetActiveProfileIds()
	assert.Len(t, activeIds, 1)
	assert.Contains(t, activeIds, "dev")
}

func TestParseFile(t *testing.T) {
	tmpDir := t.TempDir()
	settingsPath := filepath.Join(tmpDir, "settings.xml")
	err := os.WriteFile(settingsPath, []byte(testSettingsXML), 0644)
	assert.Nil(t, err)

	settings, err := ParseFile(settingsPath)
	assert.Nil(t, err)
	assert.NotNil(t, settings)
	assert.Equal(t, "/home/user/.m2/repository", settings.LocalRepository)
}

func TestParseFileNotFound(t *testing.T) {
	_, err := ParseFile("/nonexistent/path/settings.xml")
	assert.NotNil(t, err)
}

func TestParseInvalidXML(t *testing.T) {
	_, err := ParseBytes([]byte("not xml"))
	assert.NotNil(t, err)
}

func TestGetDefaultSettingsPath(t *testing.T) {
	path := GetDefaultSettingsPath()
	assert.NotEmpty(t, path)
	assert.Contains(t, path, ".m2")
	assert.Contains(t, path, "settings.xml")
}

func TestEmptySettingsCollections(t *testing.T) {
	emptyXML := `<?xml version="1.0" encoding="UTF-8"?><settings></settings>`
	settings, err := ParseBytes([]byte(emptyXML))
	assert.Nil(t, err)

	assert.NotNil(t, settings.GetServers())
	assert.Empty(t, settings.GetServers())
	assert.NotNil(t, settings.GetMirrors())
	assert.Empty(t, settings.GetMirrors())
	assert.NotNil(t, settings.GetProxies())
	assert.Empty(t, settings.GetProxies())
	assert.NotNil(t, settings.GetProfiles())
	assert.Empty(t, settings.GetProfiles())
	assert.NotNil(t, settings.GetActiveProfileIds())
	assert.Empty(t, settings.GetActiveProfileIds())
}
