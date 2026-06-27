import React from 'react';
import { Layout, Typography } from 'antd';
import theme from '../styles/theme';

const { Footer: AntFooter } = Layout;
const { Text, Link } = Typography;

const Footer: React.FC = () => {
  return (
    <AntFooter
      style={{
        background: theme.colors.bgDark,
        padding: '48px 24px 32px',
        borderTop: `1px solid ${theme.colors.borderDark}`,
      }}
    >
      <div style={{ maxWidth: theme.layout.maxWidth, margin: '0 auto' }}>
        {/* Top section */}
        <div
          style={{
            display: 'grid',
            gridTemplateColumns: 'repeat(auto-fit, minmax(180px, 1fr))',
            gap: 32,
            marginBottom: 40,
          }}
        >
          <div>
            <Text strong style={{ color: '#F8FAFC', fontSize: 15, display: 'block', marginBottom: 12 }}>
              mvn-skills
            </Text>
            <Text style={{ color: '#94A3B8', fontSize: 14, lineHeight: 1.6 }}>
              Maven operations toolkit for AI agents and Go applications.
            </Text>
          </div>

          <div>
            <Text strong style={{ color: '#E2E8F0', fontSize: 14, display: 'block', marginBottom: 12 }}>
              Resources
            </Text>
            <div style={{ display: 'flex', flexDirection: 'column', gap: 8 }}>
              <Link href="https://pkg.go.dev/github.com/scagogogo/mvn-skills" style={{ color: '#94A3B8', fontSize: 14 }}>
                Go Package Reference
              </Link>
              <Link href="https://scagogogo.github.io/mvn-skills/" style={{ color: '#94A3B8', fontSize: 14 }}>
                Documentation
              </Link>
              <Link href="https://github.com/scagogogo/mvn-skills/releases" style={{ color: '#94A3B8', fontSize: 14 }}>
                Releases
              </Link>
            </div>
          </div>

          <div>
            <Text strong style={{ color: '#E2E8F0', fontSize: 14, display: 'block', marginBottom: 12 }}>
              Community
            </Text>
            <div style={{ display: 'flex', flexDirection: 'column', gap: 8 }}>
              <Link href="https://github.com/scagogogo/mvn-skills" style={{ color: '#94A3B8', fontSize: 14 }}>
                GitHub
              </Link>
              <Link href="https://github.com/scagogogo/mvn-skills/issues" style={{ color: '#94A3B8', fontSize: 14 }}>
                Issues
              </Link>
              <Link href="https://github.com/scagogogo/mvn-skills/pulls" style={{ color: '#94A3B8', fontSize: 14 }}>
                Pull Requests
              </Link>
            </div>
          </div>
        </div>

        {/* Bottom section */}
        <div
          style={{
            borderTop: '1px solid #334155',
            paddingTop: 20,
            display: 'flex',
            justifyContent: 'space-between',
            alignItems: 'center',
            flexWrap: 'wrap',
            gap: 12,
          }}
        >
          <Text style={{ color: '#64748B', fontSize: 13 }}>
            Released under the MIT License. Copyright 2024 scagogogo.
          </Text>
        </div>
      </div>
    </AntFooter>
  );
};

export default Footer;
