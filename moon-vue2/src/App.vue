<template>
  <div id="app">
    <top-nav v-if="navFlag"></top-nav>
    <left-nav v-if="navFlag"></left-nav>
    <router-view class="view" />
    <patel v-if="navFlag"></patel>
    <load v-if="loadFlag"></load>
  </div>
</template>

<script>
import TopNav from "@/components/TopNav.vue";
import Patel from "@/components/Patel.vue";
import Load from "@/components/Load.vue";
import LeftNav from "@/components/LeftNav.vue";
export default {
  name: "App",
  components: { TopNav, Patel, Load, LeftNav },
  computed: {
    navFlag() {
      return this.$store.state.navFlag;
    },
    loadFlag() {
      return this.$store.state.loadFlag;
    },
  },
};
</script>

<style lang="scss">
.view {
  height: 100vh;
  width: 100vw;
  overflow: auto;
}

.view::-webkit-scrollbar {
  width: 5px;
}

.view::-webkit-scrollbar-button {
  display: none;
}

.view::-webkit-scrollbar-thumb {
  background-color: lightskyblue;
}
</style>
