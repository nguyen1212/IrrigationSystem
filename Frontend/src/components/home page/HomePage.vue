<template>
<div id="app" class="container-fluid d-flex flex-column vh-100 overflow-hidden">
  <!-- <div id="homepage" class="d-flex flex-column h-100 "> -->
    <!-- <div class="d-flex flex-column h-100 "> -->
    <div class="row">
      <panel @plotSelection="switchPlot" :PlotId="this.PlotId" :UserId="this.UserId"/>
    </div>
    <div class="row flex-grow-1 overflow-hidden" >
      <div class="col-3 mh-100">
        <mainMenu style="height:100%; width: 100%;" v-on:changefeature="switchFeature"></mainMenu>
      </div>
      <div class="col mh-100 overflow-auto">
        <keep-alive>
          <component v-on:changefeature="switchFeature" v-bind:is="component" :PlotId="this.PlotId" :UserId="this.UserId"/>
        </keep-alive>
      </div>
    </div>
    <!-- </div> -->
  <!-- </div> -->
</div>  
</template>

<script>
import Panel from '../panel/panel.vue';
import Menu from '../menu/menu.vue';
export default {
  name: 'HomePage',

  data() { return {
    component: '',
    UserId: '',
    PlotId: '',
    } 
  },

  created(){
    this.UserId = this.$route.params.UserId
  },

  components: {mainMenu:Menu, panel: Panel},

  methods: {
    switchFeature(feature){
      this.component = feature;
    },
    switchPlot(PlotId){
      this.PlotId = PlotId;
    }
  }
}
</script>

<style>
html, body{height:100%}

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

</style>
