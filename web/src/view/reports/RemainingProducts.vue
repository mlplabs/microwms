<template>
  <div id="detailForm" class="modal fade" tabindex="-1">
    <div class="modal-dialog modal-lg modal-dialog-centered">
      <product-form :prop-product-id="this.currentProductId"  @onUpdateData="this.generateReport"></product-form>
    </div>
  </div>

  <div class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
    <h5>{{ lng.title }}</h5>
  </div>
  <form class="row g-3 mb-3">
    <div class="col">
      <div class="row">
        <label for="inputName" class="col-sm-2 col-form-label">Товар: </label>
        <div class="col-sm-10">
          <autocomplete-input
            prop-input-id="inputName"
            v-model:prop-suggestions="conditionSuggestion"
            v-model:prop-selection-id="condition.id"
            v-model:prop-selection-val="condition.text"
            @onUpdateData="updateProductsData">
          </autocomplete-input>
        </div>
      </div>
    </div>
    <div class="col-auto">
      <button type="submit" class="btn btn-md btn-outline-secondary" @click.prevent="generateReport">Сформировать</button>
    </div>
  </form>

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
      <tr v-for="(item, index) in remainingData" :key="index">

        <td><a href="#" class="text-decoration-none" data-bs-toggle="modal" data-bs-target="#detailForm" @click="showDetailForm(item.product.id)">{{item.product.name}}<br><small>{{item.manufacturer.name}}</small></a></td>
        <td>{{item.zone.name}}</td>
        <td>{{item.cell.name}}</td>
        <td style="text-align: right">{{item.quantity}}</td>
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
import AutocompleteInput from "@/components/AutocompleteInput";
import DataProvider from "@/services/DataProvider";
import ProductForm from "@/view/refs/forms/ProductForm"
import router from "@/router";

export default {
  name: "RemainingProducts",
  components: {ProductForm, AutocompleteInput},
  data(){
    return{
      currentProductId:0,
      condition:{
        id: 0,
        text: ""
      },
      conditionSuggestion:[],
      lng: {
        title: "Остатки товаров",
      },
      remainingData: []
    }
  },
  methods: {
    // Открываем форму нового или существующего
    showDetailForm(id){
      this.currentProductId = id
    },
    generateReport(){
      this.requestRemainingData()
    },
    updateProductsData(emitData){
      DataProvider.GetSuggestionReference('products', emitData.val)
        .then((response) => {
          this.conditionSuggestion = response.data
        })
        .catch(error => { DataProvider.ErrorProcessing(error) });
    },
    requestRemainingData(){
      DataProvider.GetReport('remaining', this.condition.id)
        .then((response) => {
          this.remainingData = response.data
        })
        .catch(error => { DataProvider.ErrorProcessing(error) });
    },
    showReport(reportName, id){
      console.log(`show report for ${id}`)
      router.push(`/reports/${reportName}`)
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