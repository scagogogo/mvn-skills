import React from 'react';
import { Typography, Row, Col, Card, Tabs } from 'antd';
import { RobotOutlined, CodeOutlined, BugOutlined, ThunderboltOutlined } from '@ant-design/icons';
import theme from '../styles/theme';

const { Title, Paragraph, Text } = Typography;

const QuickStartPage: React.FC = () => {
  const goSteps = [
    {
      title: 'Install the package',
      desc: 'Add mvn-skills to your Go project.',
      code: `go get github.com/scagogogo/mvn-skills@latest`,
    },
    {
      title: 'Find Maven',
      desc: 'Auto-detect the Maven executable on your system.',
      code: `mvn, err := finder.FindMaven()
if err != nil {
    log.Fatal("Maven not found:", err)
}
fmt.Println("Found Maven at:", mvn)`,
    },
    {
      title: 'Run a build',
      desc: 'Use the fluent CommandBuilder to execute Maven commands.',
      code: `output, err := command.NewCommandBuilder().
    WithExecutable(mvn).
    WithBatchMode().
    WithSkipTests().
    CleanInstall()
if err != nil {
    log.Fatal("Build failed:", err)
}
fmt.Println(output)`,
    },
    {
      title: 'Parse a POM file',
      desc: 'Extract project metadata from pom.xml.',
      code: `project, err := pom.ParseFile("pom.xml")
if err != nil {
    log.Fatal("Parse failed:", err)
}
fmt.Printf("%s:%s:%s\\n",
    project.GroupId,
    project.ArtifactId,
    project.Version)`,
    },
  ];

  const aiSteps = [
    {
      title: 'Add marketplace',
      desc: 'Add the skill marketplace to Claude Code.',
      code: `claude plugin marketplace add scagogogo/mvn-skills`,
    },
    {
      title: 'Install the plugin',
      desc: 'Install the Maven operations skill.',
      code: `claude plugin install maven-skills@mvn-skills`,
    },
    {
      title: 'Start using',
      desc: 'Your AI agent can now run Maven commands, parse POMs, and manage dependencies.',
      code: `"Find Maven on my system and run a clean install"
"Parse the pom.xml and show me the dependencies"
"Install Maven 3.9.6"`,
    },
  ];

  const cliSteps = [
    {
      title: 'Download binary',
      desc: 'Get the latest release for your platform.',
      code: `curl -sL https://github.com/scagogogo/mvn-skills/releases/latest/download/mvn-skills-latest.tar.gz | tar -xz`,
    },
    {
      title: 'Run commands',
      desc: 'Execute Maven operations from the CLI.',
      code: `./mvn-skills find        # Find Maven
./mvn-skills version     # Check Maven version
./mvn-skills build clean install  # Run a build`,
    },
  ];

  const mcpSteps = [
    {
      title: 'Configure MCP server',
      desc: 'Add the MCP server configuration to your AI tool.',
      code: `{
  "mcpServers": {
    "mvn-skills": {
      "command": "mvn-skills",
      "args": ["mcp", "serve"]
    }
  }
}`,
    },
    {
      title: 'Use with any MCP client',
      desc: 'Connect from any MCP-compatible tool like Claude Desktop, Cursor, or Windsurf.',
      code: `// The MCP server exposes tools for:
// - Finding Maven executables
// - Running Maven commands
// - Parsing POM files
// - Managing dependencies`,
    },
  ];

  const renderStep = (step: typeof goSteps[0], idx: number) => (
    <div key={idx} style={{ marginBottom: 32 }}>
      <div style={{ display: 'flex', alignItems: 'center', gap: 12, marginBottom: 8 }}>
        <div
          style={{
            width: 28,
            height: 28,
            background: theme.colors.primary,
            borderRadius: theme.radius.sm,
            display: 'flex',
            alignItems: 'center',
            justifyContent: 'center',
            color: '#FFFFFF',
            fontSize: 14,
            fontWeight: 600,
            flexShrink: 0,
          }}
        >
          {idx + 1}
        </div>
        <Title level={4} style={{ fontSize: 18, fontWeight: 600, margin: 0, color: theme.colors.text }}>
          {step.title}
        </Title>
      </div>
      <Paragraph style={{ color: theme.colors.textSecondary, fontSize: 15, marginLeft: 40, marginBottom: 12 }}>
        {step.desc}
      </Paragraph>
      <div style={{ marginLeft: 40 }}>
        <div
          style={{
            background: theme.colors.bgDark,
            borderRadius: theme.radius.sm,
            padding: 16,
            border: `1px solid ${theme.colors.borderDark}`,
          }}
        >
          <pre style={{ color: '#E2E8F0', fontFamily: theme.font.mono, fontSize: 13, lineHeight: 1.65, margin: 0 }}>
            {step.code}
          </pre>
        </div>
      </div>
    </div>
  );

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
            Quick Start
          </Title>
          <Paragraph style={{ fontSize: 17, color: theme.colors.textSecondary, maxWidth: 560 }}>
            Get up and running with mvn-skills in minutes. Choose your preferred integration method.
          </Paragraph>
        </div>
      </section>

      {/* Steps */}
      <section style={{ padding: '64px 0' }}>
        <div style={{ maxWidth: 880, margin: '0 auto', padding: '0 24px' }}>
          <Tabs
            defaultActiveKey="go"
            size="large"
            items={[
              {
                key: 'go',
                label: (
                  <span style={{ display: 'flex', alignItems: 'center', gap: 8 }}>
                    <CodeOutlined /> Go SDK
                  </span>
                ),
                children: goSteps.map(renderStep),
              },
              {
                key: 'ai',
                label: (
                  <span style={{ display: 'flex', alignItems: 'center', gap: 8 }}>
                    <RobotOutlined /> AI Agents
                  </span>
                ),
                children: aiSteps.map(renderStep),
              },
              {
                key: 'cli',
                label: (
                  <span style={{ display: 'flex', alignItems: 'center', gap: 8 }}>
                    <BugOutlined /> CLI
                  </span>
                ),
                children: cliSteps.map(renderStep),
              },
              {
                key: 'mcp',
                label: (
                  <span style={{ display: 'flex', alignItems: 'center', gap: 8 }}>
                    <ThunderboltOutlined /> MCP Server
                  </span>
                ),
                children: mcpSteps.map(renderStep),
              },
            ]}
          />
        </div>
      </section>

      {/* Advanced Example */}
      <section
        style={{
          padding: '64px 0',
          background: theme.colors.bgSecondary,
          borderTop: `1px solid ${theme.colors.border}`,
        }}
      >
        <div style={{ maxWidth: theme.layout.maxWidth, margin: '0 auto', padding: '0 24px' }}>
          <Title level={2} style={{ fontSize: 28, fontWeight: 700, marginBottom: 16, color: theme.colors.text }}>
            Advanced Example
          </Title>
          <Paragraph style={{ fontSize: 16, color: theme.colors.textSecondary, marginBottom: 24 }}>
            A complete CI/CD pipeline example with timeout, cancellation, and error handling.
          </Paragraph>

          <Row gutter={[32, 24]}>
            <Col xs={24} lg={14}>
              <div
                style={{
                  background: theme.colors.bgDark,
                  borderRadius: theme.radius.sm,
                  overflow: 'hidden',
                  border: `1px solid ${theme.colors.borderDark}`,
                }}
              >
                <div
                  style={{
                    padding: '8px 20px',
                    background: theme.colors.bgDarkSecondary,
                    borderBottom: `1px solid ${theme.colors.borderDark}`,
                    fontSize: 13,
                    fontFamily: theme.font.mono,
                    color: '#64748B',
                  }}
                >
                  ci_pipeline.go
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
{`package main

import (
    "context"
    "errors"
    "fmt"
    "log"
    "time"

    "github.com/scagogogo/mvn-skills/pkg/command"
    "github.com/scagogogo/mvn-skills/pkg/finder"
    "github.com/scagogogo/mvn-skills/pkg/pom"
)

func main() {
    // 1. Find Maven with fallback
    mvn, err := finder.FindBestMaven(".")
    if err != nil {
        log.Fatal("Maven not found:", err)
    }

    // 2. Parse project info
    project, _ := pom.ParseFile("pom.xml")
    log.Printf("Building %s:%s:%s",
        project.GroupId,
        project.ArtifactId,
        project.Version)

    // 3. Build with timeout
    ctx, cancel := context.WithTimeout(
        context.Background(),
        10*time.Minute,
    )
    defer cancel()

    output, err := command.NewCommandBuilder().
        WithExecutable(mvn).
        WithContext(ctx).
        WithBatchMode().
        WithSkipTests().
        WithProfiles("ci").
        CleanDeploy()

    // 4. Handle errors
    if err != nil {
        var me *command.MavenError
        if errors.As(err, &me) {
            log.Printf("Exit code: %d", me.ExitCode)
            log.Printf("Stderr: %s", me.Stderr)
        }
        log.Fatal("Build failed")
    }

    fmt.Println("Build succeeded!")
    fmt.Println(output)
}`}
                </pre>
              </div>
            </Col>

            <Col xs={24} lg={10}>
              <div style={{ display: 'flex', flexDirection: 'column', gap: 16 }}>
                {[
                  { num: '1', title: 'Smart Maven Discovery', desc: 'FindBestMaven() tries Maven Wrapper first, falls back to system Maven — perfect for CI environments.' },
                  { num: '2', title: 'Project Metadata', desc: 'Parse POM files to extract GAV coordinates, dependencies, and properties before building.' },
                  { num: '3', title: 'Timeout & Cancellation', desc: 'Go context.Context integration ensures builds don\'t hang forever in CI pipelines.' },
                  { num: '4', title: 'Structured Error Handling', desc: 'MavenError type provides exit codes, command details, and stderr for precise failure diagnosis.' },
                ].map((item) => (
                  <Card
                    key={item.num}
                    className="feature-card"
                    bordered={false}
                    style={{
                      borderRadius: theme.radius.sm,
                      border: `1px solid ${theme.colors.border}`,
                    }}
                    bodyStyle={{ padding: 20 }}
                  >
                    <div style={{ display: 'flex', gap: 14 }}>
                      <div
                        style={{
                          width: 28,
                          height: 28,
                          background: theme.colors.primaryLight,
                          borderRadius: theme.radius.sm,
                          display: 'flex',
                          alignItems: 'center',
                          justifyContent: 'center',
                          color: theme.colors.primary,
                          fontSize: 14,
                          fontWeight: 600,
                          flexShrink: 0,
                        }}
                      >
                        {item.num}
                      </div>
                      <div>
                        <Text strong style={{ fontSize: 15, display: 'block', marginBottom: 4, color: theme.colors.text }}>
                          {item.title}
                        </Text>
                        <Text style={{ fontSize: 14, color: theme.colors.textSecondary, lineHeight: 1.5 }}>
                          {item.desc}
                        </Text>
                      </div>
                    </div>
                  </Card>
                ))}
              </div>
            </Col>
          </Row>
        </div>
      </section>
    </div>
  );
};

export default QuickStartPage;
