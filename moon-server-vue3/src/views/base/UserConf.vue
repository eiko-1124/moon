<template>
  <div class="conf-container">
    <section>
      <div class="conf-info">
        <div class="conf-info-main">
          <div class="conf-info-portrait"></div>
          <div class="conf-info-name">
            <h2>{{ name }}</h2>
            <p>2000-11-24</p>
          </div>
        </div>
        <el-divider />
        <div class="conf-info-message">
          <label>{{ "文章" + sum.article }}</label>
          <el-divider direction="vertical" />
          <label>{{ "心情" + sum.mood }}</label>
          <el-divider direction="vertical" />
          <label>{{ "友联" + sum.link }}</label>
        </div>
        <span class="conf-icon"></span>
      </div>
      <div class="conf-link">
        <h3>我的友联</h3>
        <ul>
          <el-scrollbar class="el-scrollbar">
            <li v-for="(link, index) in links" :key="index">
              <div>
                <el-input
                  :modelValue="link.name"
                  @update:modelValue="linkEvent($event, index, 'name')"
                  class="link-name"
                  placeholder="备注"
                  clearable
                />
                <el-button type="primary" @click="setLink(index)">
                  保存
                </el-button>
                <el-button type="danger" @click="delLink(index)"
                  >删除</el-button
                >
              </div>
              <el-input
                :modelValue="link.flink"
                @update:modelValue="linkEvent($event, index, 'flink')"
                placeholder="链接"
                clearable
              />
            </li>
            <li>
              <el-button
                style="width: 100%"
                type="primary"
                plain
                @click="addFlink()"
                >新增</el-button
              >
            </li>
          </el-scrollbar>
        </ul>
        <span class="conf-icon"></span>
      </div>
    </section>
    <section class="conf-sign">
      <el-scrollbar class="scrollbar conf-sign-box">
        <h3>网页信息</h3>
        <div class="conf-signs">
          <h4>我的首页</h4>
          <el-input
            placeholder="Please input"
            style="margin-top: 1rem"
            v-model="sentence"
          >
            <template #prepend>短句</template>
            <template #append
              ><el-button @click="setInfo('sentence', sentence)"
                >保存</el-button
              ></template
            >
          </el-input>
          <el-divider />
          <h4>签名展示</h4>
          <el-input
            v-for="(sign, index) in signs"
            :key="index"
            :modelValue="sign.content"
            placeholder="Please input"
            style="margin-top: 1rem"
            @update:modelValue="signEvent($event, index)"
          >
            <template #prepend>签名{{ 1 }}</template>
            <template #append
              ><el-button @click="setSign(index, sign.content)"
                >保存</el-button
              ></template
            >
          </el-input>
          <el-divider />
          <h4>站点公告</h4>
          <el-input
            v-model="notice"
            :autosize="{ minRows: 3, maxRows: 5 }"
            type="textarea"
            placeholder="Please input"
            style="margin-top: 1rem"
          />
          <el-button
            style="width: 100%; margin-top: 1rem"
            type="primary"
            plain
            @click="setInfo('notice', notice)"
            >保存</el-button
          >
          <el-divider />
          <h4>关于本站</h4>
          <el-input
            v-model="illustate"
            :autosize="{ minRows: 3, maxRows: 5 }"
            type="textarea"
            placeholder="Please input"
            style="margin-top: 1rem"
          />
          <el-button
            style="width: 100%; margin-top: 1rem"
            type="primary"
            plain
            @click="setInfo('illustate', illustate)"
            >保存</el-button
          >
          <el-divider />
          <h4>其他信息</h4>
          <el-input
            v-model="qq"
            placeholder="Please input"
            style="margin-top: 1rem"
          >
            <template #prepend>QQ</template>
            <template #append
              ><el-button @click="setInfo('qq', qq)">保存</el-button></template
            >
          </el-input>
          <el-input
            v-model="github"
            placeholder="Please input"
            style="margin-top: 1rem"
          >
            <template #prepend>github</template>
            <template #append
              ><el-button @click="setInfo('github', github)"
                >保存</el-button
              ></template
            >
          </el-input>
        </div>
      </el-scrollbar>
      <span class="conf-icon"></span>
    </section>
  </div>
</template>

<script>
import { useStore } from "vuex";
import { computed } from "vue";
import {
  updateLink,
  deleteLink,
  updateInfo,
  updateSign,
} from "@/api/apis/user";
export default {
  setup() {
    const store = useStore();
    store.dispatch("setInfo");
    store.dispatch("setSigns");
    store.dispatch("setLinks");
    store.dispatch("getSum");
    let [name, notice, illustate, sentence, qq, github] = [
      "name",
      "notice",
      "illustate",
      "sentence",
      "qq",
      "github",
    ].map((key) =>
      computed({
        get: () => store.state.User.info[key],
        set: (value) => store.commit("setInfoMessage", [key, value]),
      })
    );
    const signs = computed(() => store.state.User.signs);
    const signEvent = (event, index) => {
      store.commit("setSign", [event, index]);
    };
    const links = computed(() => store.state.User.links);
    const linkEvent = (event, index, key) => {
      store.commit("setLink", [event, index, key]);
    };
    const sum = computed(() => store.state.User.sum);
    const addFlink = () => {
      store.commit("addLink");
    };
    const addSign = () => {
      store.commit("addSign");
    };
    const setLink = (index) => {
      const form = new FormData();
      form.append("id", index + 1);
      form.append("name", links.value[index].name);
      form.append("link", links.value[index].flink);
      updateLink(
        form,
        () => {
          ElMessage("修改成功");
        },
        (Err) => {
          console.log(Err);
        }
      );
    };
    const delLink = (index) => {
      console.log(index);
      if (index < links.value.length) {
        const form = new FormData();
        form.append("id", index + 1);
        deleteLink(
          form,
          () => {
            ElMessage("删除成功");
            links.value.splice(index, 1);
          },
          () => {
            console.log(Err);
          }
        );
      } else {
        ElMessage("删除成功");
        links.value.splice(index, 1);
      }
    };
    const setSign = (index, sign) => {
      const form = new FormData();
      form.append("id", index + 1);
      form.append("content", sign);
      updateSign(
        form,
        () => {
          ElMessage("修改成功");
        },
        () => {
          ElMessage("修改失败");
        }
      );
    };
    const setInfo = (key, value) => {
      const form = new FormData();
      form.append("key", key);
      form.append("value", value);
      updateInfo(
        form,
        () => {
          ElMessage("修改成功");
        },
        () => {
          ElMessage("修改失败");
        }
      );
    };
    return {
      name,
      notice,
      illustate,
      sentence,
      qq,
      github,
      signs,
      links,
      sum,
      signEvent,
      linkEvent,
      addFlink,
      addSign,
      setLink,
      setSign,
      delLink,
      setInfo,
    };
  },
};
</script>

<style lang="scss" scoped>
.conf-container {
  display: grid;
  grid-template-columns: 3fr 5fr;
  column-gap: 1rem;
  height: calc(100vh - 9.5rem);
  min-height: 34.5rem;
  min-width: 48rem;
  section {
    overflow: hidden;
    min-width: 23.5rem;
  }
}
.conf-info {
  position: relative;
  padding: 1rem;
  margin-bottom: 1rem;
  min-width: 2rem;
  height: 16rem;
  background-color: white;
  box-shadow: 0px 0px 1px 0px darkgray;
  overflow: hidden;
  animation: load-bottom 1s;
}
.conf-link {
  position: relative;
  padding: 1rem;
  height: calc(100% - 17rem);
  background-color: white;
  box-shadow: 0px 0px 1px 0px darkgray;
  overflow: hidden;
  animation: load-bottom 1s;
  h3 {
    color: darkcyan;
  }
  ul {
    margin-top: 1rem;
    padding: 0rem;
    list-style: none;
    height: calc(100% - 3rem);
    overflow: auto;
    li {
      padding: 0.5rem 1rem;
      & > div {
        display: flex;
        align-items: center;
      }
    }
  }
}
.conf-sign {
  position: relative;
  height: 100%;
  background-color: white;
  box-shadow: 0px 0px 1px 0px darkgray;
  overflow: hidden;
  animation: load-bottom 1s;
}
.conf-info-main {
  display: flex;
  height: 6rem;
}
.conf-info-portrait {
  height: 6rem;
  width: 6rem;
  border-radius: 0.2rem;
  background-image: url("../../assets/image/head-portrait.png");
  background-position: center center;
  background-size: cover;
  box-shadow: 0px 0px 1px 0px black;
}
.conf-info-name {
  flex: 1;
  position: relative;
  margin-left: 2rem;
  color: darkcyan;
  letter-spacing: 0.1rem;
  p {
    position: absolute;
    bottom: 0rem;
  }
}
.conf-info-message {
  padding: 1rem;
  display: flex;
  justify-content: center;
  align-items: center;
  font-size: 1.2rem;
  color: darkcyan;
  font-weight: 300;
  label {
    margin: 0rem 1rem;
  }
}
.el-scrollbar {
  background-color: white;
}
.conf-icon {
  position: absolute;
  top: -2rem;
  right: -2rem;
  width: 4rem;
  height: 4rem;
  background-color: darkcyan;
  transform: rotateZ(45deg);
}
.conf-icon::after {
  content: "";
  position: absolute;
  display: block;
  top: 2.8rem;
  left: 1.6rem;
  width: 0.8rem;
  height: 0.8rem;
  background-color: white;
  border-radius: 0.4rem;
}
.link-name {
  margin: 0.1rem;
  width: 12rem;
}
.conf-sign-box {
  padding: 1rem;
  h3 {
    color: darkcyan;
  }
}
.conf-signs {
  padding: 1rem;
  h4 {
    text-align: center;
    color: lightseagreen;
  }
}

@media screen and (max-width: 780px) {
  .conf-container {
    display: block;
    min-width: 23.5rem;
    section {
      min-width: 100%;
    }
  }
}
</style>