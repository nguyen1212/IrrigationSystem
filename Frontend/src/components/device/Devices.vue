<template>
  <div class="devices">
    <div class="row">
      <div class="col-8" >
        <apexchart type="area" height="400"  :options="chartOptions" :series="series"></apexchart>
      </div>
      <div class="col">
        <p> <strong> Device Information </strong></p>
        <div style="width: 100%; ">
        <div>
          <p>Type: <b-form-select style="width: 80%; float: right;" v-model="selectedType" :options="types"></b-form-select></p>
        </div>
        <div>
          <p>Name: <b-form-select style="width: 80%; float: right;" v-model="selectedDevice" :options="selectedDeviceList" v-bind:disabled="selectedType === null"></b-form-select></p>
        </div>
        <div>
          <p>Date: <b-form-datepicker style="width: 80%; float: right;" id="ex-disabled-readonly" v-model="date"  :disabled="selectedDevice === null"></b-form-datepicker></p>
        </div>
        </div>
        <br>
        <hr/>
        <p> <strong> Statistics </strong> </p>
        <div style="float: left">
        <p> Average: {{this.avg}}</p>
        <p> Highest: {{this.max}} - {{this.maxTime}}</p>
        <p> Lowest: {{this.min}} - {{this.minTime}}<p></div>
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
        avg: '',
        max: '',
        maxTime: '',
        min: '',
        minTime: '',
        date: '',
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
              text: 'Time Interval'
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
    getDeviceDataMeasurement(){
      var self = this
      axios.post("http://localhost:8080/devices/data/log", {
        start_time: self.start_time,
        end_time: self.end_time,
        deviceName: 'SOIL',
        type: self.selectedType
      })
      .then((response)=>{
        this.deviceData = response.data.Data
        this.min = response.data.min
        this.max = response.data.max
        this.avg = response.data.avg
        this.maxTime = new Date(response.data.maxTime).toISOString().split("T")[1]
        this.minTime = new Date(response.data.minTime).toISOString().split("T")[1]
        console.log(response.data.Data)
        this.series = [{
          data: response.data.Data
        }]
      })
      .catch((error)=>{
        window.alert(`Cannot plot data due to error: ${error}`)
      })
    },
    clearPlotData(){
      this.series = [{
          data: []
      }]
    }
  },
  watch:{
    PlotId: function(newVal){
      this.PlotId = newVal
      this.items = []
      // this.getDevice()
    },
    selectedType: function(newType){
      this.selectedType = newType
      this.clearPlotData()
      if (newType == 'soil')
        this.selectedDeviceList = this.soilDevices
      else if (newType == 'temp')
        this.selectedDeviceList = this.tempDevices
      else
        this.selectedDeviceList = this.humidDevices
    },
    selectedDevice: function(newDevice){
      this.selectedDevice = newDevice
      this.clearPlotData()
      if (this.start_time != '' && this.end_time != ''){
        this.getDeviceDataMeasurement()}
    },
    date: function(newDate){
      this.clearPlotData()
      this.date = newDate
      let yesterday = new Date(this.date.toString())
      yesterday.setDate(yesterday.getDate() - 1)
      yesterday = yesterday.toISOString()
      yesterday = yesterday.split("T")[0]
      this.start_time = yesterday + "T" + "17:00:00" + "Z"
      this.end_time = this.date + "T" + "17:00:00" + "Z"
      this.getDeviceDataMeasurement()
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