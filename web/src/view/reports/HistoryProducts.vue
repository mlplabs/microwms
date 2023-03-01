<template>
  <div class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
    <h5>{{ lng.title }}</h5>
  </div>
  <form class="mb-3">
    <div class="row mb-3">
      <div class="col-4">
        <div class="row">
          <label for="inputName" class="col-sm-2 col-form-label">Товар:</label>
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
      <div class="col-4">
        <div class="row">
          <label for="inputDirect" class="col-sm-4 col-form-label">Вид&nbsp;движения:</label>
          <div class="col-sm-8">
            <input type="text" id="inputDirect" class="form-control">
          </div>
        </div>
      </div>
    </div>

    <div class="row">
      <div class="col-4">
        <div class="row">
          <label for="inputBegin" class="col-sm-3 col-form-label">Период&nbsp;с: </label>
          <div class="col-sm-9">
          <input type="text" id="inputBegin"  class="form-control">
          </div>
        </div>
      </div>
      <div class="col-4">
        <div class="row">
          <label for="inputEnd" class="col-sm-4 col-form-label">по: </label>
          <div class="col-sm-8">
          <input type="text" id="inputEnd"  class="form-control">
          </div>
        </div>
      </div>
      <div class="col-auto">
        <button type="submit" class="btn btn-md btn-outline-secondary" @click.prevent="generateReport">Сформировать</button>
      </div>
    </div>
  </form>

  <div id="reportData" class="table-responsive">

    <table class="table table-striped table-hover table-bordered">
      <thead>
      <tr>
        <th scope="col" class="col_head">Вид движения</th>
        <th scope="col" class="col_head">Ячейка</th>
        <th scope="col" class="col_head">Количество</th>
      </tr>
      </thead>
      <tbody>
      <tr>
        <td></td>
        <td></td>
        <td></td>
      </tr>
      </tbody>
    </table>
  </div>

</template>

<script>
import AutocompleteInput from "@/components/AutocompleteInput";
import DataProvider from "@/services/DataProvider";
//import InlineTable from "@/components/InlineTable";

export default {
  name: "HistoryProducts",
  components: {AutocompleteInput},
  data(){
    return{
      condition:{
        id: 0,
        text: ""
      },
      conditionSuggestion:[],
      lng: {
        title: "Карточка товара",
      },
    }
  },
  methods: {
    generateReport(){

    },
    updateProductsData(emitData){
      DataProvider.GetSuggestionReference('products', emitData.val)
        .then((response) => {
          this.conditionSuggestion = response.data
        })
        .catch(error => { this.errorProc(error) });
    },
  }
}
</script>

<style scoped>

</style>