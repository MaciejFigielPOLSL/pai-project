<template>
  <div class="container-fluid p-3 mb-5 header">
    <div class="row">
      <div class="col-4">
        BLOG
        <div v-show="loggedIn">
          Zalogowano jako: {{ loggedUser.username }}
        </div>
      </div>

      <div class="col-6">
      </div>

      <div class="col-2" v-show="!loggedIn">
        <router-link :to="{name: 'register'}" tag="button" class="btn btn-secondary btn-sm">REJESTRUJ</router-link>
        <router-link :to="{name: 'login'}" tag="button" class="btn btn-secondary btn-sm">LOGIN</router-link>
      </div>
      <div class="col-2" v-show="loggedIn">
        <button tag="button" class="btn btn-secondary btn-sm" v-on:click="logout">LOGOUT</button>
      </div>
    </div>
    <br>
    <div class="row">
      <navigation-panel></navigation-panel>
    </div>
  </div>
</template>

<script>
import NavigationPanel from "./NavigationPanel";
import axios from "axios";
import Vue from "vue";
import {EventBus} from "@/main";

export default {
  name: "TopBar",
  props: [
    'loggedIn',
    'loggedUser'
  ],
  components: {
    NavigationPanel
  },
  methods: {
    logout() {
      axios
          .post('http://localhost:8080/api/logout', {}, {withCredentials: true})
            // eslint-disable-next-line no-unused-vars
          .then(response => {
            console.log(response);
            Vue.prototype.$mySession.loggedIn = false;
            // this.$router.push("/");
            EventBus.$emit('refresh');
          })
          // eslint-disable-next-line no-unused-vars
          .catch(err => {});
    }
  }
}
</script>

<style scoped>
.header {
  height: 100px;
  width: 100%;
}
</style>