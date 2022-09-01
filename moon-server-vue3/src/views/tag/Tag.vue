<template>
  <div class="tag-contianer">
    <section>
      <img class="fire-keeper" src="../../assets/image/fire-keeper.png" />
    </section>
    <section class="tag-add">
      <el-input
        v-model="input3"
        placeholder="Please input"
        class="input-with-select"
      >
        <template #append>
          <el-button>添加</el-button>
        </template>
      </el-input>
    </section>
    <section class="tag-box">
      <el-tag
        style="margin: 0.2rem 0.2rem"
        v-for="(tag, index) in tags"
        :key="index"
        :type="info"
        size="large"
        >{{ tag.tag + " " + tag.num }}</el-tag
      >
    </section>
  </div>
</template>

<script>
import { computed, onMounted } from "@vue/runtime-core";
import { useStore } from "vuex";
export default {
  setup() {
    const store = useStore();
    const tags = computed(() => store.state.Article.tagsDetails);
    onMounted(() => {
      store.dispatch("getTagsDetails");
    });
    return {
      tags,
    };
  },
};
</script>

<style lang="scss" scoped>
.tag-contianer {
  padding: 1rem;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  height: calc(100vh - 9.5rem);
  min-width: 23.5rem;
  background-color: white;
  box-shadow: 0px 0px 1px 0px darkgray;
  animation: load-bottom 1s;
}
section {
  display: flex;
  justify-content: center;
  flex-wrap: wrap;
}
.tag-add {
  margin-bottom: 1rem;
  width: 30rem;
  height: 2.5rem;
}
.tag-contianer {
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
.tag-box {
  padding-bottom: 5rem;
}
.fire-keeper {
  width: 10rem;
}
.input-with-select {
  max-width: 100%;
}
</style>