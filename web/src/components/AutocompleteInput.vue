<template>
  <div>
  <input class="form-control" type="text" v-model="textValue"
         @keydown.enter = 'enter'
         @keydown.down = 'down'
         @keydown.up = 'up'
         @input = 'change'
         @click = 'click'
         @keydown = 'keypress'
  />
  <div class="dropdown" style="position:relative; min-width: inherit">
    <ul class="dropdown-menu" v-bind:class="{'show':openSuggestion}">
      <li class="dropdown-item" v-for="(item, index) in matches" :key="index" v-bind:class="{'active': isActive(index)}" @click="selectSuggestion(index)" style="padding: 5px;">
        <a href="#" @click.prevent>{{ item.val }}</a>
      </li>
    </ul>
  </div>
  </div>
</template>


<script>
/*
  AutocompleteInput component
  example:
    <autocomplete-input
      v-model:prop-suggestions="dataArray"  // массив значений  [{id: 1 val: text 1}, {id: 3 val: text 3}]
      v-model:prop-selection-val="value_id" // значение выбора, ключ
      v-model:prop-selection-val="value_str"  // значение выбора, строка
      @onUpdateData="eventArrayUpdate"> // событие родителю о введенном тексте для обновления dataArray
    </autocomplete-input>
*/

export default {
  name: "AutocompleteInput",
  data() {
    return {
      open: false,
      current: 0,
    }
  },

  props: {
    propSuggestions: {
      type: Array,
      required: true
    },
    propSelectionId:{
      type: Number,
    },
    propSelectionVal:{
      type: String,
    },
    propKey:{
      type: String
    }

  },
  computed: {
    textValue:{
      get(){
        return this.propSelectionVal
      },
      set(value){
        this.$emit('update:propSelectionVal', value)
      }
    },
    idValue:{
      get(){
        return this.propSelectionId
      },
      set(value){
        this.$emit('update:propSelectionId', value)
      }
    },
    matches() {
      return this.propSuggestions.filter((str) => {
        return str.val.indexOf(this.textValue) >= 0;
      });
    },
    openSuggestion() {
      return this.textValue !== "" &&
        this.matches.length !== 0 &&
        this.open === true;
    }
  },

  methods: {
    enter() {
      if (this.open) {
        this.textValue = this.matches[this.current].val;
        this.idValue = this.matches[this.current].id;
      }
      this.open = false;
    },
    up() {
      if(this.current > 0)
        this.current--;
    },
    down() {
      if(this.current < this.matches.length - 1)
        this.current++;
    },
    change() {
      if (this.open === false) {
        this.open = true;
        this.current = 0;
      }
    },
    click(){
      if (this.open){
        this.open = false
      }
    },
    keypress(event){
      const val = event.target.value
      console.log(this.matches)
      if (val === ''){
        console.log('text val is empty ' + val)
        return
      }
      this.$emit('onUpdateData', {
          val:val,
          key:this.propKey
        }
      )
    },
    selectSuggestion(index) {
      this.textValue = this.matches[index].val;
      this.idValue = this.matches[index].id;
      this.open = false;
    },
    isActive(index) {
      return index === this.current;
    },

  },
}
</script>

<style scoped>
/*
.dropdown-item a {
  text-decoration: none !important;
}
.dropdown-menu > .active > a,
.dropdown-menu > .active > a:hover,
.dropdown-menu > .active > a:focus {
  color: #FFFFFF;
}
 */
</style>