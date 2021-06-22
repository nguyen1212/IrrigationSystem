
<template>
  <b-tabs card>
    <b-tab title="Repeat Mode" active>
      <div class="d-flex justify-content-center">
        <b-container>
          <div  class="accordion" role="tablist">
            <b-row class=" pt-3 pb-3 border-bottom" >
              <b-col class="col-1"><b>#</b></b-col>
              <b-col class="leftalign col-3"><b>Name</b></b-col>
              <b-col class="leftalign col-5 middle"/>
              <b-col class="col-3"><b>Status</b></b-col>
            </b-row>
            <!-- magic -->
            <div :set="repeatCount = 0"/>

            <b-row class="border-bottom" role="tab" :key="o.id" v-for="o in JSON.parse(this.repeatContent)" :set="repeatCount=repeatCount+1">
              <b-link class='pt-3 pb-3 row w-100 ml-0' v-b-toggle="'repeatCollapse-' + repeatCount" href="#">
                <b-col class="col-1">
                  {{repeatCount}}
                </b-col>
                <b-col class="leftalign col-3">{{o.name}}</b-col>
                <b-col class="leftalign col-5 middle"><i style="color: gray">Click to view detail</i></b-col>
                <b-col class="pr-2 col-3">
                <div v-if="repeatSelected === o.id">
                  <span class="active-dot" v-b-popover.hover.top="'Active'"></span>
                </div>
                <div v-else>
                  <span class="inactive-dot" v-b-popover.hover.top="'Inactive'"></span>
                </div>
                </b-col>
              </b-link>
              <b-collapse class='w-100' accordion="my-accordion" :id="'repeatCollapse-' + repeatCount" role="tabpanel">
                <b-card class="border-0 ml-0">
                  <div class="row">
                    <div class="col-8">
                      <ul>
                      <li class="leftalign"><b>ID:</b> {{o.id}}</li>
                      <li class="leftalign"><b>Name:</b> {{o.name}}</li>
                      <li class="leftalign"><b>Time:</b> {{o.time}}</li>
                      <li class="leftalign"><b>Amount:</b> {{o.amount}}</li>
                      <li class="leftalign"><b>Repeat:</b> {{o.repeat}}</li>
                      </ul>
                    </div>
                  </div>
                  <div class="w-100 row">
                    <div class="ml-3"/>
                    <b-button v-if="repeatSelected !== o.id" class="col-1 ml-3" variant="success" @click="$emit('select-repeat-preset', o.id)" >Select</b-button>
                    <b-button class="col-1 ml-3" variant="dark" v-b-toggle="'repeatCollapse-inner-' + repeatCount" @click="resetEditForm">Edit</b-button>
                    <b-button class="col-1 ml-3" variant="danger" @click="$emit('delete-repeat-preset', o.id)">Delete</b-button>
                  </div>
                  <div>
                    <b-collapse class='w-100  mt-3' :id="'repeatCollapse-inner-' + repeatCount">
                      <b-card class="ml-0">
                        <b-form 
                          @submit="handleRepeatUpdatePreset" 
                          @reset="resetEditForm">
                          <b-form-group id="input-group-1" label="Preset Name" label-for="input-1">
                            <b-form-input
                            :id="'input-1-' + repeatCount"
                            v-model="repeatEditform.name"
                            placeholder="Enter name"
                            required
                            />
                          </b-form-group>

                          <b-form-group id="input-group-2" label="Preset Time" label-for="input-2">
                            <b-form-timepicker
                            :id="'input-2-' + repeatCount"
                            v-model="repeatEditform.time" 
                            />  
                          </b-form-group>

                          <b-form-group id="input-group-3" label="Amount" label-for="input-3">
                            <b-form-spinbutton 
                              :id="'input-3-' + repeatCount" 
                              v-model="repeatAmountValue"
                              :formatter-fn="repeatEditAmountFormatter"
                              min="0"
                              max="9"
                              wrap
                            />
                          </b-form-group>

                          <b-form-group id="input-group-4" label="Repeat" label-for="input-4">
                            <b-form-spinbutton 
                              :id="'input-4-' + repeatCount"
                              v-model="repeatValue"
                              :formatter-fn="repeatEditFormatter"
                              min="0"
                              max="7"
                              wrap
                            />
                          </b-form-group>
                      
                          <b-button type="submit" variant="primary" class="mr-3">Submit</b-button>
                          <b-button type="reset" variant="danger">Reset</b-button>
                        </b-form>
                      </b-card>
                    </b-collapse>
                  </div>
                </b-card>
              </b-collapse>
            </b-row>
              
            <b-row class="border-bottom" role="tab">
              <b-link class='pt-3 pb-3 row w-100 ml-0' v-b-toggle.repeat_collapse_add href="#">
                <b-col class="col-1">
                  <b-icon class="h3 mt-0" icon="plus-circle-fill" style="color: black;" variant="black"/>
                </b-col>

                <b-col class="leftalign col-3">Add</b-col>
                <b-col class="leftalign col-5 middle"><i style="color: gray">Click to add</i></b-col>
                <b-col class="pr-2 col-3"/>
              </b-link>

              <b-collapse class="w-100" id="repeat_collapse_add" accordion="my-accordion"  role="tabpanel">
                <b-card class="border-0 ml-0 w-100">
                  <b-form 
                    @submit="handleRepeatAddPreset" 
                    @reset="resetAddForm">
                    <b-form-group id="input-group-1" label="Preset Name" label-for="input-1">
                      <b-form-input
                        id="input-1"
                        v-model="repeatAddform.name"
                        placeholder="Enter name"
                        required
                      />
                    </b-form-group>

                    <b-form-group id="input-group-2" label="Preset Time" label-for="input-2">
                      <b-form-timepicker
                      id="input-2" 
                      v-model="repeatAddform.time" 
                      />  
                    </b-form-group>

                    <b-form-group id="input-group-3" label="Amount" label-for="input-3">
                      <b-form-spinbutton 
                        id="input-3"
                        v-model="repeatAmountValue"
                        :formatter-fn="repeatAddAmountFormatter"
                        min="0"
                        max="9"
                        wrap
                      />
                    </b-form-group>

                    <b-form-group id="input-group-4" label="Repeat" label-for="input-4">
                      <b-form-spinbutton 
                        id="input-4"
                        v-model="repeatValue"
                        :formatter-fn="repeatAddFormatter"
                        min="0"
                        max="7"
                        wrap
                      />
                    </b-form-group>

                    <b-button type="submit" variant="primary" class="mr-3">Submit</b-button>
                    <b-button type="reset" variant="danger">Reset</b-button>
                  </b-form>
                </b-card>
              </b-collapse>
            </b-row>
          </div>
        </b-container>
      </div>
    </b-tab>

    <b-tab title="Interval Mode">
      <div class="d-flex justify-content-center">
        <b-container>
          <div  class="accordion" role="tablist">
            <b-row class=" pt-3 pb-3 border-bottom" >
              <b-col class="col-1"><b>#</b></b-col>
              <b-col class="leftalign col-3"><b>Name</b></b-col>
              <b-col class="leftalign col-5 middle"/>
              <b-col class="col-3"><b>Status</b></b-col>
            </b-row>
            <!-- magic -->
            <div :set="intervalCount = 0"/>

            <b-row class="border-bottom" role="tab" :key="o.id" v-for="o in JSON.parse(this.intervalContent)" :set="intervalCount=intervalCount+1">
              <b-link class='pt-3 pb-3 row w-100 ml-0' v-b-toggle="'intervalCollapse-' + intervalCount" href="#">
                <b-col class="col-1">
                  {{intervalCount}}
                </b-col>
                <b-col class="leftalign col-3">{{o.name}}</b-col>
                <b-col class="leftalign col-5 middle"><i style="color: gray">Click to view detail</i></b-col>
                <b-col class="pr-2 col-3">
                  <div v-if="intervalSelected === o.id">
                    <span class="active-dot" v-b-popover.hover.top="'Active'"/>
                  </div>
                  <div v-else>
                    <span class="inactive-dot" v-b-popover.hover.top="'Inactive'"></span>
                  </div>
                </b-col>
              </b-link>
              <b-collapse class='w-100' accordion="my-accordion" :id="'intervalCollapse-' + intervalCount" role="tabpanel">
                <b-card class="border-0 ml-0">
                  <div class="row">
                    <div class="col-8">
                      <ul>
                      <li class="leftalign"><b>ID:</b> {{o.id}}</li>
                      <li class="leftalign"><b>Name:</b> {{o.name}}</li>
                      <li class="leftalign"><b>Amount:</b> {{o.amount}}</li>
                      <li class="leftalign"><b>Interval:</b> {{o.interval}}</li>
                      </ul>
                    </div>
                  </div>
                  <div class="w-100 row">
                    <div class="ml-3"/>
                    <b-button v-if="intervalSelected !== o.id" class="col-1 ml-3" variant="success" @click="$emit('select-interval-preset', o.id)" >Select</b-button>
                    <b-button class="col-1 ml-3" variant="dark" v-b-toggle="'intervalCollapse-inner-' + intervalCount" @click="resetEditForm">Edit</b-button>
                    <b-button class="col-1 ml-3" variant="danger" @click="$emit('delete-interval-preset', o.id)">Delete</b-button>
                  </div>
                  <div>
                    <b-collapse class='w-100  mt-3' :id="'intervalCollapse-inner-' + intervalCount">
                      <b-card class="ml-0">
                        <b-form 
                          @submit="handleIntervalUpdatePreset" 
                          @reset="resetEditForm">
                          <b-form-group id="input-group-1" label="Preset Name" label-for="input-1">
                            <b-form-input
                            :id="'input-1-' + intervalCount"
                            v-model="intervalEditform.name"
                            placeholder="Enter name"
                            required
                            />
                          </b-form-group>

                          <b-form-group id="input-group-2" label="Amount" label-for="input-2">
                            <b-form-spinbutton 
                              :id="'input-2-' + intervalCount" 
                              v-model="intervalAmountValue"
                              :formatter-fn="intervalEditAmountFormatter"
                              min="0"
                              max="9"
                              wrap
                            />
                          </b-form-group>

                          <b-form-group id="input-group-3" label="Interval" label-for="input-3">
                            <b-form-spinbutton 
                              :id="'input-4-' + intervalCount"
                              v-model="intervalValue"
                              :formatter-fn="intervalEditFormatter"
                              min="0"
                              max="6"
                              wrap
                            />
                          </b-form-group>
                      
                          <b-button type="submit" variant="primary" class="mr-3">Submit</b-button>
                          <b-button type="reset" variant="danger">Reset</b-button>
                        </b-form>
                      </b-card>
                    </b-collapse>
                  </div>
                </b-card>
              </b-collapse>
            </b-row>
              
            <b-row class="border-bottom" role="tab">
              <b-link class='pt-3 pb-3 row w-100 ml-0' v-b-toggle.interval_collapse_add href="#">
                <b-col class="col-1">
                  <b-icon class="h3 mt-0" icon="plus-circle-fill" style="color: black;" variant="black"></b-icon>
                </b-col>
                <b-col class="leftalign col-3">Add</b-col>
                <b-col class="leftalign col-5 middle"><i style="color: gray">Click to add</i></b-col>
                <b-col class="pr-2 col-3"/>
              </b-link>

              <b-collapse class="w-100" id="interval_collapse_add" accordion="my-accordion" role="tabpanel">
                <b-card class="border-0 ml-0 w-100">
                  <b-form 
                    @submit="handleIntervalAddPreset" 
                    @reset="resetAddForm">
                    <b-form-group id="input-group-1" label="Preset Name" label-for="input-1">
                      <b-form-input
                        id="input-1"
                        v-model="intervalAddform.name"
                        placeholder="Enter name"
                        required
                      />
                    </b-form-group>

                    <b-form-group id="input-group-2" label="Amount" label-for="input-2">
                      <b-form-spinbutton 
                        id="input-2"
                        v-model="intervalAmountValue"
                        :formatter-fn="intervalAddAmountFormatter"
                        min="0"
                        max="9"
                        wrap
                      />
                    </b-form-group>

                    <b-form-group id="input-group-3" label="Interval" label-for="input-3">
                      <b-form-spinbutton 
                        id="input-3"
                        v-model="intervalValue"
                        :formatter-fn="intervalAddFormatter"
                        min="0"
                        max="6"
                        wrap
                      />
                    </b-form-group>

                    <b-button type="submit" variant="primary" class="mr-3">Submit</b-button>
                    <b-button type="reset" variant="danger">Reset</b-button>
                  </b-form>
                </b-card>
              </b-collapse>
            </b-row>
          </div>
        </b-container>
      </div>
    </b-tab>
  </b-tabs> 
</template>

<script>
export default {
  name: 'RepeatView',
  components: {
  },
  data() {
    var repeatObj = JSON.parse(this.repeatContent)
    var intervalObj = JSON.parse(this.intervalContent)
    return {
      repeatObj,
      intervalObj,
      i: 0,
      repeatAmountValue: 0,
      repeatAmount:['5', '10', '15', '20', '25' , '30' , '35', '40', '45', '50'],
      repeatValue: 0,
      repeat:['Never', 'Everyday', '2 days', '3 days', '4 days' , '5 days' , '6 days', 'Every week'],

      intervalAmountValue: 0,
      intervalAmount:['5', '10', '15', '20', '25' , '30' , '35', '40', '45', '50'],
      intervalValue: 0,
      interval:['Never', 'Every hour', '2 hours', '3 hours', '4 hours' , '5 hours' , '6 hours'],

      repeatAddform : {
        name: '',
        time: '',
        amount: '',
        repeatValue: '',
      },
      repeatEditform : {
        name: '',
        time: '',
        amount: '',
        repeatValue: '',
      },

      intervalAddform : {
        name: '',
        amount: '',
        interval: ''
      },
      intervalEditform : {
        name: '',
        amount: '',
        interval: ''
      },
    }
  },
  props: {
    repeatContent: String,
    repeatSelected: Number,
    intervalContent: String,
    intervalSelected: Number,
  },

  methods: {
      handleRepeatAddPreset(event)
      {
        event.preventDefault()
        this.$emit('add-repeat-preset', this.repeatAddform)
      },            
      handleRepeatUpdatePreset(event)
      {
        event.preventDefault()
        this.$emit('update-repeat-preset', this.repeatEditform)
      },

      handleIntervalAddPreset(event)
      {
        event.preventDefault()
        this.$emit('add-interval-preset', this.intervalAddform)
      },            
      handleIntervalUpdatePreset(event)
      {
        event.preventDefault()
        this.$emit('update-interval-preset', this.intervalEditform)
      },   

      resetAddForm()
      {
        this.repeatAddform.name=""
        this.repeatAddform.time=""
        this.repeatAmountValue=0
        this.repeatValue=0
        this.intervalAddform.name=""
        this.intervalAmountValue=0
        this.intervalValue=0
        console.log("form reset")
      },

      resetEditForm()
      {
        this.repeatEditform.name=""
        this.repeatEditform.time=""
        this.repeatAmountValue=0
        this.repeatValue=0
        this.intervalEditform.name=""
        this.intervalAmountValue=0
        this.intervalValue=0
        console.log("form reset")
      },
      repeatAddAmountFormatter(repeatAmountValue) {
        this.repeatAddform.amount = this.repeatAmount[repeatAmountValue]
        return this.repeatAmount[repeatAmountValue]
      },
      repeatEditAmountFormatter(repeatAmountValue) {
        this.repeatEditform.amount = this.repeatAmount[repeatAmountValue]
        return this.repeatAmount[repeatAmountValue]
      },
      repeatAddFormatter(repeatValue) {
        this.repeatAddform.repeatValue = this.repeat[repeatValue]
        return this.repeat[repeatValue]
      },
      repeatEditFormatter(repeatValue) {
        this.repeatEditform.repeatValue = this.repeat[repeatValue]
        return this.repeat[repeatValue]
      },
      intervalAddAmountFormatter(intervalAmountValue) {
        this.intervalAddform.amount = this.intervalAmount[intervalAmountValue]
        return this.intervalAmount[intervalAmountValue]
      },
      intervalEditAmountFormatter(intervalAmountValue) {
        this.intervalEditform.amount = this.intervalAmount[intervalAmountValue]
        return this.intervalAmount[intervalAmountValue]
      },
      intervalAddFormatter(intervalValue) {
        this.intervalAddform.intervalValue = this.interval[intervalValue]
        return this.interval[intervalValue]
      },
      intervalEditFormatter(intervalValue) {
        this.intervalEditform.intervalValue = this.interval[intervalValue]
        return this.interval[intervalValue]
      },
  }
}
</script>

<style scoped>
.active-dot {
  height: 20px;
  width: 20px;
  background-color: #5cb85c;
  border-radius: 50%;
  display: inline-block;
}
.inactive-dot {
  height: 20px;
  width: 20px;
  background-color: #f0ad4e;
  border-radius: 50%;
  display: inline-block;
}
a {
  color: black;
}
a:hover {
  color: black;
  text-decoration: none;
  background-color: antiquewhite;
}
.middle {
  overflow: hidden;
  white-space: nowrap;
  text-overflow: ellipsis;
}
.leftalign {
  text-align: left;
}
</style>