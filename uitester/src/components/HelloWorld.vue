<template>
  <div >
  <a-menu v-model:selectedKeys="current" mode="horizontal">
    <a-menu-item key="home">
      Anki
      <template #icon>
        <home-outlined />
      </template>
    </a-menu-item>
    <a-menu-item key="mail">
      <template #icon>
        <cloud-sync-outlined />
      </template>
      Recite
    </a-menu-item>
    <a-menu-item key="app">
      <template #icon>
        <setting-outlined />
      </template>
      Settings
    </a-menu-item>
    <a-sub-menu key="sub1">
      <template #icon>
        <appstore-outlined />
      </template>
      <template #title>Tools</template>
        <a-menu-item key="setting:1">Dict</a-menu-item>
        <a-menu-item key="setting:2">Search</a-menu-item>
        <a-menu-item key="setting:3">Web Search</a-menu-item>

    </a-sub-menu>
    <a-menu-item key="alipay">
      <template #icon>
        <read-outlined />
      </template>
      <a href="https://antdv.com" target="_blank" rel="noopener noreferrer">
        Reader
      </a>
    </a-menu-item>
  </a-menu>
  </div>
</template>
<script>
import { defineComponent, ref } from 'vue';
import { ReadOutlined,HomeOutlined, AppstoreOutlined, SettingOutlined, CloudSyncOutlined } from '@ant-design/icons-vue';
export default defineComponent({
  components: {
    ReadOutlined,
    HomeOutlined,
    AppstoreOutlined,
    SettingOutlined,
    CloudSyncOutlined,
  },

  setup() {
    const current = ref(['mail']);
    return {
      current,
    };
  },

});
</script>