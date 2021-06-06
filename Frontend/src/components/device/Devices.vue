<template>
  <div class="devices">
    <div class="row">
      <div id="chart" ref="chart" class="col-8" >
        <apexchart type="area" height="400"  :options="chartOptions" :series="series"></apexchart>
      </div>
      <div class="col">
        <p> Device Information </p>
        <div> Type: <b-form-select style="width: 50%" v-model="selectedType" :options="types"></b-form-select></div>
        <div> Type: <b-form-select style="width: 50%" v-model="selectedDevice" :options="selectedDeviceList" ></b-form-select></div>
        <p>{{selectedDevice}}</p>
        <hr/>
        <p> Device Type </p>
      </div>
    </div>
  </div>
</template>
<script>
import axios from 'axios'
import VueApexCharts from 'vue-apexcharts'

export default {
  name: 'devices',
  components: {
    apexchart: VueApexCharts
  },
  props:{
    UserId: String,
    PlotId: String,
  },
  data() {
    {
      return {
        start_time: '',
        end_time: '',
        deviceName: '',
        deviceType: '',
        selectedType: null,
        selectedDevice: null,
        selectedDeviceList: [],
        types: [
          {value: 'soil', text: 'Soil Moisture'},
          {value: 'temp', text: 'Temperature'},
          {value: 'humid', text: 'Humidity'}
        ],
        soilDevices: ['1', '2', '3'],
        tempDevices: ['4', '5', '6'],
        humidDevices: ['7', '8', '9'],
        deviceData: [],
        series: [{
            name: 'percentage',
            data: [],
          }],
        chartOptions:{
          chart: {
            type: 'area',
            stacked: false,
            zoom: {
              type: 'x',
              enabled: true,
              autoScaleYaxis: true
            },
            toolbar: {
              tools:{
                reset: false,
                download: false
              },
              autoSelected: 'zoom'
            }
          },
          dataLabels: {
            enabled: true
          },
          markers: {
            size: 0,
          },
          title: {
            text: 'Data Measurement',
            align: 'left'
          },
          fill: {
            type: 'gradient',
            gradient: {
              shadeIntensity: 1,
              inverseColors: false,
              opacityFrom: 0.5,
              opacityTo: 0,
              stops: [0, 90, 100]
            },
          },
          yaxis: {
            title: {
              text: 'Percentage'
            },
            type: 'string',
          },
          xaxis: {
            type: 'datetime',
            title: {
              text: 'Date&Time Interval'
            },
          },
          tooltip: {
            shared: false,
          }
        }
        }}
  },
  
  created: function(){},
  methods: {
    getDevice(){
      var self = this
      axios.post("http://localhost:8080/devices/data/info",{
        PlotId: self.PlotId
      })
      .then((response)=>{
        this.soilDevices = response.data.soilDevices
        this.humidDevices = response.data.humidDevices
        this.tempDevices = response.data.tempDevices
        // console.log(this.items[0])
      })
      .catch((error)=>{
        window.alert(`The Database Server returned an error: ${error}`);
      })
    },
    plotData(deviceName, start_time, end_time){
      axios.post("http://localhost:8080/devices/data/log", {
        start_time: "2021-05-04T02:45:00Z",
        end_time: "2021-06-04T13:47:00Z",
        deviceName: 'SOIL',
        type: 'soilmoist'
      })
      .then((response)=>{
        this.deviceData = response.data.Data
        this.series = [{
          data: response.data.Data
        }]
      })
      .catch((error)=>{
        window.alert(`Cannot plot data due to error: ${error}`)
      })
    }
  },
  watch:{
    PlotId: function(newVal){
      this.PlotId = newVal
      this.items = []
      this.getDevice()
    },
    selectedType: function(newType){
      this.selectedType = newType
      if (newType == 'soil')
        this.selectedDeviceList = this.soilDevices
      else if (newType == 'temp')
        this.selectedDeviceList = this.tempDevices
      else
        this.selectedDeviceList = this.humidDevices
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