<template>
  <div class="common-layout">
    <el-container class="el-container">
      <el-aside class="el-aside" :class="{ 'aside-status': navStatus }">
        <div class="aside-header">
          <el-icon class="moon-night" :size="30"><MoonNight /></el-icon>
          <h3>Moon-博客管理</h3>
        </div>
        <el-divider border-style="solid" style="margin: 0rem" />
        <nav-menu class="aside-nav"></nav-menu>
      </el-aside>
      <el-container class="el-main-container">
        <el-header class="el-header">
          <nav-header></nav-header>
        </el-header>
        <el-scrollbar class="el-scrollbar">
          <el-main class="el-main">
            <el-scrollbar>
              <nav-alive></nav-alive>
              <router-view />
            </el-scrollbar>
          </el-main>
        </el-scrollbar>
      </el-container>
    </el-container>
  </div>
</template>

<script>
import NavMenu from "@/components/nav/NavMenu.vue";
import NavHeader from "@/components/nav/NavHeader.vue";
import NavAlive from "@/components/nav/NavAlive.vue";
import { useStore } from "vuex";
import { computed } from "@vue/runtime-core";
export default {
  components: { NavMenu, NavHeader, NavAlive },
  setup() {
    const store = useStore();
    return {
      navStatus: computed(() => store.state.User.navStatus),
    };
  },
};
</script>

<style lang="scss">
.common-layout {
  position: relative;
  height: 100vh;
  width: 100vw;
}
.el-container {
  height: 100%;
}
.el-aside {
  background-color: #606266;
  width: 16rem;
  height: 100%;
  transition-duration: 0.3s;
}
.el-scrollbar {
  background-color: #f2f3f5;
}
.el-header {
  height: 3rem;
  background-color: darkcyan;
}
.aside-header {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 3rem;
  width: 16rem;
  background-color: #303133;
  h3 {
    color: darkcyan;
    font-weight: 500;
    letter-spacing: 0.1rem;
  }
  .moon-night {
    margin-right: 0.5rem;
    color: white;
  }
}

.aside-nav {
  width: 16rem;
}

.aside-status {
  width: 0rem;
}

@media screen and (max-width: 960px) {
  .el-aside {
    width: 0rem;
  }
  .aside-status {
    width: 16rem;
  }
}
</style>
