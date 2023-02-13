<template>
  <loader v-if="loading" object="#ff0000" color1="#ffffff" color2="#17fd3d" size="5" speed="2" bg="#343a40" objectbg="#999793" opacity="80" disableScrolling="false" name="spinning"></loader>
</template>

<script>
import {bus} from "@/main";

export default {
  name: "PreloaderStart",
  data(){
    return {
      loading: false
    }
  },
  created() {
    bus.$on('loading', data => {
      this.loading = data;
    })
  },
}
</script>

<style scoped>

</style>