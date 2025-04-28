<template>
  <div class="modal-content">
    <div class="modal-header">
      <h5 v-if="detailItem.id === 0" class="modal-title">{{lng.title_form_create}}</h5>
      <h5 v-else class="modal-title">{{ lng.title_form_edit }}</h5>
      <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close" @click="closeDetailForm"></button>
    </div>
    <div class="modal-body">
      <form>
        <input type="hidden" value="{{detailItem.id}}" >
        <div class="mb-3">
          <autocomplete-input
            v-model:prop-suggestions="productsSuggestion"
            v-model:prop-selection-val="detailItem.name"
            v-model:prop-placeholder="lng.label_name"
            @onUpdateData="updateProductsData">
          </autocomplete-input>

          <!-- div id="nameHelp" class="form-text">We'll never share your email with anyone else.</div-->
        </div>
        <div class="row mb-3">
          <div class="col-md-8 col-6">
            <autocomplete-input
              v-model:prop-suggestions="manufacturersSuggestion"
              v-model:prop-selection-id="detailItem.manufacturer.id"
              v-model:prop-selection-val="detailItem.manufacturer.name"
              v-model:prop-placeholder="lng.label_manufacturer"
              @onUpdateData="updateManufacturersData">
            </autocomplete-input>
          </div>

          <div class="col-md-4 col-6">
            <input type="text" class="form-control" id="inputItemNumber" v-model="detailItem.item_number" :placeholder="lng.label_item_number">
          </div>
        </div>
        <div class="row mb-3">
          <div class="col-md">
            <input class="form-control tags" data-role="tagsinput" type="text" id="inputTags" v-model="detailItem.tags" :placeholder="lng.label_tags">
          </div>
        </div>
        <div class="mb-3">
          <inline-table
            :is-show-paging="false"
            :is-show-search="false"
            :rows="detailItem.barcodes"
            :columns="barcodesColumns"
            @onNewItemClick="onNewBarcodeItem"
            @onRowDelete="onDeleteBarcodeItem">
          </inline-table>

        </div>
      </form>

    </div>
    <div class="modal-footer">
      <button type="button" class="btn btn-secondary" data-bs-dismiss="modal" @click="closeDetailForm">{{ lng.btn_form_close }}</button>
      <button type="button" class="btn btn-primary" data-bs-dismiss="modal" @click="storeItem">{{ lng.btn_form_store }}</button>
    </div>
  </div>
</template>

<script>
import DataProvider from "@/services/DataProvider";
import AutocompleteInput from "@/components/AutocompleteInput";
import InlineTable from "@/components/InlineTable";

export default {
  name: "ProductForm",
  components: {AutocompleteInput, InlineTable},
  props:{
    propProductId: Number
  },
  data(){
    return{
      detailItem: {
        id: 0,
        name: "",
        item_number: "",
        manufacturer:{
          id: 0,
          name: ''
        },
        tags: "",
        barcodes:[]
      },
      barcodeTypes:[],
      barcodesColumns:[
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
          values: []
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
        label_name: "Наименование",
        label_manufacturer: "Производитель",
        label_item_number: "Артикул",
        label_tags: "Тэги",
        title_form_create: "Создание товара",
        title_form_edit: "Редактирование товара",
        btn_form_close: "Закрыть",
        btn_form_store: "Сохранить",
      },
    }
  },
  methods:{
    // Barcodes
    onNewBarcodeItem(){
      let newBc = this.detailItem.barcodes.find(item => item.id === 0);
      if (newBc !== undefined){
        return
      }
      this.detailItem.barcodes.push({id: 0, name:'', type: 0})
    },

    onDeleteBarcodeItem(emitData){
      console.log('del barcode ' + emitData)
      let idx = this.detailItem.barcodes.findIndex(item => item.id === emitData.id);
      if (emitData.id != 0) {
        this.deleteBarcodeItem(emitData.id, idx)
      }else {
        this.detailItem.barcodes.splice(idx, 1)
      }
    },

    deleteBarcodeItem(id, idx){
      DataProvider.DeleteItemReference('barcodes', id)
        .then((response) => {
          const affRows = response.data;
          if (affRows !== 1){
            console.log('delete failed')
          }
          // удалим из списка, когда из базы удалим
          this.detailItem.barcodes.splice(idx, 1)
        })
        .catch(error => { DataProvider.ErrorProcessing(error) });
    },

    getDetailItem(id){
      DataProvider.GetItemReference('products', id)
        .then((response) => {
           this.detailItem = response.data.data
        })
        .catch(error => { DataProvider.ErrorProcessing(error) });
    },
    storeItem(){
      DataProvider.StoreItemReference('products', this.detailItem)
        .then((response) => {
          const storeId = response.data.data;
          if (storeId > 0) {
            this.$emit('onUpdateData', {storeId} )
          }
        })
        .catch(error => { DataProvider.ErrorProcessing(error) });

      this.resetDetailItem()
    },
    closeDetailForm(){
    },

    updateProductsData(emitData){
      DataProvider.GetSuggestionReference('products', emitData.val)
        .then((response) => {
          this.productsSuggestion = response.data.data
        })
        .catch(error => { DataProvider.ErrorProcessing(error) });
    },

    updateManufacturersData(emitData){
      DataProvider.GetSuggestionReference('manufacturers', emitData.val)
        .then((response) => {
          this.manufacturersSuggestion = response.data.data
        })
        .catch(error => { DataProvider.ErrorProcessing(error) });
    },

    getEnumBarcodeType(){
      DataProvider.GetEnum('barcodes/types')
        .then((response) => {
          console.log(response)
          this.barcodeTypes = response.data.data
          this.updateBarcodeColumn()
        })
        .catch(error => { this.errorProc(error) });
    },

    updateBarcodeColumn(){
      for (let i=0; i<this.barcodesColumns.length;i++ ) {
        if (this.barcodesColumns[i].field === 'type') {
          this.barcodesColumns[i].values = this.barcodeTypes
        }
      }
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

      if (this.barcodeTypes?.length === 0){
        this.getEnumBarcodeType()
      }
    }

  },
  mounted() {
    this.getEnumBarcodeType()
  },
  watch:{
    propProductId(val, oldVal) {
      if (val !== oldVal) {
        // reset
        this.resetDetailItem()
        if (val !== 0)
          this.getDetailItem(val)
      }
    }
  },
}
</script>

<style scoped>

</style>