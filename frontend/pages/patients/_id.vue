<template>

  <main role="main">
    <div class="container mt-5">
      <h2>{{ fullName }}</h2>
      <hr/>

      <nav aria-label="breadcrumb">
        <ol class="breadcrumb">
          <li class="breadcrumb-item"><b-link to="/">Patienten</b-link></li>
          <li class="breadcrumb-item active" aria-current="page">{{ fullName }}</li>
        </ol>
      </nav>
    </div>

    <div class="container">
      <b-card title="Stammdaten">

        <b-card-text class="mt-5 ">
          <b-form>
            <b-row>
              <b-col>
                <p><span class="text-muted">Vorname: </span>{{ patient.first_name }}</p>
              </b-col>
              <b-col>
                <p><span class="text-muted">Nachname: </span>{{ patient.last_name }}</p>
              </b-col>
            </b-row>

            <b-row>
              <b-col>
                <p><span class="text-muted">Email: </span>
                  <b-link :href="`mailto:${patient.email}`"> {{ patient.email }}</b-link>
                </p>
              </b-col>
              <b-col>
                <p><span class="text-muted">Telefon: </span>{{ patient.phone }}</p>

              </b-col>
            </b-row>

            <b-row>
              <b-col>
                <p><span class="text-muted">Geburtstag: </span>{{ patient.birthdate }}</p>

              </b-col>
              <b-col>
                <p><span class="text-muted">Geschlecht: </span>{{ patient.sex }}</p>
              </b-col>
            </b-row>

            <h5 class="my-4">Adresse</h5>
            <b-row>
              <b-col>
                <p><span class="text-muted">Straße: </span>{{ patient.street }}</p>
              </b-col>
              <b-col>
                <p><span class="text-muted">PLZ: </span>{{ patient.zip_code }}</p>
              </b-col>

            </b-row>

            <b-row>
              <b-col>
                <p><span class="text-muted">Ort: </span>{{ patient.city }}</p>
              </b-col>
              <b-col>
                <p><span class="text-muted">Land: </span>{{ patient.country }}</p>
              </b-col>
            </b-row>

            <h5 class="my-4">Weiteres</h5>

            <b-row>
              <b-col>
                <p><span class="text-muted">Diagnosen: </span>
                  <br />
                  {{ patient.diagnosis }}
                </p>
              </b-col>
            </b-row>

            <b-form-row>
              <b-col>
                <p><span class="text-muted">Verdachtsdiagnosen: </span>
                  <br />
                  {{ patient.suspected_diagnosis }}
                </p>
              </b-col>
            </b-form-row>

            <b-row>
              <b-col>
                <p>
                  <span class="text-muted">Medikamente: </span>
                  <br />
                  {{ patient.medication }}
                </p>
              </b-col>
            </b-row>
          </b-form>
          <div class="text-center">
    <b-button   variant="outline-dark" class="mt-4">Zur Dokumentation &raquo;</b-button>
          </div>
        </b-card-text>
        <hr class="mt-5"/>
        <b-row>
          <b-col>
            <b-button :to="`/patients/update/${id}`" variant="primary">Aktualisieren</b-button>
          </b-col>
          <b-col class="text-right">
            <b-button @click="showModal" class="" variant="danger">Löschen</b-button>
          </b-col>
        </b-row>
      </b-card>
      <b-modal ref="my-modal" hide-footer title="Patient löschen?">
        <b-row>
          <b-col>
        <b-button  variant="danger" block @click="deletePatient">Löschen!</b-button>
          </b-col>
          <b-col>
        <b-button variant="secondary" block @click="hideModal">Abbrechen</b-button>
          </b-col>
        </b-row>
      </b-modal>
    </div> <!-- /container -->

  </main>
</template>

<script>

import config from '@/config';

export default {
  async asyncData({params}) {
    const id = params.id
    return {id}
  },
  data() {
    return {
      patient: {}
    }
  },
  mounted() {
    this.fetchPatient()
  },
  computed: {
    fullName() {
      return `${this.patient.first_name} ${this.patient.last_name}`;
    }
  },
  methods: {
    showModal() {
      this.$refs['my-modal'].show()
    },
    hideModal() {
      this.$refs['my-modal'].hide()
    },
    async fetchPatient() {
      const response = await this.$axios.get(`${config.baseUrl}/api/patients/${this.id}`)
      this.patient = response.data;
    },
    async deletePatient() {
      await this.$axios.delete(`${config.baseUrl}/api/patients/${this.id}`);

      this.hideModal()
      this.$router.push("/")
    }
  }
}
</script>

