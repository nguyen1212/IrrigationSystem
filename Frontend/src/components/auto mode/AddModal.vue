<template>
 <div>
<b-modal v-model="showModal" :ref="customref" title="Add new preset" hide-footer>
    <div>
    <b-form 
    @submit="handleAddPreset" 
    @reset="resetForm">
      <b-form-group id="input-group-1" label="Preset Name" label-for="input-1">
        <b-form-input
          id="input-1"
          v-model="form.name"
          placeholder="Enter name"
          required
        ></b-form-input>
      </b-form-group>
      <b-form-group id="input-group-2" label="Moisture" label-for="input-2">
        <b-form-spinbutton 
        id="input-2" 
        v-model="form.moisture"
        min="1" 
        max="100"></b-form-spinbutton>
      </b-form-group>

        <b-form-group id="input-group-3" label="Notes" label-for="input-3">
        <b-form-textarea
        id="input-3"
          v-model="form.note"
        placeholder="Default textarea"
      ></b-form-textarea>
      </b-form-group>
      <b-button type="submit" variant="primary" class="mr-3">Submit</b-button>
      <b-button type="reset" variant="danger">Reset</b-button>
    </b-form></div>
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
          moisture: 0,
          notes: ""
        },
        showModal: false,
      }
    },
    props: {
      customref: String
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
        this.form.moisture=0
        this.form.notes=""
      },
      handleSubmit()
      {
        console.log(this.form)
        this.$emit('add-preset', this.form)
      }
    }
}
</script>

<style scoped>

</style>
