<template>
  <div>
  <input class="form-control" type="text" :id="propInputId" v-model="textValue" :readonly="propReadonly" :placeholder="propPlaceholder"
         @keydown.enter = 'onEnter'
         @keydown.down = 'onKeyDown'
         @keydown.up = 'onKeyUp'
         @click = 'onClick'
         @input = 'onChange'
  />
  <div class="dropdown" style="position:relative; min-width: inherit">
    <ul class="dropdown-menu" v-bind:class="{'show':openSuggestion}">
      <li class="dropdown-item" v-for="(item, index) in matches" :key="index" v-bind:class="{'active': isActive(index)}" @click="onSelect(index)" style="padding: 5px;">
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
      eventTargetValue: ""
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
    },
    propReadonly:{
      type: Boolean
    },
    propInputId:{
      type: String
    },
    propPlaceholder:{
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
      return this.propSuggestions
      // for local search
      //return this.propSuggestions.filter((str) => {
      //  return str.val.indexOf(this.textValue) >= 0;
      //});
    },
    openSuggestion() {
      return this.textValue !== "" &&
        this.matches.length !== 0 &&
        this.open === true;
    },
  },

  methods: {
    // Выбор с клавиатуры элемента из списка подсказок
    onEnter() {
      if (this.open) {
        this.textValue = this.matches[this.current].val;
        this.idValue = this.matches[this.current].id;
        this.$emit('onSelectData', { val:this.matches[this.current].val, key:this.matches[this.current].id } )
      }
      this.open = false;
    },
    // Выбор мышью элемента из списка подсказок
    onSelect(index) {
      if (this.open) {
        this.textValue = this.matches[index].val;
        this.idValue = this.matches[index].id;
        this.$emit('onSelectData', { val:this.matches[index].val, key:this.matches[index].id } )
      }
      this.open = false;
    },

    // Перемещение по списку подсказок
    onKeyUp() {
      console.log('key up')
      if(this.current > 0)
        this.current--;
    },
    onKeyDown() {
      console.log('key down')
      if(this.current < this.matches.length - 1)
        this.current++;
    },

    onChange() {
       if (this.open === false) {
         this.open = true;
         this.current = 0;
       }
      if (this.textValue === ''){
        return
      }
      this.$emit('onUpdateData', {
            val:this.textValue,
            key:this.propKey
          }
      )
    },
    onClick(){
      if (this.open){
        this.open = false
      }
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