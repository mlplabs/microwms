<template>
      <nav id="sidebarMenu" class="col-md-3 col-lg-2 d-md-block bg-light sidebar collapse">
        <div class="position-sticky pt-3">
          <ul class="nav flex-column">
            <li class="nav-item">
              <router-link to="/refs" class="nav-link active" aria-current="page">
                <i class="bi-journal-text" style="font-size: 1rem; color: cornflowerblue;margin-right: 1px;"></i>Справочники</router-link>
            </li>
            <li class="nav-item">
              <router-link to="/receipt" class="nav-link" aria-current="page">
                <i class="bi-journal-text" style="font-size: 1rem; color: cornflowerblue;margin-right: 1px;"></i>Поступления</router-link>
            </li>
            <li class="nav-item">
              <router-link to="/shipment" class="nav-link" aria-current="page">
                <i class="bi-journal-text" style="font-size: 1rem; color: cornflowerblue;margin-right: 1px;"></i>Отгрузки</router-link>
            </li>
            <li class="nav-item">
              <a class="nav-link" href="#"></a>
            </li>
            <li class="nav-item">
              <router-link to="/reports" class="nav-link" aria-current="page">
                <i class="bi-journal-text" style="font-size: 1rem; color: cornflowerblue;margin-right: 1px;"></i>Отчеты</router-link>
            </li>
            <li class="nav-item">
              <router-link to="/props" class="nav-link" aria-current="page">
                <i class="bi-journal-text" style="font-size: 1rem; color: cornflowerblue;margin-right: 1px;"></i>Настройки</router-link>
            </li>
          </ul>

          <h6 class="sidebar-heading d-flex justify-content-between align-items-center px-3 mt-4 mb-1 text-muted">
            <span>Saved reports</span>
            <a class="link-secondary" href="#" aria-label="Add a new report">
              <span data-feather="plus-circle"></span>
            </a>
          </h6>
          <ul class="nav flex-column mb-2">
            <li class="nav-item">
              <a class="nav-link" href="#">
                <i class="bi-file-earmark-spreadsheet" style="font-size: 1rem; color: cornflowerblue;margin-right: 1px;"></i>
                Current month
              </a>
            </li>
            <li class="nav-item">
              <a class="nav-link" href="#">
                <i class="bi-file-earmark-spreadsheet" style="font-size: 1rem; color: cornflowerblue;margin-right: 1px;"></i>
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
                <i class="bi-file-earmark-spreadsheet" style="font-size: 1rem; color: cornflowerblue;margin-right: 1px;"></i>
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
