<template>
    <div class="pt-5">
      <h1><b>Auto Preset</b></h1>
      <div v-if="this.PlotId != ''" class="pt-5">
        <ListView
        :info="JSON.stringify(this.info)" 
        :selected="select" 
        @add-preset="handleAddPreset" 
        @update-preset="handleUpdatePreset"
        @delete-preset="handleDeletePreset"
        @select-preset="handleSelectPreset"
        />
      </div>
    </div>
</template>

<script>
// import axios from 'axios'
import {bus} from '../../main'
import ListView from '../auto mode/ListView.vue'
export default {
    name: 'autoMode',
    props:{
      UserId: String,
      PlotId: String,
    },
    data() {
        return {
            info: [
            {
                "id": 0,
                "name": "Daisy Preset",
                "moistureThreshold": 80,
                "notes": "my preset for my precious daisies that i care about so much i cannot live without u"
            },
            {
                "id": 1,
                "name": "Rosie Preset",
                "moistureThreshold": 30,
                "notes": "roses are red violets are blue"
            },
            {
                "id": 2,
                "name": "Rau Muong Preset",
                "moistureThreshold": 10,
                "notes": "rau muong xao toi"
            },
            {
                "id": 3,
                "name": "Bay leaves",
                "moistureThreshold": 40,
                "notes": "Bay leaves are delicious"
            }
        ],
            select: 2
        }
    },
    created() { 
      if (this.PlotId != '')
      {
        this.handleGetPresetList();
        bus.$on("select-preset", (presetId)=>{
          this.handleSelectPreset(presetId);
        })
      }
      else
      {
        window.alert("Please select a plot!")
      }
    },
    methods: {
      handleGetPresetList()
      {
        // axios
        // .get('url')
        // .then((response) => {
        //     this.info = response.presetContent
        //     this.select = response.selectedPreset
        // })
        // .catch((error) => window.alert(`Error while handling GET request: ${error}` ))
        
        console.log("new preset list fetched")
        bus.$emit('autoPresets', this.info)
      },
      handleAddPreset(form)
      {
        //   axios
        //   .post('url', form)
        //   .then((response) => {
        //   console.log(response);
        //  this.handleGetPresetList();
        //  })
        //   .catch((error) => window.alert(`Error while handling POST request: ${error}` ))
        this.info.push({
          id: this.info.length,
          name: form.content.name,
          moistureThreshold: form.content.moisture,
          notes: form.content.notes,
        })
        // console.log("form to be added:", form)

      },
      handleUpdatePreset(form)
      {
        // axios
        // .put('/api/',form)
        // .then((response) => {
        //     console.log(response);
        //     this.handleGetPresetList();
        // })
        // .catch((error) => window.alert(`Error while handling PUT request: ${error}` ))
        // // console.log("form to be updated:", form)
        if (form.content.name){this.info[form.id].name = form.content.name;}
        if (form.content.moisture>0){this.info[form.id].moistureThreshold = form.content.moisture;}
        if (form.content.notes){this.info[form.id].notes = form.content.notes;}
        console.log(this.info[form.id].name)
        console.log(this.info[form.id].moistureThreshold)
        console.log(this.info[form.id].notes)
      },
      handleDeletePreset(id)
      {
        // axios
        // .delete('url', id)
        // .then((response) => {
        //     console.log(response);
        // })
        // .catch((error) => window.alert(`Error while handling GET request: ${error}` ))
        /*console.log(`preset ${id} deleted`) */
        this.info.splice(id,1);
        if (this.select < id){this.select -= 1;}
      },
      handleSelectPreset(id)
      {
        this.select = id;
        console.log("run");
        // console.log(`preset ${id} selected`)
        //   axios
        //   .post('url', id)
        //   .then((response) => {
        //   console.log(response);
        //   this.select = id;
        //  //this.handleGetPresetList();
        //  })
        //   .catch((error) => window.alert(`Error while handling POST request: ${error}` ))
      },
    },
    watch: {
      PlotId: function(newPlot){
        this.PlotId = newPlot
        this.handleGetPresetList()
      }
    },
    components: {
        ListView
    }
}
</script>
