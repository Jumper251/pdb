<template>

  <main role="main">
    <div class="container mt-5">
      <h2>Patienten</h2>
      <hr/>
    </div>

    <div class="container">
      <div class="mt-3 mb-3 text-right">
        <b-button to="/add" variant="primary">Patient hinzufügen</b-button>
      </div>

      <b-table :per-page="perPage" :current-page="currentPage" :no-provider-paging="true"
               striped sort-icon-left hover :fields="fields" :items="patients"
      >

        <template #cell(last_name)="data">
          <b-link :to="`/patients/${data.item.ID}`">{{data.value}}</b-link>
        </template>



        <template #cell(ID)="data">
          <!-- `data.value` is the value after formatted by the Formatter -->
       <b-link :to="`/documentation/${data.value}/all`">Dokumentation</b-link>
        </template>
      </b-table>

      <b-pagination
        v-model="currentPage"
        :total-rows="rows"
        :per-page="perPage"
        aria-controls="my-table"
      ></b-pagination>

      <hr>
    </div> <!-- /container -->

  </main>
</template>

<script>

import config from '@/config';

export default {
  data() {
    return {
      patients: [],
      fields: [
        {key: 'last_name', label: "Voller Name", sortable: true, formatter: "fullName"},
        {key: 'ID', label: "Links" }
      ],
      currentPage: 1,
      perPage: 20
    }
  },
  computed: {
    rows() {
      return this.patients.length;
    }
  },
  mounted() {
    this.fetchPatients()

  },
  watch: {
    $route(to, from) {
      if (to.path === '/') {
        this.fetchPatients()
      }
    },
  },
  methods: {
    fullName(value, key, item) {
      return `${item.last_name}, ${item.first_name}`
    },
    async fetchPatients(firstname, lastname) {
      let url = `${config.baseUrl}/patients`

      const query = this.$route.query.query
      if (query) {
        url += `?query=${query}`
      }

      const response = await this.$axios.get(url);
      this.patients = response.data;
    }
  }
}
</script>

