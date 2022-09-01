<template>
  <div class="alive-nav">
    <el-scrollbar class="el-scrollbar">
      <div class="alive-nav-container">
        <label class="alive-name">{{ title }}</label>
        <el-tag
          class="tag-btn"
          size="large"
          closable
          v-for="tag in cache"
          :key="tag.name"
          @close="closeBnt(tag.name)"
          @click="routerBnt(tag.path)"
          >{{ tag.name }}</el-tag
        >
      </div>
    </el-scrollbar>
  </div>
</template>

<script>
import { computed } from "@vue/runtime-core";
import { useStore } from "vuex";
import { useRouter } from "vue-router";
export default {
  setup() {
    const store = useStore();
    const router = useRouter();
    const title = computed(() => store.state.User.viewAlive);
    const cache = computed(() => store.state.User.viewCache);
    const closeBnt = (name) => store.commit("delCache", name);
    const routerBnt = (path) => router.push(path);
    return {
      title,
      cache,
      closeBnt,
      routerBnt,
    };
  },
};
</script>

<style lang="scss" scoped>
.alive-nav {
  position: relative;
  display: flex;
  align-items: center;
  margin: 0rem 0rem 1rem;
  height: 3rem;
  background-color: white;
}
.el-scrollbar {
  background-color: white;
}
.el-scrollbar ::v-deep .el-scrollbar__view {
  height: 100%;
}
.alive-nav-container {
  padding: 0rem 1rem;
  height: 100%;
  display: flex;
  align-items: center;
}
.tag-btn {
  margin: 0rem 0.5rem;
  animation: site-load 1s;
  &:hover {
    cursor: pointer;
  }
}
.alive-name {
  display: flex;
  align-items: center;
  color: darkcyan;
  &::before {
    margin-right: 0.2rem;
    display: block;
    content: "";
    height: 0.1rem;
    width: 1rem;
    background-color: darkcyan;
  }

  &::after {
    margin-left: 0.2rem;
    display: block;
    content: "";
    height: 0.1rem;
    width: 1rem;
    background-color: darkcyan;
  }
}

@keyframes site-load {
  0% {
    opacity: 0;
    transform: scale(0, 0);
  }
  100% {
    opacity: 1;
    transform: scale(1, 1);
  }
}
</style>