<template>
  <div>
    <div id="detailForm" class="modal fade" tabindex="-1">
      <div class="modal-dialog modal-lg modal-dialog-centered">
        <div class="modal-content">
          <div class="modal-header">
            <h5 v-if="detailItem.isNew" class="modal-title">{{lng.title_form_create}}</h5>
            <h5 v-else class="modal-title">{{ lng.title_form_edit }}</h5>

            <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close" @click="closeDetailForm"></button>
          </div>
          <div class="modal-body">
            <form>
              <input type="hidden" value="{{detailItem.id}}" >
              <div class="mb-3">
                <label for="inputName" class="form-label">Наименование</label>
                <autocomplete-input
                  v-model:prop-suggestions="productsSuggestion"
                  v-model:prop-selection-id="detailItem.id"
                  v-model:prop-selection-val="detailItem.name"
                  @onUpdateData="updateProductsData">
                </autocomplete-input>

                <!-- div id="nameHelp" class="form-text">We'll never share your email with anyone else.</div-->
              </div>
              <div class="row mb-3">
                <div class="col-md-8">
                  <label for="inputMnf" class="form-label">Производитель</label>
                  <autocomplete-input
                    v-model:prop-suggestions="manufacturersSuggestion"
                    v-model:prop-selection-id="detailItem.manufacturer.id"
                    v-model:prop-selection-val="detailItem.manufacturer.name"
                    @onUpdateData="updateManufacturersData">
                  </autocomplete-input>
                </div>

                <div class="col-md-4">
                  <label for="inputItemNumber" class="form-label">Артикул</label>
                  <input type="text" class="form-control" id="inputItemNumber" v-model="detailItem.item_number">
                </div>
              </div>
              <div class="mb-3">
                <inline-table
                  :is-show-paging="false"
                  :is-show-search="false"
                  :rows="detailItem.barcodes"
                  :columns="barcodesColumns"
                  @new-item-clicked="onNewItem"
                  @row-delete="onDeleteItem">
                </inline-table>

              </div>
            </form>

          </div>
          <div class="modal-footer">
            <button type="button" class="btn btn-secondary" data-bs-dismiss="modal" @click="closeDetailForm">{{ lng.btn_form_close }}</button>
            <button type="button" class="btn btn-primary" data-bs-dismiss="modal" @click="storeItem">{{ lng.btn_form_store }}</button>
          </div>
        </div>
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
      <thead> <!-- class="table-dark" -->
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
              <li><a class="dropdown-item" href="#" @click.prevent="printLabel(item.id)">Печать этикетки {{item.name}}</a></li>
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
import InlineTable from "@/components/InlineTable";

export default {
  name: "ReferenceProducts",
  components: {AutocompleteInput, PaginationBar, InlineTable},

  data(){
    return {
      refName: 'products',
      tableData: [],
      countRows: 0,
      limitRows: 11,
      currentPage: 1,
      detailItem: {
        isNew: false,
        id: 0,
        name: "",
        item_number: "",
        manufacturer:{
          id: 0,
          name: ''
        },
        barcodes:[]
      },
      barcodesColumns:[
        {
          label: "#",
          field: "id",
          isKey: true,
          align: 2
        },
        {
          label: "Штрих-код",
          field: "name",
          isKey: false,
          align: 0
        },
        {
          label: "Тип",
          field: "type",
          isKey: false,
          align: 0,
          values:[
            {key: 1, val: 'Ean 13'},
            {key: 2, val: 'Ean 14'},
            {key: 3, val: 'Ean 8'}
          ]
        },
        {
          label: "...",
          field: "actions",
          isKey: false,
          align: 1
        },
      ],
      productsSuggestion: [],
      manufacturersSuggestion: [],
      lng: {
        title: "Товары",
        title_form_create: "Создание товара",
        title_form_edit: "Редактирование товара",
        btn_list_create: "Новый товар",
        btn_form_close: "Закрыть",
        btn_form_store: "Сохранить",
      },
    }
  },

  methods:{
    onNewItem(){
      let newBc = this.detailItem.barcodes.find(item => item.id === 0);
      if (newBc !== undefined){
        return
      }
      this.detailItem.barcodes.push({id: 0, name:'', type: 0})
    },
    onDeleteItem(emitData){
      console.log(emitData)
      let idx = this.detailItem.barcodes.findIndex(item => item.id === emitData.id);
      this.detailItem.barcodes.splice(idx, 1)
    },

    // Открываем форму нового или существующего
    showDetailForm(id){
      this.resetDetailItem()
      this.detailItem.isNew = (id === 0)
      if (this.detailItem.isNew) {
        return
      }
      this.getDetailItem(id)
    },

    closeDetailForm(){
    },

    resetDetailItem(){
      this.detailItem = {
        id: 0,
        name: '',
        item_number: '',
        manufacturer: {id: 0, name: ''},
        barcodes: [],
      }
      this.productsSuggestion = []
      this.manufacturersSuggestion = []
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
        .catch(error => { this.errorProc(error) });
    },

    getDetailItem(id){
      DataProvider.GetItemReference(this.refName, id)
        .then((response) => {
          this.detailItem = response.data
        })
        .catch(error => { this.errorProc(error) });
    },

    storeItem(){
      DataProvider.StoreItemReference(this.refName, this.detailItem)
        .then((response) => {
          const storeId = response.data;
          if (storeId > 0) {
            this.updateItemsOnPage(this.currentPage)
          }
        })
        .catch(error => { this.errorProc(error) });
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
        .catch(error => { this.errorProc(error) });
    },
    printLabel(id){
      DataProvider.PrintItemReference(this.refName, id)
        .then((response) => {
          console.log(response)
        })
        .catch(error => { this.errorProc(error) });
    },
    updateProductsData(emitData){
      DataProvider.GetSuggestionReference(this.refName, emitData.val)
        .then((response) => {
          this.productsSuggestion = response.data
        })
        .catch(error => { this.errorProc(error) });
    },

    updateManufacturersData(emitData){
      DataProvider.GetSuggestionReference('manufacturers', emitData.val)
        .then((response) => {
          this.manufacturersSuggestion = response.data
        })
        .catch(error => { this.errorProc(error) });
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
    console.log(process.env.NODE_ENV);
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