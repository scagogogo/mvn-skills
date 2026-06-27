import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import { ConfigProvider, Layout } from 'antd';
import Navbar from './components/Navbar';
import Footer from './components/Footer';
import HomePage from './pages/HomePage';
import FeaturesPage from './pages/FeaturesPage';
import QuickStartPage from './pages/QuickStartPage';
import ApiPage from './pages/ApiPage';
import './styles/global.css';

const { Content } = Layout;

// AntDesign theme — flat design, Slate palette, no purple, 4px radius
const antTheme = {
  token: {
    colorPrimary: '#2563EB',
    colorInfo: '#2563EB',
    colorSuccess: '#16A34A',
    colorWarning: '#CA8A04',
    colorError: '#DC2626',
    borderRadius: 4,
    fontFamily: `-apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, "Helvetica Neue", Arial, sans-serif`,
    colorText: '#0F172A',
    colorTextSecondary: '#475569',
    colorBorder: '#E2E8F0',
    colorBgContainer: '#FFFFFF',
    colorBgLayout: '#FFFFFF',
    boxShadow: 'none',
  },
  components: {
    Button: {
      borderRadius: 4,
      controlHeight: 40,
      boxShadow: 'none',
    },
    Card: {
      borderRadius: 4,
      boxShadow: 'none',
    },
    Input: {
      borderRadius: 4,
    },
    Menu: {
      itemBorderRadius: 4,
      horizontalItemHoverBg: '#F8FAFC',
      horizontalItemSelectedBg: '#EFF6FF',
      horizontalItemSelectedColor: '#2563EB',
    },
    Tabs: {
      itemSelectedColor: '#2563EB',
      inkBarColor: '#2563EB',
    },
    Table: {
      borderRadius: 4,
      headerBg: '#F8FAFC',
      headerColor: '#475569',
      borderColor: '#E2E8F0',
    },
    Tag: {
      borderRadiusSM: 2,
    },
  },
};

const App: React.FC = () => {
  return (
    <ConfigProvider theme={antTheme}>
      <Router basename={process.env.PUBLIC_URL}>
        <Layout style={{ minHeight: '100vh', background: '#FFFFFF' }}>
          <Navbar />
          <Content>
            <Routes>
              <Route path="/" element={<HomePage />} />
              <Route path="/features" element={<FeaturesPage />} />
              <Route path="/quickstart" element={<QuickStartPage />} />
              <Route path="/api" element={<ApiPage />} />
            </Routes>
          </Content>
          <Footer />
        </Layout>
      </Router>
    </ConfigProvider>
  );
};

export default App;
