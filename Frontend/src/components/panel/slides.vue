<template>
  <div>
    <b-carousel
      id="carousel-1"
      v-model="slide"
      :interval="0"
      controls
      indicators
      style="
        text-shadow: 1px 1px 2px #333;
        height: 6cm;
        width: 15cm;
        margin: auto;
        backdrop-filter: blur(20px);
        border-radius: 20px;
      "
      @sliding-start="onSlideStart"
      @sliding-end="onSlideEnd"
    >
      <b-carousel-slide caption="Select Plot" img-blank>
        <template>
          <div>
            <b-form-select
              style="width: 60%"
              v-model="selectedPlot"
              :options="plots"
            ></b-form-select>
          </div>
        </template>
        <br /><br /><br /><br /><br /><br /><br /><br /><br /><br /><br /><br /><br /><br /><br /><br />
      </b-carousel-slide>

      <b-carousel-slide img-blank>
        <template>
          <div>
            <div>
              <b-button variant="dark" size="lg" @click="send"
                ><b-icon
                  icon="power"
                  aria-hidden="true"
                  scale="2"
                  v-bind:variant="forcemode"
                ></b-icon
              ></b-button>
            </div>
            <div>
              <b-form-select
                style="width: 50%; margin-top: 4mm"
                v-model="selectedMode"
                :options="modes"
                :disabled="selectedPlot === ''"
              ></b-form-select>
            </div>
            <div class="info" style="display: flex; justify-content: center">
              <p>
                {{ soil }}% <br />
                Soil Moisture
              </p>
              <p>
                {{ temp + "&deg;" }}C <br />
                Temperature
              </p>
              <p>
                {{ humid }}% <br />
                Humidity
              </p>
            </div>
          </div>
        </template>
        <br /><br /><br /><br /><br /><br /><br /><br /><br /><br /><br /><br /><br /><br />
      </b-carousel-slide>
    </b-carousel>
  </div>
</template>

<script>
import { bus } from "../../main";
import axios from "axios";
export default {
  props: {
    ws: WebSocket,
    soil: String,
    temp: String,
    humid: String,
    PlotId: String,
    UserId: String,
  },
  name: "slides",
  data() {
    return {
      slide: 0,
      sliding: null,
      forceData: "",
      forcemode: "",
      selectedPlot: "",
      selectedMode: "",
      plotName: "",
      plots: [{ value: "", text: "Please select an option" }],
      modes: [
        {
          value: null,
          text: "Please select...",
        },
        {
          label: "Auto Mode",
          options: [],
        },
        {
          label: "Interval Mode",
          options: [],
        },
        {
          label: "Repeat Mode",
          options: [],
        },
      ],
    };
  },
  created() {
    this.getPlot();
    bus.$on("autoPresets", (autoPresets) => {
      this.modes[1].options = [];
      for (let i = 0; i < autoPresets.length; i++) {
        this.modes[1].options.push({
          value: { preset: autoPresets[i], mode: "auto" },
          text: "[Auto] " + autoPresets[i]["name"],
        });
      }
    });
    bus.$on("intervalPreset", (intervalPresets) => {
      this.modes[2].options = [];
      for (let i = 0; i < intervalPresets.length; i++) {
        this.modes[2].options.push({
          value: { preset: intervalPresets[i], mode: "interval" },
          text: "[Interval] " + intervalPresets[i]["name"],
        });
      }
    });
    bus.$on("timerPreset", (timerPresets) => {
      this.modes[3].options = [];
      for (let i = 0; i < timerPresets.length; i++) {
        this.modes[3].options.push({
          value: { preset: timerPresets[i], mode: "repeat" },
          text: "[Repeat] " + timerPresets[i]["name"],
        });
      }
    });
  },
  destroyed() {
    bus.$off("autoPresets");
  },
  methods: {
    onSlideStart() {
      this.sliding = true;
    },
    onSlideEnd() {
      this.sliding = false;
    },
    toggle(mode) {
      if (mode == "success") return "danger";
      else return "success";
    },
    send() {
      var self = this;
      self.forcemode = this.toggle(self.forcemode);
      if (self.forcemode == "success") self.forceData = "0";
      else self.forceData = "1";
      this.ws.send(
        JSON.stringify({
          userId: "admin@gmail.com",
          plotId: "newyork",
          name: "RELAY",
          data: self.forceData,
        })
      );
    },
    choosePreset(modeType, preset) {
      if (modeType == "auto"){
      var url = "http://127.0.0.1:8010/api/preset/" + modeType;
      axios
        .post(url,
          {
            id: preset.id,
            name: preset.name,
            moistureThreshold: preset.moistureThreshold,
            notes: preset.notes,
          }
        )
        .then((response) => {
          console.log(response);
          bus.$emit('select-preset', preset.id);
          bus.$emit('deselect-interval');
          bus.$emit('deselect-repeat');
        })
        .catch((error) =>
          window.alert(`Error while handling POST request: ${error}`)
        );
      }
      if (modeType == "interval"){
      url = "http://127.0.0.1:8010/api/preset/" + modeType;
      axios
        .post(url,
          {
            id: preset.id,
            name: preset.name,
            amount: preset.amount,
            interval: preset.interval,
          }
        )
        .then((response) => {
          console.log(response);
          bus.$emit('select-interval', preset.id);
          bus.$emit('deselect-repeat');
          bus.$emit('deselect-auto');
        })
        .catch((error) =>
          window.alert(`Error while handling POST request: ${error}`)
        );
      }
      if (modeType == "repeat"){
      url = "http://127.0.0.1:8010/api/preset/" + modeType;
      axios
        .post(url,
          {
            id: preset.id,
            name: preset.name,
            amount: preset.amount,
            time: preset.time,
            repeat: preset.repeat
          }
        )
        .then((response) => {
          console.log(response);
          bus.$emit('select-repeat', preset.id);
          bus.$emit('deselect-interval');
          bus.$emit('deselect-auto');
        })
        .catch((error) =>
          window.alert(`Error while handling POST request: ${error}`)
        );
      }
    },
    getPlot() {
      //var self = this;
      // axios
      //   .post("http://localhost:8080/devices/data/info", {
      //     UserId: self.UserId,
      //   })
      //   .then((response) => {
      //     var newPlot = [
      //       { value: response.data.Id, text: response.data.PlotName },
      //     ];
      //     this.plots = [{ value: "", text: "Please select an option" }].concat(
      //       newPlot
      //     );

      //     // console.log(this.items[0])
      //   })
      //   .catch((error) => {
      //     window.alert(`The Database Server returned an error: ${error}`);
      //   });
      var newPlot = [{ value: "plot1", text: "Plot 1" }];
      this.plots = [{ value: "", text: "Please select an option" }].concat(
        newPlot
      );
    },
  },
  watch: {
    selectedPlot: function (newPlot) {
      this.selectedPlot = newPlot;
      this.$emit("plotSelection", this.selectedPlot);
    },
    selectedMode: function (newMode) {
      this.selectedMode = newMode;
      console.log(newMode.mode, newMode.preset);
      this.choosePreset(newMode.mode, newMode.preset)
    },
  },
};
</script>
<style scoped>
.info p {
  margin-bottom: -2mm;
  margin: auto;
}
.info {
  margin-top: 0.5cm;
}
</style>