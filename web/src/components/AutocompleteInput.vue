<template>
  <div class="dropdown" style="position:relative">
    <input class="form-control" type="text" v-model="textValue"
           @keydown.enter = 'enter'
           @keydown.down = 'down'
           @keydown.up = 'up'
           @input = 'change'
           @click = 'click'
           @keydown = 'keypress'
    />
    <ul class="dropdown-menu" style="width:100%" v-bind:class="{'show':openSuggestion}">
      <li class="dropdown-item" v-for="(suggestion, index) in matches" :key="index" v-bind:class="{'active': isActive(index)}" @click="selectSuggestion(index)" style="padding: 5px;">
        <a href="#" @click.prevent>{{ suggestion }}</a>
      </li>
    </ul>
  </div>
</template>


<script>
/*
  AutocompleteInput component
  example:
    <autocomplete-input v-model:prop-suggestions="dataArray" v-model:prop-selection="value" @onUpdateData="eventArrayUpdate"></autocomplete-input>
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
    propSelection: {
      type: String,
    },
    propKey:{
      type: String
    }

  },
  computed: {
    textValue:{
      get(){
        return this.propSelection
      },
      set(value){
        this.$emit('update:propSelection', value)
      }
    },
    matches() {
      return this.propSuggestions.filter((str) => {
        return str.indexOf(this.textValue) >= 0;
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
      if (this.open)
        this.textValue = this.matches[this.current];
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
      console.log(val)
      if (val === ''){
        console.log('text val is empty ' + val)
        return
      }
      this.$emit('onUpdateData',
        {
          val:val,
          key:this.propKey
        }
      )
    },
    selectSuggestion(index) {
      console.log(this.matches[index])
      this.textValue = this.matches[index];
      this.open = false;
    },
    isActive(index) {
      return index === this.current;
    },

  },
}
</script>

<style scoped>
.dropdown-item a {
  text-decoration: none !important;
}
.dropdown-menu > .active > a,
.dropdown-menu > .active > a:hover,
.dropdown-menu > .active > a:focus {
  color: #FFFFFF;
}
</style>