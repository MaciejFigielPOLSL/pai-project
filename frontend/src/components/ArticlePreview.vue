<template>
  <div class="container-fluid shadow-lg p-3">
<!--    {{ article }}-->
    <div class="container-fluid">
      <router-link :to="{ name: 'main'}"> Back </router-link>
    </div>

    <br>

    <div class="container-fluid">
      {{ this.getDate(article.AddDate) }}
    </div>

    <div class="container-fluid">
      <u>Autor: {{ article.AuthorName }}</u>
    </div>

    <br>

    <div class="container-fluid">
      <b>{{ article.Title }}</b>
    </div>

    <hr>

    <div class="container-fluid">
      {{ article.Text }}
    </div>

    <hr>

    <div class="container-fluid shadow-sm p-1 rounded-2" v-on:mouseenter="commentFocus = true" v-on:mouseleave="commentFocus=false">
      <div class="row m-2">
        <u>Dodaj komentarz:</u>
      </div>
      <div v-show="commentFocus" class="row m-2">
        <textarea v-model="newCommentText" placeholder="Komentarz..."></textarea>
        <input v-model="commentAuthor" placeholder="Nick">
        <button type="button"  class="btn btn-secondary" v-on:click="sendComment()">Prześlij</button>
      </div>
    </div>

    <br>

    <div class="container-fluid px-5 pb-4">
      <u>Komentarze:</u>
    </div>
    <div v-if="comments.length > 0" class="container-fluid">
      <comment v-for="comment in comments" :key="comment.ID" :comment="comment"></comment>
    </div>
    <br>
    <div v-if="comments.length <= 0" class="container-fluid">
      Brak komentarzy
    </div>
  </div>
</template>

<script>
import axios from "axios";
import Comment from "./Comment";

export default {
  name: "ArticlePreview",
  components: {
    Comment
  },
  mounted() {
    var id = this.$route.params['id'];
    axios
        .get('http://localhost:8080/articles/' + id.toString())
        .then(response => {
          this.article = response.data.article;
          this.comments = response.data.comments;
        })
  },
  data () {
    return {
      article: {},
      comments: [],
      newCommentText: "",
      commentAuthor: "",
      commentFocus: false,
    }
  },
  methods: {
    refresh() {
      axios
          .get('http://localhost:8080/articles/' + this.article.ID)
          .then(response => {
            this.article = response.data.article;
            this.comments = response.data.comments;
          })
    },
    getDate(date) {
      var d = new Date(date);
      return d.toDateString();
    },
    sendComment() {
      axios.post('http://localhost:8080/comment', {
        'articleId': this.article.ID,
        'authorName': this.commentAuthor === "" ? "anonymous" : this.commentAuthor,
        'text': this.newCommentText,
        // eslint-disable-next-line no-unused-vars
      }).then(_ => this.refresh());
    }
  }
}
</script>

<style scoped>

</style>