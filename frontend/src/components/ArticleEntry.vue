<template>
  <div class="row p-3">
    <div class="container-fluid">
      {{ this.getDate(data.AddDate) }}
    </div>
    <div class="container-fluid">
      <u>Autor: {{ data.AuthorName }}</u>
    </div>
    <div class="container-fluid">
      <b>{{ data.Title }}</b>
    </div>
    <div class="container-fluid">
      <router-link :to="{ name: 'article', params: { id: data.ID }}">Czytaj więcej ></router-link>
    </div>
    <div class="container-fluid">
      <hr>
    </div>
  </div>
</template>

<script>
export default {
  name: "ArticleEntry",
  props: [
      'data'
  ],
  data: function () {
    return {
    }
  },
  methods: {
    getDate(date) {
      var d = new Date(date);
      return d.toDateString();
    }
  }
}
</script>

<style scoped>

</style>