<template>
    <li>
        <label>
        <input type="checkbox" :checked="todo.done" @change="handleCheck(todo.id)"/>
        <span v-show='!todo.isEdit'>{{todo.value}}</span>
        <input v-show='todo.isEdit' type="text" :value="todo.value" @blur="updateTodo(todo,$event)" ref='updateTodoInput'>
        </label>

        <button class="btn btn-danger" @click="handleDelete(todo.id)">删除</button>
        <button v-show='!todo.isEdit' class="btn btn-edit" @click="changeToEditable(todo)">编辑</button>
    </li>

</template>

<script>
import pubsub from 'pubsub-js'
export default {
  name: 'Item',
  props:['todo','checkTodo'],
  methods:{
      handleCheck(id){
          this.checkTodo(id)
      },
      handleDelete(id){
        pubsub.publish("deleteTodo",id)
      },
      changeToEditable(todo){
        todo.isEdit = true
        
        this.$nextTick(()=>{
          this.$refs.updateTodoInput.focus()
        })
      },
      updateTodo(todo,e){
        //更新数据
        pubsub.publish("updateTodo",[todo.id,e.target.value])

        //切换Edit状态
        todo.isEdit = false
      }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
/*item*/
li {
  list-style: none;
  height: 36px;
  line-height: 36px;
  padding: 0 5px;
  border-bottom: 1px solid #ddd;
}

li label {
  float: left;
  cursor: pointer;
}

li label li input {
  vertical-align: middle;
  margin-right: 6px;
  position: relative;
  top: -1px;
}

li button {
  float: right;
  display: none;
  margin-top: 3px;
}

li:before {
  content: initial;
}

li:last-child {
  border-bottom: none;
}

li:hover button {
  float: right;
  display: block;
  margin-top: 3px;
}
</style>
