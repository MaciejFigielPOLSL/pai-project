<template>
  <div class="container">
    <div class="row sized">
      <articles></articles>
    </div>
  </div>
</template>

<script>
import Articles from "./Articles";

export default {
  name: "MainPlace",
  components: {
    Articles,
  }
}
</script>

<style scoped>
.sized {
  padding-top: 50px;
  padding-bottom: 50px;
}
</style>