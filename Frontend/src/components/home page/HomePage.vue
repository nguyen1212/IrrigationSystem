<template>
<div id="app" class="container-fluid d-flex flex-column vh-100 overflow-hidden">
  <!-- <div id="homepage" class="d-flex flex-column h-100 "> -->
    <!-- <div class="d-flex flex-column h-100 "> -->
    <div class="row">
      <panel @plotSelection="switchPlot" :PlotId="this.PlotId" :UserId="this.UserId"/>
    </div>
    <div class="row flex-grow-1 overflow-hidden" >
      <div class="col-3 mh-100">
        <mainMenu style="height:100%;" v-on:changefeature="switchFeature" v-on:closemenu="toggleSidebar"></mainMenu>
      </div>
      <div class="col mh-100 overflow-auto">
        <b-button id= "menu-btn" style="left: 0; display: none; position: fixed" size="sm" @click="toggleSidebar"><b-icon icon="list" aria-hidden="true"></b-icon></b-button>
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
    sideBarVision: 1,
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
    },
    toggleSidebar(){
      var sidebar = document.getElementsByClassName("my-sidebar");
      var sidebarheader = document.getElementsByClassName("b-sidebar-header");
      var sidebarbody = document.getElementsByClassName("b-sidebar-body");
      var menuWrapper = document.getElementsByClassName("col-3 mh-100");
      var menuBtn = document.getElementById('menu-btn')
      if (this.sideBarVision == 1){
        sidebar[0].style.width = '0px'
        menuWrapper[0].style.maxWidth = '0px'
        menuBtn.style.display = ''
        // sidebarheader[0].style.display = 'none'
        sidebarheader[0].style.opacity = sidebarbody[0].style.opacity = '0'
        sidebarheader[0].style.visibility = sidebarbody[0].style.visibility ='hidden'
        sidebarheader[0].style.transition = sidebarbody[0].style.transition ='visibility 0s, opacity 0s linear'
      }
      else{
        menuBtn.style.display = 'none'
        menuWrapper[0].style.maxWidth = '20%'
        sidebar[0].style.width = '100%'
        // setTimeout(function() {sidebarheader[0].style.display = ''}, 500)
        sidebarheader[0].style.opacity = sidebarbody[0].style.opacity = '1'
        sidebarheader[0].style.visibility = sidebarbody[0].style.visibility ='visible'
        sidebarheader[0].style.transition = sidebarbody[0].style.transition ='visibility 0s, opacity 1.5s linear 0.5s'
      }
      this.sideBarVision = -1 * this.sideBarVision
    }
  }
}
</script>

<style>
html, body{height:100%}

/* .my-sidebar.b-sidebar-outer {
    position: absolute !important;
    height: 100% !important;
    width: 100%;
    transition: 0.5s;
} */

.my-sidebar .b-sidebar, .my-sidebar.b-sidebar-outer{
    position: absolute !important;
    height: 100% !important;
    width: 100%;
    transition: 1s;
}

.b-sidebar-header, .b-sidebar-body{
  transition: visibility 0s, opacity 1.5s linear 0.5s;
}

.col-3{
  transition: 1s;
  padding: 0;
  max-width: 20%;
}

</style>
