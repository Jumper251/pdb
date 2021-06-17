<template>

  <main role="main">
    <div class="container mt-5">
      <h2>Neue Dokumentation: {{ fullName }}</h2>
      <hr/>

      <nav aria-label="breadcrumb">
        <ol class="breadcrumb">
          <li class="breadcrumb-item">
            <b-link to="/">Patienten</b-link>
          </li>
          <li class="breadcrumb-item">
            <b-link :to="`/patients/${id}`">{{ fullName }}</b-link>
          </li>
          <li class="breadcrumb-item active" aria-current="page">Dokumentation</li>

        </ol>
      </nav>
    </div>

    <div class="container">
      <b-card title="Dokumentation">

        <b-card-text class="mt-5 ">
          <quill-editor
            ref="editor"
            v-model="content"
            :options="editorOption"

          />
          <b-form-row class="mt-5">
            <b-col>
              <b-form-group
                label="Datum:"
                label-for="date"
              >
                <b-form-datepicker ref="datepicker" :state="validDate" v-model="date" id="date"></b-form-datepicker>
              </b-form-group>

            </b-col>

            <b-col md="6" sm="12">
              <b-form-group
                label="Zeit:"
                label-for="time"
              >
                <b-input-group>
                  <b-form-input
                    id="time"
                    v-model="time"
                    type="text"
                    placeholder="HH:mm"
                    :state="validTime"

                  ></b-form-input>
                  <b-input-group-append>
                    <b-form-timepicker
                      ref="timepicker"
                      v-model="time"
                      button-only
                      right
                      locale="de"
                      aria-controls="time"
                    ></b-form-timepicker>
                  </b-input-group-append>
                </b-input-group>
              </b-form-group>

            </b-col>
          </b-form-row>
        </b-card-text>
        <hr class="mt-5"/>

        <b-button @click="createDocumentation" variant="primary">Erstellen</b-button>
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
      patient: {},
      time: '',
      date: '',
      content: '',
      editorOption: {
        // Some Quill options...
        theme: 'snow',
        modules: {
          toolbar: [
            ['bold', 'italic', 'underline', 'strike'],
            [{'list': 'ordered'}, {'list': 'bullet'}, {'header': 1}, {'header': 2}],
            ['blockquote', 'link', 'image'],
            [{'color': []}],
          ]
        }
      }
    }
  },
  mounted() {

    this.fetchPatient()
  },
  computed: {
    fullName() {
      return `${this.patient.first_name} ${this.patient.last_name}`;
    },
    validTime() {
      return !!this.time
    },
    validDate() {
      return !!this.date
    }
  },
  methods: {
    async fetchPatient() {
      const response = await this.$axios.get(`${config.baseUrl}/patients/${this.id}`)
      this.patient = response.data;
    },
    async createDocumentation() {
      if (this.content.length <= 0 || !this.validTime || !this.validDate) return;

      const timeValue = this.$refs.timepicker.formattedValue
      const split = timeValue.split(":")
      const hours = parseInt(split[0]);
      const minutes = parseInt(split[1]);

      let date = new Date(this.$refs.datepicker.activeYMD)
      date.setHours(date.getHours() + hours)
      date.setMinutes(date.getMinutes() + minutes)

      try {
        const response = await this.$axios.post(`${config.baseUrl}/documentation/${this.id}`, {
          time: date.toISOString(),
          content: this.content,
          patient_id: parseInt(this.id)
        })

        this.$router.push(`/documentation/${this.id}/all`)
      } catch (err) {
        console.log(err.response)
      }

    }
  }
}
</script>

