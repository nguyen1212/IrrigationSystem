<template>
  <div class="home" id="home">
    <h1> Implement home in this directory</h1>
    <p>../components/home/Home.vue</p>
    <p> 
      <router-link to='/login'> goto login</router-link>
      <router-link to='/home'> goto home</router-link>
    </p>
  </div>
</template>
<script>
export default {
  name: 'home',
  data () {
    return {
      msg: 'Welcome to Your Vue.js App'
    }
  }
}
</script>
<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>