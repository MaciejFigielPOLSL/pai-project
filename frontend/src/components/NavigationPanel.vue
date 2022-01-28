<template>
  <nav class="navbar navbar-expand-lg navbar-light">
    <div class="container-fluid">
      <div class="collapse navbar-collapse" id="navbarScroll">
        <ul class="navbar-nav me-auto my-2 my-lg-0 navbar-nav-scroll" style="--bs-scroll-height: 100px;">
          <li class="nav-item">
            <router-link class="nav-link active" aria-current="page" :to="{name: 'main'}">Nowości</router-link>
          </li>
          <li class="nav-item">
            <a class="nav-link disabled">Najpopularniejsze</a>
          </li>
          <li class="nav-item">
            <a class="nav-link disabled">Archiwum</a>
          </li>
          <li class="nav-item">
            <a class="nav-link">Kontakt</a>
          </li>
        </ul>
        <form class="d-flex">
          <input class="form-control me-2" type="search" placeholder="Search" aria-label="Search">
          <button class="btn btn-outline-success" type="submit">Search</button>
        </form>
      </div>
    </div>
  </nav>
</template>

<script>
export default {
  name: "NavigationPanel"
}
</script>

<style scoped>

</style>