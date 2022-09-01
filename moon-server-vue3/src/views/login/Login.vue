<template>
  <div class="login">
    <article>
      <h2>LOGIN</h2>
      <section>
        <label>账号</label>
        <el-input
          v-model="id"
          placeholder="Please input"
          maxlength="20"
          clearable
        />
      </section>
      <section>
        <label>密码</label>
        <el-input
          v-model="pswd"
          type="password"
          placeholder="Please input password"
          maxlength="20"
          show-password
          clearable
        />
      </section>
      <section>
        <el-button type="primary" text bg @click="login">登录</el-button>
      </section>
    </article>
  </div>
</template>

<script>
import { reactive, toRefs } from "@vue/reactivity";
import { goLogin } from "@/api/apis/login";
import { useRouter } from "vue-router";
export default {
  setup() {
    let data = reactive({
      id: "",
      pswd: "",
    });
    const router = useRouter();
    const login = () => {
      const form = new FormData();
      form.append("id", data.id);
      form.append("pswd", data.pswd);
      goLogin(form, loginSuccess, loginError);
    };
    const loginSuccess = (resData) => {
      ElMessage("登录成功： " + resData.res);
      let date = new Date();
      date.setTime(date.getTime() + 7 * 24 * 60 * 60 * 1000);
      let expires = "expires=" + date.toGMTString();
      document.cookie = "token=" + resData.token + ";" + expires;
      setTimeout(() => {
        router.push("/Dashboard");
      }, 1000);
    };
    const loginError = (msg) => {
      ElMessage("登录失败： " + msg);
    };
    return {
      ...toRefs(data),
      login,
    };
  },
};
</script>

<style lang="scss" scoped>
.login {
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 1rem;
  height: 100vh;
  width: 100vw;
  overflow: hidden;
  background-image: linear-gradient(120deg, lightblue, darkcyan);
}
article {
  padding: 1rem;
  height: 16rem;
  width: 30rem;
  background-color: rgba($color: white, $alpha: 0.6);
  border-radius: 0.5rem;
  overflow: hidden;
  h2 {
    margin-bottom: 1rem;
    text-align: center;
    color: darkcyan;
  }
  section {
    padding: 0.5rem 1rem;
    display: flex;
    align-items: center;
    justify-content: center;
    button {
      margin: 1rem;
      color: darkcyan;
      width: 6rem;
    }
  }
  label {
    margin-right: 1rem;
    color: darkcyan;
    flex-shrink: 0;
  }
}
</style>