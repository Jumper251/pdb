<template>

  <main role="main">
    <div class="container mt-5">
      <h2>Dokumentation - {{ fullName }}</h2>
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

      <hr />
      <div class="text-right">
        <b-button class="mr-3" :to="`/documentation/${id}/create`" variant="primary">Dokumentation erstellen</b-button>
        <b-button class="d-md-inline-block bv-d-sm-down-none" :href="`/export/${id}`"  target="_blank" variant="secondary">Exportieren</b-button>

        <hr />
      </div>
    </div>

    <div class="container">
      <b-table id="my-table" borderless :per-page="perPage" :current-page="currentPage" :no-provider-paging="true"
               :fields="fields" :items="documentations"
      >
        <template #cell(content)="data">
          <!-- `data.value` is the value after formatted by the Formatter -->

          <!--         <b-card no-body >
                     <b-card-header header-bg-variant="light" header-tag="header" class="p-1" role="tab">
                       <b-button class="text-left" block v-b-toggle="`content-${data.item.ID}`" variant="light">{{ `Am ${formatDate(data.item.time)} um ${formatTime(data.item.time)} Uhr`}}</b-button>
                     </b-card-header>
                     <b-collapse :id="`content-${data.item.ID}`" role="tabpanel">
                       <b-card-body>
                         <b-card-text v-html="data.item.content"></b-card-text>
                       </b-card-body>
                     </b-collapse>
                   </b-card> -->
          <b-overlay :show="showOverlay(data.item.ID)">
            <b-card :sub-title="`Am ${formatDate(data.item.time)} um ${formatTime(data.item.time)} Uhr`">
              <b-card-text class="mt-4">
                <span v-if="!isCollapsed(data.item.ID) && !showOverlay(data.item.ID)"
                      v-html="shortenContent(data.item.content).html"></span>
                <b-collapse :visible="isCollapsed(data.item.ID) && !showOverlay(data.item.ID)" id="collapse-3">
                  <span v-html="data.item.content"></span>
                </b-collapse>
              </b-card-text>

              <div class="text-right">
                <b-button @click="toggleCollapse(data.item.ID)" variant="light" v-if="isCollapsed(data.item.ID)"><span
                  class="mr-2">Weniger anzeigen</span> ▲
                </b-button>

                <b-button @click="toggleCollapse(data.item.ID)" variant="light"
                          v-else-if="shortenContent(data.item.content).more"><span class="mr-2">Mehr anzeigen</span> ▼
                </b-button>

              </div>
            </b-card>
          </b-overlay>
        </template>
      </b-table>

      <b-pagination
        v-model="currentPage"
        :total-rows="rows"
        :per-page="perPage"
        aria-controls="my-table"
        class="ml-3"
      ></b-pagination>

      <hr />
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
      fields: [
        {key: 'content', label: "Übersicht"},
      ],
      patient: {},
      documentations: [],
      collapsed: [],
      overlays: [],
      currentPage: 1,
      perPage: 5
    }
  },
  mounted() {
    this.fetchDocumentations()
    this.fetchPatient()
  },
  computed: {
    rows() {
      return this.documentations.length;
    },
    fullName() {
      return `${this.patient.first_name} ${this.patient.last_name}`;
    },
    isCollapsed() {
      return (id) => {
        return this.collapsed.includes(id)
      }
    },
    showOverlay() {
      return (id) => {
        return this.overlays.includes(id)
      }
    }
  },
  methods: {
    async fetchPatient() {
      const response = await this.$axios.get(`${config.baseUrl}/patients/${this.id}`)
      this.patient = response.data;
    },
    async fetchDocumentations() {
      const response = await this.$axios.get(`${config.baseUrl}/documentation/${this.id}/all`)
      this.documentations = response.data;
    },
    formatDate(time) {
      return new Date(time).toLocaleDateString([], {day: '2-digit', month: '2-digit', year: 'numeric'})
    },
    formatTime(time) {
      return new Date(time).toLocaleTimeString([], {hour: '2-digit', minute: "2-digit", timeZone: 'UTC'})
    },
    toggleCollapse(id) {
      this.overlays.push(id)

      if (this.isCollapsed(id)) {
        this.collapsed = this.collapsed.filter(fId => fId !== id);
      } else {
        this.collapsed.push(id)
      }

      setTimeout(function () {
        this.overlays = this.overlays.filter(fId => fId !== id);
      }.bind(this), 400)

    },
    shortenContent(html) {
      let options = {};

      let limit = 250,
        preserveTags = true,
        wordBreak = false,
        suffix = '...',
        moreLink = '',
        moreText = '»',
        preserveWhiteSpace = true;

      let arr = html.replace(/</g, "\n<")
        .replace(/>/g, ">\n")
        .replace(/\n\n/g, "\n")
        .replace(/^\n/g, "")
        .replace(/\n$/g, "")
        .split("\n");

      let sizeSum = 0;

      let sum = 0,
        row, cut, add, rowCut,
        tagMatch,
        tagName,
        tagStack = [],
        more = false;

      for (var i = 0; i < arr.length; i++) {

        row = arr[i];

        // count multiple spaces as one character
        if (!preserveWhiteSpace) {
          rowCut = row.replace(/[ ]+/g, ' ');
        } else {
          rowCut = row;
        }

        if (!row.length) {
          continue;
        }

        var charArr = this.getCharArr(rowCut);

        if (row === '<br>') {
          sum += 40
        }
        if (row === '<p>') {
          sum += 1
        }
        if (row === '</p>') {
          sum += 15
        }
        if (sum >= limit) {
          more = true
        }

        if (row[0] !== "<") {

          if (sum >= limit) {
            row = "";
          } else if ((sum + charArr.length) >= limit) {

            cut = limit - sum;

            if (charArr[cut - 1] === ' ') {
              while (cut) {
                cut -= 1;
                if (charArr[cut - 1] !== ' ') {
                  break;
                }
              }
            } else {

              add = charArr.slice(cut).indexOf(' ');

              // break on halh of word
              if (!wordBreak) {
                if (add !== -1) {
                  cut += add;
                } else {
                  cut = row.length;
                }
              }
            }

            row = charArr.slice(0, cut).join('') + suffix;

            sum = limit;
            more = true;
          } else {
            sum += charArr.length;
          }
        } else if (!preserveTags) {
          row = '';
        } else if (sum >= limit) {

          tagMatch = row.match(/[a-zA-Z]+/);
          tagName = tagMatch ? tagMatch[0] : '';

          if (tagName) {
            if (row.substring(0, 2) !== '</') {

              tagStack.push(tagName);
              row = '';
            } else {

              while (tagStack[tagStack.length - 1] !== tagName && tagStack.length) {
                tagStack.pop();
              }

              if (tagStack.length) {
                row = '';
              }

              tagStack.pop();
            }
          } else {
            row = '';
          }
        }

        arr[i] = row;
      }

      return {
        html: arr.join("\n").replace(/\n/g, ""),
        more: more
      }

    },
    getCharArr(rowCut) {

      let charArr = [],
        subRow,
        match,
        char;

      for (let i = 0; i < rowCut.length; i++) {

        subRow = rowCut.substring(i);
        match = subRow.match(/^&[a-z0-9#]+;/);

        if (match) {
          char = match[0];
          charArr.push(char);
          i += (char.length - 1);
        } else {
          charArr.push(rowCut[i]);
        }
      }

      return charArr;
    }
  }
}
</script>

