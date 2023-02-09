<template>
  <div class="container px-4 py-5" id="icon-grid">
    <h2 class="pb-2 border-bottom">Приходные документы</h2>

    <div class="row row-cols-1 row-cols-sm-2 row-cols-md-3 row-cols-lg-4 g-4 py-5">
      <div class="col d-flex align-items-start">
        <div>
          <router-link to="/docs/receipt" class="text-decoration-none fw-bold fs-5">Поступление</router-link>
          <p>Поступление товаров от поставщика</p>
        </div>
      </div>
      <div class="col d-flex align-items-start">
        <div>
          <router-link to="/docs/posting" class="text-decoration-none fw-bold fs-5">Оприходование</router-link>
          <p>Принятие товаров на учет</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "ReceiptView"
}
</script>
