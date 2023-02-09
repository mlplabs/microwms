<template>
  <div class="container px-4 py-5" id="icon-grid">
    <h2 class="pb-2 border-bottom">Настройки</h2>
    <router-link to="/hw/pref">Принтеры</router-link>
    <router-view></router-view>
  </div>
</template>

<script>
export default {
  name: "PropertiesView"
}
</script>

<style scoped>

</style>