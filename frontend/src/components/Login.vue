<template>
  <div class="container-fluid">
    <div class="form-group">
      <label for="login">Login</label>
      <input v-model="login" type="email" class="form-control" id="login" aria-describedby="loginHelp" placeholder="Login">
    </div>
    <div class="form-group">
      <label for="password">Password</label>
      <input v-model="password" type="password" class="form-control" id="password" placeholder="Hasło">
    </div>
    <button v-on:click="performLogin" class="btn btn-primary">Zaloguj</button>

    <div v-show="init && loginStatus != '200'">
      <p>Błąd przy logowaniu</p>
    </div>
  </div>
</template>

<script>
import axios from "axios";
import Vue from "vue";
import {EventBus} from "@/main";
// import Vue from "vue";

export default {
  name: "Login",
  data() {
    return {
      login: "",
      password: "",
      loginStatus: false,
      init: false,
    }
  },
  methods: {
    performLogin() {
      var bodyFormData = new FormData();
      bodyFormData.append("username", this.login);
      bodyFormData.append("password", this.password);
      axios
          .post('http://localhost:8080/api/login', bodyFormData, {withCredentials: true})
          // eslint-disable-next-line no-unused-vars
          .then(response => {
              console.log(response)
              this.loginStatus = true;
              Vue.prototype.$mySession.username = this.login;
              Vue.prototype.$mySession.loggedIn = true;
              this.$router.push("/");
              // this.$root.$emit('refresh');
              EventBus.$emit('refresh');
            }
          )
          // eslint-disable-next-line no-unused-vars
          .catch(_ => {
              this.init = true;
              this.loginStatus = false;
            }
          )
    }
  }
}
</script>

<style scoped>

</style>