<template>
  <div class="recovery">
    <el-tabs v-model="activeName">
      <el-tab-pane label="文章" name="first">
        <el-scrollbar
          class="el-scrollbar"
          v-infinite-scroll="getCovers"
          :infinite-scroll-disabled="aloading"
        >
          <div class="a-iamge-container">
            <div
              class="recovery-article-site"
              v-for="(cover, index) in aCovers"
              :key="index"
            >
              <el-image :src="cover.cover" fit="fill" />
              <div class="site-text">
                <label>{{ cover.title }}</label>
              </div>
            </div>
          </div>
        </el-scrollbar>
      </el-tab-pane>
      <el-tab-pane label="记录" name="second">
        <el-scrollbar
          class="el-scrollbar"
          v-infinite-scroll="getImages"
          :infinite-scroll-disabled="mloading"
        >
          <div class="a-iamge-container">
            <div
              class="recovery-article-site"
              v-for="(image, index) in mImages"
              :key="index"
            >
              <el-image :src="image.cover" fit="fill" />
              <div class="site-text">
                <label>{{ image.time }}</label>
              </div>
            </div>
          </div></el-scrollbar
        ></el-tab-pane
      >
      <el-tab-pane label="其他" name="third">
        <el-scrollbar
          class="el-scrollbar"
          v-infinite-scroll="getFiles"
          :infinite-scroll-disabled="oloading"
        >
          <ul class="another-container">
            <li v-for="(file, index) in oFiles" :key="index">
              <el-icon class="icon"><Files /></el-icon>{{ file.name
              }}<a :href="file.link">{{ file.link }}</a>
            </li>
          </ul></el-scrollbar
        ></el-tab-pane
      >
    </el-tabs>
  </div>
</template>

<script>
import { ref } from "@vue/reactivity";
import { computed, onMounted, watch } from "@vue/runtime-core";
import { useStore } from "vuex";
export default {
  setup() {
    const store = useStore();
    const activeName = ref("first");
    const aloading = ref(false);
    const mloading = ref(false);
    const oloading = ref(false);
    const aCovers = computed(() => store.state.File.articleCovers);
    const aSum = computed(() => store.state.File.articleCoverSum);
    const mImages = computed(() => store.state.File.moodImages);
    const mSum = computed(() => store.state.File.moodImageSum);
    const oSum = computed(() => store.state.File.otherFileSum);
    const oFiles = computed(() => store.state.File.otherFiles);
    const getCovers = () => {
      if (aCovers.value.length < aSum.value) {
        aloading.value = true;
        store.dispatch("getArticleCovers", [aCovers.value.length, aloading]);
      }
    };
    const getImages = () => {
      if (mImages.value.length < mSum.value) {
        mloading.value = true;
        store.dispatch("getMoodImages", [mImages.value.length, mloading]);
      }
    };
    const getFiles = () => {
      if (oFiles.value.length < oSum.value) {
        oloading.value = true;
        store.dispatch("getOtherFiles", [oFiles.value.length, oloading]);
      }
    };
    watch(aSum, () => {
      getCovers();
    });
    onMounted(() => {
      store.dispatch("getArticleCoverSum");
      store.dispatch("getMoodImageSum");
      store.dispatch("getOtherFileSum");
    });
    return {
      activeName,
      aCovers,
      aSum,
      mImages,
      mSum,
      oSum,
      oFiles,
      aloading,
      mloading,
      oloading,
      getCovers,
      getImages,
      getFiles,
    };
  },
};
</script>

<style lang="scss" scoped>
.recovery {
  position: relative;
  padding: 1rem;
  height: 80vh;
  min-width: 23.5rem;
  background-color: white;
  animation: load-bottom 1s;
}
.el-scrollbar {
  background-color: white;
}
.a-iamge-container {
  position: relative;
  display: grid;
  grid-template-columns: repeat(5, 1fr);
  gap: 0.5rem 0.5rem;
}
.recovery-article-site {
  position: relative;
  animation: site-load 1s;
}
.site-text {
  padding: 0rem 0.5rem;
  text-align: center;
  label {
    white-space: 0.1rem;
    word-break: break-all;
    font-size: 0.8rem;
    color: gray;
  }
}
.another-container {
  padding: 0rem;
  list-style: none;
  li {
    display: flex;
    margin-bottom: 0.2rem;
    align-items: center;
    color: gray;
    a {
      margin-left: 0.5rem;
      text-decoration: none;
      color: gray;
    }
    &:hover {
      color: darkcyan;
      cursor: pointer;
    }
  }
}
.icon {
  margin-right: 0.2rem;
}

@media screen and (max-width: 780px) {
  .a-iamge-container {
    grid-template-columns: repeat(3, 1fr);
  }
}

@media screen and (max-width: 560px) {
  .a-iamge-container {
    grid-template-columns: repeat(2, 1fr);
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