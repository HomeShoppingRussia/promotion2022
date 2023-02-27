<template>
  <div v-if="isDebug">
    <button v-if="isDebug" @click="fillTickets">Fill ticktes from CSV</button>
    <div v-if="fillDbStatusMsg">{{ fillDbStatusMsg }}</div>]
    <button v-if="isDebug" @click="fillPrizes">Fill prizes</button>
    <div v-if="fillPrizesStatusMsg">{{ fillPrizesStatusMsg }}</div>
  </div>
</template>

<script>
import axios from "axios";

export default {
  computed: {
    isDebug() {
      return this.$route.query.debug !== undefined
    }
  },
  data() {
    return {
      showModal: false,
      fillDbStatusMsg: '',
      fillPrizesStatusMsg: '',
    }
  },
  methods: {
    async fillTickets() {
      const headers = {
        'Content-Type': 'application/json',
        'Access-Control-Allow-Origin': '*',
        'Host': 'http://192.168.1.110:8182'
      }
      try {
        const response = await axios.get('http://192.168.1.110:8082/api/uploadData', headers);
        const data = await response.json()
        if (data.fillDbStatusMsg) {
          this.fillDbStatusMsg = 'Data successfully loaded into the database!'
        }
      } catch (error) {
        console.error(error)
        this.fillDbStatusMsg = 'Error loading data into the database.'
      }
    },
    async fillPrizes() {
      const headers = {
        'Content-Type': 'application/json',
        'Access-Control-Allow-Origin': '*',
        'Host': 'http://192.168.1.110:8182'
      }
      try {
        const response = await axios.get('http://192.168.1.110:8082/api/fillPrizesTable', headers);
        const data = await response.json()
        if (data.fillDbStatusMsg) {
          this.fillDbStatusMsg = 'Data successfully loaded into the database!'
        }
      } catch (error) {
        console.error(error)
        this.fillDbStatusMsg = 'Error loading data into the database.'
      }
    }
  }
}
</script>