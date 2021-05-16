<template>
<div  class="d-flex justify-content-center">
  <AddModal myid="mymodal"> </AddModal>
   <b-container>
      <b-row class=" pt-3 pb-3 border-bottom" >
        <b-col class="col-1"><b>#</b></b-col>
        <b-col class="leftalign col-3"><b>Name</b></b-col>
        <b-col class="leftalign col-5 description"><b>Description</b></b-col>
        <b-col class="col-3"><b>Status</b></b-col>
      </b-row>
      <!-- magic -->
      <div :set="count = 0"></div>
      <b-row class="border-bottom" :key="o.id" v-for="o in obj" :set="count=count+1">
        <b-link class='pt-3 pb-3 row w-100 ml-0' v-b-toggle="'collapse-' + count" href="#">
            <b-col class="col-1">
              <b-avatar size="2em" icon="flower2" variant="info" :src="o.img"></b-avatar>
            </b-col>
            <b-col class="leftalign col-3">{{o.name}}</b-col>
            <b-col class="leftalign col-5 description">{{o.note}}</b-col>
            <b-col class="pr-2 col-3">
              <div v-if="selected === o.id">
                <b-button variant="success">Select</b-button>
              </div>
              <div v-else>
                <b-button variant="outline-primary">Select</b-button>
              </div>
            </b-col>
        </b-link>
        <b-collapse class='w-100' :id="'collapse-' + count">
          <b-card class="border-0 ml-0">
            <div class="row">
              <div class="col-8">
              <ul>
              <li class="leftalign"><b>ID:</b> {{o.id}}</li>
              <li class="leftalign"><b>Name:</b> {{o.name}}</li>
              <li class="leftalign"><b>Humid Level:</b> {{o.humidityThreshold}}</li>
              <li class="leftalign"><b>Moisture Level:</b> {{o.moistureThreshold}}</li>
              <li class="leftalign"><b>Temperature Level:</b> {{o.temperatureThreshold}}</li>
              <li class="leftalign"><b>Note:</b> {{o.note}}</li>
            </ul>
              </div>
              <div class="col-4">
                <div v-if="o.img !== ''">
                  <b-img :src="o.img" rounded="circle" alt="Circle image">
                </b-img></div>
                
              </div>
          </div>
          <div class="row">
            <b-button class="col-1" variant="dark">Edit</b-button>
          </div>
          </b-card>
        </b-collapse>
      </b-row>
      
      <b-row class="border-bottom">
        <b-link class='pt-3 pb-3 row w-100 ml-0' @click="$bvModal.show('mymodal')">
            <b-col class="col-1">
            </b-col>
            <b-col class="leftalign col-3">#</b-col>
            <b-col class="leftalign col-5 description">#</b-col>
            <b-col class="pr-2 col-3">
              
            </b-col>
        </b-link>
      </b-row>
  </b-container>

</div>
</template>

<script>
import AddModal from "../components/auto mode/AddModal"
export default {
  name: 'ListView',
  components: {
    AddModal
  },
  data() {
      var obj = JSON.parse(this.info)
    return {
      obj,
      i: 0,
      currentOpen: null
    }
  },
  props: {
    info: String,
    selected: Number
  },

  methods: {
      onClick(id){
        var s = id.toString();
        if (this.currentOpen === null || this.currentOpen !== s)
        {
          this.currentOpen = s
        }
        else
        {
          this.currentOpen = null
        }
      },
      showModal() {
        this.$refs.mymodal.show()
      },
  }
}
</script>

<style scoped>
a {
  color: black;
}
a:hover {
  color: black;
  text-decoration: none;
  background-color: antiquewhite;
}
.description {
  overflow: hidden;
  white-space: nowrap;
  text-overflow: ellipsis;
}
.leftalign {
  text-align: left;
}
img {
  max-width: 200px;
  max-height: 200px;
}
</style>