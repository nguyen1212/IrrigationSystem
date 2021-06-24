<template>
  <div class="devices">
    <div class="row">
      <div class="col-8" >
        <apexchart type="area" height="400"  :options="chartOptions" :series="series"></apexchart>
      </div>
      <div class="col">
        <p> <strong> Device Information </strong></p>
        <div style="width: 100%; text-align: left;">
          <p>Type: <b-form-select style="width: 80%; float: right; margin-right: 5%" v-model="selectedType" :options="types"></b-form-select></p>
          <p>Name: <b-form-select style="width: 80%; float: right; margin-right: 5%" v-model="selectedDevice" :options="selectedDeviceList" v-bind:disabled="selectedType === null"></b-form-select></p>
          <p>Date: <b-form-datepicker style="width: 80%; float: right; margin-right: 5%" v-model="date"  :disabled="selectedDevice === null"></b-form-datepicker></p>
        </div>
        <br>
        <hr/>
        <p> <strong> Statistics </strong> </p>
        <div style="text-align: left">
          <p id="highest"> Highest:</p>
          <p id="lowest"> Lowest:<p>
          <p id="avg">Average: </p>
        </div>
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
        unit: '',
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
        soilDevices: ['SOIL'],
        tempDevices: ['TEMP'],
        humidDevices: ['HUMID'],
        deviceData: [],
        series: [{
            name: '',
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
                download: false,
                pan: false,
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
  
  created: function(){
    if (this.PlotId != ''){
      this.getDevice()
    }
  },
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
        deviceName: self.deviceName,
        type: self.selectedType,
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
        if (this.selectedType == 'temp')
          document.getElementById("avg").innerHTML = 'Average:' + this.avg + '&deg;' + 'C';
        else
          document.getElementById("avg").innerHTML = 'Average:' + this.avg + '%';
        document.getElementById("highest").innerHTML = 'Highest: ' + this.max  + ' - ' + this.maxTime.slice(0, -5);
        document.getElementById("lowest").innerHTML = 'Lowest: ' + this.min  + ' - ' + this.minTime.slice(0, -5);
      })
      .catch((error)=>{
        window.alert(`Cannot plot data due to error: ${error}`)
      })
    },
    clearPlotData(){
      this.series = [{
          data: []
      }]
    },
    clearAvgData(){
      document.getElementById("avg").innerHTML = 'Average:';
      document.getElementById("highest").innerHTML = 'Highest:';
      document.getElementById("lowest").innerHTML = 'Lowest:';
    }
  },
  watch:{
    PlotId: function(newVal){
      this.PlotId = newVal
      this.getDevice()
    },
    selectedType: function(newType){
      this.selectedType = newType
      if (newType == 'temp'){
        this.clearAvgData();
        this.chartOptions = {
          yaxis:{
            title:{
              text: 'Celcius',
            }
          }
        },
        this.series = [{
          name: 'Degree',
          data: []
        }],
        this.selectedDeviceList = this.tempDevices;
      }
      else{
        this.clearAvgData();
        this.chartOptions = {
          yaxis:{
            title:{
              text: 'Percentage'
            }
          }
        },
        this.series = [{
          name: 'Percentage',
          data: []
        }]
      if (newType == 'soil')
        this.selectedDeviceList = this.soilDevices
      else
        this.selectedDeviceList = this.humidDevices
      }
    },
    selectedDevice: function(newDevice){
      this.clearAvgData();
      this.selectedDevice = newDevice
      this.deviceName = newDevice
      if (newDevice != "SOIL")
        this.deviceName = 'TEMP-HUMID'
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
      console.log(this.start_time, this.end_time, this.type)
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