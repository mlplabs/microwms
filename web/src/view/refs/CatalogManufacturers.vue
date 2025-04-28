<template>
  <div>
    <div id="detailForm" class="modal fade" tabindex="-1">
      <div class="modal-dialog modal-lg modal-dialog-centered">
        <div class="modal-content">
          <div class="modal-header">
            <h5 v-if="detailItem.id === 0" class="modal-title">{{lng.title_form_create}}</h5>
            <h5 v-else class="modal-title">{{lng.title_form_edit}}</h5>

            <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close" @click="closeDetailForm"></button>
          </div>
          <div class="modal-body">

            <form>
              <div class="mb-3">
                <autocomplete-input
                  v-model:prop-suggestions="manufacturersSuggestion"
                  v-model:prop-selection-val="detailItem.name"
                  v-model:prop-placeholder="lng.label_name"
                  @onUpdateData="updateManufacturersData">
                </autocomplete-input>
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
    <h5>{{ lng.title }} {{globalSearch}}</h5>
    <div class="btn-toolbar mb-2 mb-md-0">
      <div class="btn-group me-2">
        <button type="button" class="btn btn-sm btn-outline-secondary" data-bs-toggle="modal" data-bs-target="#detailForm"  @click="showDetailForm(0) ">
          <i class="bi bi-plus"></i>
        </button>
      </div>
    </div>
  </div>

  <h2></h2>
  <div class="table-responsive">
    <table class="table table-striped table-hover table-bordered">
      <thead>
      <tr>
        <!-- th scope="col" class="col_head col_id">#</th -->
        <th scope="col" class="col_head">Наименование</th>
        <th scope="col" class="col_head col_action">...</th>
      </tr>
      </thead>
      <tbody>
      <tr v-for="(item, index) in tableData" :key="index">
        <!-- td class="col_id">{{ item.id }}</td -->
        <td><a href="#" class="text-decoration-none" data-bs-toggle="modal" data-bs-target="#detailForm" @click="showDetailForm(item.id)">{{ item.name }}</a></td>
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
import { inject } from 'vue'
import DataProvider from "@/services/DataProvider";
import PaginationBar from "@/components/PaginationBar";
import AutocompleteInput from "@/components/AutocompleteInput";

export default {
  name: "CatalogManufacturers",
  components: {AutocompleteInput, PaginationBar},

  data(){
    return {
      refName: 'manufacturers',
      tableData: [],
      countRows: 0,
      limitRows: 7,
      currentPage: 1,
      detailItem: {
        id: 0,
        name: "",
      },
      manufacturersSuggestion: [],
      columns: [
        {
          label: "#",
          field: "id",
          isKey: true,
        },
        {
          label: "Наименование",
          field: "name",
          isKey: false,
        },
        {
          label: "Действия",
          field: "actions",
          isKey: false,
        }
      ],
      lng: {
        title: "Производители",
        title_form_create: "Создание производителя",
        title_form_edit: "Редактирование производителя",
        btn_list_create: "Новый производитель",
        btn_form_close: "Закрыть",
        btn_form_store: "Сохранить",
        label_name: "Наименование",
      },
    }
  },

  methods:{
    showDetailForm(id){
      this.resetDetailItem()
      if (id === 0) {
        return
      }
      this.getDetailItem(id)
    },

    closeDetailForm(){
    },

    resetDetailItem(){
      this.detailItem.id = 0
      this.detailItem.name = ''
      this.manufacturersSuggestion = []
    },

    onSelectPage(eventData){
      this.currentPage = eventData.page
      this.updateItemsOnPage(eventData.page)
    },

    updateItemsOnPage(page){
      let offset = ( page -1 ) * this.limitRows
      DataProvider.GetItemsReference(this.refName, page, this.limitRows, offset, this.globalSearch)
        .then((response) => {
          this.tableData = response.data.data
          this.countRows = response.data.count
        })
        .catch(error => { this.errorProc(error) });
    },

    getDetailItem(id){
      DataProvider.GetItemReference(this.refName, id)
        .then((response) => {
          this.detailItem = response.data.data
        })
        .catch(error => { this.errorProc(error) });
    },

    storeItem(){
      DataProvider.StoreItemReference(this.refName, this.detailItem)
        .then((response) => {
          const storeId = response.data.data;
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

    updateManufacturersData(emitData){
      DataProvider.GetSuggestionReference(this.refName, emitData.val)
        .then((response) => {
          this.manufacturersSuggestion = response.data.data
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
    this.updateItemsOnPage(this.currentPage)
  },
  setup() {
    const globalSearch = inject('global_search')
    return { globalSearch }
  },
  watch:{
    globalSearch(newMsg, oldMsg) {
      console.log('old :'+ oldMsg + ', new:' + newMsg)
      this.updateItemsOnPage(this.currentPage)
    }
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