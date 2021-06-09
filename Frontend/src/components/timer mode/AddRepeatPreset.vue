<template>
 <div>
  <b-modal 
  :id="myid" 
  button-size="md"
  title="Add new preset"
  @show="resetModal"
  @hidden="resetModal"
  @ok="handleSubmit">
    <div class="d-block text-left">
    <b-form v-if="show">
      <b-form-group id="repeat-input-group" label="Name" label-for="repeat-input">
        <b-form-input
          id="repeat-input"
          v-model="form.name"
          placeholder="Input name here"
          required
        />
      </b-form-group>

      <b-form-group id="repeat-input-group" label="Time" label-for="repeat-input">
        <b-form-timepicker
        v-model="form.time" 
        required
        />  
      </b-form-group>

      <b-form-group id="repeat-input-group" label="Amount" label-for="repeat-input">
        <b-form-spinbutton 
          v-model="form.amountValue"
          :formatter-fn="amountFormatter"
          min="0"
          max="9"
          wrap
        />
      </b-form-group>

      <b-form-group id="repeat-input-group" label="Repeat" label-for="repeat-input">
        <b-form-spinbutton 
          v-model="form.repeatValue"
          :formatter-fn="repeatFormatter"
          min="0"
          max="7"
          wrap
        />
      </b-form-group>

    </b-form>
    </div>
    <template #modal-footer="{ ok, cancel }">
      <b-button variant="success" @click="ok()">
        Add
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
    name: 'AddRepeatPreset',
    data() {
      return {
        form: {
          name: '',
          time: '',
          amountValue: 0,
          amount:['5', '10', '15', '20', '25' , '30' , '35', '40', '45', '50'],
          repeatValue: 0,
          repeat:['Never', 'Everyday', '2 days', '3 days', '4 days' , '5 days' , '6 days', 'Every week'],
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
      repeatFormatter(repeatValue) {
        return this.form.repeat[repeatValue]
      },
      amountFormatter(amountValue) {
        return this.form.amount[amountValue]
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
        this.form.amountValue=0
        this.form.repeatValue=0
      },
      handleSubmit()
      {
        
      }
    }
}
</script>

<style scoped>
</style>
