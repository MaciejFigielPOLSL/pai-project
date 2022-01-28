<template>
  <div id="app" class="container-fluid main">
    <top-bar v-on:refresh="refresh()" :loggedUser="loggedUser" :loggedIn="loggedUser.loggedIn"></top-bar>
    <main-place :loggedUser="loggedUser" :loggedIn="loggedUser.loggedIn"></main-place>
  </div>
</template>

<script>
import TopBar from "./components/TopBar";
import MainPlace from "./components/MainPlace";
import Vue from "vue";
import axios from "axios";
import {EventBus} from "@/main";

export default {
  name: 'App',
  components: {
    TopBar,
    MainPlace,
  },
  data() {
    return {
      loggedUser: {
        username: "",
        loggedIn: false
      },
      loggedIn: false
    }
  },
  mounted() {
    this.isLoggedIn();
    EventBus.$on('refresh', () => {
      this.refresh();
    })
  },
  methods: {
    refresh() {
      this.loggedUser = { username: Vue.prototype.$mySession.username, loggedIn: Vue.prototype.$mySession.loggedIn};
      this.loggedIn = Vue.prototype.$mySession.loggedIn;
    },
    isLoggedIn() {
      axios
          .get('http://localhost:8080/api/status', {withCredentials: true})
          .then(response => {
            if (response.data.logged) {
              Vue.prototype.$mySession.loggedIn = true;
            } else {
              Vue.prototype.$mySession.loggedIn = false;
            }
            this.loggedIn = Vue.prototype.$mySession.loggedIn;
          })
          // eslint-disable-next-line no-unused-vars
          .catch(err => console.log(err));
    },
  }
}

</script>

<style>
html, body {
  height: 100%;
}

.main {
  padding-top: 10px;
}
</style>
