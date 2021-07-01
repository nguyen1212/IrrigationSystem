<template>
  <div>
  <mdb-card class="card-body">
    <b-sidebar class="my-sidebar" id="sidebar-1" title="Menu" shadow visible=True bg-variant="dark" text-variant="light" >
      <template #header="{}">
       <div style="width: 100%;">
        <strong style="float:left;">Menu</strong>
        <b-button style="float:right;" size="sm" @click="hide"><b-icon icon="list" aria-hidden="true"></b-icon></b-button>
       </div>
      </template>
      <template>
          <nav class="mr-auto">
            <b-button-group vertical style="width: 100%;">
              <b-button variant="dark" class="menu-button" v-on:click="selectFeature(home)"> 
                <b-container> 
                  <b-row>
                  <b-col cols="2"/>
                  <b-col cols="1"> <b-icon icon="house-door-fill" aria-hidden="true"></b-icon></b-col>
                  <b-col class="text-left"> <b>Home</b> </b-col>
                  </b-row>
                </b-container>
              </b-button>
              <b-button variant="dark" class="menu-button" v-on:click="selectFeature(autoMode)">
                <b-container> 
                  <b-row>
                  <b-col cols="2"/>
                  <b-col cols="1"> 
                  <span class="auto-icon"></span>
                  </b-col>
                  <b-col class="text-left"> <b>Auto Mode</b></b-col>
                  </b-row>
                </b-container>
              </b-button>
              <b-button variant="dark" class="menu-button" v-on:click="selectFeature(timerMode)">
                <b-container> 
                  <b-row>
                  <b-col cols="2"/>
                  <b-col cols="1"> <b-icon icon="clock-history" aria-hidden="true"></b-icon></b-col>
                  <b-col class="text-left"><b>Timer Mode</b></b-col>
                  </b-row>
                </b-container>
              </b-button>
              <b-button variant="dark" class="menu-button" v-on:click="selectFeature(waterLog)">
                <b-container> 
                  <b-row>
                  <b-col cols="2"/>
                  <b-col cols="1"> <b-icon icon="graph-up" aria-hidden="true"></b-icon></b-col>
                  <b-col class="text-left"><b>Water Log</b></b-col>
                  </b-row>
                </b-container>
              </b-button>
              <b-button variant="dark" class="menu-button" v-on:click="selectFeature(devices)">
                <b-container> 
                  <b-row>
                  <b-col cols="2"/>
                  <b-col cols="1"><b-icon icon="cpu" aria-hidden="true"></b-icon></b-col>
                  <b-col class="text-left"><b>Devices</b></b-col>
                  </b-row>
                </b-container>
              </b-button>
            </b-button-group>
          </nav>
      </template>
    </b-sidebar>
  </mdb-card>
  </div>
</template>
<script>
import {mdbCard} from 'mdbvue';
import home from '../home/Home.vue';
import autoMode from '../auto mode/AutoMode.vue';
import waterLog from '../water log/WaterLog.vue';
import timerMode from '../timer mode/TimerMode.vue';
import devices from '../device/Devices.vue';
export default{
  name: 'Menu',
  data(){return{
    home: home,
    autoMode: autoMode,
    waterLog: waterLog,
    timerMode: timerMode,
    devices: devices
  }},
  components: mdbCard,
  methods: {
    selectFeature(feature){
        this.$emit('changefeature', feature)
    },
    hide(){
      this.$emit('closemenu')
      // var btnGroup = document.getElementsByClassName("my-sidebar");
    }
  }
}
</script>
<style scoped>
.ml-0{
  width: 100%;
}
.menu-button{
  padding: 10px;
}
.auto-icon {
    background: url('~@/assets/icon/automode/letter-a.svg');
    height: 21px;
    width: 21px;
    display: flex;
    /* Other styles here */
}
</style>