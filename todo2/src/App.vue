<template>
<div id="root">
  <el-card class="box-card">
    <div class="todo-wrap">
    <Top @addTodo="addTodo"/>
    <List :todos='todos' :checkTodo="checkTodo"/>
    <Bottom :todos='todos' :checkAllTodo="checkAllTodo" :clearAllTodo="clearAllTodo"/>

    </div>
  </el-card>
</div>
</template>

<script>
import pubsub from 'pubsub-js'
import Top from './components/Top.vue'
import Bottom from './components/Bottom.vue'
import List from './components/List.vue'



export default {
  name: 'App',
  components: {
    Top,Bottom,List
  },
    data(){
    return {
        todos:JSON.parse(localStorage.getItem('todo'))||[]
    }
  },
  methods:{
    //添加todo
    addTodo(todoObj){
      this.todos.unshift(todoObj)
    },
    //勾选或者取消勾选todo
    checkTodo(id){
      this.todos.forEach(todo => {
        if(todo.id===id) todo.done =!todo.done
      });
    },
    deleteTodo(_,id){
      this.todos = this.todos.filter((todo)=>{
        return todo.id !==id
        }
      )
    },
    updateTodo(_,id_title){

      this.todos.forEach(todo => {
        if(todo.id===id_title[0]) todo.value = id_title[1]
      });
      
    },
    //全部勾选或者全不勾选
    checkAllTodo(done){
        this.todos.forEach(todo => {
          todo.done = done
      });
    },
    //清除所有已经完成的todo
    clearAllTodo(){
      this.todos = this.todos.filter((todo)=>{
        return !todo.done
      })
    },
  },
  watch:{
    todos:{
      immediate:true,
      deep:true,
      handler(newval){
        localStorage.setItem('todo',JSON.stringify(newval))
      }
    }
  },
  mounted(){
    this.pubid = pubsub.subscribe('deleteTodo',this.deleteTodo)
    this.pubid2 = pubsub.subscribe('updateTodo',this.updateTodo)
//     axios.get(`http://localhost:8000/flashcards/dict_query_get/find/`).then(
//     response => {
//         console.log('请求成功了')
//         //请求成功后更新List的数据
//         console.log('数据为：'+response.data['html_result'])
//     },
//     error => {
//         //请求后更新List的数据
//         console.log('请求失败'+error)
//     }
// )
    
  },
  beforeUnmount() {
    pubsub.unsubscribe(this.pubid)
  },
}
</script>

<style>
/*base*/
body {
  background: #fff;
}

.btn {
  display: inline-block;
  padding: 4px 12px;
  margin-bottom: 0;
  font-size: 14px;
  line-height: 20px;
  text-align: center;
  vertical-align: middle;
  cursor: pointer;
  box-shadow: inset 0 1px 0 rgba(255, 255, 255, 0.2), 0 1px 2px rgba(0, 0, 0, 0.05);
  border-radius: 4px;
}

.btn-edit {
  color: #fff;
  background-color: #9edbff;
  border: 1px solid #47a3ee;
}

.btn-danger {
  color: #fff;
  background-color: #da4f49;
  border: 1px solid #bd362f;
}

.btn-danger:hover {
  color: #fff;
  background-color: #bd362f;
}

.btn:focus {
  outline: none;
}

.todo-container {
  width: 600px;
  margin: 0 auto;
}
.todo-container .todo-wrap {
  padding: 10px;
  border: 1px solid #ddd;
  border-radius: 5px;
}







</style>
