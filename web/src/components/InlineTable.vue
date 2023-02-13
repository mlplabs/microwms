<template>
<div>
  <div class="btn-toolbar mb-2 mb-md-0 justify-content-between" style="padding-top:5px; padding-bottom: 1px;">
    <div class="btn-group me-2">
      <button type="button" class="btn btn-sm btn-outline-secondary" @click="$emit('new-item-clicked') ">
        <i class="bi bi-plus"></i>
      </button>
    </div>
    <div class="input-group">
      <input type="text" class="form-control" placeholder="поиск" v-show="isShowSearch">
    </div>
  </div>
    <!--div class="" style="overflow-x: initial;"-->
    <table class="table table-striped table-hover table-bordered" >
      <!-- table head -->
      <thead>
      <tr>
        <th scope="col" class="col_head"
          v-for="(col, index) in columns"
          :key="index"
          :class="{'col_id': col.isKey, col_action: col.field === 'actions'}"
         >{{ col.label }}</th>
      </tr>
      </thead>

      <!-- table body -->
      <template v-if="rows.length > 0">
        <tbody>
        <tr v-for="(row, j) in rows"  :key="j" @click="$emit('row-clicked', row)">

          <!-- set column css class -->
          <td v-for="(col, index) in columns" :key="index" :class="{'col_id': col.isKey, 'col_action': col.field === 'actions', 'text-start': col.align === 0, 'text-center': col.align === 1, 'text-end': col.align === 2 }">
            <template v-if="row.id === 0 && (col.isKey === false && col.field !== 'actions')">

              <!-- select - if values > 0 -->
              <select class="form-select" v-if="col.values !== undefined && col.values.length > 0">
                <option v-for="(op, h) in col.values" :key="h" value="{{op.key}}">{{op.val}}</option>
              </select>

              <!-- suggest - if suggestion > 0 -->
              <autocomplete-input v-else-if="col.suggestion === true"
                v-model:propSelection=row[col.field]
                v-model:propSuggestions="suggestionValue"
                v-model:prop-key="col.field"
                @onUpdateData="$emit('update-suggestion', $event)">
              </autocomplete-input>

              <!-- input integer if isNum is true $emit('update-suggestion') -->
              <input v-else-if="col.isNum" class="form-control text-end" type="number" v-model.number="row[col.field]" />

              <!--  input text -->
              <input v-else class="form-control" type="text" v-model="row[col.field]" />
            </template>

            <template v-else-if="col.field === 'actions'">
              <div class="dropdown">
                <button type="button" class="btn btn-sm btn-outline-secondary" data-bs-toggle="dropdown">
                  <i class="bi bi-three-dots-vertical"></i>
                </button>
                <ul class="dropdown-menu dropdown-menu-end">
                  <li><a class="dropdown-item" href="#" @click.prevent="$emit('row-delete', j)">Удалить {{j}}</a></li>
                  <li><a class="dropdown-item" href="#">Another action {{row.id}}</a></li>
                  <li><a class="dropdown-item" href="#">Something else here</a></li>
                </ul>
              </div>
            </template>

            <template v-else>
              <span v-if="col.values !== undefined && col.values.length > 0">{{getEnumValue(row, col)}}</span>
              <span v-else>{{row[col.field]}}</span>
            </template>

          </td>
        </tr>
        </tbody>
      </template>

    </table>

    <pagination-bar v-show="isShowPaging" v-bind:param-current-page=currentPage v-bind:param-count-rows=countRows v-bind:param-limit-rows=limitRows v-on:selectPage="$emit('page-selected', 1)"></pagination-bar>
  </div>
</template>

<script>
import {computed, reactive} from "vue";
import PaginationBar from "@/components/PaginationBar";
import AutocompleteInput from "@/components/AutocompleteInput";


export default {
  name: "InlineTable",
  components: {AutocompleteInput, PaginationBar},
  props:{
    columns: {
      type: Array,
      default: () => {
        return [];
      },
    },
    rows: {
      type: Array,
      default: () => {
        return [];
      },
    },
    isShowPaging: {
      type: Boolean,
      default: true,
    },
    isShowSearch:{
      type: Boolean,
      default: true,
    },
    suggestionData:{
      type: Array,
      default: () => {
        return [];
      },
    }

  },
  data(){
    return{
      currentPage: 1,
      countRows: 30,
      limitRows: 5,
    }
  },
  computed:{
    suggestionValue:{
      get(){
        return this.suggestionData
      },
    },
  },
  methods:{
    getEnumValue(row, col){
      let i = col.values.find(item => item.key === row[col.field])
      if (i===undefined ){
        return ''
      }
      return i.val
    },
  },
  setup(props){
    const setting = reactive({
      keyColumn: computed(() => {
        let key = "";
        Object.assign(props.columns).forEach((col) => {
          if (col.isKey) {
            key = col.field;
          }
        });
        return key;
      }),
      // current page number
      page: props.page,
      // Display count per page
      //pageSize: defaultPageSize.value,
      // Maximum number of pages
      maxPage: computed(() => {
        if (props.total <= 0) {
          return 0;
        }
        let maxPage = Math.floor(props.total / setting.pageSize);
        let mod = props.total % setting.pageSize;
        if (mod > 0) {
          maxPage++;
        }
        return maxPage;
      }),
      // The starting value of the page number
      offset: computed(() => {
        return (setting.page - 1) * setting.pageSize + 1;
      }),
      // Maximum number of pages
      limit: computed(() => {
        let limit = setting.page * setting.pageSize;
        return props.total >= limit ? limit : props.total;
      }),
      // Paging array
      paging: computed(() => {
        let startPage = setting.page - 2 <= 0 ? 1 : setting.page - 2;
        if (setting.maxPage - setting.page <= 2) {
          startPage = setting.maxPage - 4;
        }
        startPage = startPage <= 0 ? 1 : startPage;
        let pages = [];
        for (let i = startPage; i <= setting.maxPage; i++) {
          if (pages.length < 5) {
            pages.push(i);
          }
        }
        return pages;
      }),
    });
  }
}
</script>

<style scoped>
tr{
  cursor: pointer;
}
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
/* (-) стрелки input type number */
input::-webkit-outer-spin-button,
input::-webkit-inner-spin-button {
   -webkit-appearance: none;
   margin: 0;
}
input[type=number] {
  -moz-appearance: textfield;
}

</style>