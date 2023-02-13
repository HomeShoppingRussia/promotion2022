<template>
  <div class="wrapper" id="app">
    <PrizeTable/>
    <PreloaderStart/>
    <WinnersTable/>
  </div>
</template>

<script>
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap/dist/js/bootstrap.js'
import PrizeTable from './components/PrizeTable.vue'
import WinnersTable from './components/WinnersTable.vue'
import PreloaderStart from './components/PreloaderStart.vue'

export default {
  name: 'App',
  components: {
    PrizeTable,
    PreloaderStart,
    WinnersTable
  },
  props: ['filename'],
}
</script>

<style>
.wrapper {
  padding: 5px;
}
</style>