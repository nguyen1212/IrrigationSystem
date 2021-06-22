<template>
    <div>
      <div class="pt-5">
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
import axios from 'axios'
import ListView from '../auto mode/ListView.vue'
export default {
    data() {
        return {
            info: [],
            select: 1
        }
    },
    created() { 
      this.handleGetPresetList();
    },
    name: 'autoMode',
    methods: {
      handleGetPresetList()
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
            this.info = response.presetContent
            this.select = response.selectedPreset
        })
        .catch((error) => window.alert(`Error while handling GET request: ${error}` ))
        /* this.info = [
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
                "id": 4,
                "name": "Bay leaves",
                "moistureThreshold": 40,
                "notes": "Bay leaves are delicious"
            }
        ]
        this.select=1
        console.log("new preset list fetched")*/
      },
      handleAddPreset(form)
      {
          axios
          .post('url', form)
          .then((response) => {
          console.log(response);
         this.handleGetPresetList();
         })
          .catch((error) => window.alert(`Error while handling POST request: ${error}` ))
        // console.log("form to be added:", form)

      },
      handleUpdatePreset(form)
      {
        axios
        .put('url',form)
        .then((response) => {
            console.log(response);
            this.handleGetPresetList();
        })
        .catch((error) => window.alert(`Error while handling PUT request: ${error}` ))
        // console.log("form to be updated:", form)
      },
      handleDeletePreset(id)
      {
        axios
        .delete('url', id)
        .then((response) => {
            console.log(response);
        })
        .catch((error) => window.alert(`Error while handling GET request: ${error}` ))
        /*console.log(`preset ${id} deleted`) */
      },
      handleSelectPreset(id)
      {
        // console.log(`preset ${id} selected`)
          axios
          .post('url', id)
          .then((response) => {
          console.log(response);
         this.handleGetPresetList();
         })
          .catch((error) => window.alert(`Error while handling POST request: ${error}` ))
      }
    },
    components: {
        ListView
    }
}
</script>
