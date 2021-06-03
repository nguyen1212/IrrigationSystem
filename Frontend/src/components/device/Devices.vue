<template>
  <div class="devices">
    <div class="row">
      <div class="col-8">
        <p> Device Location </p>
      </div>
      <div class="col">
        <p> Statistic </p>
        <b-table striped hover :items="items" :fields="fields"></b-table>
        <hr/>
        <p> Device Type </p>
      </div>
    </div>
  </div>
</template>
<script>
import axios from 'axios'
export default {
  name: 'devices',
  props:{
    UserId: '',
    PlotId: '',
  },
  data() {
      return {
        fields: ['Name', 'Type', 'Position'],
        items: []
      }
  },
  created: this.getData(),
  methods: {
    getData(){
      var self = this
      axios.post("http://localhost:8081/devices/data",{
        UserId: self.UserId,
        PlotId: self.PlotId
      })
      .then((response)=>{
        for(var i = 0; i < response.data.length; i++){
          this.items.push(response.data[i])
        }
        console.log(this.items[0])
      })
      .catch((error)=>{
        window.alert(`The Database Server returned an error: ${error}`);
      })
    }
  },
  watch:{
    PlotId: function(newVal){
      this.PlotId = newVal
      this.items = []
      this.getData()
    }
  }

}
</script>
<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.col-8{
  border-right: 1px solid black;
}
</style>