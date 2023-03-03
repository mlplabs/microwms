<template>
  <div>
    <div id="detailForm" class="modal fade" tabindex="-1">
      <div class="modal-dialog modal-xl modal-dialog-centered">
        <div class="modal-content">
          <div class="modal-header">
            <h5 class="modal-title">{{ lng.title_form_create }}</h5>
            <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close" @click="onClickFormClose"></button>
          </div>
          <div class="modal-body">

            <form class="row g-3">
              <div class="col-md-3">
                <label for="inputNumber" class="form-label">Номер</label>
                <input type="text" class="form-control" id="inputNumber" v-model="detailItem.number" :readonly="detailItem.id !== 0">
              </div>
              <div class="col-md-3">
                <label for="inputDate" class="form-label">Дата</label>
                <input type="text" class="form-control" id="inputDate" v-model="detailItem.date" :readonly="detailItem.id !== 0">
              </div>
              <div class="col-md-3">
                <label for="inputWhs" class="form-label">Склад</label>
                <input type="text" class="form-control" id="inputWhs" v-model="detailItem.whs" :readonly="detailItem.id !== 0">
              </div>

              <div class="col-12 table-responsive-xl">

                <inline-table
                  :columns="productColumns"
                  :rows="detailItem.items"
                  :suggestionData="suggestion"
                  :is-show-paging="false"
                  :is-show-search="false"
                  @onRowClick="onClickTableRow"
                  @onNewItem-clicked="onClickTableNewItem"
                  @onRowDelete="onClickTableDelRow"
                  @onUpdateSuggestion="onUpdateSuggestionTable"
                  @onSelectSuggestion="onSelectSuggestionTable"
                ></inline-table>

              </div>

            </form>

          </div>
          <div class="modal-footer">
            <button type="button" class="btn btn-secondary" data-bs-dismiss="modal" @click="onClickFormClose">{{lng.btn_form_close}}</button>
            <button type="button" class="btn btn-primary" data-bs-dismiss="modal" @click="onClickFormStore">{{lng.btn_form_store}}</button>
          </div>
        </div>
      </div>
    </div>
  </div>
  <div class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
    <h5>{{lng.title}}</h5>
    <div class="btn-toolbar mb-2 mb-md-0">
      <div class="btn-group me-2">
        <button type="button" class="btn btn-sm btn-outline-secondary" data-bs-toggle="modal" data-bs-target="#detailForm"  @click="showForm(0) ">
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
        <th scope="col" class="col_head">Тип</th>
        <th scope="col" class="col_head">Номер</th>
        <th scope="col" class="col_head">Дата</th>
        <th scope="col" class="col_head col_action">...</th>
      </tr>
      </thead>
      <tbody>
      <tr v-for="(item, index) in tableData" :key="index">
        <td class="col_id">{{ item.id }}</td>
        <td><span>{{ this.getDocType(item.doc_type) }}</span></td>
        <td><a href="#" data-bs-toggle="modal" data-bs-target="#detailForm" @click="showForm(item.id)">{{ item.number }}</a></td>
        <td><a href="#" data-bs-toggle="modal" data-bs-target="#detailForm" @click="showForm(item.id)">{{ item.date }}</a></td>
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

import InlineTable from "@/components/InlineTable";
import DataProvider from "@/services/DataProvider";
import PaginationBar from "@/components/PaginationBar";

export default {
  name: "DocShipment",
  components: {InlineTable, PaginationBar},
  data(){
    return{
      countRows: 0,
      limitRows: 11,
      currentPage: 1,
      detailItem:{
        id: 0,
        number: "",
        date: "08.02.2023",
        whs: '',
        items: [],
      },
      tableData: [],
      suggestion: [],


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
        .catch(error => { this.errorProc(error) });
    },
    onClickFormClose(){

    },
    showForm(id){
      this.resetDetailItem()

      for(let i=0; i< this.productColumns.length; i++){
        if (this.productColumns[i].field === 'product_name'
          || this.productColumns[i].field === 'product_manufacturer'
          || this.productColumns[i].field === "quantity") {
          this.productColumns[i].readonly = this.detailItem.id !== 0
        }
      }

      if (this.detailItem.id === 0) {
        return
      }
      this.getDetailItem(id)
    },
    resetDetailItem(){
      this.detailItem = {
        id: 0,
        number: '',
        date: '',
        items: []
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
    // Update suggestions list
    onUpdateSuggestionTable(emitData){
      console.log('suggestion update for ' + emitData.key + ' ' + emitData.val)
      if(emitData.key === "product_name"){
        this.updateProductsData(emitData)
        return
      }
      if(emitData.key === "product_manufacturer"){
        this.updateManufacturersData(emitData)
        return
      }
      this.suggestion = []
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
      DataProvider.GetReceiptDocs("receipt", page, this.limitRows, offset)
        .then((response) => {
          console.log(response.data)
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
    fillProductRow(row){
      DataProvider.GetItemReference('products', row.id)
        .then((response) => {
          row.product_manufacturer =response.data.manufacturer.name
        })
        .catch(error => { DataProvider.ErrorProcessing(error) });
    },
    // Products suggestions
    updateProductsData(emitData){
      DataProvider.GetSuggestionReference('products', emitData.val)
        .then((response) => {
          console.log('updateProductsData: ' + response.data)
          this.suggestion = response.data
        })
        .catch(error => { DataProvider.ErrorProcessing(error) });
    },
    // Manufacturers suggestions
    updateManufacturersData(emitData){
      DataProvider.GetSuggestionReference('manufacturers', emitData.val)
        .then((response) => {
          console.log('updateManufacturersData: ' + response.data)
          this.suggestion = response.data
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