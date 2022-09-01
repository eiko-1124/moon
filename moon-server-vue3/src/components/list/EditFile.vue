<template>
  <h3>资料文件及背景选择</h3>
  <el-tabs class="el-tabs" type="border-card">
    <el-tab-pane label="上传">
      <el-button type="primary" @click="uploadFile">上传</el-button>
      <el-scrollbar class="el-scrollbar-1">
        <ul>
          <li
            class="file"
            v-for="(file, index) in hadUploadFiles"
            :key="index"
            @click="setCover(file.link)"
          >
            <div class="file-base">
              <p>
                <el-icon class="file-icon"><Connection /></el-icon
                >{{ file.fileName }}[{{ file.link }}]
              </p>
              <el-icon
                class="file-option-icon"
                @click.stop="copyLink(file.link)"
                ><CopyDocument
              /></el-icon>
            </div>
          </li>
          <li class="file" v-for="(file, index) in uploadingFiles" :key="index">
            <div class="file-upload">
              <p>
                <el-icon class="file-icon"><Connection /></el-icon
                >{{ file.fileName }}[waiting...]
              </p>
              <label class="file-upload-bar">{{ file.percentage }}%</label>
            </div>
          </li>
        </ul>
      </el-scrollbar>
    </el-tab-pane>
    <el-tab-pane label="全部">
      <el-scrollbar class="el-scrollbar-2" v-infinite-scroll="loadMoreFiles">
        <ul>
          <li
            class="file"
            v-for="(file, index) in allUploadFiles"
            :key="index"
            @click="setCover(file.link)"
          >
            <div class="file-base">
              <p>
                <el-icon class="file-icon"><Connection /></el-icon
                >{{ file.fileName }}[{{ file.link }}]
              </p>
              <el-icon
                class="file-option-icon"
                @click.stop="copyLink(file.link)"
                ><CopyDocument
              /></el-icon>
            </div>
          </li>
        </ul>
      </el-scrollbar>
    </el-tab-pane>
  </el-tabs>
  <el-input
    placeholder="Please input link"
    style="margin-top: 1rem"
    v-model="cover"
  >
    <template #prepend>封面</template>
  </el-input>
</template>

<script>
import { computed } from "@vue/runtime-core";
import { useStore } from "vuex";
export default {
  setup() {
    const store = useStore();
    const hadUploadFiles = computed(() => store.state.File.hadUpload);
    const uploadingFiles = computed(() => store.state.File.uploading);
    const allUploadFiles = computed(() => store.state.File.allUpload);
    const cover = computed({
      get: () => {
        return store.state.Article.article.cover;
      },
      set: (value) => {
        store.commit("setCover", value);
      },
    });
    let page = 0;
    const loadMoreFiles = () => {
      if (store.state.File.allUpload.length == 0) {
        store.dispatch("getAllFiles", page);
        page++;
      }
    };
    const uploadFile = () => {
      const input = document.createElement("input");
      input.type = "file";
      input.click();
      input.addEventListener("change", () => {
        const size = Math.ceil(input.files[0].size / (1024 * 1024));
        if (size > 20) ElMessage("文件限制为20m");
        else store.dispatch("uploadFile", input.files[0]);
      });
    };
    const setCover = (link) => {
      store.commit("setCover", link);
    };
    const copyLink = async (value) => {
      await navigator.clipboard.writeText(value);
    };
    return {
      uploadFile,
      hadUploadFiles,
      uploadingFiles,
      allUploadFiles,
      cover,
      setCover,
      loadMoreFiles,
      copyLink,
    };
  },
};
</script>

<style lang="scss" scoped>
h3 {
  padding-bottom: 1rem;
}
.el-tabs {
  position: relative;
}
.el-scrollbar-1 {
  background-color: white;
  ul {
    padding: 0rem;
    list-style: none;
    overflow: auto;
  }
  &::v-deep .el-scrollbar__view {
    height: 14rem;
  }
}
.el-scrollbar-2 {
  background-color: white;
  &::v-deep .el-scrollbar__view {
    height: 16rem;
  }
  ul {
    padding: 0rem;
    list-style: none;
    overflow: auto;
  }
}
.file {
  padding: 0.2rem;
  color: gray;
}
.file-base {
  padding-right: 0.5rem;
  overflow: hidden;
  .file-icon,
  .file-option-icon {
    margin: 0.1rem;
    height: 1rem;
  }
  p {
    float: left;
    word-break: break-all;
  }
  .file-option-icon {
    float: right;
    cursor: pointer;
    &:hover {
      color: darkcyan;
    }
  }
  &:hover {
    p {
      color: deepskyblue;
      cursor: default;
    }
  }
}
.file-upload {
  padding-right: 0.5rem;
  overflow: hidden;
  .file-icon {
    margin: 0.1rem;
    height: 1rem;
  }
  p {
    float: left;
    word-break: break-all;
  }
  &:hover {
    cursor: not-allowed;
    color: deepskyblue;
  }
  .file-upload-bar {
    float: right;
  }
}
</style>