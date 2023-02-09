<template>
  <nav aria-label="Page navigation">
    <ul class="pagination justify-content-end">
      <li v-bind:class="[CurrentPage === 1 ? disabledPageClass : '', defaultPageClass]">
        <a class="page-link" href="#" aria-label="Previous" @click.prevent.stop="selectPage(CurrentPage-1)">
          <span aria-hidden="true">&laquo;</span>
        </a>
      </li>

      <li v-for="idx in CountPages" :key="idx" v-bind:class="[idx === CurrentPage ? activePageClass : '', defaultPageClass]">
        <a class="page-link" href="#" @click.prevent="selectPage(idx)">{{idx}}</a>
      </li>

      <li v-bind:class="[CurrentPage === CountPages ? disabledPageClass : '', defaultPageClass]">
        <a class="page-link" href="#" aria-label="Next" @click.prevent.stop="selectPage(CurrentPage+1)">
          <span aria-hidden="true">&raquo;</span>
        </a>
      </li>
    </ul>
  </nav>
</template>

<script>
export default {
  name: "PaginationBar",
  props:{
    paramCurrentPage: Number,
    paramCountRows: Number,
    paramLimitRows: Number
  },
  data(){
    return{
      CurrentPage: this.paramCurrentPage,
      activePageClass: 'active',
      disabledPageClass: 'disabled',
      defaultPageClass: 'page-item'
    }
  },
  computed:{
    CountPages(){
      return  Math.ceil(this.paramCountRows / this.paramLimitRows)
    }
  },
  methods:{
    selectPage(page){
      this.CurrentPage = page
      this.$emit('selectPage', {page})
    }
  }

}
</script>

