<template>
  <a-card hoverable style="width:80%" title="abandon">
    <template #extra><a-tag color="blue">blue</a-tag></template>
    <a-collapse v-model:activeKey="activeKey" ghost>
      <a-collapse-panel key="1" header="This is panel header 1">
        <p>111</p>
      </a-collapse-panel>
      <a-collapse-panel key="2" header="This is panel header 2">
        <p>222</p>
      </a-collapse-panel>
      <a-collapse-panel key="3" header="This is panel header 3">
        <p>333</p>
      </a-collapse-panel>
    </a-collapse>
    <template #actions>
      <setting-outlined key="setting" />
      <edit-outlined key="edit" />
      <ellipsis-outlined key="ellipsis" />
    </template>
    <a-row id='answer-buttons' justify="center">
      <a-button danger :size="size">Undo</a-button>
      <a-button :size="size">ForGot</a-button>
      <a-button :size="size">Diff</a-button>
      <a-button :size="size">Mid</a-button>
      <a-button :size="size">Easy</a-button>
    </a-row>
  </a-card>
</template>
<script>
import { SettingOutlined, EditOutlined, EllipsisOutlined } from '@ant-design/icons-vue';
import { defineComponent } from 'vue';
export default defineComponent({
  components: {
    SettingOutlined,
    EditOutlined,
    EllipsisOutlined,
  },
});
</script>

<style scoped>
#answer-buttons{
  position: fixed;
  left:0;right:0;bottom:0;
  margin:auto;
}
</style>
