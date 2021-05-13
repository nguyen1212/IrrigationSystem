<template>
  <div style="width:100%;">
  <mdb-card class="card-body" style="width: 100%; ">
    <b-sidebar class="my-sidebar" id="sidebar-1" title="Menu" shadow visible=True bg-variant="dark" text-variant="light" >
      <template #header="{}">
       <div style="width: 100%;">
        <strong style="float:left;">Menu</strong>
        <b-button style="float:right;" size="sm" @click="hide"><b-icon icon="list" aria-hidden="true"></b-icon></b-button>
       </div>
      </template>
      <template>
          <nav class="mr-auto">
            <b-button-group vertical style="width: 100%;">
              <b-button variant="dark" v-on:click="selectFeature(home)">Home</b-button>
              <b-button variant="dark">Middle</b-button>
              <b-button variant="dark" v-on:click="selectFeature()">Bottom</b-button>
            </b-button-group>
          </nav>
      </template>
    </b-sidebar>
  </mdb-card>
  </div>
</template>
<script>
    import mdbCard from 'mdbvue';
    import home from '../HelloWorld.vue';
    export default{
        name: 'Menu',
        data(){return{
          home: home,
        }},
        components: mdbCard,
        methods: {
          selectFeature(feature){
              this.$emit('changefeature', feature)
          }
        }
    }
</script>
<style scoped>

.ml-0{
  width: 100%;
}

</style>