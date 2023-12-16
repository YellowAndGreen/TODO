<template>
  <router-view></router-view>
</template>
<script>
export default {
  name: 'App',
}
</script>
<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}
html,
body,
div,
ul,
h1,
h2,
h3,
p {
  padding: 0;
  margin: 0;
}
.w100p {
  width: 100%;
}
.txt-r {
  text-align: right;
}
.flex-center {
  display: flex;
  align-items: center;
  justify-content: center;
}
</style>