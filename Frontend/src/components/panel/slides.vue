<template>
  <div>
    <b-carousel
      id="carousel-1"
      v-model="slide"
      :interval="0"
      controls
      indicators
      style="text-shadow: 1px 1px 2px #333; height: 6cm; width: 15cm; margin: auto; backdrop-filter: blur(20px); border-radius: 20px;"
      @sliding-start="onSlideStart"
      @sliding-end="onSlideEnd"
    >
      <!-- Text slides with image -->
      <b-carousel-slide caption="Select Plot" img-blank>
        <template>
            <div>
                <b-form-select style="width: 60%" v-model="selectedPlot" :options="plots"></b-form-select>
            </div>
        </template>
        <br><br><br><br><br><br><br><br><br><br><br><br><br><br><br><br>
      </b-carousel-slide>

      <!-- Slides with custom text -->
      <b-carousel-slide img-blank>
        <template>
          <div>
            <div><b-button variant="dark" size="lg" @click="send" ><b-icon icon="power" aria-hidden="true" scale="2" variant="success"></b-icon>{{buttonVariant}}</b-button></div>
            <div><b-form-select style="width: 40%; margin-top: 4mm;" v-model="selectedMode" :options="modes"></b-form-select></div>
            <div class="info" style="display: flex; justify-content: center;">
              <p> {{soiMoisture}}%
              <br> Soil Moisture </p>
              <!-- <p > {{temp}}C
              <br> Temperature </p>
              <p > {{humid}}%
              <br> Humidity </p> -->
            </div>
          </div>
        </template>
          <br><br><br><br><br><br><br><br><br><br><br><br><br><br>
      </b-carousel-slide>

    </b-carousel>
  </div>
</template>

<script>
  export default {
    props: {
      ws: WebSocket,
      buttonVariant: String,
      soiMoisture: '',
      temp: '',
      humid: '',
    },
    name: 'slides',
    data() {
      return {
        slide: 0,
        sliding: null,
        selectedPlot: null,
        selectedMode: null,
        plots: [
          { value: null, text: 'Please select an option' },
          { value: '1', text: 'NewYork' },
          { value: '2', text: 'London' },
          { value: '3', text: 'Beijing' },
          { value: '4', text: 'Add...'}
        ],
        modes: [
          {value: null, text: 'Select Mode' },
          {value: 'auto', text: 'Auto Mode'},
          {value: 'manual', text: 'Manual Mode'}
        ]
      }
    },
    methods: {
      onSlideStart() {
        this.sliding = true
      },
      onSlideEnd() {
        this.sliding = false
      },
      send: function(){
        this.ws.send(
          JSON.stringify({
            userId: 'admin@gmail.com',
            plotId: 'california',
            name: 'RELAY',
            data: '0'
          })
        )
      }
    }
  }
</script>
<style scoped>
.info p{
  margin-bottom: -2mm; 
  margin: auto
}
.info {
  margin-top: 0.5cm;
}
</style>