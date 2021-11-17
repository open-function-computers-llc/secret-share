import Vue from 'vue'
import App from './App.vue'
import axios from "axios"

Vue.config.productionTip = false
Vue.prototype.$api = axios.create({
  baseURL: "/api/"
})

new Vue({
  render: h => h(App),
}).$mount('#app')
