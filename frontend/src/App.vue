<template>
  <div id="app" class="container-fluid main">
    <top-bar></top-bar>
    <v-divider></v-divider>
    <main-place></main-place>
  </div>
</template>

<script>
import TopBar from "./components/TopBar";
import MainPlace from "./components/MainPlace";

export default {
  name: 'App',
  components: {
    TopBar,
    MainPlace,
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
