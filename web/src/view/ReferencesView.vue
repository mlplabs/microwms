<template>
  <div class="container px-4 py-5" id="icon-grid">
    <h2 class="pb-2 border-bottom">Справочники</h2>

    <div class="row row-cols-1 row-cols-sm-2 row-cols-md-3 row-cols-lg-4 g-4 py-5">
      <div class="col d-flex align-items-start">
        <div>
          <router-link to="/refs/products" class="text-decoration-none fw-bold fs-5">Товары</router-link>
          <p>Товарный справочник</p>
        </div>
      </div>
      <div class="col d-flex align-items-start">
        <div>
          <router-link to="/refs/manufacturers" class="text-decoration-none fw-bold fs-5">Производители</router-link>
          <p>Производители товаров</p>
        </div>
      </div>
      <div class="col d-flex align-items-start">
        <div>
          <router-link to="/refs/barcodes" class="text-decoration-none fw-bold fs-5">Штрих-коды</router-link>
          <p>Штрих-коды товаров</p>
        </div>
      </div>
      <div class="col d-flex align-items-start">
        <div>
          <router-link to="/refs/whs" class="text-decoration-none fw-bold fs-5">Склады</router-link>
          <p>Склады, места хранения</p>
        </div>
      </div>
      <div class="col d-flex align-items-start">
        <div>
          <router-link to="#" class="text-decoration-none fw-bold fs-5">Ячейки</router-link>
          <p>Ячейки склада</p>
        </div>
      </div>
      <div class="col d-flex align-items-start">
        <div>
          <router-link to="/refs/users" class="text-decoration-none fw-bold fs-5">Пользователи</router-link>
          <p>Пользователи системы</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default{
  name: 'ReferencesView'
}
</script>