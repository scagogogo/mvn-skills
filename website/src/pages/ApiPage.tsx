import React from 'react';
import { Typography, Row, Col, Card } from 'antd';
import theme from '../styles/theme';

const { Title, Paragraph, Text } = Typography;

const ApiPage: React.FC = () => {
  const packages = [
    {
      name: 'pkg/command',
      desc: 'Maven command execution with fluent builder pattern and standalone convenience functions.',
      functions: [
        { sig: 'Clean(mvn string) (string, error)', desc: 'Execute mvn clean' },
        { sig: 'Compile(mvn string) (string, error)', desc: 'Execute mvn compile' },
        { sig: 'Test(mvn string) (string, error)', desc: 'Execute mvn test' },
        { sig: 'Package(mvn string) (string, error)', desc: 'Execute mvn package' },
        { sig: 'Verify(mvn string) (string, error)', desc: 'Execute mvn verify' },
        { sig: 'Install(mvn string) (string, error)', desc: 'Execute mvn install' },
        { sig: 'Deploy(mvn string) (string, error)', desc: 'Execute mvn deploy' },
        { sig: 'Version(mvn string) (string, error)', desc: 'Execute mvn version' },
        { sig: 'DependencyGet(mvn, groupId, artifactId, version string) (string, error)', desc: 'Download an artifact' },
        { sig: 'ParseVersion(output string) (*MavenVersion, error)', desc: 'Parse version from mvn -version output' },
        { sig: 'NewCommandBuilder() *CommandBuilder', desc: 'Create a new fluent command builder' },
      ],
    },
    {
      name: 'pkg/finder',
      desc: 'Find Maven executable on the system using multiple detection strategies.',
      functions: [
        { sig: 'FindMaven() (string, error)', desc: 'Search PATH and M2_HOME for Maven' },
        { sig: 'FindMavenWrapper(projectDir string) (string, error)', desc: 'Find Maven Wrapper in project directory' },
        { sig: 'FindBestMaven(projectDir string) (string, error)', desc: 'Wrapper preferred, system Maven fallback' },
      ],
    },
    {
      name: 'pkg/pom',
      desc: 'Parse and analyze Maven POM files with full XML structure support.',
      functions: [
        { sig: 'ParseFile(path string) (*Project, error)', desc: 'Parse POM from file path' },
        { sig: 'ParseReader(r io.Reader) (*Project, error)', desc: 'Parse POM from io.Reader' },
      ],
      types: [
        { name: 'Project', fields: 'GroupId, ArtifactId, Version, Packaging, Dependencies, Properties, Parent, Repositories, Profiles' },
        { name: 'Dependency', fields: 'GroupId, ArtifactId, Version, Scope, Type, Classifier, Optional' },
      ],
    },
    {
      name: 'pkg/settings',
      desc: 'Parse Maven settings.xml for mirrors, servers, and profile configurations.',
      functions: [
        { sig: 'ParseFile(path string) (*Settings, error)', desc: 'Parse settings.xml from file path' },
        { sig: 'ParseReader(r io.Reader) (*Settings, error)', desc: 'Parse settings.xml from io.Reader' },
      ],
    },
    {
      name: 'pkg/installer',
      desc: 'Download and install Apache Maven on Linux, macOS, and Windows.',
      functions: [
        { sig: 'Install(version string) (string, error)', desc: 'Download and install a specific Maven version' },
        { sig: 'InstallLatest() (string, error)', desc: 'Install the latest Maven version' },
      ],
    },
    {
      name: 'pkg/local_repository',
      desc: 'Navigate and search the local Maven repository (~/.m2/repository).',
      functions: [
        { sig: 'GetLocalRepositoryPath() (string, error)', desc: 'Find the local repository path' },
        { sig: 'FindArtifact(groupId, artifactId string) ([]string, error)', desc: 'Find installed versions of an artifact' },
      ],
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
            API Reference
          </Title>
          <Paragraph style={{ fontSize: 17, color: theme.colors.textSecondary, maxWidth: 560 }}>
            Complete Go API reference for all mvn-skills packages.
          </Paragraph>
        </div>
      </section>

      {/* Packages */}
      <section style={{ padding: '64px 0' }}>
        <div style={{ maxWidth: theme.layout.maxWidth, margin: '0 auto', padding: '0 24px' }}>
          <Row gutter={[24, 24]}>
            {packages.map((pkg, idx) => (
              <Col xs={24} key={idx}>
                <Card
                  className="feature-card"
                  bordered={false}
                  style={{
                    borderRadius: theme.radius.sm,
                    border: `1px solid ${theme.colors.border}`,
                  }}
                  bodyStyle={{ padding: 24 }}
                >
                  <div style={{ marginBottom: 20 }}>
                    <code style={{ color: theme.colors.primary, fontSize: 18, fontWeight: 600 }}>
                      {pkg.name}
                    </code>
                    <Paragraph style={{ color: theme.colors.textSecondary, fontSize: 14, marginTop: 8, marginBottom: 0 }}>
                      {pkg.desc}
                    </Paragraph>
                  </div>

                  {/* Functions */}
                  <div style={{ marginBottom: pkg.types ? 20 : 0 }}>
                    <Text strong style={{ fontSize: 13, color: theme.colors.textMuted, marginBottom: 8, display: 'block', textTransform: 'uppercase', letterSpacing: '0.05em' }}>
                      Functions
                    </Text>
                    <div
                      style={{
                        background: theme.colors.bgSecondary,
                        borderRadius: theme.radius.sm,
                        border: `1px solid ${theme.colors.border}`,
                        overflow: 'hidden',
                      }}
                    >
                      {pkg.functions.map((fn, fIdx) => (
                        <div
                          key={fIdx}
                          style={{
                            padding: '8px 16px',
                            borderBottom: fIdx < pkg.functions.length - 1 ? `1px solid ${theme.colors.border}` : 'none',
                            display: 'flex',
                            justifyContent: 'space-between',
                            alignItems: 'flex-start',
                            gap: 16,
                          }}
                        >
                          <code style={{ fontSize: 13, fontFamily: theme.font.mono, color: theme.colors.text, flexShrink: 0 }}>
                            {fn.sig}
                          </code>
                          <Text style={{ fontSize: 13, color: theme.colors.textSecondary, textAlign: 'right' }}>
                            {fn.desc}
                          </Text>
                        </div>
                      ))}
                    </div>
                  </div>

                  {/* Types */}
                  {pkg.types && (
                    <div>
                      <Text strong style={{ fontSize: 13, color: theme.colors.textMuted, marginBottom: 8, display: 'block', textTransform: 'uppercase', letterSpacing: '0.05em' }}>
                        Types
                      </Text>
                      <div
                        style={{
                          background: theme.colors.bgSecondary,
                          borderRadius: theme.radius.sm,
                          border: `1px solid ${theme.colors.border}`,
                          overflow: 'hidden',
                        }}
                      >
                        {pkg.types.map((t, tIdx) => (
                          <div
                            key={tIdx}
                            style={{
                              padding: '8px 16px',
                              borderBottom: tIdx < pkg.types.length - 1 ? `1px solid ${theme.colors.border}` : 'none',
                            }}
                          >
                            <code style={{ fontSize: 14, fontWeight: 600, color: theme.colors.primary }}>{t.name}</code>
                            <div style={{ marginTop: 4 }}>
                              <Text style={{ fontSize: 13, color: theme.colors.textSecondary, fontFamily: theme.font.mono }}>
                                {t.fields}
                              </Text>
                            </div>
                          </div>
                        ))}
                      </div>
                    </div>
                  )}
                </Card>
              </Col>
            ))}
          </Row>
        </div>
      </section>
    </div>
  );
};

export default ApiPage;
