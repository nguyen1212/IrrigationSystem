<template>
<div class="d-flex justify-content-center">

<div class="d-flex card-holder" :style="{ width: cardholderWidth + 'px', height: cardHeight + 'px' }">
    <b-card 
    class="mb-2 front-card-class" 
    :id="'front-card-' + myinfo.id" 
    :style="{ width: cardWidth + 'px', height: cardHeight + 'px',  left: halfWidth + 'px'  }"
    :title="myinfo.name" 
    :img-src="myinfo.img" 
    img-alt="Image" 
    img-height="80%"
    img-top 
    @click="toggleSideView" 
    >
    <b-link @click="toggleSideView"  href="#" style="height: 100%; width: 100%"></b-link>
    <!-- <b-button >Open</b-button> -->
    </b-card>
    <b-card
    :id="'back-card-' + myinfo.id" 
    class="mb-2 back-card-class" 
    :style="{ width: cardWidth + 'px', height: cardHeight + 'px' }"
    >
      <b-card-text class="w-100">
        <div>
              <ul>
              <li ><b>ID:</b> {{myinfo.id}}</li>
              <li ><b>Name:</b> {{myinfo.name}}</li>
              <li ><b>Humidity:</b> {{myinfo.humidityThreshold}}</li>
              <li ><b>Moisture:</b> {{myinfo.moistureThreshold}}</li>
              <li ><b>Temperature:</b> {{myinfo.temperatureThreshold}}</li>
              <li ><b>Note:</b> {{myinfo.note}}</li>
            </ul>
            <b-button variant="dark">Edit</b-button>

        </div>
      </b-card-text>
      
    </b-card>
</div>
</div>

</template>

<script>
export default {
  name: 'Card',
  data() {
    var myinfo = JSON.parse(this.info)
    return {
      myinfo,
      view: false,
      cardholderWidth: this.width,
      cardWidth: (parseInt(this.width)/2).toString(),
      halfWidth: (parseInt(this.width)/4).toString(),
      cardHeight: this.height
    }
  },
  props: {
    info: String,
    width: String,
    height: String
  },

  methods: {
    toggleSideView()
    {
        this.view = !this.view;
        let frontId = "front-card-" + this.myinfo.id
        let backId = "back-card-" + this.myinfo.id
        console.log(frontId)
        document.getElementById(frontId).style.left = ((this.view)? "0px": this.halfWidth + 'px')
        document.getElementById(backId).style.opacity = ((this.view)? "1": "0")
    }
  }
}
</script>

<style scoped>
.card-holder {
  background-color: transparent;
}
.front-card-class {
  transition: left 0.1s;
  background-color: white;
  color: black;
  text-shadow: none;
  z-index: 10;
}
.back-card-class {
  opacity: 0;
  transition: opacity 0.1s;
  background-color: white;
  color: black;
  text-shadow: none;
}
li {
  text-align: left;
}
</style>