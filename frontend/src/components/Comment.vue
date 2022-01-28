<template>
  <div class="row">
    <div class="col-1">
    </div>
    <div class="col-10">
      <div class="container-fluid">
        {{ this.getDate(comment.CreatedAt) }}
      </div>
      <div class="container-fluid">
        <u>Autor: {{ comment.AuthorName }}</u>
      </div>
      <br>
      <div class="container-fluid">
        {{ comment.Text }}
      </div>
      <div class="container-fluid">
        <div class="row">
          <div class="col-9">
          </div>
          <div class="col-1">
            <div @click="addLike">
              <img src="@/assets/like.png"/>
            </div>
          </div>
          <div class="col-1">
            <div @click="addDislike">
              <img src="@/assets/dislike.png"/>
            </div>
          </div>
        </div>
        <div class="row">
          <div class="col-9">
          </div>
          <div class="col-1">
            {{ currentLikes }}
          </div>
          <div class="col-1">
            {{ currentDislikes }}
          </div>
        </div>
      </div>
      <div class="container-fluid">
        <hr>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "Comment",
  props: [
      'comment'
  ],
  data() {
    return {
      likeGiven: false,
      currentLikes: 0,
      currentDislikes: 0,
    }
  },
  mounted() {
    this.currentLikes = this.comment.Likes;
    this.currentDislikes = this.comment.Dislikes;
  },
  methods: {
    getDate(date) {
      var d = new Date(date);
      return d.toDateString();
    },
    addLike() {
      if (this.likeGiven) { return; }
      this.currentLikes++;
      axios.post('http://localhost:8080/comment/like/' + this.comment.ID);

      this.likeGiven = true;
    },
    addDislike() {
      if (this.likeGiven) { return; }
      this.currentDislikes++;
      axios.post('http://localhost:8080/comment/dislike/' + this.comment.ID);

      this.likeGiven = true;
    }
  }
}
</script>

<style scoped>

</style>