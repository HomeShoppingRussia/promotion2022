<template>
  <section>
    <div v-if="isError" class="alert alert-danger alert-dismissible fade show" role="alert">
      {{errorMsg}}
      <button type="button" class="btn-close" data-bs-dismiss="alert" aria-label="Close"></button>
    </div>
    <div v-if="showTables === false" class="btn-wrapper">
      <div class="btn-wrapper-inner">
        <button class="btn btn-danger w-50 h-50 mb-5" @click="getData()">Розыгрыш!</button>
      </div>
    </div>
    <div v-if="showTables === true" class="d-flex flex-column align-items-center wrapper-winners-table">
      <div v-if="showTables" class="winners-table">
        <vue-good-table :columns="columns" :rows="rows"/>
      </div>
    </div>
  </section>
</template>

<script>
import axios from 'axios';
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
      showTables: false,
      loading: false,
      isError: false,
      errorMsg: "",
      rows: [],
      rightItems: [],
      columns: [
        {
          label: 'Номер билета',
          field: 'ticketNumber',
        },
        {
          label: 'Имя',
          field: 'name',
        },
        {
          label: 'Отчество',
          field: 'middleName',
        },
        {
          label: 'Фамилия',
          field: 'surname',
        },
        {
          label: 'Телефон',
          field: 'phone',
        },
        {
          label: 'ШПИ',
          field: 'spi',
        },
        {
          label: 'Приз',
          field: 'prize',
        },
      ],
      options: {
        headings: {
          name: 'Country Name',
          code: 'Country Code',
          uri: 'View Record'
        },
        editableColumns:['name'],
        sortable: ['name', 'code'],
        filterable: ['name', 'code']
      }
    }
  },
  methods: {
    async getData() {
      try {
        this.loading = true;
        bus.$emit('loading',this.loading);
        const headers = {
          'Content-Type': 'application/json',
          'Access-Control-Allow-Origin': '*',
          'Host': 'http://localhost:8182'
        }
        const response = await axios.get('http://localhost:8082/api/getDraw', headers);
        if (response.data.status === false) {
          this.isError = true;
          bus.$emit('errorMsg',this.isError);
          this.errorMsg = response.data.error;
          bus.$emit('isError',this.errorMsg);
          this.loading = false;
          bus.$emit('loading',this.loading);
          debugger;
          return;
        }
        this.rows = response.data.list.map(item => {
          item.surname = item.surname.substring(0, 1);
          return item;
        });
        this.rightItems = response.data.rightItems;
        this.showTables = true;
        bus.$emit('showTables',this.showTables);
        this.loading = false;
        bus.$emit('loading',this.loading);
      } catch (error) {
        console.log(error);
      }
    }
  }
}
</script>

<style>
.btn-wrapper {
  display: flex;
  height: 100vh;
  width: 100%;
  align-items: center;
  justify-content: center;
}
.btn-wrapper-inner {
  height: 170px;
  width: 100%;
  justify-content: center;
  display: flex;
}
.btn-danger {
  font-size: 2rem;
}
.wrapper-winners-table {
  width: 100%;
}
.winners-table {
  justify-content: right;
  display: flex;
  width: 100%;
}
.table.vgt-table td {
  max-width: 194px;
}
.vgt-wrap {
  position: relative;
  max-width: 1030px;
}
</style>