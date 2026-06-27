import React from 'react';
import { Typography, Button, Space, Tag, Row, Col, Card } from 'antd';
import {
  SearchOutlined,
  BuildOutlined,
  FileTextOutlined,
  SettingOutlined,
  FolderOutlined,
  DownloadOutlined,
  ApiOutlined,
  LaptopOutlined,
  RobotOutlined,
  CodeOutlined,
  ThunderboltOutlined,
  ArrowRightOutlined,
  CheckCircleOutlined,
} from '@ant-design/icons';
import { Link } from 'react-router-dom';
import theme from '../styles/theme';

const { Title, Paragraph, Text } = Typography;

const HomePage: React.FC = () => {
  return (
    <div>
      {/* ===== Hero Section — 纯色背景，无渐变 ===== */}
      <section
        style={{
          background: '#F8FAFC',
          padding: '80px 0 64px',
          borderBottom: `1px solid ${theme.colors.border}`,
        }}
      >
        <div style={{ maxWidth: theme.layout.maxWidth, margin: '0 auto', padding: '0 24px' }}>
          <Row align="middle" gutter={[48, 32]}>
            <Col xs={24} lg={14}>
              <Space direction="vertical" size={16}>
                <Tag
                  style={{
                    background: theme.colors.primaryLight,
                    color: theme.colors.primary,
                    border: 'none',
                    borderRadius: theme.radius.xs,
                    padding: '4px 12px',
                    fontSize: 13,
                    fontWeight: 500,
                  }}
                >
                  Open Source · MIT License
                </Tag>

                <Title
                  style={{
                    fontSize: 44,
                    fontWeight: 700,
                    lineHeight: 1.2,
                    color: theme.colors.text,
                    margin: 0,
                    letterSpacing: '-0.02em',
                  }}
                >
                  Maven Operations{' '}
                  <span style={{ color: theme.colors.primary }}>Toolkit</span>
                </Title>

                <Paragraph
                  style={{
                    fontSize: 18,
                    color: theme.colors.textSecondary,
                    lineHeight: 1.6,
                    maxWidth: 540,
                  }}
                >
                  Execute builds, parse POM files, manage dependencies, install Maven, and automate Java project workflows — for AI agents and Go applications.
                </Paragraph>

                <Space size={12} style={{ marginTop: 8 }}>
                  <Button
                    type="primary"
                    size="large"
                    href="/quickstart"
                    style={{
                      borderRadius: theme.radius.sm,
                      fontWeight: 600,
                      height: 44,
                      paddingInline: 24,
                      fontSize: 15,
                      boxShadow: 'none',
                    }}
                  >
                    Get Started
                    <ArrowRightOutlined />
                  </Button>
                  <Button
                    size="large"
                    href="https://github.com/scagogogo/mvn-skills"
                    target="_blank"
                    style={{
                      borderRadius: theme.radius.sm,
                      fontWeight: 500,
                      height: 44,
                      paddingInline: 24,
                      fontSize: 15,
                    }}
                  >
                    View on GitHub
                  </Button>
                </Space>

                {/* Install command — 终端风格 */}
                <div
                  style={{
                    background: theme.colors.bgDark,
                    borderRadius: theme.radius.sm,
                    padding: '12px 20px',
                    display: 'flex',
                    alignItems: 'center',
                    gap: 12,
                    marginTop: 8,
                    maxWidth: 520,
                    border: `1px solid ${theme.colors.borderDark}`,
                  }}
                >
                  <span style={{ color: '#94A3B8', fontSize: 13 }}>$</span>
                  <code style={{ color: '#E2E8F0', fontSize: 14, fontFamily: theme.font.mono, background: 'none', padding: 0 }}>
                    go get github.com/scagogogo/mvn-skills@latest
                  </code>
                </div>
              </Space>
            </Col>

            <Col xs={24} lg={10}>
              <div
                style={{
                  background: theme.colors.bgDark,
                  borderRadius: theme.radius.sm,
                  border: `1px solid ${theme.colors.borderDark}`,
                  overflow: 'hidden',
                }}
              >
                {/* 终端标题栏 */}
                <div
                  style={{
                    padding: '10px 16px',
                    borderBottom: `1px solid ${theme.colors.borderDark}`,
                    display: 'flex',
                    alignItems: 'center',
                    gap: 8,
                  }}
                >
                  <div style={{ width: 12, height: 12, borderRadius: '50%', background: '#EF4444' }} />
                  <div style={{ width: 12, height: 12, borderRadius: '50%', background: '#F59E0B' }} />
                  <div style={{ width: 12, height: 12, borderRadius: '50%', background: '#22C55E' }} />
                  <span style={{ marginLeft: 8, color: '#64748B', fontSize: 12, fontFamily: theme.font.mono }}>main.go</span>
                </div>
                <pre
                  style={{
                    color: '#E2E8F0',
                    fontFamily: theme.font.mono,
                    fontSize: 13,
                    lineHeight: 1.65,
                    margin: 0,
                    padding: '16px 20px',
                    whiteSpace: 'pre-wrap',
                  }}
                >
{`package main

import (
    "fmt"
    "${"github.com/scagogogo/mvn-skills"}/pkg/command"
    "${"github.com/scagogogo/mvn-skills"}/pkg/finder"
    "${"github.com/scagogogo/mvn-skills"}/pkg/pom"
)

func main() {
    // Find Maven
    mvn, _ := finder.FindMaven()

    // Run a build
    output, _ := command.NewCommandBuilder().
        WithExecutable(mvn).
        WithBatchMode().
        WithSkipTests().
        CleanInstall()

    // Parse a POM file
    project, _ := pom.ParseFile("pom.xml")
    fmt.Printf("%s:%s:%s\\n",
        project.GroupId,
        project.ArtifactId,
        project.Version)
}`}
                </pre>
              </div>
            </Col>
          </Row>
        </div>
      </section>

      {/* ===== Features Grid — 8 功能卡片 ===== */}
      <section style={{ padding: '80px 0' }}>
        <div style={{ maxWidth: theme.layout.maxWidth, margin: '0 auto', padding: '0 24px' }}>
          <div style={{ textAlign: 'center', marginBottom: 48 }}>
            <Title
              level={2}
              style={{ fontSize: 32, fontWeight: 700, color: theme.colors.text, marginBottom: 8, letterSpacing: '-0.01em' }}
            >
              Everything You Need for Maven
            </Title>
            <Paragraph style={{ fontSize: 16, color: theme.colors.textSecondary, maxWidth: 560, margin: '0 auto' }}>
              A comprehensive toolkit that covers the full Maven workflow — from finding the executable to deploying artifacts.
            </Paragraph>
          </div>

          <Row gutter={[24, 24]}>
            {[
              { icon: <SearchOutlined />, title: 'Maven Finder', desc: 'Auto-detect Maven from PATH, M2_HOME, or Maven Wrapper with intelligent fallback.' },
              { icon: <BuildOutlined />, title: 'Command Builder', desc: 'Fluent API with 30+ Maven CLI options. Build complex commands with method chaining.' },
              { icon: <FileTextOutlined />, title: 'POM Parser', desc: 'Parse and analyze pom.xml files. Extract GAV coordinates, dependencies, and properties.' },
              { icon: <SettingOutlined />, title: 'Settings Parser', desc: 'Read Maven settings.xml configurations. Access mirrors, profiles, and server credentials.' },
              { icon: <FolderOutlined />, title: 'Local Repository', desc: 'Navigate and search the local Maven repository. Find installed artifacts by GAV.' },
              { icon: <DownloadOutlined />, title: 'Maven Installer', desc: 'Download and install Maven automatically on Linux, macOS, and Windows.' },
              { icon: <ApiOutlined />, title: 'Context Support', desc: 'Cancel and timeout Maven commands via Go\'s context.Context for safe concurrency.' },
              { icon: <LaptopOutlined />, title: 'Cross-Platform', desc: 'Full Windows, macOS, and Linux support. Tested on all major platforms.' },
            ].map((feature, idx) => (
              <Col xs={24} sm={12} lg={6} key={idx}>
                <Card
                  className="feature-card"
                  bordered={false}
                  style={{
                    height: '100%',
                    borderRadius: theme.radius.sm,
                    border: `1px solid ${theme.colors.border}`,
                  }}
                  bodyStyle={{ padding: 24 }}
                >
                  <div
                    style={{
                      width: 44,
                      height: 44,
                      background: theme.colors.primaryLight,
                      borderRadius: theme.radius.sm,
                      display: 'flex',
                      alignItems: 'center',
                      justifyContent: 'center',
                      marginBottom: 16,
                      color: theme.colors.primary,
                      fontSize: 20,
                    }}
                  >
                    {feature.icon}
                  </div>
                  <Title level={4} style={{ fontSize: 16, fontWeight: 600, marginBottom: 8, color: theme.colors.text }}>
                    {feature.title}
                  </Title>
                  <Paragraph style={{ color: theme.colors.textSecondary, fontSize: 14, lineHeight: 1.6, marginBottom: 0 }}>
                    {feature.desc}
                  </Paragraph>
                </Card>
              </Col>
            ))}
          </Row>
        </div>
      </section>

      {/* ===== Multi-Platform Section — 浅灰背景 ===== */}
      <section
        style={{
          padding: '80px 0',
          background: theme.colors.bgSecondary,
          borderTop: `1px solid ${theme.colors.border}`,
          borderBottom: `1px solid ${theme.colors.border}`,
        }}
      >
        <div style={{ maxWidth: theme.layout.maxWidth, margin: '0 auto', padding: '0 24px' }}>
          <div style={{ textAlign: 'center', marginBottom: 48 }}>
            <Title level={2} style={{ fontSize: 32, fontWeight: 700, color: theme.colors.text, marginBottom: 8 }}>
              Works Everywhere You Need It
            </Title>
            <Paragraph style={{ fontSize: 16, color: theme.colors.textSecondary, maxWidth: 560, margin: '0 auto' }}>
              Use mvn-skills as a Go SDK, CLI tool, AI agent plugin, or MCP server.
            </Paragraph>
          </div>

          <Row gutter={[24, 24]}>
            {[
              {
                icon: <RobotOutlined />,
                title: 'AI Agents',
                tag: 'Claude Code / OpenCode',
                desc: 'Install as a skill plugin and let your AI agent run Maven commands, parse POMs, and manage dependencies autonomously.',
                install: 'claude plugin install maven-skills@mvn-skills',
              },
              {
                icon: <CodeOutlined />,
                title: 'Go SDK',
                tag: 'Go Applications',
                desc: 'Import the Go package and build Maven automation into your applications with a fluent, type-safe API.',
                install: 'go get github.com/scagogogo/mvn-skills@latest',
              },
              {
                icon: <CodeOutlined />,
                title: 'CLI',
                tag: 'Standalone Binary',
                desc: 'Download a single binary with zero dependencies. Run Maven operations from any shell script or CI pipeline.',
                install: 'curl -sL .../mvn-skills-latest.tar.gz | tar -xz',
              },
              {
                icon: <ThunderboltOutlined />,
                title: 'MCP Server',
                tag: 'Any AI Tool',
                desc: 'Wrap the Go SDK as an MCP server to provide Maven operations to any MCP-compatible AI tool or editor.',
                install: 'See documentation for setup details',
              },
            ].map((platform, idx) => (
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
                  <div style={{ display: 'flex', alignItems: 'center', gap: 16, marginBottom: 16 }}>
                    <div
                      style={{
                        width: 48,
                        height: 48,
                        background: theme.colors.primaryLight,
                        borderRadius: theme.radius.sm,
                        display: 'flex',
                        alignItems: 'center',
                        justifyContent: 'center',
                        color: theme.colors.primary,
                        fontSize: 22,
                      }}
                    >
                      {platform.icon}
                    </div>
                    <div>
                      <Title level={4} style={{ fontSize: 18, fontWeight: 600, margin: 0, color: theme.colors.text }}>
                        {platform.title}
                      </Title>
                      <Tag
                        style={{
                          background: theme.colors.accentLight,
                          color: theme.colors.accent,
                          border: 'none',
                          borderRadius: theme.radius.xs,
                          fontSize: 12,
                          marginTop: 4,
                        }}
                      >
                        {platform.tag}
                      </Tag>
                    </div>
                  </div>

                  <Paragraph style={{ color: theme.colors.textSecondary, fontSize: 14, lineHeight: 1.6, marginBottom: 16 }}>
                    {platform.desc}
                  </Paragraph>

                  <div
                    style={{
                      background: theme.colors.bgDark,
                      borderRadius: theme.radius.sm,
                      padding: '10px 16px',
                      display: 'flex',
                      alignItems: 'center',
                      gap: 8,
                      border: `1px solid ${theme.colors.borderDark}`,
                    }}
                  >
                    <span style={{ color: '#64748B', fontSize: 12 }}>$</span>
                    <code style={{ color: '#E2E8F0', fontSize: 13, fontFamily: theme.font.mono, background: 'none', padding: 0 }}>
                      {platform.install}
                    </code>
                  </div>
                </Card>
              </Col>
            ))}
          </Row>
        </div>
      </section>

      {/* ===== Command Builder Example — 左文右码 ===== */}
      <section style={{ padding: '80px 0' }}>
        <div style={{ maxWidth: theme.layout.maxWidth, margin: '0 auto', padding: '0 24px' }}>
          <Row gutter={[48, 32]} align="middle">
            <Col xs={24} lg={10}>
              <Title level={2} style={{ fontSize: 32, fontWeight: 700, marginBottom: 16, color: theme.colors.text }}>
                Powerful Command Builder
              </Title>
              <Paragraph style={{ fontSize: 16, color: theme.colors.textSecondary, lineHeight: 1.6, marginBottom: 24 }}>
                Chain method calls to build complex Maven commands. The fluent API makes it easy to compose exactly the build you need.
              </Paragraph>

              <Space direction="vertical" size={12}>
                {[
                  'Fluent builder pattern with method chaining',
                  '30+ Maven CLI options as type-safe methods',
                  'Context support for cancellation and timeouts',
                  'Cross-platform path and executable handling',
                ].map((item, idx) => (
                  <div key={idx} style={{ display: 'flex', alignItems: 'center', gap: 10 }}>
                    <CheckCircleOutlined style={{ color: theme.colors.success, fontSize: 16 }} />
                    <Text style={{ fontSize: 15, color: theme.colors.textSecondary }}>{item}</Text>
                  </div>
                ))}
              </Space>
            </Col>

            <Col xs={24} lg={14}>
              <div
                style={{
                  background: theme.colors.bgDark,
                  borderRadius: theme.radius.sm,
                  overflow: 'hidden',
                  border: `1px solid ${theme.colors.borderDark}`,
                }}
              >
                {/* 文件标签栏 */}
                <div
                  style={{
                    display: 'flex',
                    borderBottom: `1px solid ${theme.colors.borderDark}`,
                    background: theme.colors.bgDarkSecondary,
                  }}
                >
                  {['main.go', 'command.go', 'finder.go'].map((tab, idx) => (
                    <div
                      key={tab}
                      style={{
                        padding: '8px 20px',
                        fontSize: 13,
                        fontFamily: theme.font.mono,
                        color: idx === 0 ? '#E2E8F0' : '#64748B',
                        borderBottom: idx === 0 ? `2px solid ${theme.colors.primary}` : '2px solid transparent',
                        background: idx === 0 ? theme.colors.bgDark : 'transparent',
                      }}
                    >
                      {tab}
                    </div>
                  ))}
                </div>

                <pre
                  style={{
                    color: '#E2E8F0',
                    fontFamily: theme.font.mono,
                    fontSize: 13,
                    lineHeight: 1.65,
                    margin: 0,
                    padding: 20,
                  }}
                >
{`// Simple one-liners
output, err := command.Clean("mvn")
output, err := command.Version("mvn")

// Builder pattern for complex builds
output, err := command.NewCommandBuilder().
    WithWorkingDirectory("/path/to/project").
    WithBatchMode().
    WithSkipTests().
    WithProfiles("ci", "release").
    WithProperty("skipITs", "true").
    CleanDeploy()

// With timeout and cancellation
ctx, cancel := context.WithTimeout(
    context.Background(), 5*time.Minute,
)
defer cancel()

output, err := command.NewCommandBuilder().
    WithContext(ctx).
    CleanInstall()`}
                </pre>
              </div>
            </Col>
          </Row>
        </div>
      </section>

      {/* ===== CTA Section — 深色背景 ===== */}
      <section
        style={{
          padding: '80px 0',
          background: theme.colors.bgDark,
        }}
      >
        <div style={{ maxWidth: 640, margin: '0 auto', padding: '0 24px', textAlign: 'center' }}>
          <Title level={2} style={{ fontSize: 32, fontWeight: 700, color: '#F8FAFC', marginBottom: 16 }}>
            Ready to Automate Maven?
          </Title>
          <Paragraph style={{ fontSize: 16, color: '#94A3B8', marginBottom: 32, lineHeight: 1.6 }}>
            Get started in seconds. Install mvn-skills and start building, parsing, and deploying with Maven.
          </Paragraph>
          <Space size={12}>
            <Button
              type="primary"
              size="large"
              style={{
                borderRadius: theme.radius.sm,
                fontWeight: 600,
                height: 44,
                paddingInline: 24,
                fontSize: 15,
                boxShadow: 'none',
              }}
            >
              <Link to="/quickstart" style={{ color: '#fff' }}>
                Quick Start Guide
              </Link>
            </Button>
            <Button
              size="large"
              ghost
              href="https://github.com/scagogogo/mvn-skills"
              target="_blank"
              style={{
                borderRadius: theme.radius.sm,
                fontWeight: 500,
                height: 44,
                paddingInline: 24,
                fontSize: 15,
                borderColor: '#475569',
                color: '#E2E8F0',
              }}
            >
              Star on GitHub
            </Button>
          </Space>
        </div>
      </section>
    </div>
  );
};

export default HomePage;
