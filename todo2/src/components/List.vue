<template>
  <el-container class="todo-main">

    <el-row>
        <Item v-for="todoObj in todos" :key='todoObj.id' :todo="todoObj" :checkTodo='checkTodo'
              :deleteTodo='deleteTodo'/>
      </el-row>
  </el-container>
</template>

<script>
import Item from './Item.vue'
export default {
  // eslint-disable-next-line vue/multi-word-component-names
  name: 'List',
  components:{Item,},
  props:['todos','checkTodo','deleteTodo']
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
/*main*/
.todo-main {
  margin-left: 0px;
  border: 1px solid #ddd;
  border-radius: 2px;
  padding: 0px;
}

.todo-empty {
  height: 40px;
  line-height: 40px;
  border: 1px solid #ddd;
  border-radius: 2px;
  padding-left: 5px;
  margin-top: 10px;
}
</style>
