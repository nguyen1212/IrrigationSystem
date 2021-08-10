<template>
    <div>
      <div class="mt-5">
      <h1><b>Timer Preset</b></h1>
        <div v-if="this.PlotId != ''" class="pt-5">
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
        bus.$on("select-repeat", (presetId)=>{
          this.repeatSelect = presetId;
        })
        bus.$on("select-interval", (presetId)=>{
          this.intervalSelect = presetId;
        })
        bus.$on('deselect-interval', () => {
          this.intervalSelect = -1
        });
        bus.$on('deselect-repeat', () => {
          this.repeatSelect = -1
        });
      }
      else
      {
        window.alert("Please select a plot!")
      }
    },
    name: 'timerMode',
    methods: {
      handleGetRepeatPresetList()
      {
        this.repeatContent = [
          {
              "id": 0,
              "name": "Preset 1",
              "time": "11:11:00",
              "amount": "20",
              "repeat": "Everyday",
          },
          {
              "id": 1,
              "name": "Preset 2",
              "time": "03:09:00",
              "amount": "30",
              "repeat": "Every week",
          },
          {
              "id": 2,
              "name": "Preset 3",
              "time": "10:09:00",
              "amount": "10",
              "repeat": "Never",
          },
          {
              "id": 3,
              "name": "Preset 4",
              "time": "5:03:00",
              "amount": "40",
              "repeat": "3 days",
          }
      ],
      this.repeatSelect = -1
      console.log("new preset list fetched")
      bus.$emit('timerPreset', this.repeatContent)
      },
      handleGetIntervalPresetList()
      {
        this.intervalContent = [
          {
              "id": 0,
              "name": "Preset 1",
              "amount": "80",
              "interval": "1 hour",
          },
          {
              "id": 1,
              "name": "Preset 2",
              "amount": "30",
              "interval": "5 hours",
          },
          {
              "id": 2,
              "name": "Preset 3",
              "amount": "10",
              "interval": "2 hours",
          },
          {
              "id": 3,
              "name": "Preset 4",
              "amount": "40",
              "interval": "5 hours",
          }
        ],
        this.intervalSelect = -1;
        console.log("new preset list fetched");
        bus.$emit("intervalPreset", this.intervalContent);
      },
      handleRepeatAddPreset(form)
      {
         this.repeatContent.push({
           id: this.repeatContent.length,
           name: form.content.name,
           time: form.content.time,
           amount: form.content.amount,
           repeat: form.content.repeatValue
         })
         bus.$emit('timerPreset', this.repeatContent)
      },
      handleRepeatUpdatePreset(form)
      {
        if (form.content.name){this.repeatContent[form.id].name = form.content.name}
        if (form.content.time){this.repeatContent[form.id].time = form.content.time}
        if (form.content.amount){this.repeatContent[form.id].amount = form.content.amount}
        if (form.content.repeatValue){this.repeatContent[form.id].repeat = form.content.repeatValue}
        bus.$emit('timerPreset', this.repeatContent)
      },
      handleRepeatDeletePreset(id)
      {
        this.repeatContent.splice(id,1)
        if (this.repeatSelect < id){this.repeatSelect -= 1;}
        bus.$emit('timerPreset', this.repeatContent)
      },
      handleRepeatSelectPreset(id)
      {
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
         this.intervalContent.push({
           id: this.intervalContent.length,
           name: form.content.name,
           amount: form.content.amount,
           interval: form.content.interval
         })
         bus.$emit("intervalPreset", this.intervalContent);
      },
      handleIntervalUpdatePreset(form)
      {
        if (form.content.name){this.repeatContent[form.id].name = form.content.name}
        if (form.content.interval){this.repeatContent[form.id].interval = form.content.interval}
        if (form.content.amount){this.repeatContent[form.id].amount = form.content.amount}
        bus.$emit("intervalPreset", this.intervalContent);
      },
      handleIntervalDeletePreset(id)
      {
        this.intervalContent.splice(id,1)
        if (this.intervalSelect < id){this.intervalSelect -= 1;}
        bus.$emit("intervalPreset", this.intervalContent);
      },
      handleIntervalSelectPreset(id)
      {
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