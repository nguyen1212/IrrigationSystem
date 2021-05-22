<template>
  <div style="margin-top: 20px">
    <b-card no-body>
      <b-tabs card>
      <!-- TIME MODE -->
        <b-tab title="Timer Mode" active>
          <b-card-text style="display: flex; color: #c0c0c0;">Current Mode: Timer Mode
            <div style="width: 471px"/>
            <b-button-group>
              <b-button style="background-color: #c0c0c0; color: black">Timer</b-button>
              <b-button style="background-color: #ffffff; color: black">Repeat</b-button>
            </b-button-group>
          </b-card-text>

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
              >
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
</template>

<script>
export default {
  name: 'timerMode',
  data() {
      return {
        intervalValue: 0,
        interval: ['1 hours', '2 hours','3 hours','4 hours','5 hours','6 hours','7 hours','8 hours','9 hours','10 hours','11 hours','12 hours','13 hours','14 hours','15 hours','16 hours','17 hours','18 hours','19 hours','20 hours','21 hours','22 hours','23 hours', '24 hours'],
        switchValue1: false,
        switchValue2: true,
        form:{
          label: '',
          time: '',
          wateredInValue: 0,
          wateredIn:['5 mins', '10 mins', '15 mins', '20 mins', '30 mins'],
          repeatValue: 0,
          repeat:['Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday' , 'Saturday' , 'Sunday'],
        },
        showModal: false,
      }
    },
    methods: {
      intervalFormatter(intervalValue) {
        return this.interval[intervalValue]
      },
      wateredInFormatter(wateredInValue) {
        return this.form.wateredIn[wateredInValue]
      },
      repeatFormatter(repeatValue) {
        return this.form.repeat[repeatValue]
      },
      show() {
        this.showModal = true
      },
      hide(){
        this.showModal = false
      },
      resetModal(){
        this.resetForm()
      },
      resetForm()
      {
        this.form.name=""
        this.form.time=""
        this.form.wateredInValue=0
        this.form.repeatValue=0
      },
      handleSubmit() {
        
      }
    }
}
</script>
<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>