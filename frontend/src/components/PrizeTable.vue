<template>
  <div v-if="showTables" class="prize-table">
    <vue-good-table :columns="columns" :rows="rows"/>
  </div>
</template>
<script>
import { VueGoodTable } from 'vue-good-table';
// import the styles
import 'vue-good-table/dist/vue-good-table.css'
import { bus } from '../main';

export default {
  components: {
    VueGoodTable,
  },
  name: "my-component",
  data() {
    return {
      rows: [
        {
          prize: "Сертификат на 1000 руб",
          count: 100,
          photo: "/1000.png",
        },
        {
          prize: "Сертификат на 4000 руб",
          count: 10,
          photo: "/4000.png",
        },
        {
          prize: "Смартфон",
          count: 5,
          photo: "/phone.png",
        },
        {
          prize: "Планшет",
          count: 3,
          photo: "/tablet.png",
        },
        {
          prize: "Суперприз телевизор (100 000)",
          count: 1,
          photo: "TV.png",
        },
      ].map(o=>({
        ...o,
        photo:`<img height="110" src="${o.photo}"/>`
      })),
      columns: [
        {
          label: 'Приз',
          field: 'prize',
        },
        {
          label: 'Количество',
          field: 'count',
        },
        {
          label: 'Изображение',
          field: 'photo',
          html: true,
          filterOptions: {
            enabled: false,
          },
        }
      ],
      options: {

      },
      showTables: false,
    }
  },
  created() {
    bus.$on('showTables', data => {
      this.showTables = data;
    })
  },
  methods: {
    async getData() {
    }
  }
}
</script>

<style>
.prize-table {
  width: 625px;
  position: fixed;
}
</style>