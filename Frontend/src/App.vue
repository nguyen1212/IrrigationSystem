<template>
  <div id="app" class="container-fluid h-100">
    <div class="d-flex flex-column h-100 ">
    <div class="row ">
      <panel/>
      <!-- <div class="col-md-6 offset-md-3 py-5">
        <h1>Generate a thumbnail of a website</h1>     
        <form v-on:submit.prevent="makeWebsiteThumbnail">
          <div class="form-group">
            <input v-model="websiteUrl" type="text" id="website-input" placeholder="Enter a website" class="form-control">
          </div>
          <div class="form-group">
            <button class="btn btn-primary">Generate!</button>
          </div>
        </form>
        <img :src="thumbnailUrl"/>
      </div> -->
    </div>
    <div class="row flex-grow-1" >
      <!-- <component v-bind:is="component" />
      <button v-on:click="toggle">Toggle</button> -->
      <div class="col-4">
        <testmenu style="height:100%; width: 100%;" v-on:changefeature="switchFeature"></testmenu>
      </div>
      <div class="col">
        <component v-bind:is="component" />
      </div>
    </div>
    </div>
</template>

<script>

import axios from 'axios';
import home from './components/HelloWorld.vue';
import Panel from './components/panel/panel.vue';
import Menu from './components/menu/menu.vue';
export default {
  name: 'App',

  data() { return {
    websiteUrl: '',
    thumbnailUrl: '',
    component: '',
    // panel: Panel,
  } },

  components: {testmenu:Menu, panel: Panel},

  methods: {
    makeWebsiteThumbnail() {
      axios.post("http://localhost:8081/api/thumbnail", {
        url: this.websiteUrl,
      })
      .then((response) => {
        this.thumbnailUrl = response.data.screenShotUrl;
      })
      .catch((error) => {
        window.alert(`The API returned an error: ${error}`);
      })
    },
    switchFeature(feature){
        this.component = feature;
    }
  }
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
.my-sidebar.b-sidebar-outer {
    position: absolute !important;
    height: 100% !important;
    width: 100% !important;
}

.my-sidebar .b-sidebar {
    position: absolute !important;
    height: 100% !important;
    width: 100% !important;
}
html,body{height:100%;}

</style>
