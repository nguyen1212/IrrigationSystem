<template>
<div  class="d-flex justify-content-center">
   <b-container>
     <div  class="accordion" role="tablist">
      <b-row class=" pt-3 pb-3 border-bottom" >
        <b-col class="col-1"><b>#</b></b-col>
        <b-col class="leftalign col-3"><b>Name</b></b-col>
        <b-col class="leftalign col-5 description"><b>Description</b></b-col>
        <b-col class="col-3"><b>Status</b></b-col>
      </b-row>
      <!-- magic -->
      <div :set="count = 0"></div>
      <b-row class="border-bottom" role="tab" :key="o.id" v-for="o in JSON.parse(this.info)" :set="count=count+1">
        <b-link class='pt-3 pb-3 row w-100 ml-0' v-b-toggle="'collapse-' + count" href="#">
            <b-col class="col-1">
              {{count}}
            </b-col>
            <b-col class="leftalign col-3">{{o.name}}</b-col>
            <b-col class="leftalign col-5 description"><i style="color: gray">{{o.notes}}</i></b-col>
            <b-col class="pr-2 col-3">
              <div v-if="selected === o.id">
                <span class="active-dot" v-b-popover.hover.top="'Active'"></span>
              </div>
              <div v-else>
                <span class="inactive-dot" v-b-popover.hover.top="'Inactive'"></span>
              </div>
            </b-col>
        </b-link>
        <b-collapse class='w-100' accordion="my-accordion" :id="'collapse-' + count" role="tabpanel">
          <b-card class="border-0 ml-0">
            <div class="row">
              <div class="col-8">
              <ul>
              <li class="leftalign"><b>ID:</b> {{o.id}}</li>
              <li class="leftalign"><b>Name:</b> {{o.name}}</li>
              <li class="leftalign"><b>Moisture Level:</b> {{o.moistureThreshold}}</li>
              <li class="leftalign"><b>Note:</b> {{o.notes}}</li>
            </ul>
              </div>
          </div>
          <div class="w-100 row">
            <div class="ml-3"></div>
            <!-- <b-button v-if="selected !== o.id" class="col-1 ml-3" variant="primary" @click="$emit('select-preset', o.id)" >Select</b-button> -->
            <b-button class="col-1 ml-3" variant="warning" v-b-toggle="'collapse-inner-' + count" @click="resetEditForm">Edit</b-button>
            <b-button v-if="selected !== o.id" class="col-1 ml-3" variant="dark" @click="handleDeletePreset(o.id, o.name)" >Delete</b-button>
          </div>
          <div>
          <b-collapse class='w-100  mt-3' :id="'collapse-inner-' + count">
          <b-card class="ml-0">
            <b-form 
              @submit="handleUpdatePreset" 
              @reset="resetEditForm">
            <b-form-group style="width: 300px;">
              <label :for="'input-1-' + count" style="width: 300px;" class="leftalign"><b>Preset Name</b></label>
              <b-form-input
                :id="'input-1-' + count"
                :placeholder="o.name"
                v-model="editform.name"
                required
              >
              </b-form-input>
            </b-form-group>
            <b-form-group style="width: 300px;">
              <label :for="'input-2-' + count"  style="width: 300px;" class="leftalign"><b>Moisture</b></label>
              <b-form-spinbutton 
              :id="'input-2-' + count" 
              :value="o.moistureThreshold"
              v-model="editform.moisture"
              min="1" 
              max="100"></b-form-spinbutton>
            </b-form-group>

              <b-form-group>
              <label :for="'input-3-' + count" class="leftalign w-100"><b>Notes</b></label>
              <b-form-textarea
              :id="'input-3-' + count"
              :placeholder="o.notes"
              v-model="editform.notes"
            ></b-form-textarea>
            </b-form-group>
            <div class="leftalign w-100">           
            <b-button type="submit" variant="primary" class="mr-3">Submit</b-button>
            <b-button type="reset" variant="danger">Reset</b-button>
            </div>
            </b-form>
          </b-card>
        </b-collapse></div>
          </b-card>
        </b-collapse>
      </b-row>
      
      <b-row class="border-bottom"  role="tab">
        <b-link class='pt-3 pb-3 row w-100 ml-0' v-b-toggle.collapse_add href="#">
            <b-col class="col-1">
            <b-icon class="h3 mt-0" icon="plus-circle-fill" style="color: black;" variant="black"></b-icon>

            </b-col>
            <b-col class="leftalign col-3">Add</b-col>
            <b-col class="leftalign col-5 description"><i style="color: gray">Click to add</i></b-col>
            <b-col class="pr-2 col-3">
              
            </b-col>
        </b-link>
        <b-collapse class="w-100" id="collapse_add" accordion="my-accordion"  role="tabpanel">
          <b-card class="ml-0 mt-3 w-100">
            <b-form 
              @submit="handleAddPreset" 
              @reset="resetAddForm">
            <b-form-group style="width: 300px;">
              <label for="input-1" style="width: 300px;" class="leftalign"><b>Preset Name</b></label>
              <b-form-input
                id="input-1"
                v-model="addform.name"
                placeholder="Enter name"
                required
              ></b-form-input>
            </b-form-group>
            <b-form-group style="width: 300px;">
              <label for="input-2"  style="width: 300px;" class="leftalign"><b>Moisture</b></label>
              <b-form-spinbutton 
              id="input-2" 
              v-model="addform.moisture"
              min="1" 
              max="100"></b-form-spinbutton>
            </b-form-group>

            <b-form-group>
              <label :for="input-3" class="leftalign w-100"><b>Notes</b></label>
              <b-form-textarea
              id="input-3"
                v-model="addform.notes"
              placeholder="Default textarea"
            ></b-form-textarea>
            </b-form-group>
            <div class="leftalign w-100">           
            <b-button type="submit" variant="primary" class="mr-3">Submit</b-button>
            <b-button type="reset" variant="danger">Reset</b-button>
            </div>
            </b-form>
        </b-card>
        
      </b-collapse>
    </b-row>
    </div>
  </b-container>

</div>
</template>

<script>
export default {
  name: 'ListView',
  components: {
  },
  data() {
      // var obj = JSON.parse(this.info)
    return {
      obj: '',
      i: 0,
      addform : {
        name: '',
        moisture: 0,
        notes: ""
      },
      editform : {
        name: '',
        moisture: 0,
        notes: ""
      },
    }
  },
  props: {
    info: String,
    selected: Number
  },

  methods: {
      handleAddPreset(event)
      {
        event.preventDefault()
        this.$emit('add-preset', this.addform)
        // console.log(this.addform.name)
        // console.log(this.addform.moisture)
        // console.log(this.addform.notes)
        
      },            
      handleUpdatePreset(event)
      {
        event.preventDefault()
        this.$emit('update-preset', this.editform)
        // console.log(this.addform.name)
        // console.log(this.addform.moisture)
        // console.log(this.addform.notes)
        
      },   
      handleDeletePreset(id, name)
      {
        
        let cnf = window.confirm(`Delete ${name}?`)
        if (cnf)
        {
          // console.log("ok")
          this.$emit('delete-preset', id)
        }
        
      },   
      resetAddForm()
      {
        this.addform.name=""
        this.addform.moisture=0
        this.addform.notes=""
        // console.log("form reset")
      },
      resetEditForm()
      {
        this.editform.name=""
        this.editform.moisture=0
        this.editform.notes=""
        // console.log("form reset")
      },
  },
}
</script>

<style scoped>
.active-dot {
    height: 20px;
  width: 20px;
  background-color: #00BFA1;
  border-radius: 50%;
  display: inline-block;
}
.inactive-dot {
  height: 20px;
  width: 20px;
  background-color: rgb(255, 166, 0);
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
.description {
  overflow: hidden;
  white-space: nowrap;
  text-overflow: ellipsis;
}
.leftalign {
  text-align: left;
}
img {
  max-width: 200px;
  max-height: 200px;
}
</style>