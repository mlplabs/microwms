<template>
  <div>
    <div id="detailForm" class="modal fade" tabindex="-1">
      <div class="modal-dialog modal-lg modal-dialog-centered">
        <div class="modal-content">
          <div class="modal-header">
            <h5 class="modal-title">Поступление товаров</h5>
            <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close" @click="closeDetailForm"></button>
          </div>
          <div class="modal-body">

            <form class="row g-3">
              <div class="col-md-3">
                <label for="inputNumber" class="form-label">Номер</label>
                <input type="text" class="form-control" id="inputNumber">
              </div>
              <div class="col-md-3">
                <label for="inputDate" class="form-label">Дата</label>
                <input type="text" class="form-control" id="inputDate">
              </div>
              <div class="col-12">
                <label for="inputShipper" class="form-label">Поставщик</label>
                <input type="text" class="form-control" id="inputShipper">
              </div>
              <!-- div class="col-12">
                <div class="btn-toolbar mb-2 mb-md-0">
                  <div class="btn-group me-2">
                    <button type="button" class="btn btn-sm btn-outline-secondary" data-bs-toggle="" data-bs-target=""  @click="addRow"><i class="bi bi-journal-plus text-primary" style="padding-right: 3px;"></i>Новый товар</button>
                  </div>
                </div>
              </div -->

              <div class="col-12">

                <inline-table :columns="columns" :rows="rows" :is-show-paging="false"
                  @row-clicked="onRowClick"
                  @new-item-clicked="onNewClick"
                ></inline-table>

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
            <button type="button" class="btn btn-secondary" data-bs-dismiss="modal" @click="closeDetailForm">Закрыть</button>
            <button type="button" class="btn btn-primary" data-bs-dismiss="modal" @click="storeItem">Сохранить</button>
          </div>
        </div>
      </div>
    </div>
  </div>
  <div class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
    <h1 class="h2">Поступления</h1>
    <div class="btn-toolbar mb-2 mb-md-0">
      <div class="btn-group me-2">
        <button type="button" class="btn btn-sm btn-outline-secondary" data-bs-toggle="modal" data-bs-target="#detailForm"  @click="showDetailForm(0) "><i class="bi bi-journal-plus text-primary" style="padding-right: 3px;"></i>Новое поступление</button>
      </div>
    </div>
  </div>
</template>

<script>

import InlineTable from "@/components/InlineTable";
export default {
  name: "DocReceipt",
  components: {InlineTable},
  data(){
    return{
      tableData: [],
      tableItem:{
        id: 0,
        name: ''
      },
      columns: [
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
      rows: [
        {
          id: 1,
          name: 'jack',
          action: 'example@example.com'
        },
        {
          id: 2,
          name: 'rose',
          action: 'example@example.com'
        }
      ]
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
      this.rows.push({id:99, name: 'asdfasd ', action: ' asdfasdf'})
    },
    deleteRow(idx){
      console.log('delete ' + idx)
      this.tableData.splice(idx+1, 1)
    },
    onRowClick(eventData){
      console.log(eventData)
    },
    onNewClick(){
      console.log('new click')
    },
  }

}
</script>

<style scoped>

</style>