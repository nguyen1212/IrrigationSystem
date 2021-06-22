import Vue from 'vue'
import App from './App.vue'
import VueRouter from 'vue-router'
import firebase from 'firebase'
import router from './router'

import { BootstrapVue, IconsPlugin} from 'bootstrap-vue'
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'

var firebaseConfig = {
  apiKey: "AIzaSyBBVqNFIYbFjh7hMxl0-lz2Tck8B0pZil8",
  authDomain: "irrigation-system-43da6.firebaseapp.com",
  projectId: "irrigation-system-43da6",
  storageBucket: "irrigation-system-43da6.appspot.com",
  messagingSenderId: "165584049106",
  appId: "1:165584049106:web:c1e4d16a5a6985be276e0b",
  measurementId: "G-E8PR0FFMVR"
};
firebase.initializeApp(firebaseConfig);

Vue.config.productionTip = false
Vue.use(BootstrapVue)
Vue.use(IconsPlugin)
Vue.use(VueRouter)

new Vue({
  router,
  render: h => h(App),
}).$mount('#app')
