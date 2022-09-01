<template>
  <el-collapse class="el-collapse" modelValue="activeName" accordion>
    <el-collapse-item
      class="el-collapse-item"
      :class="{ 'nav-select': aIndex == activeName }"
      :name="aIndex"
      v-for="(nav, aIndex) in menu"
      :key="aIndex"
    >
      <template #title>
        <el-icon class="nav-title-icon"
          ><component :is="nav.icon"></component
        ></el-icon>
        {{ nav.mainNav }}
      </template>
      <div v-for="(subNav, bIndex) in nav.subNav" :key="bIndex">
        <button
          class="nav-sub-bnt"
          :class="{ 'bnt-select': bIndex == activeBtn && aIndex == activeName }"
          @click="menuClick(aIndex, bIndex)"
        >
          <router-link :to="subNav.path">{{ subNav.title }}</router-link>
        </button>
      </div>
    </el-collapse-item>
  </el-collapse>
</template>

<script>
import { computed, onMounted } from "vue";
import { menu, routerIndex } from "@/conf/nav.js";
import { useRoute } from "vue-router";
import { useStore } from "vuex";
export default {
  setup() {
    const route = useRoute();
    const store = useStore();
    const activeName = computed(() => store.state.User.activeName);
    const activeBtn = computed(() => store.state.User.activeBtn);
    const menuClick = (aIndex, bIndex) => {
      store.commit("setActiveName", aIndex);
      store.commit("setActiveBtn", bIndex);
    };
    onMounted(() => {
      store.commit("setActiveName", routerIndex[route.name][0]);
      store.commit("setActiveBtn", routerIndex[route.name][1]);
    });
    return {
      activeName,
      activeBtn,
      menuClick,
      menu,
    };
  },
};
</script>

<style lang="scss" scoped>
.el-collapse {
  padding: 0rem 1rem;
  background-color: transparent;
  border: none;
}
.el-collapse-item {
  background-color: transparent;
}
.el-collapse-item ::v-deep .el-collapse-item__header {
  background-color: transparent;
  border: none;
  color: white;
  font-size: 1rem;
  letter-spacing: 0.1rem;
}
.el-collapse-item ::v-deep .el-collapse-item__wrap {
  background-color: transparent;
  border: none;
}
.el-collapse-item ::v-deep .el-collapse-item__content {
  padding: 0rem;
}
.nav-title-icon {
  margin: 0rem 0.5rem;
}
.nav-sub-bnt {
  margin-left: 1.5rem;
  background-color: transparent;
  border: none;
  letter-spacing: 0.2rem;
  a {
    color: white;
    text-decoration: none;
  }
}
.el-collapse-item ::v-deep .el-collapse-item__header,
.nav-sub-bnt {
  &:hover {
    color: lightblue;
    a {
      color: lightblue;
    }
  }
}
.nav-select {
  & ::v-deep .el-collapse-item__header {
    color: skyblue;
  }
}
.bnt-select {
  a {
    color: skyblue;
  }
}
</style>