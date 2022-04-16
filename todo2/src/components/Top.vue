<template>
    <div class="todo-header">
      <el-input v-model="value" placeholder="请输入你的任务名称，按回车键确认" @keyup.enter="add" />
    </div>
</template>

<script>
    import {nanoid} from 'nanoid'
import { reactive } from '@vue/reactivity'
    export default {
      // eslint-disable-next-line vue/multi-word-component-names
    name: 'Top',
    data(){
        return {
            value:'',
        }
        },

    methods:{
        add(){
            //不能为全空格
            if(!this.value.trim()) return alert('输入不能为空')
            //将输入值包装成todo对象
            const todoObj = reactive({
                id:nanoid(),
                value:this.value,
                done:false,
                content:'',
                order:-1,
            })
            this.value = ""
            this.$emit('addTodo',todoObj)
        }
    }
    }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
/*header*/
.todo-header input {
  width: 560px;
  height: 28px;
  font-size: 14px;
  border: 1px solid #ccc;
  border-radius: 4px;
  padding: 4px 7px;
}

.todo-header input:focus {
  outline: none;
  border-color: rgba(82, 168, 236, 0.8);
  box-shadow: inset 0 1px 1px rgba(0, 0, 0, 0.075), 0 0 8px rgba(82, 168, 236, 0.6);
}
</style>
