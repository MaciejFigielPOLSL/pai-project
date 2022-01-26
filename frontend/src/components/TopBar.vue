<template>
  <div class="container-fluid p-3 mb-5 header">
    <div class="row">
      <div class="col-1">
        BLOG
      </div>
      <div class="col-9">

      </div>
      <div class="col-2">
        <button type="button" class="btn btn-secondary btn-sm">LOGIN</button>
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
export default {
  name: "TopBar",
  components: {NavigationPanel}
}
</script>

<style scoped>

.header {
  height: 100px;
  width: 100%;
}
</style>