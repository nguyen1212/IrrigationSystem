<template>
  <div>
    <b-carousel
      id="carousel-1"
      v-model="slide"
      :interval="0"
      controls
      indicators
      style="text-shadow: 1px 1px 2px #333; height: 5cm; width: 8cm; margin: auto; backdrop-filter: blur(20px); border-radius: 20px;"
      @sliding-start="onSlideStart"
      @sliding-end="onSlideEnd"
    >
      <!-- Text slides with image -->
      <b-carousel-slide caption="Select Plot" img-blank>
        <template>
            <div>
                <b-form-select v-model="selected" :options="options"></b-form-select>
            </div>
        </template>
        <br><br><br><br><br><br>
      </b-carousel-slide>

      <!-- Slides with custom text -->
      <b-carousel-slide caption="Force Mode" img-blank>
          <br><br><br><br><br><br><br>
      </b-carousel-slide>

    </b-carousel>
  </div>
</template>

<script>
  export default {
    name: 'slides',
    data() {
      return {
        slide: 0,
        sliding: null,
        selected: null,
        options: [
          { value: null, text: 'Please select an option' },
          { value: '1', text: 'NewYork' },
          { value: '2', text: 'London' },
          { value: '3', text: 'Beijing' },
          { value: '4', text: 'Add...'}
        ]
      }
    },
    methods: {
      onSlideStart(slide) {
        this.sliding = true
      },
      onSlideEnd(slide) {
        this.sliding = false
      }
    }
  }
</script>
<style scoped>
</style>