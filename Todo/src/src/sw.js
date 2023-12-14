// sw.js 文件
self.addEventListener('fetch', function(event) {
    // 使用fetch方法从网络获取资源
    event.respondWith(fetch(event.request));
  });