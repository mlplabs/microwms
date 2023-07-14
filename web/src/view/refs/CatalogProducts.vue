<template>
  <div>
    <div id="detailForm" class="modal fade" tabindex="-1">
      <div class="modal-dialog modal-lg modal-dialog-centered">
        <product-form :prop-product-id="this.currentProductId" @onUpdateData="this.updateItemsOnPage(this.currentPage)"></product-form>
      </div>
    </div>
  </div>

  <div class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
    <h5>{{ lng.title }}</h5>
    <div class="btn-toolbar mb-2 mb-md-0">
      <div class="btn-group me-2">
        <button type="button" class="btn btn-sm btn-outline-secondary" data-bs-toggle="modal" data-bs-target="#detailForm"  @click="showDetailForm(0) ">
          <i class="bi bi-plus"></i>
        </button>
      </div>
    </div>
  </div>

  <div class="table-responsive">
    <table class="table table-striped table-hover table-bordered">
      <thead>
      <tr>
        <th scope="col" class="col_head col_id">#</th>
        <th scope="col" class="col_head">Наименование</th>
        <th scope="col" class="col_head">Артикул</th>
        <th scope="col" class="col_head">Производитель</th>
        <th scope="col" class="col_head col_action">...</th>
      </tr>
      </thead>
      <tbody>
      <tr v-for="(item, index) in tableData" :key="index">
        <td class="col_id">{{ item.id }}</td>
        <td><a href="#" class="text-decoration-none" data-bs-toggle="modal" data-bs-target="#detailForm" @click="showDetailForm(item.id)">{{ item.name }}</a></td>
        <td>{{ item.item_number }}</td>
        <td>{{ item.manufacturer.name }}</td>
        <td class="col_action">
          <div class="dropdown">
            <button type="button" class="btn btn-sm btn-outline-secondary" data-bs-toggle="dropdown">
              <i class="bi bi-three-dots-vertical"></i>
            </button>
            <ul class="dropdown-menu dropdown-menu-end">
              <li><a class="dropdown-item" href="#" @click.prevent="showReport('remaining', item.id)">Остатки товара: {{item.name}}</a></li>
              <li><a class="dropdown-item" href="#" @click.prevent="showReport('history', item.id)">Движения товара: {{item.name}}</a></li>
              <li><a class="dropdown-item" href="#" @click.prevent="printLabel(item.id)">Печать этикетки: {{item.name}}</a></li>
              <li><a class="dropdown-item" href="#" @click.prevent="deleteItem(item.id)">Удалить: {{item.name}}</a></li>
            </ul>
          </div>
        </td>
      </tr>
      </tbody>
    </table>
  </div>
  <pagination-bar v-bind:param-current-page="currentPage" v-bind:param-count-rows=countRows v-bind:param-limit-rows=limitRows v-on:selectPage="onSelectPage"></pagination-bar>
</template>

<script>
import DataProvider from "@/services/DataProvider";
import PaginationBar from "@/components/PaginationBar";
import ProductForm from "@/view/refs/forms/ProductForm";
import router from "@/router";

export default {
  name: "CatalogProducts",
  components: {PaginationBar, ProductForm},

  data(){
    return {
      refName: 'products',
      currentProductId: 0,
      tableData: [],
      countRows: 0,
      limitRows: 11,
      currentPage: 1,
      lng: {
        title: "Товары",
      },
    }
  },
  methods:{
    // Открываем форму нового или существующего
    showDetailForm(id){
      this.currentProductId = id
    },

    onSelectPage(eventData){
      this.currentPage = eventData.page
      this.updateItemsOnPage(eventData.page)
    },

    updateItemsOnPage(page){
      let offset = ( page -1 ) * this.limitRows
      DataProvider.GetItemsReference(this.refName, page, this.limitRows, offset)
        .then((response) => {
          this.tableData = response.data.data
          this.countRows = response.data.header.count
        })
        .catch(error => { DataProvider.ErrorProcessing(error) });
    },

    deleteItem(id){
      DataProvider.DeleteItemReference(this.refName, id)
        .then((response) => {
            const affRows = response.data;
            if (affRows !== 1){
              console.log('delete failed')
            }
            this.updateItemsOnPage(this.currentPage)
        })
        .catch(error => { DataProvider.ErrorProcessing(error) });
    },
    showReport(reportName, id){
      console.log(`show report for ${id}`)
      router.push({name: 'RemainingProducts', params: {propProductId: id}})
    },

    printLabel(id){
      DataProvider.PrintItemReference(this.refName, id)
        .then((response) => {
          console.log(response)
        })
        .catch(error => { DataProvider.ErrorProcessing(error) });
    },
  },
  mounted() {
    this.updateItemsOnPage(this.currentPage)
  }
}
</script>

<style scoped>
th.col_head{
  text-align: center;
}
th.col_id {
  width: 30px;
}
th.col_action {
  width: 20px;
}
td.col_id{
  text-align: right;
}
td.col_action{
  text-align: center;
}
</style>