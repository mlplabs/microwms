<template>
  <div class="container px-4 py-5" id="icon-grid">
    <h2 class="pb-2 border-bottom">Принтеры</h2>
    <form class="row g-3">
      <div class="col-12">
        <label for="inputPrinterDoc" class="form-label">Принтер для документов</label>
        <select id="inputPrinterDoc" class="form-select" aria-label="Default select example">
          <option selected>Выберите принтер...</option>
          <option v-for="(printer, index) in printersList" :key="index">{{printer.name}}</option>
        </select>
      </div>
      <div class="col-12">
        <label for="inputPrinterLabel" class="form-label">Принтер этикеток</label>
        <select id="inputPrinterLabel" class="form-select" aria-label="Default select example">
          <option selected>Выберите принтер...</option>
          <option v-for="(printer, index) in printersList" :key="index">{{printer.name}}</option>
        </select>
      </div>
    </form>
  </div>
</template>

<script>
import DataProvider from "@/services/DataProvider";

export default {
  name: "HardwarePrinters",
  data(){
    return {
      printersList: [],
      printerDoc: '',
      printerLabel: '',
    }
  },
  methods:{

    updatePrintersData(emitData){
      DataProvider.GetHwPrinters(emitData.val)
        .then((response) => {
          this.printersList = response.data
        })
        .catch(error => { this.errorProc(error) });
    },

    getPrintersData(){
      DataProvider.GetHwPrinters()
        .then((response) => {
          this.printersList = response.data
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
    this.getPrintersData()
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