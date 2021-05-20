<template>
<div   class="d-flex justify-content-center">
    <!-- <Card :info="JSON.stringify(obj[0])" width="500" height="300" /> -->
  <AddModal myid="mymodal"> </AddModal>
<div
style="width: 800px; height: 500px">
    <b-carousel
      id="carousel-1"
      ref="mycarousel"
      v-model="slide"
      :interval="0"
      controls
      background="transparent"
      style="width: 800px; height: 400px"
    >
      <!-- Slide with blank fluid image to maintain slide aspect ratio -->

      <b-carousel-slide img-blank v-for="o in obj" :key="o.id" style="width: 800px; height: 400px">
        <Card :info="JSON.stringify(o)" width="500" height="300" />
      </b-carousel-slide>

      <b-carousel-slide img-blank style="width: 800px; height: 400px">
        <div class="d-flex justify-content-center">
        <div 
        class="d-flex justify-content-center" 
        style="width: 500px; height: 300px; background-color: transparent">
            <b-card 
            class="mb-2" 
            style="width: 250px; height: 300px;  color: black"
            title="Add Preset"
            @click="$bvModal.show('mymodal')"
            >
            <b-icon class="mt-5" icon="plus-circle-fill" style="color: black; width: 120px; height: 120px;" variant="black"></b-icon>
            <b-card-text>

            <b-link @click="$bvModal.show('mymodal')"  href="#" style="height: 100%; width: 100%"></b-link>
            </b-card-text>
            
            <!-- <b-button >Open</b-button> -->
            </b-card>
        </div>
        </div>
      </b-carousel-slide>
    </b-carousel>
  </div>
</div>
</template>

<script>
import Card from "../auto mode/Card"
import AddModal from "../auto mode/AddModal"
export default {
  name: 'ListView',
  data() {
      var obj = JSON.parse(this.info)
    
    return {
      obj,
      i: 0,
      currentOpen: null,
      slide: 0,
      sliding: null,
    }
  },
  props: {
    info: String,
    selected: Number
  },
  components:{
      Card,
      AddModal
  },
  methods: {
      onSlideStart(slide) {
        this.sliding = true
      },
      onSlideEnd(slide) {
        this.sliding = false
      },
      prev() {
        this.$refs.mycarousel.prev()
      },
      next() {
        this.$refs.mycarousel.next()
      }
  }
}
</script>

<style>
.carousel-control-next,
.carousel-control-prev /*, .carousel-indicators */ {
    filter: invert(100%);
}
</style>