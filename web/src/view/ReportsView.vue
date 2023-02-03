<template>
  <div class="container px-4 py-5" id="icon-grid">
    <h2 class="pb-2 border-bottom">Отчеты</h2>

    <div class="row row-cols-1 row-cols-sm-2 row-cols-md-3 row-cols-lg-4 g-4 py-5">
      <div class="col d-flex align-items-start">
        <i class="bi-journal-text" style="font-size: 1.5rem; color: cornflowerblue;margin-right: 5px;"></i>
        <div>
          <router-link to="/reports/stock"><h4 class="fw-bold mb-0">Остатки</h4></router-link>
          <p>Отчет о товарных запасах</p>
        </div>
      </div>
      <div class="col d-flex align-items-start">
        <i class="bi-journal-text" style="font-size: 1.5rem; color: cornflowerblue;margin-right: 5px;"></i>
        <div>
          <router-link to="/reports/receipt"><h4 class="fw-bold mb-0">Карточка товара</h4></router-link>
          <p>Информация о движениях товара</p>
        </div>
      </div>
      <div class="col d-flex align-items-start">
        <i class="bi-journal-text" style="font-size: 1.5rem; color: cornflowerblue;margin-right: 5px;"></i>
        <div>
          <router-link to="/reports/shipment"><h4 class="fw-bold mb-0">Отгрузки</h4></router-link>
          <p>Отчет об отгрузках</p>
        </div>
      </div>
      <div class="col d-flex align-items-start">
        <i class="bi-journal-text" style="font-size: 1.5rem; color: cornflowerblue;margin-right: 5px;"></i>
        <div>
          <router-link to="/reports/receipt"><h4 class="fw-bold mb-0">Поступления</h4></router-link>
          <p>Отчет о поступлениях</p>
        </div>
      </div>
    </div>
  </div>

</template>
<script>
export default {
  name: "ReportsView"
}
</script>
Stock report
Sales Report