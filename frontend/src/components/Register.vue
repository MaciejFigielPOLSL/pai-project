<template>
  <div class="container-fluid">
    <div class="form-group">
      <label for="login">Login</label>
      <input v-model="login" type="email" class="form-control" id="login" aria-describedby="loginHelp" placeholder="Login">
    </div>
    <div class="form-group">
      <label for="email">Login</label>
      <input v-model="email" type="email" class="form-control" id="email" aria-describedby="loginHelp" placeholder="Email">
    </div>
    <div class="form-group">
      <label for="password">Password</label>
      <input v-model="password" type="password" class="form-control" id="password" placeholder="Hasło">
    </div>
    <button v-on:click="performRegister" class="btn btn-primary">Register</button>
  </div>
</template>

<script>
import axios from "axios";
import {EventBus} from "@/main";

export default {
  name: "Register",
  data() {
    return {
      login: "",
      email: "",
      password: "",
    }
  },
  methods: {
    performRegister() {
      var bodyFormData = new FormData();
      bodyFormData.append("username", this.login);
      bodyFormData.append("email", this.email);
      bodyFormData.append("password", this.password);
      axios
          .post('http://localhost:8080/api/register', bodyFormData, {withCredentials: true})
          // eslint-disable-next-line no-unused-vars
          .then(_ => {
                this.$router.push("/");
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