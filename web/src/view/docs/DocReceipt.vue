<template>
  <div class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
    <h5>{{lng.title}}</h5>
    <div class="btn-toolbar mb-2 mb-md-0">
      <div class="btn-group me-2">
        <button type="button" class="btn btn-sm btn-outline-secondary" @click="showForm()">
          <i class="bi bi-plus"></i>
        </button>
      </div>
    </div>
  </div>

  <div id="inlineForm" v-show="inlineFormIsOpen" class="container-md p-3 border mb-2 start-0">
    <form class="row g-2" >
      <div class="col-md-3" v-show="receiptAsDocument">
        <label for="inputNumber" class="form-label">Номер</label>
        <input type="text" class="form-control" id="inputNumber" v-model="docItem.number" :readonly="docItem.id !== 0">
      </div>
      <div class="col-md-3" v-show="receiptAsDocument">
        <label for="inputDate" class="form-label">Дата</label>
        <input type="text" class="form-control" id="inputDate" v-model="docItem.date" :readonly="docItem.id !== 0">
      </div>
      <div class="col-md-3" v-show="receiptAsDocument">
        <label for="inputWhs" class="form-label">Склад</label>
        <input type="text" class="form-control" id="inputWhs" v-model="docItem.whs" :readonly="docItem.id !== 0">
      </div>

      <div class="col-md-6">
        <label for="inputName" class="form-label">Наименование</label>
        <autocomplete-input
            v-model:prop-suggestions="itemNameSuggestion"
            v-model:prop-selection-id="docItem.items[0].product.id"
            v-model:prop-selection-val="docItem.items[0].product.name"
            @onUpdateData="updateProductsData"
            @onSelectData="fillProductFields"
        ></autocomplete-input>
      </div>
      <div class="col-md-6">
        <label for="inputMnf" class="form-label">Производитель</label>
        <autocomplete-input
            v-model:prop-suggestions="itemMnfSuggestion"
            v-model:prop-selection-id="docItem.items[0].product.manufacturer.id"
            v-model:prop-selection-val="docItem.items[0].product.manufacturer.name"
            @onUpdateData="updateManufacturersData"
        ></autocomplete-input>
      </div>
      <div class="col-md-4">
        <label for="inputNumber" class="form-label">Штрих-код</label>
        <input type="text" class="form-control" id="inputNumber" v-model="docItem.items[0].barcode">
      </div>
      <div class="col-md-4">
        <label for="inputNumber" class="form-label">Ячейка</label>
        <input type="text" class="form-control" id="inputNumber" v-model="docItem.items[0].cell">
      </div>
      <div class="col-md-4">
        <label for="inputNumber" class="form-label">Количество</label>
        <input type="number" class="form-control" id="inputNumber" min="0" v-model="docItem.items[0].quantity">
      </div>
      <div class="col-md-12 text-end">
        <button type="button" class="btn btn-secondary me-2" @click="onClickFormClose">{{lng.btn_form_close}}</button>
        <button type="button" class="btn btn-primary" @click="onClickFormStore">{{lng.btn_form_store}}</button>
      </div>

    </form>
  </div>
  <div class="table-responsive">
    <table class="table table-striped table-hover table-bordered">
      <thead>
      <tr>
        <th scope="col" class="col_head">Дата</th>
        <th scope="col" class="col_head">Товар</th>
        <th scope="col" class="col_head">Количество</th>
        <th scope="col" class="col_head col_action">...</th>
      </tr>
      </thead>
      <tbody>
      <tr v-for="(item, index) in tableData" :key="index">
        <td><span>{{ item.doc.date }}</span></td>
        <td>{{ item.product.name }}<br><small class="text-secondary">{{ item.product.manufacturer.name}} </small></td>

        <td class="col_id">{{ item.quantity }}</td>
        <td class="col_action">
          <div class="dropdown">
            <button type="button" class="btn btn-sm btn-outline-secondary" data-bs-toggle="dropdown">
              <i class="bi bi-three-dots-vertical"></i>
            </button>
            <ul class="dropdown-menu dropdown-menu-end">
              <li><a class="dropdown-item" href="#" @click.prevent="deleteItem(item.id)">Удалить {{item.name}}</a></li>
              <!--li><a class="dropdown-item" href="#">Another action {{item.id}}</a></li>
              <li><a class="dropdown-item" href="#">Something else here</a></li-->
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
import AutocompleteInput from "@/components/AutocompleteInput";

export default {
  name: "DocReceipt",
  components: {AutocompleteInput, PaginationBar},
  data(){
    return{
      countRows: 0,
      limitRows: 11,
      currentPage: 1,

      tableData: [],
      suggestion: [],

      // register the receipt of goods as a document
      receiptAsDocument: false,
      inlineFormIsOpen: false,

      docItem:{
        id: 0,
        number: '',
        date: '',
        whs: '',
        items: [
          {
            product:{
              id: 0,
              name: '',
              manufacturer:{
                id: 0,
                name: ''
              }
            },
            barcode: '',
            cell: '',
            quantity: 0
          }
        ]
      },

      itemNameSuggestion: [],
      itemMnfSuggestion: [],

      lng: {
        title: "Поступления",
        title_form_create: "Поступление/Оприходование товара",
        btn_list_create: "Новый товар",
        btn_form_close: "Закрыть",
        btn_form_store: "Сохранить",
      },
    }
  },
  methods:{
    /*
    * List
    *   DetailForm (Form)
    *     DetailTable (Table)
    */

    // DETAIL FORM METHODS
    onClickFormStore(){
      DataProvider.StoreDocument("receipt/docs", this.docItem)
        .then((response) => {
          const storeId = response.data;
          if (storeId > 0) {
            this.resetInlineForm()
            this.inlineFormIsOpen = false
            this.updateListItems(this.currentPage)
          }
        })
        .catch(error => { DataProvider.ErrorProcessing(error) });
    },
    onClickFormClose(){
      this.resetInlineForm()
      this.inlineFormIsOpen = false
    },

    resetInlineForm(){
      this.docItem.id = 0
      this.docItem.number = ''
      this.docItem.date = ''
      this.docItem.whs = ''
      this.docItem.items = [
            {
              product:{
                id: 0,
                name: '',
                manufacturer:{
                  id: 0,
                  name: ''
                }
              },
              barcode: '',
              cell: '',
              quantity: 0
            }
      ]
    },

    showForm(){
      if (this.inlineFormIsOpen === false) {
        this.inlineFormIsOpen = true
      } else {
        this.inlineFormIsOpen = false
      }
      if (this.docItem.items.length === 0){
        this.docItem.items.push()
      }
    },

    // LIST ITEMS METHODS
    // selection page on pagination bar
    onSelectPage(eventData){
      this.currentPage = eventData.page
      this.updateListItems(eventData.page)
    },

    // COMMUNICATIONS METHODS
    // List items update
    updateListItems(page){
      let offset = ( page -1 ) * this.limitRows
      console.log(process.env.VUE_APP_API)
      DataProvider.GetDocuments("receipt/products", page, this.limitRows, offset)
        .then((response) => {
          this.tableData = response.data.data
          this.countRows = response.data.header.count
        })
        .catch(error => { DataProvider.ErrorProcessing(error) });
    },
    // Getting product info
    getDetailItem(id){
      DataProvider.GetReceiptDoc("receipt", id)
        .then((response) => {
          this.detailItem = response.data
        })
        .catch(error => { DataProvider.ErrorProcessing(error) });
    },
    // Getting product and // fill row data //
    fillProductFields(emitData){
      DataProvider.GetItemReference('products', emitData.key)
        .then((response) => {
          this.docItem.items[0].product.name = response.data.name
          this.docItem.items[0].product.manufacturer.id = response.data.manufacturer.id
          this.docItem.items[0].product.manufacturer.name = response.data.manufacturer.name
          this.docItem.items[0].barcode = response.data.barcodes.length > 0 ? response.data.barcodes[0].name: ''
        })
        .catch(error => { DataProvider.ErrorProcessing(error) });
    },
    // Products suggestions
    updateProductsData(emitData){
      DataProvider.GetSuggestionReference('products', emitData.val)
        .then((response) => {
          this.itemNameSuggestion = response.data
        })
        .catch(error => { DataProvider.ErrorProcessing(error) });
    },
    // Manufacturers suggestions
    updateManufacturersData(emitData){
      DataProvider.GetSuggestionReference('manufacturers', emitData.val)
        .then((response) => {
          this.itemMnfSuggestion = response.data
        })
        .catch(error => { DataProvider.ErrorProcessing(error) });
    },
    // OTHER
  },
  mounted() {
    this.updateListItems(this.currentPage)
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