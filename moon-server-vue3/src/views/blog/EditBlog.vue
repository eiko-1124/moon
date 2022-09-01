<template>
  <section class="edit-message">
    <div class="edit-message-box">
      <div class="edit-message-main">
        <h3>基本信息</h3>
        <div>
          <label>标题</label>
          <el-input
            :modelValue="article.title"
            @update:modelValue="baseInput($event, 'title')"
            placeholder="title"
            type="text"
            clearable
          />
        </div>
        <div>
          <label>类型</label>
          <el-radio-group
            :modelValue="article.aType"
            @update:modelValue="baseInput($event, 'aType')"
          >
            <el-radio :label="1">技术</el-radio>
            <el-radio :label="2">杂文</el-radio>
            <el-radio :label="3">故事</el-radio>
          </el-radio-group>
        </div>
        <div>
          <label>标签</label>
          <el-input
            placeholder="split by #"
            type="text"
            clearable
            :modelValue="article.tag"
            @update:modelValue="baseInput($event, 'tag')"
          />
        </div>
        <div class="edit-message-tag-box">
          <el-tag
            class="edit-message-tag"
            type="success"
            effect="dark"
            v-for="(tag, index) in tags"
            :key="index"
            >{{ tag }}</el-tag
          >
        </div>
      </div>
      <div class="edit-message-other">
        <h3>相关说明</h3>
        <el-input
          :autosize="{ minRows: 5, maxRows: 5 }"
          type="textarea"
          placeholder="illustrate"
          :modelValue="article.illustrate"
          @update:modelValue="baseInput($event, 'illustrate')"
        />
      </div>
    </div>
    <div class="edit-source-box">
      <edit-file></edit-file>
    </div>
  </section>
  <section class="edit-text-box">
    <h3>正文</h3>
    <mavonEditor
      class="mavon-editer"
      :toolbars="mdOption"
      :subfield="false"
      :modelValue="article.rContent"
      @change="editInput"
      @imgAdd="editImgAdd"
      ref="md"
    ></mavonEditor>
  </section>
  <section class="edit-submit">
    <el-button type="primary" @click="submit">提交</el-button>
    <el-button type="danger" @click="cancel">取消</el-button>
  </section>
</template>

<script>
import { mavonEditor } from "mavon-editor";
import "mavon-editor/dist/css/index.css";
import { mdOption } from "@/conf/editConfig";
import EditFile from "@/components/list/EditFile";
import { computed, onMounted, ref } from "@vue/runtime-core";
import { useStore } from "vuex";
import { minorFile } from "@/api/apis/file";
import { submitArticle } from "@/api/apis/article";
import { useRoute, useRouter } from "vue-router";
export default {
  components: {
    EditFile,
    mavonEditor,
  },
  setup() {
    const store = useStore();
    const router = useRouter();
    const route = useRoute();
    const article = computed(() => store.state.Article.article);
    const baseInput = (value, key) => {
      store.commit("editArticle", [key, value]);
    };
    const editInput = (rander, value) => {
      store.commit("editArticle", ["rContent", rander]);
      store.commit("editArticle", ["fContent", value]);
    };
    let tags = computed(() =>
      store.state.Article.article.tag.split("#").filter((i) => i && i.trim())
    );
    const md = ref(null);
    const editImgAdd = (pos, $file) => {
      const size = Math.ceil($file.size / (1024 * 1024));
      if (size > 20) {
        ElMessage("文件限制为20m");
      } else {
        const id = Symbol();
        const fileName = $file.name;
        let percentage = 0;
        let per = {
          set Value(value) {
            store.commit("setPercentage", [id, value]);
          },
        };
        store.commit("waitingUpload", { id, fileName, percentage });
        const form = new FormData();
        form.append("file", $file);
        form.append("fileName", fileName);
        minorFile(
          form,
          per,
          (Res) => {
            if (Res.res == "1") {
              store.commit("overUpload", [
                id,
                {
                  fileName: Res.file.fileName,
                  link: Res.file.link,
                },
              ]);
              md.value.$img2Url(pos, Res.file.link);
            } else {
              store.commit("failUpload", id);
            }
          },
          (Err) => {
            console.log(Err);
            store.commit("failUpload", id);
          }
        );
      }
    };
    const submit = () => {
      const form = checkForm();
      if (form != null) {
        submitArticle(
          form,
          (Res) => {
            if (Res.res == 1) {
              ElMessage("上传成功");
              setTimeout(() => {
                store.commit("setActiveName", 1);
                store.commit("setActiveBtn", 1);
                router.push("/BlogList");
              }, 1000);
            } else ElMessage("网络错误");
          },
          (Err) => {
            ElMessage("网络错误");
            console.log(Err);
          }
        );
      } else {
        ElMessage("资料不完整");
      }
    };
    const cancel = () => {
      router.push("BlogList");
    };
    const checkForm = () => {
      if (article.value.aType === 0) return null;
      const form = new FormData();
      for (const [key, value] of Object.entries(article.value)) {
        if (value === "") return null;
        form.append(key, value);
      }
      return form;
    };
    onMounted(() => {
      if (route.params.id == 0)
        store.commit("initArticle", {
          id: 0,
          title: "",
          aType: 0,
          tag: "",
          illustrate: "",
          cover: "",
          fContent: "",
          rContent: "",
        });
      else store.dispatch("getArticle", route.params.id);
    });
    return {
      mdOption,
      article,
      baseInput,
      editInput,
      tags,
      md,
      editImgAdd,
      submit,
      cancel,
    };
  },
};
</script>

<style lang="scss" scoped>
section {
  margin-bottom: 1rem;
  min-width: 23.5rem;
}
.edit-message {
  display: flex;
  animation: edit-load 1s;
}
.edit-message-main {
  margin-bottom: 1rem;
  padding: 1rem;
  background-color: white;
  box-shadow: 0px 0px 1px 0px darkgray;
  & > div {
    margin: 1rem;
    display: flex;
    align-items: center;
    label {
      width: 4rem;
    }
  }
}
.edit-message-box {
  flex: 4;
}
.edit-source-box {
  margin-left: 1rem;
  padding: 1rem;
  flex: 5;
  background-color: white;
  box-shadow: 0px 0px 1px 0px darkgray;
}
.edit-message-tag-box {
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
  padding-left: 3rem;
}
.edit-message-tag {
  margin: 0.1rem 0.4rem;
}
.edit-message-other {
  padding: 1rem 1rem 2rem;
  background-color: white;
  box-shadow: 0px 0px 1px 0px darkgray;
  h3 {
    margin-bottom: 1rem;
  }
}
.edit-text-box {
  height: 25rem;
  background-color: white;
  box-shadow: 0px 0px 1px 0px darkgray;
  animation: edit-load 1s;
}
.edit-submit {
  text-align: center;
  animation: edit-load 1s;
}
.edit-message-main,
.edit-message-other,
.edit-source-box,
.edit-text-box {
  position: relative;
  overflow: hidden;
  &::after {
    content: "";
    position: absolute;
    display: block;
    top: 0.5rem;
    right: 0.5rem;
    width: 0.8rem;
    height: 0.8rem;
    background-color: white;
    border-radius: 0.4rem;
  }
  &::before {
    content: "";
    position: absolute;
    display: block;
    top: -2rem;
    right: -2rem;
    width: 4rem;
    height: 4rem;
    background-color: darkcyan;
    transform: rotateZ(45deg);
  }
}
.edit-text-box {
  padding: 1rem 0rem 0rem;
  display: flex;
  flex-direction: column;
  height: 90vh;
  h3 {
    padding: 0rem 1rem 1rem;
  }
}
.mavon-editer {
  flex: 1;
}

@media screen and (max-width: 780px) {
  .edit-message {
    display: block;
  }
  .edit-source-box {
    margin: 1rem 0rem;
  }
}

@keyframes edit-load {
  0% {
    opacity: 0;
  }
  100% {
    opacity: 1;
  }
}
</style>