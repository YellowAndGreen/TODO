<template>
    <div class="todo-footer" v-show="totalTodo">
    <label>
        <input type="checkbox" v-model="isAll"/>
    </label>
    <span>
        <span>{{doneTotal}}</span> / {{totalTodo}}
    </span>
    <button class="btn btn-danger" @click="clearAll">清除已完成任务</button>
    </div>
</template>

<script>
export default {
  // eslint-disable-next-line vue/multi-word-component-names
    name: 'Bottom',
    props:['todos','checkAllTodo','clearAllTodo'],
    computed:{
        //todo总数
        totalTodo(){
            return this.todos.length
        },
        // 完成的总数
        doneTotal(){
            return this.todos.reduce((pre,todo)=>{
                return todo.done?pre+1:pre
            },0)
        },
        // 是否全部完成
        isAll:{
            get(){
                return this.doneTotal === this.totalTodo && this.totalTodo>0 
            },
            set(val){
                this.checkAllTodo(val)
            }
        },
    },
    methods:{
        clearAll(){
            this.clearAllTodo()
        }
    }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
/*footer*/
.todo-footer {
  height: 40px;
  line-height: 40px;
  padding-left: 6px;
  margin-top: 5px;
}

.todo-footer label {
  display: inline-block;
  margin-right: 20px;
  cursor: pointer;
}

.todo-footer label input {
  position: relative;
  top: -1px;
  vertical-align: middle;
  margin-right: 5px;
}

.todo-footer button {
  float: right;
  margin-top: 5px;
}
</style>
