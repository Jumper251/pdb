<template>
  <b-navbar toggleable="lg" type="dark"  class="navbar navbar-expand-md navbar-dark fixed-top bg-dark">
    <b-navbar-brand to="/">PDB</b-navbar-brand>
    <b-navbar-toggle target="nav-collapse"></b-navbar-toggle>


    <b-collapse id="nav-collapse" is-nav>
      <ul class="navbar-nav mr-auto">
        <li class="nav-item active">
          <b-link class="nav-link" to="/">Patienten <span class="sr-only">(current)</span></b-link>
        </li>


      </ul>
      <form class="form-inline my-2 my-lg-0">
        <input v-model="searchQuery" class="form-control mr-sm-2" type="text" placeholder="Suche" aria-label="Search">
        <button @click.prevent="search" class="btn btn-outline-warning my-2 my-sm-0">Suchen</button>
      </form>
    </b-collapse>
  </b-navbar>
</template>

<script>
export default {
  name: "Navbar",
  data() {
    return {
      searchQuery: ''
    }
  },
  methods: {
    search() {
      this.$router.push({path: '/', query: {query: this.searchQuery}, replace: true})

    }
  }
}
</script>

