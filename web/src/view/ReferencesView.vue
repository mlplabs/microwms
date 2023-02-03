<template>
  <div class="container px-4 py-5" id="icon-grid">
    <h2 class="pb-2 border-bottom">Справочники</h2>

    <div class="row row-cols-1 row-cols-sm-2 row-cols-md-3 row-cols-lg-4 g-4 py-5">
      <div class="col d-flex align-items-start">
        <i class="bi-journal-text" style="font-size: 1.5rem; color: cornflowerblue;margin-right: 5px;"></i>
        <div>
          <router-link to="/refs/products"><h4 class="fw-bold mb-0">Товары</h4></router-link>
          <p>Товарный справочник</p>
        </div>
      </div>
      <div class="col d-flex align-items-start">
        <i class="bi-journal-text" style="font-size: 1.5rem; color: cornflowerblue;margin-right: 5px;"></i>
        <div>
          <router-link to="/refs/manufacturers"><h4 class="fw-bold mb-0">Производители</h4></router-link>
          <p>Производители товаров</p>
        </div>
      </div>
      <div class="col d-flex align-items-start">
        <i class="bi-journal-text" style="font-size: 1.5rem; color: cornflowerblue;margin-right: 5px;"></i>
        <div>
          <router-link to="/refs/barcodes"><h4 class="fw-bold mb-0">Штрих-коды</h4></router-link>
          <p>Штрих-коды товаров</p>
        </div>
      </div>
      <div class="col d-flex align-items-start">
        <i class="bi-journal-text" style="font-size: 1.5rem; color: cornflowerblue;margin-right: 5px;"></i>
        <div>
          <router-link to="/refs/products" para><h4 class="fw-bold mb-0">Еще один справочник</h4></router-link>
          <p></p>
        </div>
      </div>
      <div class="col d-flex align-items-start">
        <i class="bi-journal-text" style="font-size: 1.5rem; color: cornflowerblue;margin-right: 5px;"></i>
        <div>
          <a href="#"><h4 class="fw-bold mb-0">Склады</h4></a>
          <p>Склады, места хранения</p>
        </div>
      </div>
      <div class="col d-flex align-items-start">
        <i class="bi-journal-text" style="font-size: 1.5rem; color: cornflowerblue;margin-right: 5px;"></i>
        <div>
          <a href="#"><h4 class="fw-bold mb-0">Зоны</h4></a>
          <p>Зоны складв</p>
        </div>
      </div>
      <div class="col d-flex align-items-start">
        <i class="bi-journal-text" style="font-size: 1.5rem; color: cornflowerblue;margin-right: 5px;"></i>
        <div>
          <a href="#"><h4 class="fw-bold mb-0">Ячейки</h4></a>
          <p>Ячейки склада</p>
        </div>
      </div>
      <div class="col d-flex align-items-start">
        <i class="bi-journal-text" style="font-size: 1.5rem; color: cornflowerblue;margin-right: 5px;"></i>
        <div>
          <a href="#"><h4 class="fw-bold mb-0">Пользователи</h4></a>
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