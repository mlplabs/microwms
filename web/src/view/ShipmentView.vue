<template>
  <div class="container px-4 py-5" id="icon-grid">
    <h2 class="pb-2 border-bottom">Расходные документы</h2>

    <div class="row row-cols-1 row-cols-sm-2 row-cols-md-3 row-cols-lg-4 g-4 py-5">
      <div class="col d-flex align-items-start">
        <i class="bi-journal-text" style="font-size: 1.5rem; color: cornflowerblue;margin-right: 5px;"></i>
        <div>
          <router-link to="/docs/shipment"><h4 class="fw-bold mb-0">Отгрузка</h4></router-link>
          <p>Paragraph of text beneath the heading to explain the heading.</p>
        </div>
      </div>
      <div class="col d-flex align-items-start">
        <i class="bi-journal-text" style="font-size: 1.5rem; color: cornflowerblue;margin-right: 5px;"></i>
        <div>
          <router-link to="/docs/write-off"><h4 class="fw-bold mb-0">Списание</h4></router-link>
          <p>Paragraph of text beneath the heading to explain the heading.</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "ShipmentView"
}
</script>
