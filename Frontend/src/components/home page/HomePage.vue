<template>
<div id="app" class="container-fluid h-100">
  <div id="homepage" class="d-flex flex-column h-100 ">
    <!-- <div class="d-flex flex-column h-100 "> -->
    <div class="row ">
      <panel @plotSelection="switchPlot" :PlotId="this.PlotId" :UserId="this.UserId"/>
    </div>
    <div class="row flex-grow-1" >
      <div class="col-3">
        <testmenu style="height:100%; width: 100%;" v-on:changefeature="switchFeature"></testmenu>
      </div>
      <div class="col">
        <keep-alive>
          <component style="height:100%; width: 100%;" v-bind:is="component" :PlotId="this.PlotId" :UserId="this.UserId"/>
        </keep-alive>
      </div>
    </div>
    <!-- </div> -->
  </div>
</div>  
</template>

<script>
import Panel from '../panel/panel.vue';
import Menu from '../menu/menu.vue';
export default {
  name: 'HomePage',

  data() { return {
    component: '',
    UserId: 'pod-yone',
    PlotId: '',
  } },

  components: {testmenu:Menu, panel: Panel},

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
