<template>
  <div class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
    <h5>{{lng.title}}</h5>
    <div class="btn-toolbar mb-2 mb-md-0">
      <div class="btn-group me-2">
        <button type="button" class="btn btn-sm btn-outline-secondary" @click="showForm(0) ">
          <i class="bi bi-plus"></i>
        </button>
      </div>
    </div>
  </div>

  <div id="inlineForm" v-show="inlineFormIsOpen" class="container-md p-3 border mb-2 start-0">
    <form class="row g-2" >
      <div class="col-md-3" v-show="receiptAsDocument">
        <label for="inputNumber" class="form-label">Номер</label>
        <input type="text" class="form-control" id="inputNumber" v-model="detailItem.number" :readonly="detailItem.id !== 0">
      </div>
      <div class="col-md-3" v-show="receiptAsDocument">
        <label for="inputDate" class="form-label">Дата</label>
        <input type="text" class="form-control" id="inputDate" v-model="detailItem.date" :readonly="detailItem.id !== 0">
      </div>
      <div class="col-md-3" v-show="receiptAsDocument">
        <label for="inputWhs" class="form-label">Склад</label>
        <input type="text" class="form-control" id="inputWhs" v-model="detailItem.whs" :readonly="detailItem.id !== 0">
      </div>

      <div class="col-md-6">
        <label for="inputName" class="form-label">Наименование</label>
        <autocomplete-input
            v-model:prop-suggestions="itemNameSuggestion"
            v-model:prop-selection-id="detailItem.Id"
            v-model:prop-selection-val="detailItem.Name"
            @onUpdateData="updateProductsData"
            @onSelectData="fillProductFields"
        ></autocomplete-input>
      </div>
      <div class="col-md-6">
        <label for="inputMnf" class="form-label">Производитель</label>
        <autocomplete-input
            v-model:prop-suggestions="itemMnfSuggestion"
            v-model:prop-selection-id="detailItem.MnfId"
            v-model:prop-selection-val="detailItem.MnfName"
            @onUpdateData="updateManufacturersData"
        ></autocomplete-input>
      </div>
      <div class="col-md-4">
        <label for="inputNumber" class="form-label">Штрих-код</label>
        <input type="text" class="form-control" id="inputNumber">
      </div>
      <div class="col-md-4">
        <label for="inputNumber" class="form-label">Ячейка</label>
        <input type="text" class="form-control" id="inputNumber">
      </div>
      <div class="col-md-4">
        <label for="inputNumber" class="form-label">Количество</label>
        <input type="number" class="form-control" id="inputNumber" min="0">
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
        <th scope="col" class="col_head col_id">#</th>
        <th scope="col" class="col_head">Тип</th>
        <th scope="col" class="col_head">Номер</th>
        <th scope="col" class="col_head">Дата</th>
        <th scope="col" class="col_head col_action">...</th>
      </tr>
      </thead>
      <tbody>
      <tr v-for="(item, index) in tableData" :key="index">
        <td class="col_id">{{ item.id }}</td>
        <td><span>{{ item.doc.date }}</span></td>
        <td><a href="#" @click="showForm(item.product.id)">{{ item.product.name }}</a></td>
        <td><a href="#" @click="showForm(item.product.manufacturer.id)">{{ item.product.manufacturer.name }}</a></td>
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

      detailItem:{
        Id: 0,
        Name: "",
        MnfId: 0,
        MnfName: "",
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
  computed:{
    productColumns(){
      return [
        {
          label: "#",
          field: "product_id",
          isKey: true,
          isNum: true,
          align: 2
        },
        {
          label: "Наименование",
          field: "product_name",
          field_id: "id",
          isKey: false,
          suggestion: true,
          align: 0,
        },
        {
          label: "Производитель",
          field: "product_manufacturer",
          isKey: false,
          suggestion: true,
          align: 0
        },
        {
          label: "Штрих-код",
          field: "product_barcode",
          isKey: false,
          suggestion: true,
          align: 0
        },
        {
          label: "Ячейка",
          field: "cell_name",
          isKey: false,
          suggestion: true,
          align: 0
        },
        {
          label: "Количество",
          field: "quantity",
          isKey: false,
          isNum: true,
          readonly: false,
          align: 2
        },
        {
          label: "Действия",
          field: "actions",
          isKey: false,
          align: 1
        }
      ]
    },

  },
  methods:{
    /*
    * List
    *   DetailForm (Form)
    *     DetailTable (Table)
    */

    // DETAIL FORM METHODS
    onClickFormStore(){
      DataProvider.StoreReceiptDoc("receipt", this.detailItem)
        .then((response) => {
          const storeId = response.data;
          if (storeId > 0) {
            this.updateListItems(this.currentPage)
          }
        })
        .catch(error => { DataProvider.ErrorProcessing(error) });
    },
    onClickFormClose(){
      this.resetDetailItem()
      this.inlineFormIsOpen = false
    },
    showForm(id){

      if (this.inlineFormIsOpen === false) {
        this.inlineFormIsOpen = true
      } else {
        this.inlineFormIsOpen = false
      }
      console.log(id)
    },
    resetDetailItem(){
      this.detailItem = {
        id: 0,
        Name: '',
        MnfId: '',
        MnfName: ''
      }
    },

    onClickTableNewItem(){
      let newBc = this.detailItem.items.find(item => item.name === "");
      if (newBc !== undefined){
        return
      }
      this.detailItem.items.push({id: 0, product_name: "", product_manufacturer: "", product_barcode: "", cell:"", quantity: 0})
    },
    onClickTableDelRow(idx){
      console.log('delete ' + idx)
      this.detailItem.items.splice(idx, 1)
    },
    onClickTableRow(eventData){
      console.log(eventData)
    },

    // Select suggestion from table in detail form
    onSelectSuggestionTable(emitRow){
      if (emitRow.id !== 0) {
        this.fillProductRow(emitRow)
      }
    },

    // LIST ITEMS METHODS
    // selection page on pagination bar
    onSelectPage(eventData){
      this.currentPage = eventData.page
      this.updateListItems(eventData.page)
    },
    getDocType(doc_type){
      if(doc_type === 1){
        return "Поступление"
      }
      if(doc_type === 2){
        return "Оприходование"
      }
    },

    // COMMUNICATIONS METHODS
    // List items update
    updateListItems(page){
      let offset = ( page -1 ) * this.limitRows
      DataProvider.GetReceiptDocs("receipt/products", page, this.limitRows, offset)
        .then((response) => {
          this.tableData = response.data.data
          this.countRows = response.data.header.count
          console.log(this.tableData)
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
          this.detailItem.Name = response.data.name
          this.detailItem.MnfId = response.data.manufacturer.id
          this.detailItem.MnfName = response.data.manufacturer.name
          this.detailItem.Barcode = !response.data.barcodes[0].name
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
    deleteItem(id){
      DataProvider.DeleteReceiptDoc('receipt', id)
        .then((response) => {
          const affRows = response.data;
          if (affRows !== 1){
            console.log('delete failed')
          }
          this.updateListItems(this.currentPage)
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