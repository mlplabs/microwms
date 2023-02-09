<template>
  <div>
    <div id="detailForm" class="modal fade" tabindex="-1">
      <div class="modal-dialog modal-lg modal-dialog-centered">
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

              <div class="col-12">

                <inline-table
                  :columns="productColumns"
                  :rows="detailItem.products"
                  :is-show-paging="false"
                  :is-show-search="false"
                  @row-clicked="onRowClick"
                  @new-item-clicked="onNewItem"

                ></inline-table>


                <!--
                                <inline-table
                  :is-show-paging="false"
                  :is-show-search="false"
                  :rows="detailItem.barcodes"
                  :columns="barcodesColumns"
                  @new-item-clicked="onNewItem"
                  @row-delete="onDeleteItem">
                </inline-table>
                -->

                <!-- div class="table-responsive">
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
                      <td><input type="text"></td>
                      <td class="col_action"><button type="button" class="btn" @click="deleteRow(index)"><i class="bi bi-journal-x text-danger"></i></button> </td>
                    </tr>
                    </tbody>
                  </table>
                </div -->

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
</template>

<script>

import InlineTable from "@/components/InlineTable";
export default {
  name: "DocPosting",
  components: {InlineTable},
  data(){
    return{
      detailItem:{
        number: "",
        date: "08.02.2023",
        products: [],
      },
      tableData: [],
      tableItem:{
        id: 0,
        name: ''
      },
      productColumns: [
        {
          label: "#",
          field: "id",
          isKey: true,
          editable: false,
          align: 2
        },
        {
          label: "Наименование",
          field: "name",
          isKey: false,
          editable: false,
          align: 0
        },
        {
          label: "Количество",
          field: "quantity",
          isKey: false,
          editable: true,
          align: 2
        },
        {
          label: "Действия",
          field: "actions",
          isKey: false,
          editable: false,
          align: 1
        }
      ],

      lng: {
        title: "Оприходования",
        title_form_create: "Оприходование товара",
        btn_list_create: "Новый товар",
        btn_form_close: "Закрыть",
        btn_form_store: "Сохранить",
      },
    }
  },

  methods:{
    storeItem(){

    },
    closeDetailForm(){

    },
    showDetailForm(){
    },
    addRow(){
      this.rows.push({id:99, name: 'asdfasd '})
    },
    deleteRow(idx){
      console.log('delete ' + idx)
      this.tableData.splice(idx+1, 1)
    },
    onRowClick(eventData){
      console.log(eventData)
    },
    onNewItem(){
      let newBc = this.detailItem.products.find(item => item.name === "");
      if (newBc !== undefined){
        return
      }
      this.detailItem.products.push({id: 0, name: ""})
    },
  }

}
</script>

<style scoped>

</style>