<template>
  <div class="modal-content">
    <div class="modal-header">
      <h5 v-if="propModeForm === 0" class="modal-title">{{lng.title_form_get}}</h5>
      <h5 v-else-if="propModeForm === 1" class="modal-title">{{ lng.title_form_put }}</h5>
      <h5 v-else class="modal-title">{{ lng.title_form_move }}</h5>
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
        </div>
        <div class="row mb-3">
          <div class="col-md-5">
            <autocomplete-input
                v-model:prop-suggestions="cellsSuggestionFrom"
                v-model:prop-selection-id="detailItem.cellFrom.id"
                v-model:prop-selection-val="detailItem.cellFrom.name"
                v-model:prop-placeholder="lng.label_cell_from"
                @onUpdateData="updateCellsDataFrom">
            </autocomplete-input>
            <div id="nameHelp" class="form-text fw-lighter fst-italic">откуда</div>
          </div>

          <div class="col-md-5">
            <autocomplete-input
                v-model:prop-suggestions="cellsSuggestionTo"
                v-model:prop-selection-id="detailItem.cellTo.id"
                v-model:prop-selection-val="detailItem.cellTo.name"
                v-model:prop-placeholder="lng.label_cell_to"
                @onUpdateData="updateCellsDataTo">
            </autocomplete-input>
          </div>

          <div class="col-md-2">
            <input class="form-control"/>
            <div id="nameHelp" class="form-text fw-lighter fst-italic">количество</div>
          </div>

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

export default {
  name: "ProductMovementForm",
  components: {AutocompleteInput},
  props:{
    propModeForm: Number
  },
  data(){
    return{
      detailItem: {
        id: 0,
        name: "",
        item_number: "",
        cellFrom: {
          id: 0,
          name: ''
        },
        cellTo: {
          id: 0,
          name: ''
        }
      },
      productsSuggestion: [],
      cellsSuggestionFrom: [],
      cellsSuggestionTo: [],
      lng: {
        label_name: "Наименование",
        label_cell_from: "Откуда",
        label_cell_to: "Куда",
        label_item_number: "Артикул",
        label_tags: "Тэги",
        title_form_get: "Взять",
        title_form_put: "Положить",
        title_form_move: "Переместить",
        btn_form_close: "Закрыть",
        btn_form_store: "Сохранить",
      },
    }
  },
  methods:{

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

    updateCellsDataFrom(emitData){
      DataProvider.GetSuggestionReference('cells', emitData.val)
          .then((response) => {
            this.cellsSuggestionFrom = response.data.data
          })
          .catch(error => { DataProvider.ErrorProcessing(error) });
    },
    updateCellsDataTo(emitData){
      DataProvider.GetSuggestionReference('cells', emitData.val)
          .then((response) => {
            this.cellsSuggestionTo = response.data.data
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