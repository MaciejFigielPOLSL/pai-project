import Vue from 'vue'
import App from './App.vue'
import BootstrapVue from 'bootstrap-vue/dist/bootstrap-vue.esm';
import 'bootstrap-vue/dist/bootstrap-vue.css';
import 'bootstrap/dist/css/bootstrap.css';
import "./styles/skin_light.css";
import VueFullPage from 'vue-fullpage';

Vue.use(BootstrapVue);
Vue.use(VueFullPage);
Vue.config.productionTip = false;

Vue.component('main-place')
Vue.component('top-bar')
Vue.component('side-menu')
Vue.component('articles')
Vue.component('bottom')

new Vue({
  el: '#app',
  render: h => h(App)
}).$mount('#app')
