<template>
  <div>
    <div id="detailForm" class="modal fade" tabindex="-1">
      <div class="modal-dialog modal-lg modal-dialog-centered">
        <product-movement-form :prop-mode-form="this.modeForm" @onUpdateData="this.updateItemsOnPage(this.currentPage)"></product-movement-form>
      </div>
    </div>
  </div>

  <div class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
    <h5>{{ lng.title }}</h5>
    <div class="btn-toolbar mb-2 mb-md-0">
      <div class="btn-group me-2">
        <button type="button" class="btn btn-sm btn-outline-secondary" title="Взять" data-bs-toggle="modal" data-bs-target="#detailForm"  @click="showDetailForm(0) ">
          <i class="bi bi-cart-dash-fill"></i>
        </button>
        <button type="button" class="btn btn-sm btn-outline-secondary" title="Положить" data-bs-toggle="modal" data-bs-target="#detailForm"  @click="showDetailForm(1) ">
          <i class="bi bi-cart-plus-fill"></i>
        </button>
        <button type="button" class="btn btn-sm btn-outline-secondary" title="Переместить" data-bs-toggle="modal" data-bs-target="#detailForm"  @click="showDetailForm(2) ">
          <i class="bi bi-cart-check-fill"></i>
        </button>
      </div>
    </div>
  </div>

<div id="reportData" class="table-responsive">
<table class="table table-striped table-hover table-bordered">
  <thead>
  <tr>
    <th scope="col" class="col_head">Товар</th>
    <th scope="col" class="col_head">Зона</th>
    <th scope="col" class="col_head">Ячейка</th>
    <th scope="col" class="col_head">Остаток</th>
    <th scope="col" class="col_head col_action">...</th>
  </tr>
  </thead>
  <tbody>
  <tr v-for="(item, index) in stockData" :key="index">

    <td><a href="#" class="text-decoration-none" data-bs-toggle="modal" data-bs-target="#detailForm" @click="showDetailForm(item.product.id)">{{item.product.name}}</a></td>
    <td>{{ item.zone.name }}</td>
    <td>{{ item.cells[0].Name }}</td>
    <td style="text-align: right">{{item.quantity }}</td>
    <td class="col_action">
      <div class="dropdown">
        <button type="button" class="btn btn-sm btn-outline-secondary" data-bs-toggle="dropdown">
          <i class="bi bi-three-dots-vertical"></i>
        </button>
        <ul class="dropdown-menu dropdown-menu-end">
          <li><a class="dropdown-item" href="#" @click.prevent="showReport('history', item.product.id)">Движения товара: {{item.product.name}}</a></li>
        </ul>
      </div>
    </td>

  </tr>
  </tbody>
</table>
</div>

</template>

<script>

import DataProvider from "@/services/DataProvider";
import router from "@/router";
import ProductMovementForm from "@/view/reports/forms/ProductMovementForm.vue";

export default {
  name: "StockProducts",
  components: {ProductMovementForm},
  data(){
    return{
      modeForm: 0,
      currentProductId:0,
      condition:{
        id: this.propProductId,
        text: ""
      },
      conditionSuggestion:[],
      lng: {
        title: "Остатки товаров",
      },
      stockData: []
    }
  },
  methods: {
    // Действие Сформировать отчет
    generateReport(){
      this.requestStockData()
    },
    // Процедура запроса данных
    requestStockData(){
      DataProvider.GetReport('stock', this.condition.id)
          .then((response) => {
            this.stockData = response.data.data.rows
          })
          .catch(error => { DataProvider.ErrorProcessing(error) });
    },
    showReport(reportName, id){
      // TODO: посмотреть друой отчет
      console.log(`show report for ${id}`)
      router.push(`/reports/${reportName}`)
    },
    showDetailForm(id){
      this.modeForm = id
    },
  },
  mounted() {
    this.generateReport()
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