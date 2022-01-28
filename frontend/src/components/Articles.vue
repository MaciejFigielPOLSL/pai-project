<template>
  <div class="container shadow p-3 mb-5 bg-gradient">

    <div class="container-fluid" v-show="loggedUser.loggedIn">
      <router-link :to="{name: 'add'}" tag="button" class="btn btn-secondary btn-sm">DODAJ ARTYKUŁ</router-link>
    </div>

    <article-entry v-for="article in getArticlesForPage()" :key="article.ID" :data="article" d="article"></article-entry>
    <br>
    <br>

    <div class="d-flex justify-content-center">
      <div class="row">
        <nav aria-label="...">
          <ul class="pagination">
            <li class="page-item" :class="{ disabled: currentPage < 2}">
              <a class="page-link" @click="setPage(currentPage - 1)">Previous</a>
            </li>

            <li v-for="page in amountOfPages" :key="page" class="page-item" :class="{ active: page === currentPage}"><a class="page-link" @click="setPage(page)">{{ page }}</a></li>

            <li class="page-item" :class="{ disabled: currentPage >= amountOfPages}">
              <a class="page-link" @click="setPage(currentPage + 1)">Next</a>
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
      articles: null,
      currentPage: 1,
      amountOfPages: 0,
      amountPerPage: 10,
    }
  },
  mounted () {
    axios
        .get('http://localhost:8080/articles/')
        .then(response => {
          this.articles = response.data
          this.amountOfPages = Math.ceil(this.articles.length / this.amountPerPage);
        })
  },
  methods: {
    setPage(page) {
      this.currentPage = page;
    },
    getArticlesForPage() {
      return this.articles.slice(this.amountPerPage * (this.currentPage-1), this.amountPerPage * (this.currentPage));
    }
  }
}
</script>

<style scoped>

</style>