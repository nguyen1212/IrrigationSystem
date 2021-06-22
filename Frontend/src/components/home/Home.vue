<template>
  <div id="home">
    <h1 id="headline"><b>{{webpage}}</b></h1>
    <div class="d-flex justify-content-center">
      <div id="intro-1">      
      <p class="h5">
      <b>Smart Irrigation System <br>
       intergrates automatic watering and advance tracking,
       highly customizable design and IoT Soil Monitor
      </b></p>
       </div>
    </div>
    <!-- <p> 
      <router-link to='/login'> goto login</router-link>
      <router-link to='/home'> goto home</router-link>
    </p> -->
    <div  id="intro-2" class="mt-5">
    <b-container class="w-100">
      <b-row style="max-height: 300px;" class="w-100">
        <b-col cols=6 > <img class="img-fluid" src="~@/assets/watering.jpg" style="height: 300px"> </b-col>
        <b-col cols=6 >
          <h4 class="text-left mb-4"><b>Our Mission</b></h4>
          <p class="text-left" style="line-height: 2.0">
            If you are someone looking for a better solution to water your plants when you are away or just
            looking for a better way to manage this watering task, {{webpage}} is the solution for you.
            At {{webpage}}, we thrive to deliver a better, smarter solution than primitive watering system
            to your problems. Now you can rest assured knowing that your plants are always watered on time
            and with just the right amount of water.
          </p>
        </b-col>
      </b-row>
    </b-container>
    </div>
    <div class="mt-5">
    <h2><b>Explore</b></h2>
    <div class="mt-3">

    
    <b-container>
      <b-row class="mt-3">
        <b-col cols=6>

          <div class="card my-card">
            <div class="card-header" style="background-color: white; border: none;">
              <p class="h1 text-left"><span class="auto-icon"></span></p>
            </div>
            <div class="card-body pt-1">
              <h5 class="card-title text-left"><b>Auto Mode</b></h5>
              <p class="card-text text-left">
                Our smart irrigation system comes with soil moisture sensor that measure and find out 
                the neccessary amount of water for your plants
                each time. Users only need to setup preset for their plant once, sit back and enjoy.
              </p>
              <div class="d-flex justify-content-left"><b-link class="ml-0" @click="selectFeature(autoMode)">Click to explore</b-link></div>
            </div>
          </div>

        </b-col>
        <b-col cols=6>

          <div class="card my-card">
            <div class="card-header" style="background-color: white; border: none;">
              <p class="h1 text-left"><b-icon icon="clock-history"></b-icon></p>
            </div>
            <div class="card-body pt-1">
              <h5 class="card-title text-left"><b>Timer Mode</b></h5>
              <p class="card-text text-left">
                Are you constantly checking your time to water the plants? No more worries! With 
                our smart timer mode, we make sure that your plants always get the right amount of
                water it needs every single time. Set your time once and you are ready to go.
              </p>
              <div class="d-flex justify-content-left"><b-link class="ml-0" @click="selectFeature(timerMode)">Click to explore</b-link></div>
            </div>
          </div>

        </b-col>
      </b-row>
      <b-row class="mt-3">
        <b-col cols=6>

          <div class="card my-card">
            <div class="card-header" style="background-color: white; border: none;">
              <p class="h1 text-left"><b-icon icon="graph-up" aria-hidden="true"></b-icon></p>
            </div>
            <div class="card-body pt-1">
              <h5 class="card-title text-left"><b>Smart monitoring</b></h5>
              <p class="card-text text-left">
                We carefully measure and record our watering amount so you never have to worry about over-watering your plant.
                With a tap of button, you can view water log history of days, months of usage for your plants.
              </p>
              <div class="d-flex justify-content-left"><b-link class="ml-0"  @click="selectFeature(waterLog)">Click to explore</b-link></div>
            </div>
          </div>
        </b-col>
        <b-col cols=6>
          <div class="card my-card">
            <div class="card-header" style="background-color: white; border: none;">
              <p class="h1 text-left"><b-icon icon="cpu" aria-hidden="true"></b-icon></p>
            </div>
            <div class="card-body pt-1">
              <h5 class="card-title text-left"><b>Device management</b></h5>
              <p class="card-text text-left">
                Advance devices monitoring for passionate users.
                Deep dive to your plants' info with soil moisture, temperature, even humidity with our
                various sensors' extensive record.
              </p>
              <div class="d-flex justify-content-left"><b-link class="ml-0" @click="selectFeature(devices)">Click to explore</b-link></div>
            </div>
          </div>
        </b-col>
      </b-row>
    </b-container>
    </div>
    </div>
  </div>
</template>


<script>
import autoMode from '../auto mode/AutoMode.vue';
import waterLog from '../water log/WaterLog.vue';
import timerMode from '../timer mode/TimerMode.vue';
import devices from '../device/Devices.vue';
export default {
  name: 'home',
  data () {
    return {
      webpage: 'ĐẶT CÁI TÊN TRANG WEB DÙM ĐI',
      autoMode: autoMode,
      waterLog: waterLog,
      timerMode: timerMode,
      devices: devices
  }},
  methods: {
    selectFeature(feature){
        this.$emit('changefeature', feature)
    }
  }
}
</script>
<style scoped>
#home {
  width: 100%;
  height: 100%;
}
#headline {
  margin: 30px;
  padding-left: 150px;
  text-align: left;
}
#intro-1 {
  max-width: 500px;
  width: 500px;
  text-align: left;
  margin: 0px;
}
#intro-2 {
  max-height: 300px;
  height: 300px;
  margin: 0px;
}
.my-card {
  width: 400px;
  border: none;
}
.auto-icon {
    background: url('~@/assets/icon/automode/letter-a.svg');
    height: 40px;
    width: 40px;
    display: flex;
    background-color: black;
    -webkit-filter: invert(100%);
    filter: invert(100%);
    /* Other styles here */
}
</style>