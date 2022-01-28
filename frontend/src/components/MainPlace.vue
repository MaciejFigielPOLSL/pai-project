<template>
  <div class="container">
    <div class="row sized">
      <router-view :loggedUser="loggedUser"></router-view>
<!--      <router-view name="main"></router-view>-->
<!--      <router-view name="articles"></router-view>-->

<!--      <articles></articles>-->
    </div>
  </div>
</template>

<script>
// import Articles from "./Articles";

export default {
  name: "MainPlace",
  props: [
    'loggedIn',
    'loggedUser'
  ],
  components: {
    // Articles,
  }
}
</script>

<style scoped>
.sized {
  padding-top: 50px;
  padding-bottom: 50px;
}
</style>