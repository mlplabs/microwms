<template>
  <div>
  <nav-bar @on-global-search="onNavGlobalSearch"/>
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
import {ref, provide} from 'vue'
import NavBar from "@/components/NavBar";
//import SideBar from "@/components/SideBar";
const message = ref('')

export default {
  name: 'App',
  components: {NavBar},
  methods:{
    onNavGlobalSearch(emitData){
      message.value = emitData
    }
  },
  setup() {
    provide('global_search', message)
  }
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
