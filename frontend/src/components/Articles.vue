<template>
  <div class="container shadow p-3 mb-5 bg-gradient">

    <div class="container-fluid" v-show="loggedUser.loggedIn">
      <router-link :to="{name: 'add'}" tag="button" class="btn btn-secondary btn-sm">DODAJ ARTYKUŁ</router-link>
    </div>

    <article-entry v-for="article in articles" :key="article.ID" :data="article" d="article"></article-entry>
    <br>
    <br>

    <div class="d-flex justify-content-center">
      <div class="row">
        <nav aria-label="...">
          <ul class="pagination">
            <li class="page-item disabled">
              <a class="page-link">Previous</a>
            </li>
            <li class="page-item"><a class="page-link" href="#">1</a></li>
            <li class="page-item active" aria-current="page">
              <a class="page-link" href="#">2</a>
            </li>
            <li class="page-item"><a class="page-link" href="#">3</a></li>
            <li class="page-item">
              <a class="page-link" href="#">Next</a>
            </li>
          </ul>
        </nav>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";
import ArticleEntry from "@/components/ArticleEntry";

export default {
  name: "Articles",
  props: [
    'loggedUser'
  ],
  components: {
    ArticleEntry
  },

  data () {
    return {
      articles: null
    }
  },

  mounted () {
    axios
        .get('http://localhost:8080/articles/')
        .then(response => this.articles = response.data)
  }
}
</script>

<style scoped>

</style>