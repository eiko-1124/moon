<template>
  <div class="nav-header">
    <el-button type="primary" plain @click="putNav">开/关</el-button>
    <el-badge class="nav-message" is-dot>
      <el-button type="primary" plain @click="goJournal">消息</el-button>
    </el-badge>
    <div class="nav-option">
      <el-dropdown class="nav-user">
        <span class="el-dropdown-link">
          {{ userName }}
          <el-icon class="el-icon--right">
            <arrow-down />
          </el-icon>
        </span>
        <template #dropdown>
          <el-dropdown-menu>
            <el-dropdown-item @click="cancellation">注销</el-dropdown-item>
          </el-dropdown-menu>
        </template>
      </el-dropdown>
    </div>
  </div>
</template>

<script>
import { computed, onMounted } from "@vue/runtime-core";
import { useStore } from "vuex";
import { Base64 } from "js-base64";
import { useRouter } from "vue-router";
export default {
  setup() {
    const store = useStore();
    const router = useRouter();
    const getUserName = () => {
      if (store.state.User.name.length == 0) {
        let token;
        const cookies = document.cookie.split(";");
        for (let i = 0; i < cookies.length; i++) {
          const cookie = cookies[i].trim();
          if (cookie.indexOf("token=") == 0) {
            token = cookie.substring(6, cookie.length);
            break;
          }
        }
        const str = Base64.decode(token).split('"');
        for (let i = 0; i < str.length; i++) {
          if (str[i] == "name") {
            store.commit("setName", str[i + 2]);
            break;
          }
        }
      }
    };
    const cancellation = () => {
      document.cookie = "token=; expires=Thu, 01 Jan 1970 00:00:00 GMT";
      store.commit("setName", "");
      router.push("/Login");
    };
    const putNav = () => {
      store.commit("putNav");
    };
    const goJournal = () => {
      router.push("/Journal");
    };
    onMounted(() => {
      getUserName();
    });
    return {
      userName: computed(() => store.state.User.name),
      cancellation,
      putNav,
      goJournal,
    };
  },
};
</script>

<style lang="scss" scoped>
.nav-header {
  position: relative;
  display: flex;
  align-items: center;
  height: 100%;
  min-width: 23.5rem;
}
.nav-option {
  display: flex;
  align-items: center;
  position: absolute;
  right: 0rem;
}
.nav-message {
  margin-left: 1rem;
}
.nav-user {
  margin: 0rem 1.5rem;
  color: white;
}
.example-showcase .el-dropdown-link {
  cursor: pointer;
  display: flex;
  align-items: center;
}
</style>