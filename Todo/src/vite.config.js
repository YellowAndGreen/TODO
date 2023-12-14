import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import { VitePWA } from 'vite-plugin-pwa'

// https://vitejs.dev/config/
export default defineConfig({
  base: './',
  plugins: [vue(),
    VitePWA({ // 使用PWA插件
      manifest: { // 定义manifest.json
        name: 'TODO LIST', // 应用的名称
        short_name: 'TODO', // 应用的简短名称
        description: 'My awesome PWA App', // 应用的描述
        theme_color: '#0000FF', // 应用的主题颜色
        icons: [ // 应用的图标
          {
            src: '/public/favicon.ico',
            sizes: '32x32',
            type: 'image/x-icon',
          }
        ],
        registerType: "autoUpdate",
        devOptions: {
          enabled: true,
        },
        workbox: {
          globPatterns: ["**/*.{js,css,html,ico,png,svg}"],
        },
      },
    }),
  ]
})
