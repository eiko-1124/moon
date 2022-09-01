<template>
  <div class="about-container">
    <div class="about-preview">
      <h3>预览</h3>
      <el-button type="primary" class="about-sumbit" @click="submit"
        >提交</el-button
      >
      <div class="preview-box" v-html="fText"></div>
    </div>
    <div class="about-edit">
      <mavonEditor
        class="mavon-editer"
        :modelValue="rText"
        :toolbars="aMdOption"
        :subfield="false"
        ref="md"
        @change="edit"
      ></mavonEditor>
    </div>
  </div>
</template>

<script>
import { mavonEditor } from "mavon-editor";
import "mavon-editor/dist/css/index.css";
import { aMdOption } from "@/conf/editConfig";
import { useStore } from "vuex";
import { computed, onMounted } from "@vue/runtime-core";
import { UpdateAbout } from "@/api/apis/user";
export default {
  components: { mavonEditor },
  setup() {
    const store = useStore();
    const fText = computed(() => store.state.User.about.format);
    const rText = computed(() => store.state.User.about.rander);
    const edit = (rander, value) => {
      store.commit("setRander", rander);
      store.commit("setFormat", value);
    };
    const submit = () => {
      const form = new FormData();
      form.append("rander", rText.value);
      form.append("format", fText.value);
      UpdateAbout(
        form,
        (Res) => {
          if (Res == 0) ElMessage("更新成功");
          else ElMessage("更新失败");
        },
        (Err) => {
          console.log(Err);
          ElMessage("更新失败");
        }
      );
    };
    onMounted(() => {
      store.dispatch("getAbout");
    });
    return {
      fText,
      rText,
      aMdOption,
      edit,
      submit,
    };
  },
};
</script>

<style lang="scss" scoped>
.about-container {
  animation: about-load 2s;
}
.about-preview {
  margin-bottom: 0.5rem;
  padding: 1rem;
  position: relative;
  background-color: white;
  box-shadow: 0px 0px 1px 0px darkgray;
  min-width: 23.5rem;
  h3 {
    margin-bottom: 1rem;
  }
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
    box-shadow: 0px 0px 1px 0px darkgray;
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
.preview-box {
  position: relative;
  background-color: rgba($color: gray, $alpha: 0.1);
  padding: 1rem;
  border-radius: 0.2rem;
  word-break: break-all;
  min-height: 40vh;
}
.about-sumbit {
  position: absolute;
  top: 1rem;
  left: 4rem;
}

@keyframes about-load {
  0% {
    opacity: 0;
  }
  100% {
    opacity: 1;
  }
}
</style>