<template>



    <el-col :span="24">
      <li>
        <label>
          <el-checkbox v-model="done" size="large"/>

          <span v-show='!todo.isEdit'>{{ todo.value }}</span>
          <!--          无法输入的问题-->
          <!--          <el-input v-show='todo.isEdit' :value="todo.value" @blur="updateTodo(todo,$event)"-->
          <!--                    ref='updateTodoInput'/>-->

          <input v-show='todo.isEdit' type="text" :value="todo.value" @blur="updateTodo(todo,$event)"
                 ref='updateTodoInput'>
        </label>

        <button class="btn btn-danger" @click="handleDelete(todo.id)">删除</button>
        <button v-show='!todo.isEdit' class="btn btn-edit" @click="changeToEditable(todo)">编辑</button>
      </li>
    </el-col>


</template>

<script>
import pubsub from 'pubsub-js'
export default {
  // eslint-disable-next-line vue/multi-word-component-names
  name: 'Item',
  props:['todo','checkTodo'],
  data(){
    return {
      done:false
    }
  },
  methods:{
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
  },
  mounted() {
    // 在开始时赋予其todo的值
    this.done = this.todo.done
  },
  watch:{
    // 已完成时变化
    done(){
      this.checkTodo(this.todo.id)
    }
  },
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
