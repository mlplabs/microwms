<template>
  <div>
    <div id="detailForm" class="modal fade" tabindex="-1">
      <div class="modal-dialog modal-lg modal-dialog-centered">
        <div class="modal-content">
          <div class="modal-header">
            <h5 v-if="detailItem.isNew" class="modal-title">{{lng.title_form_create}}</h5>
            <h5 v-else class="modal-title">{{lng.title_form_edit}}</h5>

            <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close" @click="closeDetailForm"></button>
          </div>
          <div class="modal-body">

            <form>
              <div class="mb-3">
                <label for="inputMnf" class="form-label">Наименование</label>
                <autocomplete-input
                  v-model:prop-suggestions="whsSuggestion"
                  v-model:prop-selection-id="detailItem.id"
                  v-model:prop-selection-val="detailItem.name"
                  @onUpdateData="updateManufacturersData">
                </autocomplete-input>
              </div>
              <div class="mb-3">
                <label for="inputAddr" class="form-label">Адрес</label>
                <input id="inputAddr" class="form-control" type="text" v-model="detailItem.address"/>
              </div>
              <div class="mb-3">
              </div>
              <div class="mb-3">
              </div>

              <ul class="nav nav-tabs" id="myTab" role="tablist">
                <li class="nav-item" role="presentation">
                  <button class="nav-link active" id="home-tab" data-bs-toggle="tab" data-bs-target="#home-tab-accept" type="button" role="tab" aria-controls="home-tab-pane" aria-selected="true">Приемка</button>
                </li>
                <li class="nav-item" role="presentation">
                  <button class="nav-link" id="profile-tab" data-bs-toggle="tab" data-bs-target="#profile-tab-storage" type="button" role="tab" aria-controls="profile-tab-pane" aria-selected="false">Хранение</button>
                </li>
                <li class="nav-item" role="presentation">
                  <button class="nav-link" id="contact-tab" data-bs-toggle="tab" data-bs-target="#contact-tab-ship" type="button" role="tab" aria-controls="contact-tab-pane" aria-selected="false">Отгрузка</button>
                </li>
              </ul>
              <div class="tab-content" id="myTabContent">
                <div class="tab-pane fade show active pt-2" id="home-tab-accept" role="tabpanel" aria-labelledby="" tabindex="0">
                    <label for="inputZoneIn" class="form-label">Зона приемки</label>
                    <input id="inputZoneIn" class="form-control" disabled type="text" v-model="detailItem.acceptance_zone.name"/>
                </div>
                <div class="tab-pane fade pt-2" id="profile-tab-storage" role="tabpanel" aria-labelledby="" tabindex="0">
                  <label for="inputZoneOut" class="form-label">Зона хранения</label>
                  <input id="inputZoneOut" class="form-control" disabled type="text" />
                </div>
                <div class="tab-pane fade pt-2" id="contact-tab-ship" role="tabpanel" aria-labelledby="" tabindex="0">
                  <label for="inputZoneStore" class="form-label">Зона отгрузки</label>
                  <input id="inputZoneStore" class="form-control" disabled type="text" v-model="detailItem.shipping_zone.name" /> <!-- v-model="detailItem.storage_zones[0].name" -->
                </div>
              </div>
            </form>

          </div>
          <div class="modal-footer">
            <button type="button" class="btn btn-secondary" data-bs-dismiss="modal" @click="closeDetailForm">{{lng.btn_form_close}}</button>
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
      <thead>
      <tr>
        <th scope="col" class="col_head col_id">#</th>
        <th scope="col" class="col_head">Наименование</th>
        <th scope="col" class="col_head col_action">...</th>
      </tr>
      </thead>
      <tbody>
      <tr v-for="(item, index) in tableData" :key="index">
        <td class="col_id">{{ item.id }}</td>
        <td><a href="#" data-bs-toggle="modal" data-bs-target="#detailForm" @click="showDetailForm(item.id)">{{ item.name }}</a></td>
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
  name: "ReferenceWhs",
  components: {AutocompleteInput, PaginationBar},

  data(){
    return {
      refName: 'warehouses',
      tableData: [],
      countRows: 0,
      limitRows: 7,
      currentPage: 1,
      detailItem: {
        isNew: false,
        id: 0,
        name: "",
        address: "",
        acceptance_zone: {
          id: 0,
          name: "",
        },
        shipping_zone: {
          id: 0,
          name: "",
        },
        storage_zones: [
          {
            id: 0,
            name: "",
          }
        ],
      },
      whsSuggestion: [],
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
        title: "Склады",
        title_form_create: "Создание склада",
        title_form_edit: "Редактирование склада",
        btn_list_create: "Новый склад",
        btn_form_close: "Закрыть",
        btn_form_store: "Сохранить",
      },
    }
  },

  methods:{
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
        address: '',
        acceptance_zone: {id: 0, name: ''},
        shipping_zone:  {id: 0, name: ''},
        storage_zones: [ {id: 0, name: ''} ]
      }
      this.whsSuggestion = []
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

    updateManufacturersData(emitData){
      DataProvider.GetSuggestionReference(this.refName, emitData.val)
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