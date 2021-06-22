<template>
  <mdb-card class="card-body" style="width: 100%;">
    <mdb-card-title>
        <span style="float:left; color: white;">{{this.date}}</span>
        <span style="float:right; color: white;">
          <p> {{UserId}}
          <b-button style="float:right;" class="ml-3" size="sm" @click="logout" ><b-icon icon="power" aria-hidden="true"></b-icon>Log out</b-button>
          </p>
          <!-- <p v-html='content'></p> -->
        </span>
        <br/>
        <span style="float:left; color: white;">{{this.time}}</span>
    </mdb-card-title>
    <div class="flex-row">
      <slides v-on="$listeners" v-bind:ws='this.ws' :soil="this.soil" :temp="this.temp" :humid="this.humid"/>
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
    props:{
      PlotId: String,
      UserId: String,
    },
    data(){ return{
      ws:null,
      msg: Object,
      soil: '',
      temp: '',
      humid: '',
      dateTimeFunc: null,
      date: '',
      time: '',
    }},
    created: function(){
      var self = this
      var tmp = ''
      this.connectWebsocket()
      this.dateTimeFunc = setInterval(()=>{
        tmp = self.currentDateTime()
        this.date = tmp.split("|")[0]
        this.time = tmp.split("|")[1]
      }, 1000)
    },
    methods: {
      classifyMsg(msg){
        if (msg.name == 'SOIL') this.soil = msg.data
        if (msg.name == 'TEMP-HUMID'){
          var data = msg.data.split("-")
          this.temp = data[0]
          this.humid = data[1]
        }
    },
      currentDateTime() {
        const current = new Date();
        const date = current.getDate() + '-' + (current.getMonth() + 1) + '-' + current.getFullYear();
        const time = current.getHours() + ":" + ((current.getMinutes() < 10) ? '0':'')+ current.getMinutes() + ":" + ((current.getSeconds() < 10) ? '0':'') + current.getSeconds();
        const dateTime = 'Date: ' + date +'|'+ 'Time: ' + time;
        return dateTime;
      },
      async logout() {
        await this.closeWebsocket();
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
      connectWebsocket(){
        var self = this;
        console.log("Starting connection to WebSocket Server")
        var url = 'ws://' + "127.0.0.1:8080" + '/devices/ws'
        this.ws= new WebSocket(url);
        this.ws.onopen = function(event) {
          console.log(event)
          console.log("Successfully connected to the echo websocket server...")
        }
        this.ws.addEventListener('message', function(e){
          var msg = JSON.parse(e.data);
          self.msg = msg
          self.classifyMsg(msg)
        })
        this.ws.addEventListener('close', function(){
          self.ws.close()}
        )
      },  
      async closeWebsocket(){
        try{
          await this.ws.close();
          console.log("Successfully closed websocket...")
        }
        catch(error){
          console.log(error.message)
        }
      }
    },
    beforeDestroy () {
      clearInterval(this.dateTimeFunc)
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