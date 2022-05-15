<template>
  <a-collapse v-model:activeKey="activeKey" ghost>
    <a-collapse-panel key="1" header="This is panel header 1">
      <p>{{ text }}</p>
    </a-collapse-panel>
    <a-collapse-panel key="2" header="This is panel header 2">
      <p>{{ text }}      <a-tag color="blue">blue</a-tag></p>

    </a-collapse-panel>
    <a-collapse-panel key="3" header="This is panel header 3">
      <p>{{ text }}</p>
    </a-collapse-panel>
  </a-collapse>
  <a-row id='answer-buttons' justify="center">
    <a-button danger :size="size">Undo</a-button>
    <a-button :size="size">ForGot</a-button>
    <a-button :size="size">Diff</a-button>
    <a-button :size="size">Mid</a-button>
    <a-button :size="size">Easy</a-button>
  </a-row>
</template>
<script>
import { defineComponent, ref, watch } from 'vue';
export default defineComponent({
  setup() {
    const text = `A dog is a type of domesticated animal.Known for its loyalty and faithfulness,it can be found as a welcome guest in many households across the world.`;
    const activeKey = ref(['1']);
    watch(activeKey, val => {
      console.log(val);
    });
    return {
      text,
      activeKey,
    };
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