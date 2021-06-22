<template>
  <div class="waterlog">
    <div class="col">
      <div class="date-select">
        <date-picker
          v-model="range"
          lang="en"
          range
          type="date"
          format="DD/MM/YYYY"
          range-separator=" - "
          width="500"
        ></date-picker>
      </div>
      <!-- DISPLAY GRAPH -->
      <div class="row graph-content">
        <div class="col-12">
          <apexchart
            ref="waterlogchart"
            type="bar"
            height="400"
            :options="chartOptions"
            :series="series"
          ></apexchart>
        </div>
      </div>
      <div class="row info-content">
        <div class="col-2">
          <div class="row justify-content-center"><b>Average</b></div>
          <div class="row justify-content-center">
            <b-icon icon="droplet-fill" style="color: #2e93fa"></b-icon>
            {{ this.avg }}ml/Day
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";
import VueApexCharts from "vue-apexcharts";
import DatePicker from "vue2-datepicker";
import "vue2-datepicker/index.css";
export default {
  name: 'waterLog',
  components: {
    DatePicker,
    apexchart: VueApexCharts,
  },
  props: {
    UserId: String,
    PlotId: String,
  },
  data() {
    {
      return {
        range: '',
        avg: '',
        series: [
          {
            name: "Ammount of Water",
            data: [],
          },
        ],
        chartOptions: {
          chart: {
            type: "bar",
            height: 350,
          },
          plotOptions: {
            bar: {
              horizontal: false,
              columnWidth: "55%",
              endingShape: "rounded",
            },
          },
          dataLabels: {
            enabled: false,
          },
          stroke: {
            show: true,
            width: 2,
            colors: ["transparent"],
          },
          xaxis: {
            categories: [],
          },
          yaxis: {
            title: {
              text: "ml",
              style: {
                color: "#2e93fa",
                fontSize: "16px",
                fontFamily: "Helvetica, Arial, sans-serif",
                fontWeight: 600,
                cssClass: "apexcharts-yaxis-title",
              },
            },
          },
          fill: {
            opacity: 1,
          },
          tooltip: {
            y: {
              formatter: function (val) {
                return val + "ml";
              },
            },
          },
        },
      };
    }
  },
  created: function () {},
  methods: {
    getWaterLogData() {
      var self = this;
      // console.log(this.extractDate(self.range[0]));
      // console.log(this.extractDate(self.range[1]));
      axios
        .post("url", {
          startDate: this.extractDate(self.range[0]),
          endDate: this.extractDate(self.range[1]),
        })
        .then((response) => {
          this.avg = response.data.avg;
          this.series = [
            {
              data: response.data.Data,
            },
          ];
        });
    },
    clearPlotData() {
      // console.log("clear");
      this.series = [
        {
          data: [],
        },
      ];
    },
    getLabelArray(start, end) {
      var arr = new Array();
      var dt = new Date(start);
      while (dt <= end) {
        arr.push(new Date(dt));
        dt.setDate(dt.getDate() + 1);
      }
      for (var i = 0; i < arr.length; i++) {
        arr[i] = this.extractLabel(arr[i]);
      }
      return arr;
    },
    extractLabel(date) {
      const rdate = date.getDate() + "/" + (date.getMonth() + 1);
      return rdate;
    },
    extractDate(date) {
      const rdate =
        date.getFullYear() +
        "-" +
        ("0" + (date.getMonth() + 1)).slice(-2) +
        "-" +
        ("0" + date.getDate()).slice(-2);
      return rdate;
    },
  },
  watch: {
    range: function () {
      this.clearPlotData();
      var labels = this.getLabelArray(this.range[0], this.range[1]);
      this.chartOptions = {
        xaxis: {
          categories: labels,
        },
      };
      // console.log(this.chartOptions.xaxis.categories);
      this.getWaterLogData();
    },
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.date-select {
  padding-top: 15px;
  margin: auto;
}
.info-content {
  margin-top: 10px;
  margin-bottom: 10px;
}
</style>