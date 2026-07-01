import { withMermaid } from 'vitepress-plugin-mermaid'

// withMermaid injects a markdown-it rule that turns ```mermaid fenced blocks
// into a client-rendered <Mermaid> component, plus the Vite/theme wiring the
// component needs. Everything else is standard VitePress config.
export default withMermaid({
  title: 'Maven SDK Go',
  description: 'A Go SDK for Maven operations',
  base: process.env.NODE_ENV === 'production' ? '/mvn-sdk/' : '/',

  // Mermaid rendering options. `theme` follows the site's light/dark mode via
  // the plugin; 'default'/'dark' are swapped automatically at runtime.
  mermaid: {
    theme: 'default',
    flowchart: { curve: 'basis', htmlLabels: true },
    securityLevel: 'loose'
  },

  // 多语言配置
  locales: {
    root: {
      label: 'English',
      lang: 'en-US'
    },
    zh: {
      label: '简体中文',
      lang: 'zh-CN',
      themeConfig: {
        nav: [
          { text: '首页', link: '/zh/' },
          { text: '架构设计', link: '/zh/architecture' },
          { text: 'API 参考', link: '/zh/api' }
        ],
        sidebar: [
          {
            text: '指南',
            items: [
              { text: '快速开始', link: '/zh/' },
              { text: '架构设计', link: '/zh/architecture' },
              { text: 'API 参考', link: '/zh/api' }
            ]
          },
          {
            text: '模块',
            items: [
              { text: 'Finder', link: '/zh/api#finder' },
              { text: 'Command', link: '/zh/api#command' },
              { text: 'POM 解析器', link: '/zh/api#pom-解析器' },
              { text: 'Settings 解析器', link: '/zh/api#settings-解析器' },
              { text: 'Local Repository', link: '/zh/api#local-repository' },
              { text: 'Installer', link: '/zh/api#installer' }
            ]
          }
        ]
      }
    }
  },

  themeConfig: {
    nav: [
      { text: 'Home', link: '/' },
      { text: 'Architecture', link: '/architecture' },
      { text: 'API Reference', link: '/api' }
    ],

    sidebar: [
      {
        text: 'Guide',
        items: [
          { text: 'Getting Started', link: '/' },
          { text: 'Architecture', link: '/architecture' },
          { text: 'API Reference', link: '/api' }
        ]
      },
      {
        text: 'Packages',
        items: [
          { text: 'Finder', link: '/api#finder' },
          { text: 'Command', link: '/api#command' },
          { text: 'POM Parser', link: '/api#pom-parser' },
          { text: 'Settings Parser', link: '/api#settings-parser' },
          { text: 'Local Repository', link: '/api#local-repository' },
          { text: 'Installer', link: '/api#installer' }
        ]
      }
    ],

    socialLinks: [
      { icon: 'github', link: 'https://github.com/scagogogo/mvn-skills' }
    ],

    footer: {
      message: 'Released under the MIT License.',
      copyright: 'Copyright © 2024 scagogogo'
    }
  }
})
