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

<<<<<<< HEAD
        @add-interval-preset="handleIntervalAddPreset" 
        @update-interval-preset="handleIntervalUpdatePreset"
        @delete-interval-preset="handleIntervalDeletePreset"
        @select-interval-preset="handleIntervalSelectPreset"
        />
      </div>
    </div>
=======
          <b-card-text style="display: flex; color: #c0c0c0; margin-bottom: 0">start at</b-card-text>
          <template>
            <div>
              <b-form-timepicker style="width: 280px" v-model="value" locale="en"></b-form-timepicker>
            </div>
          </template>

          <b-card-text style="display: flex; color: #c0c0c0; margin-top: 1rem; margin-bottom: 0">set interval</b-card-text>
          <template>
            <div>
              <b-form-spinbutton
                v-model="intervalValue"
                :formatter-fn="intervalFormatter"
                min="0"
                max="23"
                wrap
                style="width: 280px"
              ></b-form-spinbutton>
            </div>
          </template>
        </b-tab>

      <!-- REPEAT MODE -->
        <b-tab title="Repeat Mode">
          <b-card-text style="display: flex; color: #c0c0c0">Current Mode: Repeat Mode
            <div style="width: 460px"/>
            <b-button-group>
              <b-button style="background-color: #ffffff; color: black">Timer</b-button>
              <b-button style="background-color: #c0c0c0; color: black">Repeat</b-button>
            </b-button-group>
          </b-card-text>

          <b-card-text style="display: flex">
            <div style= "width: 700px"/>
            <div>
              <b-button v-b-modal.repeat-box style="background-color: #ffffff; border-color: #ffffff">
                <img style = "width: 25px" src="./plus_icon.png"/>
              </b-button>
              <b-modal
              id="repeat-box"
              title="Add new preset"
              button-size="md"
              @show="resetModal"
              @hidden="resetModal"
              @ok="handleSubmit"
              />
              <div class="d-block text-center">
                <b-form v-if="show">
                  <b-form-group id="input-preset" label="Preset Name" label-for="input-preset">
                    <b-form-input
                      id="input-preset"
                      v-model="form.label"
                      placeholder="Enter name"
                      required
                    />
                  </b-form-group>

                  <b-form-group id="input-preset" label="Time" label-for="input-preset">
                    <b-form-input
                      id="input-preset"
                      v-model="form.time"
                      placeholder="Enter time"
                      required
                    />
                  </b-form-group>

                  <b-form-group id="input-preset" label="Watered In" label-for="input-preset">
                    <b-form-spinbutton 
                      v-model="form.wateredInValue"
                      :formatter-fn="wateredInFormatter"
                      min="0"
                      max="4"
                      wrap
                    />
                  </b-form-group>

                  <b-form-group id="input-preset" label="Repeat" label-for="input-preset">
                    <b-form-spinbutton 
                      v-model="form.repeatValue"
                      :formatter-fn="repeatFormatter"
                      min="0"
                      max="6"
                      wrap
                    />
                  </b-form-group>
                </b-form>
            </div>
            </div>
            <template #modal-footer="{ ok, cancel }">
              <b-button variant="success" @click="ok()">
                Submit
              </b-button>
              <b-button variant="danger" @click="cancel()">
                Cancel
              </b-button>
            </template>
          </b-card-text>
          <hr style="margin-top: 1rem" width="100%" size= "5px" align="center"/>

          <b-card-text>
            <div style="margin: 1rem 1rem 1rem 1rem; display: flex">
              <h1 style="margin: 0 0 0 0; width: 101px">11:11</h1>
              <div style= "width: 238px"/>
              <p style="margin: 0 0 0 0; padding-top: 10px; padding-bottom:10px; color: #c0c0c0">Click here to edit</p>
              <div style= "width: 200px"/>
              <b-form-checkbox switch size="lg" style="margin: 0 0 0 0; padding-top: 10px; padding-bottom:10px" v-model="switchValue1" name="check-button"></b-form-checkbox>
            </div>
            <hr style="margin-top: 1rem" width="100%" size= "5px" align="center"/>

            <div style="margin: 1rem 1rem 1rem 1rem; display: flex">
              <h1 style="margin: 0 0 0 0; width: 101px">17:03</h1>
              <div style= "width: 238px"/>
              <p style="margin: 0 0 0 0; padding-top: 10px; padding-bottom:10px; color: #c0c0c0">Click here to edit</p>
              <div style= "width: 200px"/>
              <b-form-checkbox switch size="lg" style="margin: 0 0 0 0; padding-top: 10px; padding-bottom:10px" v-model="switchValue2" name="check-button"></b-form-checkbox>
            </div>
            <hr style="margin-top: 1rem" width="100%" size= "5px" align="center"/>
          </b-card-text>
        </b-tab>
      </b-tabs>
    </b-card>
  </div>
>>>>>>> c157c11e132864da587d0330de90ba601cce497b
</template>

<script>
import axios from 'axios'
import TimerView from '../timer mode/TimerView.vue'
export default {
    data() {
        return {
            repeatContent: [],
            repeatSelect: -1,

            intervalContent: [],
            intervalSelect: -1

        }
    },
    created() { 
      this.handleGetRepeatPresetList();
      this.handleGetIntervalPresetList();
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
        ],
        this.intervalSelect = 2
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
    
    components: {
        TimerView
    }
}
</script>
