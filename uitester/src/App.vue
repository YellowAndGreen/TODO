<template>
  <HelloWorld msg="Welcome to Your Vue.js App"/>
  <a-row justify="center" id="card">
  <CardView />
  </a-row>
</template>

<script>
import HelloWorld from './components/HelloWorld.vue'
import CardView from "@/components/CardView";

export default {
  name: 'App',
  components: {
    HelloWorld,
    CardView
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
#card{
  position: fixed;
  height: 100%;
  width: 100%;
  align-self: center;
}
html,body{height:100%;margin:0;}
</style>
