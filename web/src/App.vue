<template>
  <div>
  <nav-bar />
  <div class="container-fluid">
    <div class="row">
      <side-bar />
      <main class="col-md-9 ms-sm-auto col-lg-10 px-md-4 h-100">
        <router-view></router-view>
      </main>
    </div>


  </div>
  </div>
</template>
<script>
import NavBar from "@/components/NavBar";
import SideBar from "@/components/SideBar";
export default {
  name: 'App',
  components: {SideBar, NavBar},
}
</script>

<style>
.sidebar {
  top: 0;
  bottom: 0;
  left: 0;
  z-index: 100; /* Behind the navbar */
  padding: 48px 0 0; /* Height of navbar */
  box-shadow: inset -1px 0 0 rgba(0, 0, 0, .1);
}
@media (max-width: 767.98px) {
  .sidebar {
    top: 5rem;
  }
}

.sidebar .nav-link {
  font-weight: 500;
  color: #333;
}

.sidebar .nav-link .feather {
  margin-right: 4px;
  color: #727272;
}

.sidebar .nav-link.active {
  color: #2470dc;
}

.sidebar .nav-link:hover .feather,
.sidebar .nav-link.active .feather {
  color: inherit;
}

.sidebar-heading {
  font-size: .75rem;
  text-transform: uppercase;
}

</style>
