<template>
  <HelloWorld msg="Welcome to Your Vue.js App"/>
  <CollaspList/>
</template>

<script>
import HelloWorld from './components/HelloWorld.vue'
import CollaspList from './components/CollaspList'

export default {
  name: 'App',
  components: {
    CollaspList,
    HelloWorld
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
}
</style>
