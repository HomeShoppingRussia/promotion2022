<template>
  <div v-if="isDebug">
    <button v-if="isDebug" @click="fillDB">Fill DB from CSV</button>
    <div v-if="message">{{ message }}</div>
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
      message: ''
    }
  },
  methods: {
    async fillDB() {
      const headers = {
        'Content-Type': 'application/json',
        'Access-Control-Allow-Origin': '*',
        'Host': 'http://localhost:8182'
      }
      try {
        const response = await axios.get('http://localhost:8082/api/uploadData', headers);
        const data = await response.json()
        if (data.message) {
          this.message = 'Data successfully loaded into the database!'
        }
      } catch (error) {
        console.error(error)
        this.message = 'Error loading data into the database.'
      }
    }
  }
}
</script>