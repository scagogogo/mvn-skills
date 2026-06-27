import React from 'react';
import { Layout, Menu, Button, Typography, Space } from 'antd';
import { GithubOutlined } from '@ant-design/icons';
import { Link, useLocation } from 'react-router-dom';
import theme from '../styles/theme';

const { Header } = Layout;
const { Text } = Typography;

const Navbar: React.FC = () => {
  const location = useLocation();

  const menuItems = [
    { key: '/', label: <Link to="/">Home</Link> },
    { key: '/features', label: <Link to="/features">Features</Link> },
    { key: '/quickstart', label: <Link to="/quickstart">Quick Start</Link> },
    { key: '/api', label: <Link to="/api">API</Link> },
  ];

  return (
    <Header
      style={{
        background: '#FFFFFF',
        borderBottom: `1px solid ${theme.colors.border}`,
        padding: '0 24px',
        position: 'sticky',
        top: 0,
        zIndex: 1000,
        display: 'flex',
        alignItems: 'center',
        justifyContent: 'space-between',
        height: 64,
        boxShadow: 'none',
      }}
    >
      <div style={{ display: 'flex', alignItems: 'center', gap: 32 }}>
        <Link to="/" style={{ display: 'flex', alignItems: 'center', gap: 10 }}>
          <div
            style={{
              width: 30,
              height: 30,
              background: theme.colors.primary,
              borderRadius: theme.radius.sm,
              display: 'flex',
              alignItems: 'center',
              justifyContent: 'center',
              color: '#FFFFFF',
              fontWeight: 700,
              fontSize: 15,
            }}
          >
            M
          </div>
          <Text strong style={{ fontSize: 17, color: theme.colors.text }}>
            mvn-skills
          </Text>
        </Link>

        <Menu
          mode="horizontal"
          selectedKeys={[location.pathname]}
          items={menuItems}
          style={{
            border: 'none',
            background: 'transparent',
            fontWeight: 500,
            minWidth: 0,
          }}
        />
      </div>

      <Space size={12}>
        <Button
          type="primary"
          href="https://github.com/scagogogo/mvn-skills"
          target="_blank"
          icon={<GithubOutlined />}
          style={{
            borderRadius: theme.radius.sm,
            fontWeight: 500,
            boxShadow: 'none',
          }}
        >
          GitHub
        </Button>
      </Space>
    </Header>
  );
};

export default Navbar;
