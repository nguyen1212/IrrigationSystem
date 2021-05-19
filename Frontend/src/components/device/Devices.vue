<template>
  <div class="devices">
    <div class="row">
      <!-- <div class="col-8">
        <p> MAP </p>
      </div> -->
      <div class="col">
        <p> Statistic </p>
        <b-table striped hover :items="items"></b-table>
        <hr/>
        <p> Device Type </p>
      </div>
    </div>
  </div>
</template>
<script>
export default {
  name: 'devices',
  data() {
      return {
        items: [
          { age: 40, first_name: 'Dickerson', last_name: 'Macdonald' },
          { age: 21, first_name: 'Larsen', last_name: 'Shaw' },
          { age: 89, first_name: 'Geneva', last_name: 'Wilson' },
          { age: 38, first_name: 'Jami', last_name: 'Carney' }
        ]
      }
  }
}
</script>
<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.col-8{
  border-right: 1px solid black;
}
</style>