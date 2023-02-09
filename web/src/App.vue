<template>
  <div>
  <nav-bar />
  <div class="container-fluid">
    <div class="row">
      <!-- side-bar / -->
      <main class="h-100"> <!--px-md-4 ms-sm-auto col-lg-10  col-md-9 -->
        <router-view></router-view>
      </main>
    </div>


  </div>
  </div>
</template>
<script>
import NavBar from "@/components/NavBar";
//import SideBar from "@/components/SideBar";
export default {
  name: 'App',
  components: {NavBar},
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

.sidebar-heading {
  font-size: .75rem;
  text-transform: uppercase;
}

</style>
