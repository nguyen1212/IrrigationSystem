<template>
  <mdb-card class="card-body" style="width: 100%; ">
    <mdb-card-title>
        <span style="float:left; color: white;">{{currentDateTime()}}</span>
        <span style="float:right; color: white;">
          <p> User Name
          <b-button style="float:right;" class="ml-3" size="sm" @click="logout"><b-icon icon="power" aria-hidden="true"></b-icon>Logout</b-button>
          </p>
        </span>
    </mdb-card-title>
    <div class="flex-row">
      <slides v-bind:ws='this.ws'/>
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
      ws:null
    }},
    created: function(){
      // var self = this;
      console.log("Starting connection to WebSocket Server")
      this.ws = new WebSocket('ws://' + window.location.host + '/devices/ws');
      // this.ws.addEventListener()
      this.ws.onopen = function(event) {
        console.log(event)
        console.log("Successfully connected to the echo websocket server...")
      }

    },
    methods: {
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
    send: function(){
        this.ws.send(
          JSON.stringify({
            userId: 'admin@gmail.com',
            plotId: 'california',
            state: 'On'
          })
        )
      }
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