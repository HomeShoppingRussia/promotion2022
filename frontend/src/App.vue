<template>
  <div class="wrapper" id="app">
    <DebugPanel v-if="debugMode"/>
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
import DebugPanel from './components/DebugPanel.vue'

export default {
  name: 'App',
  components: {
    DebugPanel,
    PrizeTable,
    PreloaderStart,
    WinnersTable
  },
  props: ['filename'],
  computed: {
    debugMode() {
      return this.$route.query.debug !== undefined
    }
  }
}
</script>

<style>
.wrapper {
  padding: 5px;
}
</style>