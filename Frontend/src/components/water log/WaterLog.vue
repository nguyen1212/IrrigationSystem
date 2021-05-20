<template>
        <div class="container">
            <!-- DISPLAY DAY -->
            <div class="row justify-content-center">
            <b-carousel
            id="carousel-1"
            ref="mycarousel"
            :interval="0"
            controls
            v-model="slide"
            img-width="300px"
            img-height="25px"
            background="transparent"
            >
            <b-carousel-slide img-blank>
                <p>APRIL 12-18</p>
            </b-carousel-slide>
            </b-carousel>
            </div>
            <!-- DISPLAY GRAPH -->
            <div class="row graph-content">
                <b-img class="w-100 h-100" blank-color="gray"></b-img>
            </div>
            <!-- DISPLAY INFO -->
            <div class="row info-content">
                <div class="col-2">
                <div class="row justify-content-center"><b>Total</b></div>
                <div class="row justify-content-center">
                    <b-icon icon="droplet-fill"></b-icon>
                    13.3L
                </div>
                </div>
            </div>
            <div class="row info-content justify-content-between">
                <div class="col-2 ">
                <div class="row justify-content-center"><b>Average</b></div>
                <div class="row justify-content-center">
                    <b-icon icon="droplet-fill"></b-icon>
                    1.9L/Day
                </div>
                </div>
                <div class="col-2">
                    <b-button @click="handleCompare">Compare</b-button>
                </div>
            </div>
        </div>
</template>

<script>
export default {
    data() {
        return {
            slide: 0,
            sliding: null,
        }
    },
    name: 'Water-Log',
    methods: {
        handleCompare()
        {},
              onSlideStart(slide) {
            this.sliding = true
        },
        onSlideEnd(slide) {
            this.sliding = false
        },
    }
}
</script>
<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.carousel-caption {
    position: static;
    right: 0%;
    bottom: 0px;
    left: 0%;
    z-index: 10;
    padding-top: 0px;
    padding-bottom: 0px;
    color: #000;
    text-align: center;
    font-size: 150%;
}
.graph-content {
    height: 450px;
    margin-top: 20px;
    margin-bottom: 20px
}
.info-content {
    margin-top: 10px;
    margin-bottom: 10px
}
.carousel-control-next,
.carousel-control-prev /*, .carousel-indicators */ {
    filter: invert(100%);
}
</style>