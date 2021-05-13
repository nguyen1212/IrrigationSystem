<template>
  <mdb-card class="card-body" style="width: 100%; ">
    <mdb-card-title>
        <span style="float:left; color: white;">{{currentDateTime()}}</span>
        <span style="float:right; color: white;">Weather Forcast</span>
    </mdb-card-title>
    <div class="flex-row">
      <slides/>
    </div>
  </mdb-card>
</template>
<script>
  import { mdbCard, mdbCardTitle, mdbCardText } from 'mdbvue';
  import slides from '../panel/slides.vue';
  export default {
    name: 'Panel',
    components: {
      mdbCard,
      mdbCardTitle,
      mdbCardText,
      slides
    },
    methods: {
    currentDateTime() {
      const current = new Date();
      const date = current.getFullYear()+'-'+(current.getMonth()+1)+'-'+current.getDate();
      const time = current.getHours() + ":" + current.getMinutes() + ":" + current.getSeconds();
      const dateTime = 'Date: ' + date +'\n'+ 'Time: ' + time;
      
      return 'Date: ' + date;
    }
  }
  }
</script>
<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
  .card-body{
    max-height: 11cm!important;
    background-image: url("https://i.picsum.photos/id/159/1024/480.jpg?hmac=6PcKV3l3tqzR177maUXur9a_NlTfrPAyf5PmpvZ5Q6M");
    background-size: cover;
  }
</style>