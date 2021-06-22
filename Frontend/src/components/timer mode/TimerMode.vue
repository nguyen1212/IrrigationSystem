<template>
    <div>
      <div class="mt-5">
        <TimerView
        :repeatContent="JSON.stringify(this.repeatContent)" 
        :repeatSelected="repeatSelect"
        :intervalContent="JSON.stringify(this.intervalContent)" 
        :intervalSelected="intervalSelect"  
        @add-repeat-preset="handleRepeatAddPreset" 
        @update-repeat-preset="handleRepeatUpdatePreset"
        @delete-repeat-preset="handleRepeatDeletePreset"
        @select-repeat-preset="handleRepeatSelectPreset"

        @add-interval-preset="handleIntervalAddPreset" 
        @update-interval-preset="handleIntervalUpdatePreset"
        @delete-interval-preset="handleIntervalDeletePreset"
        @select-interval-preset="handleIntervalSelectPreset"
        />
      </div>
    </div>
</template>

<script>
import axios from 'axios'
import TimerView from '../timer mode/TimerView.vue'
import {bus} from '../../main'
export default {
    props:{
      PlotId: String,
    },
    data() {
        return {
            repeatContent: [],
            repeatSelect: -1,

            intervalContent: [],
            intervalSelect: -1

        }
    },
    created() { 
      if (this.PlotId != ''){
        this.handleGetRepeatPresetList();
        this.handleGetIntervalPresetList();
      }
    },
    name: 'timerMode',
    methods: {
      handleGetRepeatPresetList()
      {
        axios
        .get('url')
        .then((response) => {
            /**
             * response: {
              * presetContent: {
              *      ****** content
              * }
              * selectedPreset: Number
             * } 
             */
            this.repeatContent = response.presetRepeatContent
            this.repeatSelect = response.selectedRepeatPreset
        })
        .catch((error) => window.alert(`Error while handling GET request: ${error}` ))
        this.repeatContent = [
          {
              "id": 0,
              "name": "Preset 1",
              "time": "11:11 AM",
              "amount": 80,
              "repeat": "Everyday",
          },
          {
              "id": 1,
              "name": "Preset 2",
              "time": "03:09 PM",
              "amount": 30,
              "repeat": "Every week",
          },
          {
              "id": 2,
              "name": "Preset 3",
              "time": "10:09 AM",
              "amount": 10,
              "repeat": "Never",
          },
          {
              "id": 3,
              "name": "Preset 4",
              "time": "5:03 PM",
              "amount": 40,
              "repeat": "3 days",
          }
      ],
      this.repeatSelect = 1
      console.log("new preset list fetched")
      bus.$emit('timerPreset', this.repeatContent)
      },

      handleGetIntervalPresetList()
      {
        axios
        .get('url')
        .then((response) => {
            /**
             * response: {
              * presetContent: {
              *      ****** content
              * }
              * selectedPreset: Number
             * } 
             */
            this.intervalContent = response.presetIntervalContent
            this.intervalSelect = response.selectedIntervalPreset
        })
        .catch((error) => window.alert(`Error while handling GET request: ${error}` ))
        this.intervalContent = [
          {
              "id": 0,
              "name": "Preset 1",
              "amount": 80,
              "interval": "1 hour",
          },
          {
              "id": 1,
              "name": "Preset 2",
              "amount": 30,
              "interval": "5 hours",
          },
          {
              "id": 2,
              "name": "Preset 3",
              "amount": 10,
              "interval": "2 hours",
          },
          {
              "id": 3,
              "name": "Preset 4",
              "amount": 40,
              "interval": "5 hours",
          }
        ],
        this.intervalSelect = 3
        console.log("new preset list fetched")
      },

      handleRepeatAddPreset(form)
      {
          axios
          .post('url', form)
          .then((response) => {
          console.log(response);
         this.handleGetRepeatPresetList();
         })
          .catch((error) => window.alert(`Error while handling POST request: ${error}` ))
        // console.log("form to be added:", form)

      },
      handleRepeatUpdatePreset(form)
      {
        axios
        .put('url',form)
        .then((response) => {
            console.log(response);
            this.handleGetRepeatPresetList();
        })
        .catch((error) => window.alert(`Error while handling PUT request: ${error}` ))
        // console.log("form to be updated:", form)
      },
      handleRepeatDeletePreset(id)
      {
        axios
        .delete('url', id)
        .then((response) => {
            console.log("ok");
        })
        .catch((error) => window.alert(`Error while handling GET request: ${error}` ))
        /*console.log(`preset ${id} deleted`) */
      },
      handleRepeatSelectPreset(id)
      {
        // console.log(`preset ${id} selected`)
        axios
        .post('url', id)
        .then((response) => {
        console.log(response);
        this.handleGetRepeatPresetList();
        })
        .catch((error) => window.alert(`Error while handling POST request: ${error}` ))
      },

      handleIntervalAddPreset(form)
      {
          axios
          .post('url', form)
          .then((response) => {
          console.log(response);
         this.handleGetIntervalPresetList();
         })
          .catch((error) => window.alert(`Error while handling POST request: ${error}` ))
        // console.log("form to be added:", form)

      },
      handleIntervalUpdatePreset(form)
      {
        axios
        .put('url',form)
        .then((response) => {
            console.log(response);
            this.handleGetIntervalPresetList();
        })
        .catch((error) => window.alert(`Error while handling PUT request: ${error}` ))
        // console.log("form to be updated:", form)
      },
      handleIntervalDeletePreset(id)
      {
        axios
        .delete('url', id)
        .then((response) => {
            console.log("ok");
        })
        .catch((error) => window.alert(`Error while handling GET request: ${error}` ))
        /*console.log(`preset ${id} deleted`) */
      },
      handleIntervalSelectPreset(id)
      {
        // console.log(`preset ${id} selected`)
        axios
        .post('url', id)
        .then((response) => {
        console.log(response);
        this.handleGetIntervalPresetList();
        })
        .catch((error) => window.alert(`Error while handling POST request: ${error}` ))
      }
    },
    watch:{
      repeatContent: function(newContent){
        this.repeatContent = newContent
        bus.$emit('timerPresets', newContent.concat(this.intervalContent))
      },
      intervalContent: function(newContent){
        this.intervalContent = newContent
        bus.$emit('timerPresets', this.intervalContent.concat(newContent))
      },
      PlotId: function(newPlot){
        this.PlotId = newPlot
        this.handleGetRepeatPresetList();
        this.handleGetIntervalPresetList();
      }
    },
    components: {
        TimerView
    }
}
</script>

