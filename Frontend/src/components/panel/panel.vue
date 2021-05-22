<template>
  <mdb-card class="card-body" style="width: 100%; ">
    <mdb-card-title>
        <span style="float:left; color: white;">{{currentDateTime()}}</span>
        <span style="float:right; color: white;">
          <p> User Name
          <b-button style="float:right;" class="ml-3" size="sm" @click="logout" ><b-icon icon="power" aria-hidden="true"></b-icon></b-button>
          </p>
          <p v-html='content'></p>
        </span>
    </mdb-card-title>
    <div class="flex-row">
      <slides v-bind:ws='this.ws' :buttonVariant='this.buttonVariant'/>
    </div>
  </mdb-card>
</template>
<script>
  import { mdbCard, mdbCardTitle } from 'mdbvue';
  import slides from '../panel/slides.vue';
  import firebase from 'firebase'
  export default {
    name: 'Panel',
    components: {
      mdbCard,
      mdbCardTitle,
      slides
    },
    data(){ return{
      ws:null,
      msg: Object,
      buttonVariant: '',
      soiMoisture: '',
      temp: '',
      humid: '',
    }},
    created: function(){
      var self = this;
      console.log("Starting connection to WebSocket Server")
      this.ws = new WebSocket('ws://' + "127.0.0.1:8080" + '/devices/ws');
      // this.ws.addEventListener()
      this.ws.onopen = function(event) {
        console.log(event)
        console.log("Successfully connected to the echo websocket server...")
      }
      this.ws.addEventListener('message', function(e){
        var msg = JSON.parse(e.data);
        if (msg.name == 'SOIL')
          this.soiMoisture = msg.data
      })
    },
    mounted: function () {
      window.setInterval(() => {
        this.changeValue()
      }, 500)
    },
    methods: {
    changeValue(){
      this.buttonVariant = 'success'
    },
    currentDateTime() {
      const current = new Date();
      const date = current.getFullYear()+'-'+(current.getMonth()+1)+'-'+current.getDate();
      // const time = current.getHours() + ":" + current.getMinutes() + ":" + current.getSeconds();
      // const dateTime = 'Date: ' + date +'\n'+ 'Time: ' + time;
      return 'Date: ' + date;
    },
    logout() {
      firebase
          .auth()
          .signOut()
          .then(() => {
              this.$router.push('/');
          })
          .catch(error => {
              alert(error.message);
              this.$router.push('/');
          });
    },
  }
  }
</script>
<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
  .card-body{
    max-height: 11cm!important;
    background-image: url('~@/assets/greenery-main-page.jpg');
    background-size: cover;
    background-position: center;
  }
</style>