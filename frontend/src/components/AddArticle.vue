<template>
  <div class="container-fluid">


    <div class="row m-2">
      <input v-model="articleTitle" placeholder="Tytuł">
      <textarea v-model="articleText" placeholder="Treść..."></textarea>
    </div>
    <div class="d-flex justify-content-center">
      <button type="button"  class="btn btn-secondary" v-on:click="sendArticle()">PRZEŚLIJ</button>
      <router-link :to="{name: 'main'}" tag="button" class="btn btn-secondary">ANULUJ</router-link>
    </div>
  </div>
</template>

<script>
import axios from "axios";
import {EventBus} from "@/main";

export default {
  name: "AddArticle",
  data() {
    return {
      articleTitle: "",
      articleText: "",
    }
  },
  methods: {
    sendArticle() {
      axios
          .post('http://localhost:8080/articles', {
            'title': this.articleTitle,
            'text': this.articleText,
          }, {withCredentials: true})
          // eslint-disable-next-line no-unused-vars
          .then(_ => {
            this.$router.push("/");
            EventBus.$emit('refresh');
          })
          .catch();
    }
  }
}
</script>

<style scoped>

</style>