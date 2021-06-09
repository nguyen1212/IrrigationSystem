<template>
    <b-tabs card>
        <b-tab title="Repeat Mode" active>
            <div class="d-flex justify-content-center">
                <AddRepeatPreset myid="addRepeatContent"/>
                <b-container>
                    <b-row class=" pt-3 pb-3 border-bottom" >
                        <b-col class="col-1"><b>Name</b></b-col>
                        <b-col class="leftalign col-8 description"></b-col>
                        <b-col class="col-3"><b>Status</b></b-col>
                    </b-row>
                    <div :set="count = 0"/>
                    <b-row class="border-bottom" :key="o.id" v-for="o in this.repeatContent" :set="count=count+1">
                        <b-link class='pt-3 pb-3 row w-100 ml-0' v-b-toggle="'collapse-' + count" href="#">
                            <b-col class="leftalign col-4">{{o.name}}</b-col>
                            <b-col class="leftalign col-5 description"> Click to view detail</b-col>
                            <b-col class="pr-2 col-3">
                            <div v-if="repeatSelected === o.id">
                                <b-button variant="success">Select</b-button>
                            </div>
                            <div v-else>
                                <b-button variant="outline-primary">Select</b-button>
                            </div>
                            </b-col>
                        </b-link>
                        <b-collapse class='w-100' :id="'collapse-' + count">
                            <b-card class="border-0 ml-0">
                                <div class="row">
                                <div class="col-8">
                                    <ul>
                                    <li class="leftalign"><b>ID:</b> {{o.id}}</li>
                                    <li class="leftalign"><b>Name:</b> {{o.name}}</li>
                                    <li class="leftalign"><b>Time:</b> {{o.time}}</li>
                                    <li class="leftalign"><b>Amount:</b> {{o.amount}} liters</li>
                                    <li class="leftalign"><b>Repeat:</b> {{o.repeat}}</li>
                                    </ul>
                                </div>
                            </div>
                            <div class="row">
                                <b-link @click="$bvModal.show('addRepeatContent')">
                                    <b-button variant="dark">Edit</b-button>
                                </b-link>
                            </div>
                            </b-card>
                        </b-collapse>
                    </b-row>
                    <b-row class="border-bottom">
                        <b-link class='pt-3 pb-3 row w-100 ml-0' @click="$bvModal.show('addRepeatContent')">
                            <b-col class="col-1">
                                <b-icon class="h3 mt-0" icon="plus-circle-fill" style="color: black;" variant="black"></b-icon>
                            </b-col>

                            <b-col class="leftalign col-3"/>
                            <b-col class="leftalign col-5 description">Click to add new preset</b-col>
                        </b-link>
                    </b-row>
                </b-container>
            </div>
        </b-tab>

        <b-tab title="Interval Mode">
             <div class="d-flex justify-content-center">
                <AddIntervalPreset myid="addIntervalContent"/>
                <b-container>
                    <b-row class=" pt-3 pb-3 border-bottom" >
                        <b-col class="col-1"><b>Name</b></b-col>
                        <b-col class="leftalign col-8 description"></b-col>
                        <b-col class="col-3"><b>Status</b></b-col>
                    </b-row>
                    <div :set="count = 0"/>
                    <b-row class="border-bottom" :key="o.id" v-for="o in this.intervalContent" :set="count=count+1">
                        <b-link class='pt-3 pb-3 row w-100 ml-0' v-b-toggle="'collapse-' + count" href="#">
                            <b-col class="leftalign col-4">{{o.name}}</b-col>
                            <b-col class="leftalign col-5 description"> Click to view detail</b-col>
                            <b-col class="pr-2 col-3">
                            <div v-if="intervalSelected === o.id">
                                <b-button variant="success">Select</b-button>
                            </div>
                            <div v-else>
                                <b-button variant="outline-primary">Select</b-button>
                            </div>
                            </b-col>
                        </b-link>
                        <b-collapse class='w-100' :id="'collapse-' + count">
                            <b-card class="border-0 ml-0">
                                <div class="row">
                                <div class="col-8">
                                    <ul>
                                    <li class="leftalign"><b>ID:</b> {{o.id}}</li>
                                    <li class="leftalign"><b>Name:</b> {{o.name}}</li>
                                    <li class="leftalign"><b>Amount:</b> {{o.amount}} liters</li>
                                    <li class="leftalign"><b>Interval:</b> {{o.interval}}</li>
                                    </ul>
                                </div>
                            </div>
                            <div class="row">
                                <b-link @click="$bvModal.show('addIntervalContent')">
                                    <b-button variant="dark">Edit</b-button>
                                </b-link>
                            </div>
                            </b-card>
                        </b-collapse>
                    </b-row>
                    <b-row class="border-bottom">
                        <b-link class='pt-3 pb-3 row w-100 ml-0' @click="$bvModal.show('addIntervalContent')">
                            <b-col class="col-1">
                                <b-icon class="h3 mt-0" icon="plus-circle-fill" style="color: black;" variant="black"></b-icon>
                            </b-col>

                            <b-col class="leftalign col-3"/>
                            <b-col class="leftalign col-5 description">Click to add new preset</b-col>
                        </b-link>
                    </b-row>
                </b-container>
            </div>
        </b-tab>
    </b-tabs card> 
</template>

<script>
import AddRepeatPreset from "../timer mode/AddRepeatPreset"
import AddIntervalPreset from "../timer mode/AddIntervalPreset"
export default {
    data() {
        return {
            repeatContent: [],
            intervalContent: [],
            repeatSelected: 1,
            intervalSelected: 2
        }
    },
    created() {
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
            {
                "id": 3,
                "name": "Preset 4",
                "amount": 40,
                "interval": "5 hours",
            }
        ]
    },
    name: 'timerMode',
    methods: {
    },
    components: {
        AddRepeatPreset,
        AddIntervalPreset
    }
}
</script>

<style scoped>
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
</style>