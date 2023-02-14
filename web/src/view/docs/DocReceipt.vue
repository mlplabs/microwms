<template>
  <div>
    <div id="detailForm" class="modal fade" tabindex="-1">
      <div class="modal-dialog modal-xl modal-dialog-centered">
        <div class="modal-content">
          <div class="modal-header">
            <h5 class="modal-title">{{ lng.title_form_create }}</h5>
            <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close" @click="closeDetailForm"></button>
          </div>
          <div class="modal-body">

            <form class="row g-3">
              <div class="col-md-3">
                <label for="inputNumber" class="form-label">Номер</label>
                <input type="text" class="form-control" disabled id="inputNumber" v-model="detailItem.number">
              </div>
              <div class="col-md-3">
                <label for="inputDate" class="form-label">Дата</label>
                <input type="text" class="form-control" id="inputDate" v-model="detailItem.date">
              </div>

              <div class="col-12 table-responsive-xl">

                <inline-table
                  :columns="productColumns"
                  :rows="detailItem.items"
                  :suggestionData="suggestion"
                  :is-show-paging="false"
                  :is-show-search="false"
                  @row-clicked="onClickFormRow"
                  @new-item-clicked="onNewItemForm"
                  @row-delete="onDeleteFormRow"
                  @update-suggestion="onUpdateSuggestion"
                  @select-suggestion="onSelectSuggestion"
                ></inline-table>

              </div>
            </form>
          </div>
          <div class="modal-footer">
            <button type="button" class="btn btn-secondary" data-bs-dismiss="modal" @click="closeDetailForm">{{lng.btn_form_close}}</button>
            <button type="button" class="btn btn-primary" data-bs-dismiss="modal" @click="storeItem">{{lng.btn_form_store}}</button>
          </div>
        </div>
      </div>
    </div>
  </div>
  <div class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
    <h5>{{lng.title}}</h5>
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
        <th scope="col" class="col_head">Номер</th>
        <th scope="col" class="col_head">Дата</th>
        <th scope="col" class="col_head col_action">...</th>
      </tr>
      </thead>
      <tbody>
      <tr v-for="(item, index) in tableData" :key="index">
        <td class="col_id">{{ item.id }}</td>
        <td><a href="#" data-bs-toggle="modal" data-bs-target="#detailForm" @click="showDetailForm(item.id)">{{ item.number }}</a></td>
        <td><a href="#" data-bs-toggle="modal" data-bs-target="#detailForm" @click="showDetailForm(item.id)">{{ item.date }}</a></td>
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
  name: "DocReceipt",
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
        items: [],
      },
      tableData: [],
      suggestion: [],
      productColumns: [
        {
          label: "#",
          field: "id",
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
          field: "cell",
          isKey: false,
          suggestion: true,
          align: 0
        },
        {
          label: "Количество",
          field: "quantity",
          isKey: false,
          isNum: true,
          align: 2
        },
        {
          label: "Действия",
          field: "actions",
          isKey: false,
          align: 1
        }
      ],

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
    closeDetailForm(){

    },
    storeItem(){
      console.log(this.detailItem)
      DataProvider.StoreReceiptDoc("receipt", this.detailItem)
        .then((response) => {
          const storeId = response.data;
          if (storeId > 0) {
            this.updateItemsOnPage(this.currentPage)
          }
        })
        .catch(error => { this.errorProc(error) });
    },
    resetDetailItem(){
      this.detailItem = {
        id: 0,
        number: '',
        date: '',
        items: []
      }

    },

    showDetailForm(id){
      this.resetDetailItem()
      if (id === 0) {
        return
      }
      this.getDetailItem(id)
    },
    deleteRow(idx){
      console.log('delete ' + idx)
      this.tableData.splice(idx+1, 1)
    },

    getDetailItem(id){
      DataProvider.GetReceiptDoc("receipt", id)
        .then((response) => {
          this.detailItem = response.data
        })
        .catch(error => { this.errorProc(error) });
    },
    onSelectPage(){

    },

    updateItemsOnPage(page){
      let offset = ( page -1 ) * this.limitRows
      DataProvider.GetReceiptDocs("receipt", page, this.limitRows, offset)
        .then((response) => {
          console.log(response.data)
          this.tableData = response.data.data
          this.countRows = response.data.header.count
        })
        .catch(error => { this.errorProc(error) });
    },

     fillProductRow(row){
       DataProvider.GetItemReference('products', row.id)
         .then((response) => {
           //this.detailItem = response.data
           console.log(response.data)
           row.product_manufacturer =response.data.manufacturer.name
         })
         .catch(error => { this.errorProc(error) });
     },

    onSelectSuggestion(emitRow){
      if (emitRow.id !== 0) {
        this.fillProductRow(emitRow)
      }
    },


    onUpdateSuggestion(emitData){
      //console.log('suggestion update for ' + emitData.key + ' ' + emitData.val)
      if(emitData.key === "product_name"){
        this.updateProductsData(emitData)
      }
      if(emitData.key === "product_manufacturer"){
        this.updateManufacturersData(emitData)
      }

    },
    updateProductsData(emitData){
      DataProvider.GetSuggestionReference('products', emitData.val)
        .then((response) => {
          console.log(response.data)
          this.suggestion = response.data
        })
        .catch(error => { this.errorProc(error) });
    },

    updateManufacturersData(emitData){
      DataProvider.GetSuggestionReference('manufacturers', emitData.val)
        .then((response) => {
          this.suggestion = response.data
        })
        .catch(error => { this.errorProc(error) });
    },

    onNewItemForm(){
      let newBc = this.detailItem.items.find(item => item.name === "");
      if (newBc !== undefined){
        return
      }
      this.detailItem.items.push({id: 0, product_name: "", product_manufacturer: "", product_barcode: "", cell:"", quantity: 0})
    },
    onDeleteFormRow(idx){
      console.log('delete ' + idx)
      this.detailItem.items.splice(idx, 1)
    },
    onClickFormRow(eventData){
      console.log(eventData)
    },
    errorProc(error){
      console.log(error)
      if (error.response) {
        // client received an error response (5xx, 4xx)
        if (error.response.status === 404){
          this.statusText = "Ничего не найдено ("
        }else {
          this.statusText = "Произошла ошибка ("+error.response.status+")"
        }
      } else if (error.request) {
        // client never received a response, or request never left
      } else {
        // anything else
      }

    }
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