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
      <b-form-group id="interval-input-group" label="Name" label-for="interval-input">
        <b-form-input
          id="interval-input"
          v-model="form.name"
          placeholder="Input name here"
          required
        />
      </b-form-group>

      <b-form-group id="interval-input-group" label="Amount" label-for="interval-input">
        <b-form-spinbutton 
          v-model="form.amountValue"
          :formatter-fn="amountFormatter"
          min="0"
          max="9"
          wrap
        />
      </b-form-group>

      <b-form-group id="interval-input-group" label="Interval" label-for="interval-input">
        <b-form-spinbutton 
          v-model="form.intervalValue"
          :formatter-fn="intervalFormatter"
          min="0"
          max="6"
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
    name: 'AddIntervalPreset',
    data() {
      return {
        form: {
          name: '',
          amountValue: 0,
          amount:['5', '10', '15', '20', '25' , '30' , '35', '40', '45', '50'],
          intervalValue: 0,
          interval:['Never', 'Every hour', '2 hours', '3 hours', '4 hours' , '5 hours' , '6 hours'],
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
      intervalFormatter(intervalValue) {
        return this.form.interval[intervalValue]
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
