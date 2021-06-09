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
                <b-form-select style="width: 60%" v-model="selectedPlot" :options="plots" @change="switchPlot(selectedPlot)"></b-form-select>
            </div>
        </template>
        <br><br><br><br><br><br><br><br><br><br><br><br><br><br><br><br>
      </b-carousel-slide>

      <!-- Slides with custom text -->
      <b-carousel-slide img-blank>
        <template>
          <div>
            <div><b-button variant="dark" size="lg" @click="send" ><b-icon icon="power" aria-hidden="true" scale="2" v-bind:variant="forcemode"></b-icon></b-button></div>
            <div><b-form-select style="width: 40%; margin-top: 4mm;" v-model="selectedMode" :options="modes"></b-form-select></div>
            <div class="info" style="display: flex; justify-content: center;">
              <p> {{soil}}%
              <br> Soil Moisture </p>
              <p > {{temp + '&deg;'}}C
              <br> Temperature </p>
              <p > {{humid}}%
              <br> Humidity </p>
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
      soil: '',
      temp: '',
      humid: '',
      
    },
    name: 'slides',
    data() {
      return {
        slide: 0,
        sliding: null,
        forceData: '',
        forcemode: '',
        selectedPlot: '',
        selectedMode: '',
        plots: [
          { value: '', text: 'Please select an option' },
          { value: 'New York', text: 'New York' },
          { value: 'California', text: 'California' },
          { value: '', text: 'Beijing' },
          { value: '', text: 'Add...'}
        ],
        modes: [
          {value: 'none', text: 'Select Mode' },
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
      toggle(mode){
        if (mode == 'success')
          return 'danger'
        else
          return 'success'
      },
      send(){
        var self = this
        self.forcemode = this.toggle(self.forcemode)
        if (self.forcemode == 'success')
          self.forceData = '0'
        else
          self.forceData = '1'
        this.ws.send(
          JSON.stringify({
            userId: 'admin@gmail.com',
            plotId: 'newyork',
            name: 'RELAY',
            data: self.forceData
          })
        )
      },
      switchPlot(selectedPlot){
        this.$emit('plotSelection', selectedPlot)
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