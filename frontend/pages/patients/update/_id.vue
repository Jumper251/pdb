<template>

  <main role="main">
    <div class="container mt-5">
      <h2>Patient Aktualisieren</h2>
      <hr/>
    </div>

    <div class="container">

      <b-card title="Stammdaten" align="">
        <b-card-text class="mt-5 ">
          <b-form>
          <b-form-row>
            <b-col>
              <b-form-group
                label="Vorname:"
                label-for="firstName"
                label-cols-sm="3"
                label-align-sm="right"
              >
                <b-form-input :state="firstValidation" v-model="patient.first_name" id="firstName"></b-form-input>

              </b-form-group>
            </b-col>
            <b-col>
              <b-form-group
                label="Nachname:"
                label-for="lastName"
                label-cols-sm="3"
                label-align-sm="right"
              >
                <b-form-input :state="lastValidation" v-model="patient.last_name" id="lastName"></b-form-input>

              </b-form-group>
            </b-col>
          </b-form-row>

          <b-form-row>
            <b-col>
              <b-form-group
                label="Email:"
                label-for="email"
                label-cols-sm="3"
                label-align-sm="right"
              >
                <b-form-input v-model="patient.email" id="email"></b-form-input>
              </b-form-group>
            </b-col>
            <b-col>
              <b-form-group
                label="Telefon:"
                label-for="phone"
                label-cols-sm="3"
                label-align-sm="right"
              >
                <b-form-input v-model="patient.phone" id="phone"></b-form-input>
              </b-form-group>
            </b-col>
          </b-form-row>

          <b-form-row>
            <b-col>
              <b-form-group
                label="Geburtstag:"
                label-for="birthday"
                label-cols-sm="3"
                label-align-sm="right"
              >
                <b-input-group class="mb-3">
                  <b-form-input
                    id="birthday"
                    v-model="patient.birthdate"
                    type="text"
                    placeholder="YYYY-MM-DD"
                    autocomplete="off"
                  ></b-form-input>
                  <b-input-group-append>
                    <b-form-datepicker
                      v-model="patient.birthdate"
                      button-only
                      right
                      locale="de"
                      aria-controls="birthday"
                    ></b-form-datepicker>
                  </b-input-group-append>
                </b-input-group>
              </b-form-group>
            </b-col>
            <b-col>
              <b-form-group
                label="Geschlecht:"
                label-for="sex"
                label-cols-sm="3"
                label-align-sm="right"
              >
                <b-form-select id="sex" v-model="patient.sex" :options="options"></b-form-select>
              </b-form-group>
            </b-col>
          </b-form-row>

          <h5 class="my-4">Adresse</h5>
          <b-form-row>
            <b-col>
              <b-form-group
                label="Straße:"
                label-for="street"
                label-cols-sm="3"
                label-align-sm="right"
              >
                <b-form-input v-model="patient.street" id="street"></b-form-input>
              </b-form-group>
            </b-col>
            <b-col>
              <b-form-group
                label="PLZ:"
                label-for="zipCode"
                label-cols-sm="3"
                label-align-sm="right"
              >
                <b-form-input v-model="patient.zip_code" id="zipCode"></b-form-input>
              </b-form-group>
            </b-col>

          </b-form-row>

          <b-form-row>
            <b-col>
              <b-form-group
                label="Ort:"
                label-for="city"
                label-cols-sm="3"
                label-align-sm="right"
              >
                <b-form-input v-model="patient.city" id="city"></b-form-input>
              </b-form-group>

            </b-col>
            <b-col>
              <b-form-group
                label="Land:"
                label-for="country"
                label-cols-sm="3"
                label-align-sm="right"
              >
                <b-form-input v-model="patient.country" id="country"></b-form-input>
              </b-form-group>
            </b-col>
          </b-form-row>

          <h5 class="my-4">Weiteres</h5>

          <b-form-row>
            <b-col>
              <b-form-group
                label="Diagnosen:"
                label-for="diagnosis"
              >
                <b-form-input v-model="patient.diagnosis" id="diagnosis"></b-form-input>
              </b-form-group>

            </b-col>
          </b-form-row>

          <b-form-row>
            <b-col>
              <b-form-group
                label="Verdachtsdiagnosen:"
                label-for="suspected_diagnosis"
              >
                <b-form-input v-model="patient.suspected_diagnosis" id="suspected_diagnosis"></b-form-input>
              </b-form-group>

            </b-col>
          </b-form-row>

          <b-form-row>
            <b-col>
              <b-form-group
                label="Medikamente:"
                label-for="medication"
              >
                <b-form-input v-model="patient.medication" id="medication"></b-form-input>
              </b-form-group>

            </b-col>
          </b-form-row>
          </b-form>

        </b-card-text>
        <hr class="mt-5"/>
        <b-button @click="addPatient" variant="primary">Aktualisieren</b-button>

      </b-card>
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
      patient: {
        first_name: '',
        last_name: '',
        street: '',
        zip_code: '',
        country: '',
        city: '',
        phone: '',
        email: '',
        sex: '',
        birthdate: '',
        diagnosis: '',
        suspected_diagnosis: '',
        medication: ''
      },
      options: [
        { value: 'Männlich', text: 'Männlich' },
        { value: 'Weiblich', text: 'Weiblich' },
        { value: 'Divers', text: 'Divers' }
      ],
    }
  },
  computed: {
    firstValidation() {
      return !!this.patient.first_name
    },
    lastValidation() {
      return !!this.patient.last_name
    }
  },
  mounted() {
    this.fetchPatient()
  },
  methods: {
    async fetchPatient() {
      const response = await this.$axios.get(`${config.baseUrl}/patients/${this.id}`)
      this.patient = response.data;
    },
    async addPatient() {
      if (!this.firstValidation || ! this.lastValidation) {
        return;
      }

      try {
        await this.$axios.put(`${config.baseUrl}/patients/${this.id}`, this.patient)
        this.$router.push(`/patients/${this.id}`)
      } catch (err) {
        console.log(err.response)
      }
    }
  }
}
</script>

