<template>
 <div>
  <!-- <b-button id="show-btn" @click="$bvModal.show(myid)">Open Modal</b-button> -->

  <b-modal 
  :id="myid" 
  button-size="md"
  @show="resetModal"
  @hidden="resetModal"
  @ok="handleSubmit">
    <template #modal-title>
      Add new preset!
    </template>
    <div class="d-block text-center">
    <b-form v-if="show">
      <b-form-group id="input-group-2" label="Preset Name" label-for="input-2">
        <b-form-input
          id="input-2"
          v-model="form.name"
          placeholder="Enter name"
          required
        ></b-form-input>
      </b-form-group>

      <b-form-group id="input-group-2" label="Humidity" label-for="input-2">

        <b-form-spinbutton 
        id="demo-sb" 
        v-model="form.humidity"
        min="1" 
        max="100"></b-form-spinbutton>
      </b-form-group>

      <b-form-group id="input-group-2" label="Moisture" label-for="input-2">
        <b-form-spinbutton 
        id="demo-sb" 
        v-model="form.moisture"
        min="1" 
        max="100"></b-form-spinbutton>
      </b-form-group>

      <b-form-group id="input-group-2" label="Temperature" label-for="input-2">
        <b-form-spinbutton 
        id="demo-sb" 
        v-model="form.temperature"
        min="1" 
        max="100"></b-form-spinbutton>
      </b-form-group>
        <b-form-group id="input-group-2" label="Notes" label-for="input-2">
        <b-form-textarea
        id="input-2"
          v-model="form.name"
        placeholder="Default textarea"
      ></b-form-textarea>
      </b-form-group>
    </b-form>
    </div>
    <template #modal-footer="{ ok, cancel }">
      <!-- Emulate built in modal footer ok and cancel button actions -->
      <b-button variant="success" @click="ok()">
        Submit
      </b-button>
      <b-button variant="danger" @click="cancel()">
        Cancel
      </b-button>
    </template>
  </b-modal>
</div>
</template>

<script>
export default {
    name: 'AddModal',
    data() {
      return {
        form: {
          name: '',
          humidity: 0,
          moisture: 0,
          temperature: 0,
          notes: ""
        },
        showModal: false,
      }
    },
    props: {
      myid: String
    },
    methods: {
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
        this.form.humidity=0
        this.form.moisture=0
        this.form.temperature=0
        this.form.notes=""
        // this.show = false
        // this.$nextTick(() => {
        //   this.show = true
        // })
      },
      handleSubmit()
      {
        
      }
    }
}
</script>

<style scoped>

</style>
