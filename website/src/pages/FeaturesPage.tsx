import React from 'react';
import { Typography, Row, Col, Card, Table } from 'antd';
import {
  SearchOutlined,
  BuildOutlined,
  FileTextOutlined,
  SettingOutlined,
  FolderOutlined,
  DownloadOutlined,
} from '@ant-design/icons';
import theme from '../styles/theme';

const { Title, Paragraph, Text } = Typography;

const FeaturesPage: React.FC = () => {
  const commandOptions = [
    { option: 'WithBatchMode()', flag: '-B', desc: 'Non-interactive mode (CI/CD)' },
    { option: 'WithSkipTests()', flag: '-DskipTests', desc: 'Skip test execution' },
    { option: 'WithSkipTestsCompletely()', flag: '-Dmaven.test.skip=true', desc: 'Skip tests entirely' },
    { option: 'WithOffline()', flag: '-o', desc: 'Offline mode' },
    { option: 'WithUpdateSnapshots()', flag: '-U', desc: 'Force update SNAPSHOTs' },
    { option: 'WithProfiles(...)', flag: '-P', desc: 'Activate profiles' },
    { option: 'WithProperty(k, v)', flag: '-Dk=v', desc: 'Set system property' },
    { option: 'WithPomFile(path)', flag: '-f', desc: 'Specify POM file' },
    { option: 'WithSettingsFile(path)', flag: '-s', desc: 'Specify settings.xml' },
    { option: 'WithProjects(...)', flag: '-pl', desc: 'Build specific modules' },
    { option: 'WithAlsoMake()', flag: '-am', desc: 'Also build dependencies' },
    { option: 'WithDebug()', flag: '-X', desc: 'Debug output' },
    { option: 'WithQuiet()', flag: '-q', desc: 'Quiet mode' },
    { option: 'WithThreads(n)', flag: '-T', desc: 'Parallel threads' },
    { option: 'WithFailAtEnd()', flag: '-fae', desc: 'Fail at end' },
    { option: 'WithNoTransferProgress()', flag: '-ntp', desc: 'No download progress' },
    { option: 'WithEnv(...)', flag: '—', desc: 'Set environment variables' },
    { option: 'WithContext(ctx)', flag: '—', desc: 'Cancellation/timeout support' },
  ];

  const columns = [
    {
      title: 'Option',
      dataIndex: 'option',
      key: 'option',
      render: (text: string) => <code style={{ color: theme.colors.primary, fontSize: 13 }}>{text}</code>,
    },
    {
      title: 'Flag',
      dataIndex: 'flag',
      key: 'flag',
      render: (text: string) => <code style={{ fontSize: 13, color: theme.colors.text }}>{text}</code>,
    },
    {
      title: 'Description',
      dataIndex: 'desc',
      key: 'desc',
    },
  ];

  return (
    <div>
      {/* Header */}
      <section
        style={{
          background: theme.colors.bgSecondary,
          padding: '56px 0 40px',
          borderBottom: `1px solid ${theme.colors.border}`,
        }}
      >
        <div style={{ maxWidth: theme.layout.maxWidth, margin: '0 auto', padding: '0 24px' }}>
          <Title level={1} style={{ fontSize: 36, fontWeight: 700, marginBottom: 8, color: theme.colors.text }}>
            Features
          </Title>
          <Paragraph style={{ fontSize: 17, color: theme.colors.textSecondary, maxWidth: 560 }}>
            A comprehensive Maven toolkit with everything you need to automate Java project workflows.
          </Paragraph>
        </div>
      </section>

      {/* Core Features */}
      <section style={{ padding: '64px 0' }}>
        <div style={{ maxWidth: theme.layout.maxWidth, margin: '0 auto', padding: '0 24px' }}>
          <Row gutter={[24, 24]}>
            {[
              {
                icon: <SearchOutlined />,
                title: 'Maven Finder',
                desc: 'Auto-detect Maven from PATH, M2_HOME, or Maven Wrapper. Intelligent fallback from wrapper to system Maven ensures reliability.',
                details: [
                  'FindMaven() — search system PATH and M2_HOME',
                  'FindMavenWrapper() — detect Maven Wrapper in project',
                  'FindBestMaven() — wrapper preferred, system fallback',
                ],
              },
              {
                icon: <BuildOutlined />,
                title: 'Command Builder',
                desc: 'Fluent API with 30+ Maven CLI options. Build complex Maven commands with clean, type-safe method chaining.',
                details: [
                  'Simple one-liners: Clean(), Version(), etc.',
                  'Builder pattern for complex command composition',
                  'Context support for cancellation and timeouts',
                  'Structured error handling with MavenError type',
                ],
              },
              {
                icon: <FileTextOutlined />,
                title: 'POM Parser',
                desc: 'Parse and analyze pom.xml files. Extract GAV coordinates, dependencies, properties, and repository configurations.',
                details: [
                  'ParseFile() and ParseReader() APIs',
                  'Full GAV extraction (GroupId, ArtifactId, Version)',
                  'Dependency and parent POM resolution',
                  'Properties and profile parsing',
                ],
              },
              {
                icon: <SettingOutlined />,
                title: 'Settings Parser',
                desc: 'Read Maven settings.xml configurations. Access mirrors, profiles, servers, and local repository settings.',
                details: [
                  'Parse global and user settings.xml',
                  'Access mirror configurations',
                  'Read server credentials',
                  'Profile and activation parsing',
                ],
              },
              {
                icon: <FolderOutlined />,
                title: 'Local Repository',
                desc: 'Navigate and search the local Maven repository. Find installed artifacts by GAV coordinates or group patterns.',
                details: [
                  'Browse repository structure',
                  'Search artifacts by GAV coordinates',
                  'List versions of installed artifacts',
                  'Detect repository metadata',
                ],
              },
              {
                icon: <DownloadOutlined />,
                title: 'Maven Installer',
                desc: 'Download and install Maven automatically. Supports Linux, macOS, and Windows with platform-specific handling.',
                details: [
                  'Automatic download from Apache mirrors',
                  'Platform-specific installation paths',
                  'Version selection and verification',
                  'Checksum validation',
                ],
              },
            ].map((feature, idx) => (
              <Col xs={24} md={12} key={idx}>
                <Card
                  className="feature-card"
                  bordered={false}
                  style={{
                    borderRadius: theme.radius.sm,
                    border: `1px solid ${theme.colors.border}`,
                    height: '100%',
                  }}
                  bodyStyle={{ padding: 24 }}
                >
                  <div style={{ display: 'flex', alignItems: 'center', gap: 14, marginBottom: 16 }}>
                    <div
                      style={{
                        width: 40,
                        height: 40,
                        background: theme.colors.primaryLight,
                        borderRadius: theme.radius.sm,
                        display: 'flex',
                        alignItems: 'center',
                        justifyContent: 'center',
                        color: theme.colors.primary,
                        fontSize: 18,
                      }}
                    >
                      {feature.icon}
                    </div>
                    <Title level={3} style={{ fontSize: 18, fontWeight: 600, margin: 0, color: theme.colors.text }}>
                      {feature.title}
                    </Title>
                  </div>

                  <Paragraph style={{ color: theme.colors.textSecondary, fontSize: 14, lineHeight: 1.6, marginBottom: 16 }}>
                    {feature.desc}
                  </Paragraph>

                  <div style={{ display: 'flex', flexDirection: 'column', gap: 6 }}>
                    {feature.details.map((detail, dIdx) => (
                      <div key={dIdx} style={{ display: 'flex', alignItems: 'flex-start', gap: 8 }}>
                        <Text style={{ color: theme.colors.primary, fontSize: 13, marginTop: 1 }}>•</Text>
                        <Text style={{ fontSize: 13, color: theme.colors.textSecondary, fontFamily: theme.font.mono }}>
                          {detail}
                        </Text>
                      </div>
                    ))}
                  </div>
                </Card>
              </Col>
            ))}
          </Row>
        </div>
      </section>

      {/* Command Builder Reference Table */}
      <section
        style={{
          padding: '64px 0',
          background: theme.colors.bgSecondary,
          borderTop: `1px solid ${theme.colors.border}`,
          borderBottom: `1px solid ${theme.colors.border}`,
        }}
      >
        <div style={{ maxWidth: theme.layout.maxWidth, margin: '0 auto', padding: '0 24px' }}>
          <Title level={2} style={{ fontSize: 28, fontWeight: 700, marginBottom: 8, color: theme.colors.text }}>
            Command Builder Options
          </Title>
          <Paragraph style={{ fontSize: 15, color: theme.colors.textSecondary, marginBottom: 24 }}>
            All available options for the fluent CommandBuilder API.
          </Paragraph>

          <Table
            dataSource={commandOptions}
            columns={columns}
            pagination={false}
            size="middle"
            style={{
              background: '#FFFFFF',
              borderRadius: theme.radius.sm,
              border: `1px solid ${theme.colors.border}`,
            }}
          />
        </div>
      </section>

      {/* Lifecycle Methods */}
      <section style={{ padding: '64px 0' }}>
        <div style={{ maxWidth: theme.layout.maxWidth, margin: '0 auto', padding: '0 24px' }}>
          <Row gutter={[48, 32]}>
            <Col xs={24} lg={12}>
              <Title level={3} style={{ fontSize: 22, fontWeight: 600, marginBottom: 16, color: theme.colors.text }}>
                Lifecycle Methods
              </Title>
              <Paragraph style={{ color: theme.colors.textSecondary, marginBottom: 20 }}>
                Call any Maven lifecycle phase directly on the builder:
              </Paragraph>
              <div
                style={{
                  background: theme.colors.bgDark,
                  borderRadius: theme.radius.sm,
                  padding: 20,
                  border: `1px solid ${theme.colors.borderDark}`,
                }}
              >
                <pre style={{ color: '#E2E8F0', fontFamily: theme.font.mono, fontSize: 13, lineHeight: 1.7, margin: 0 }}>
{`builder.Clean()          // mvn clean
builder.Compile()        // mvn compile
builder.Test()           // mvn test
builder.Package()        // mvn package
builder.Verify()         // mvn verify
builder.Install()        // mvn install
builder.Deploy()         // mvn deploy`}
                </pre>
              </div>
            </Col>

            <Col xs={24} lg={12}>
              <Title level={3} style={{ fontSize: 22, fontWeight: 600, marginBottom: 16, color: theme.colors.text }}>
                Multi-Phase Shortcuts
              </Title>
              <Paragraph style={{ color: theme.colors.textSecondary, marginBottom: 20 }}>
                Common multi-phase combinations built-in:
              </Paragraph>
              <div
                style={{
                  background: theme.colors.bgDark,
                  borderRadius: theme.radius.sm,
                  padding: 20,
                  border: `1px solid ${theme.colors.borderDark}`,
                }}
              >
                <pre style={{ color: '#E2E8F0', fontFamily: theme.font.mono, fontSize: 13, lineHeight: 1.7, margin: 0 }}>
{`builder.CleanInstall()  // mvn clean install  ← CI
builder.CleanPackage()  // mvn clean package
builder.CleanDeploy()   // mvn clean deploy
builder.CleanVerify()   // mvn clean verify
builder.CleanTest()     // mvn clean test`}
                </pre>
              </div>
            </Col>
          </Row>
        </div>
      </section>

      {/* Packages */}
      <section
        style={{
          padding: '64px 0',
          background: theme.colors.bgSecondary,
          borderTop: `1px solid ${theme.colors.border}`,
        }}
      >
        <div style={{ maxWidth: theme.layout.maxWidth, margin: '0 auto', padding: '0 24px' }}>
          <Title level={2} style={{ fontSize: 28, fontWeight: 700, marginBottom: 24, color: theme.colors.text }}>
            Go Packages
          </Title>
          <Row gutter={[16, 16]}>
            {[
              { pkg: 'pkg/command', desc: 'Maven command execution (builder pattern + standalone functions)' },
              { pkg: 'pkg/finder', desc: 'Find Maven executable on the system' },
              { pkg: 'pkg/installer', desc: 'Download and install Maven' },
              { pkg: 'pkg/pom', desc: 'Parse and analyze pom.xml files' },
              { pkg: 'pkg/settings', desc: 'Parse Maven settings.xml' },
              { pkg: 'pkg/local_repository', desc: 'Navigate and search local Maven repository' },
            ].map((item, idx) => (
              <Col xs={24} sm={12} md={8} key={idx}>
                <Card
                  className="feature-card"
                  bordered={false}
                  style={{
                    borderRadius: theme.radius.sm,
                    border: `1px solid ${theme.colors.border}`,
                    height: '100%',
                    background: '#FFFFFF',
                  }}
                  bodyStyle={{ padding: 20 }}
                >
                  <code style={{ color: theme.colors.primary, fontSize: 14, fontWeight: 600 }}>{item.pkg}</code>
                  <Paragraph style={{ color: theme.colors.textSecondary, fontSize: 14, marginTop: 8, marginBottom: 0 }}>
                    {item.desc}
                  </Paragraph>
                </Card>
              </Col>
            ))}
          </Row>
        </div>
      </section>
    </div>
  );
};

export default FeaturesPage;
