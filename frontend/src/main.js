import Vue from 'vue'
import App from './App.vue'
import BootstrapVue from 'bootstrap-vue/dist/bootstrap-vue.esm';
import 'bootstrap-vue/dist/bootstrap-vue.css';
import 'bootstrap/dist/css/bootstrap.css';
import "./styles/skin_light.css";
import VueRouter from 'vue-router';

import ArticlePreview from './components/ArticlePreview'
import Articles from "@/components/Articles";
import Login from "@/components/Login";
import VueSession from 'vue-session';
import Register from "@/components/Register";
import AddArticle from "@/components/AddArticle";

Vue.use(VueSession);
Vue.use(VueRouter);
Vue.use(BootstrapVue);

Vue.prototype.$mySession = {
  username: "",
  loggedIn: false
}
Vue.config.productionTip = false;
//
// Vue.component('main-place')
// Vue.component('top-bar')
// Vue.component('side-menu')
// Vue.component('articles')
// Vue.component('bottom')

const router = new VueRouter({
  routes: [
    // dynamic segments start with a colon
    { path: '/', component: Articles, name: 'main' },
    { path: '/register', component: Register, name: 'register' },
    { path: '/add', component: AddArticle, name: 'add' },
    { path: '/login', component: Login, name: 'login' },
    { path: '/article/:id', component: ArticlePreview, name: 'article' }
  ]
});

export const EventBus = new Vue();

new Vue({
  el: '#app',
  render: h => h(App),
  router
}).$mount('#app')
