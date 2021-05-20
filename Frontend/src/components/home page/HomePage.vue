<template>
<div id="app" class="container-fluid h-100">
  <div id="homepage" class="d-flex flex-column h-100 ">
    <!-- <div class="d-flex flex-column h-100 "> -->
    <div class="row ">
      <panel/>
    </div>
    <div class="row flex-grow-1" >
      <div class="col-4">
        <testmenu style="height:100%; width: 100%;" v-on:changefeature="switchFeature"></testmenu>
      </div>
      <div class="col">
        <component style="height:100%; width: 100%;" v-bind:is="component" />
      </div>
    </div>
    <!-- </div> -->
  </div>
</div>  
</template>

<script>

import axios from 'axios';
import Panel from '../panel/panel.vue';
import Menu from '../menu/menu.vue';
export default {
  name: 'HomePage',

  data() { return {
    websiteUrl: '',
    thumbnailUrl: '',
    component: '',
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
