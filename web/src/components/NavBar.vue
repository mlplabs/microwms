<template>
  <header class="navbar navbar-expand-lg" aria-label="Fifth navbar example navbar-dark bg-dark">
    <div class="container-fluid">
      <a class="navbar-brand fw-bold fs-4" href="#"><!-- img src="assets/logo2.png" class="border-primary" style="width: 46px; right: 46px; margin-right: 3px;"-->MicroWMS</a>
      <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarsToggler" aria-controls="" aria-expanded="false" aria-label="Toggle navigation">
        <span class="navbar-toggler-icon"></span>
      </button>

      <div class="collapse navbar-collapse" id="navbarsToggler">
        <ul class="navbar-nav me-auto mb-2 mb-lg-0">
          <!--
          <li class="nav-item">
            <a class="nav-link active nav-link fw-bold" aria-current="page" href="#">Home</a>
          </li>
          <li class="nav-item">
            <a class="nav-link" href="#">Link</a>
          </li>
          <li class="nav-item">
            <a class="nav-link disabled">Disabled</a>
          </li>
          -->
          <li class="nav-item dropdown">
            <a class="nav-link dropdown-toggle" href="#" id="dropdown05" data-bs-toggle="dropdown" aria-expanded="false">{{ lng.label_navigation }}</a>
            <ul class="dropdown-menu" aria-labelledby="dropdown05">
              <li><router-link to="/refs" class="dropdown-item">Справочники</router-link></li>
              <!-- li><router-link to="/docs/receipt" class="dropdown-item">Поступления</router-link></li>
              <li><router-link to="/docs/shipment" class="dropdown-item">Отгрузки</router-link></li -->
              <li class="dropdown-divider"></li>
              <li><router-link class="dropdown-item" to="/reports/stock" href="#">Остатки товаров</router-link></li>
              <!-- li><router-link class="dropdown-item" to="/reports/history/0" href="#">Движение товара</router-link></li -->
              <li class="dropdown-divider"></li>
              <li><router-link to="/props" class="dropdown-item">Настройки</router-link></li>
              <!--
              <li><a class="dropdown-item" href="#">Action</a></li>
              <li><a class="dropdown-item" href="#">Another action</a></li>
              <li><a class="dropdown-item" href="#">Something else here</a></li>
              -->
            </ul>
          </li>
        </ul>
        <form>
          <input class="form-control" type="text" :placeholder="lng.label_search" aria-label="Search" @input="onInputSearch">
        </form>
      </div>
    </div>
  </header>
</template>

<script>
export default {
  name: 'NavBar',
  data() {
    return {
      lng:{
        label_navigation: "Навигация",
        label_search: "Поиск"
      }
    }
  },
  methods: {
    onInputSearch(event){
      //this.$root.$emit('onGlobalSearch', event.target.values)
      this.$emit('onGlobalSearch', event.target.value)
    }
  }
}
</script>