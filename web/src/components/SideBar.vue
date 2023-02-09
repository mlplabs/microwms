<template>
      <nav id="sidebarMenu" class="col-md-3 col-lg-2 d-md-block sidebar collapse">
        <div class="position-sticky pt-3">
          <ul class="nav flex-column">
            <li class="nav-item">
              <router-link to="/refs" class="nav-link fw-bold fs-5">
                Справочники
              </router-link>
            </li>
            <li class="nav-item">
              <router-link to="/receipt" class="nav-link fw-bold fs-5" aria-current="page">
                Поступления</router-link>
            </li>
            <li class="nav-item">
              <router-link to="/shipment" class="nav-link fw-bold fs-5" aria-current="page">
                Отгрузки</router-link>
            </li>
            <li class="nav-item">
              <a class="nav-link fw-bold fs-5" href="#"></a>
            </li>
            <li class="nav-item">
              <router-link to="/reports" class="nav-link fw-bold fs-5" aria-current="page">
                Отчеты</router-link>
            </li>
            <li class="nav-item">
              <router-link to="/props" class="nav-link fw-bold fs-5" aria-current="page">
                Настройки</router-link>
            </li>
          </ul>

          <h6 class="sidebar-heading d-flex justify-content-between align-items-center px-3 mt-4 mb-1">
            <span>Saved reports</span>
            <a class="link-secondary" href="#" aria-label="Add a new report">
              <span data-feather="plus-circle"></span>
            </a>
          </h6>
          <ul class="nav flex-column mb-2">
            <li class="nav-item">
              <a class="nav-link" href="#">
                Current month
              </a>
            </li>
            <li class="nav-item">
              <a class="nav-link" href="#">
                Last quarter
              </a>
            </li>
            <li class="nav-item">
              <a class="nav-link" href="#">
                <i class="bi-file-earmark-spreadsheet" style="font-size: 1rem; color: cornflowerblue;margin-right: 1px;"></i>
                Social engagement
              </a>
            </li>
            <li class="nav-item">
              <a class="nav-link" href="#">
                Year-end sale
              </a>
            </li>
          </ul>
        </div>
      </nav>
</template>

<script>
export default {
  name: "SideBar"
}
</script>
