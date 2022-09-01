<template>
  <el-scrollbar style="height: 100%">
    <div class="mood-site" v-for="(mood, index) in moods" :key="index">
      <div class="mood-site-main">
        <p>{{ mood.text }}</p>
        <div
          class="mood-image"
          :style="`grid-template-rows: repeat(${Math.ceil(
            mood.links.length / 4
          )}, 1fr);grid-template-columns: repeat(${
            mood.links.length > 4 ? 4 : mood.links.length
          }, 1fr);`"
        >
          <div
            class="mood-image-box"
            :class="{ 'image-one': mood.links.length == 1 }"
            v-for="(link, index) in mood.links"
            :key="index"
            :style="'background-image:' + 'url(' + link + ')'"
          ></div>
        </div>
        <el-divider content-position="right">{{ mood.time }}</el-divider>
        <div class="mood-site-option">
          <label>{{ "点赞 " + mood.like }}</label>
          <el-divider direction="vertical" />
          <label>{{ "评论 " + mood.comment }}</label>
        </div>
        <span class="del-btn">
          <el-icon @click="deleteMood(mood.id)"><Delete /></el-icon>
        </span>
      </div>
      <div
        class="mood-site-comment"
        :class="{ 'site-first': cindex == 0 }"
        v-for="(comment, cindex) in mood.comments"
        :key="cindex"
      >
        <label class="label-1"
          ><label>{{ comment.name + " " }}</label
          ><label v-if="comment.cFlag == 1">{{
            "回复 " + comment.cname + " "
          }}</label>
          <label>:</label> </label
        >{{ " " + comment.content
        }}<label class="label-2" @click="deleteComment(mood.id, comment.cid)"
          >删除</label
        >
      </div>
      <div class="mood-comment-load" v-if="mood.comment > mood.comments.length">
        <label @click="loadComment(mood.id, mood.comments.length)"
          >查看更多</label
        >
      </div>
    </div>
    <div class="mood-load" v-if="moods.length < moodSum" @click="loadMood">
      加载
    </div>
  </el-scrollbar>
</template>

<script>
import { computed, onMounted } from "@vue/runtime-core";
import { useStore } from "vuex";
export default {
  setup() {
    const store = useStore();
    const moods = computed(() => store.state.Mood.moods);
    const moodSum = computed(() => store.state.Mood.moodSum);
    const loadMood = () => {
      store.dispatch("loadMood", moods.value.length);
    };
    const loadComment = (id, page) => {
      store.dispatch("loadComments", { id, page });
    };
    const deleteComment = (id, cid) => {
      store.dispatch("deleteComment", [
        { id, cid },
        () => {
          ElMessage("删除成功");
        },
        () => {
          ElMessage("删除失败");
        },
      ]);
    };
    const deleteMood = (id) => {
      store.dispatch("deleteMood", [
        id,
        () => {
          ElMessage("删除成功");
        },
        () => {
          ElMessage("删除失败");
        },
      ]);
    };
    onMounted(() => {
      store.dispatch("getMoodList");
      store.dispatch("getMoodSum");
    });
    return {
      moods,
      moodSum,
      loadMood,
      loadComment,
      deleteComment,
      deleteMood,
    };
  },
};
</script>

<style lang="scss" scoped>
.mood-site {
  margin-bottom: 1rem;
  margin-right: 1rem;
  animation: site-load 1s;
  &:nth-last-child(1) {
    margin-bottom: 0rem;
  }
}
.mood-site-main {
  padding: 1rem;
  background-color: white;
  p {
    padding: 1rem 0rem;
    text-align: center;
  }
  &:hover {
    .del-btn {
      opacity: 1;
    }
  }
}
.mood-image {
  padding: 0rem 1rem;
  display: grid;
  gap: 0.5rem;
}
.mood-image-box {
  position: relative;
  padding-bottom: 100%;
  height: 0rem;
  overflow: hidden;
  background-position: center center;
  background-size: cover;
}
.image-one {
  margin: 0rem auto;
  padding-bottom: 50%;
  width: 75%;
}
.mood-site-option {
  display: flex;
  justify-content: center;
  label {
    margin: 0rem auto;
  }
}
.del-btn {
  padding: 0.5rem;
  position: absolute;
  display: block;
  top: 0rem;
  left: 0rem;
  width: 2.5rem;
  height: 2.5rem;
  background-color: white;
  color: rgba($color: black, $alpha: 0.7);
  z-index: 10;
  font-size: 1.2rem;
  opacity: 0;
  transition-duration: 0.5s;
  &:hover {
    color: crimson;
    cursor: pointer;
  }
}
.mood-site-comment {
  padding: 1rem;
  background-color: white;
  word-break: break-all;
  label {
    color: darkcyan;
  }
}
.site-first {
  margin-top: 0.2rem;
}
.label-1 {
  margin-right: 0.2rem;
}
.label-2 {
  margin-left: 0.5rem;
  font-size: 0.8rem;
  cursor: pointer;
}
.mood-site {
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
.mood-load {
  margin-right: 1rem;
  background-color: white;
  padding: 0.6rem;
  text-align: center;
  color: darkcyan;
  box-shadow: 0px 0px 1px 0px darkgray;
  animation: site-load 1s;
  &:hover {
    background-color: darkcyan;
    color: white;
    cursor: pointer;
  }
}
.mood-comment-load {
  margin-top: -0.5rem;
  padding: 0rem 1rem 0.5rem;
  color: darkcyan;
  background-color: white;
  text-align: right;
  font-size: 0.8rem;
  label {
    &:hover {
      cursor: pointer;
    }
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