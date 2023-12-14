import { createApp } from 'vue';
import App from './App.vue'
import store from './store/store.js'
import router from './router/router.js'


// 检查浏览器是否支持Service Worker
if ('serviceWorker' in navigator) {
    window.addEventListener('load', function() {
      // 注册Service Worker
      navigator.serviceWorker.register('./sw.js').then(function(registration) {
        // 注册成功
        console.log('ServiceWorker registration successful with scope: ', registration.scope);
      }, function(err) {
        // 注册失败
        console.log('ServiceWorker registration failed: ', err);
      });
    });
  }

const app = createApp(App)

app.use(store)
app.use(router)
app.mount('#app')

