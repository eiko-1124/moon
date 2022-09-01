<template>
  <header>
    <label>分类</label>
    <el-radio-group class="el-radio-group">
      <el-radio-button
        class="el-radio-button"
        :class="{ 'el-select': typeFlag == index }"
        v-for="(sType, index) in typeDate"
        :key="index"
        :label="sType"
        @click="typeSelect(index)"
      />
    </el-radio-group>
  </header>
  <main>
    <el-table :data="listData" :border="true" style="width: 100%; height: auto">
      <el-table-column type="expand">
        <template v-slot="props">
          <section class="comment-main">
            <div>
              <el-tag style="margin-right: 0.5rem"
                >id:{{ props.row.id }}</el-tag
              >
              <el-tag class="ml-2" type="success"
                ><label>{{ props.row.name }}</label>
                <label v-if="props.row.cFlag == 1">{{
                  "回复" + props.row.cName
                }}</label></el-tag
              >
            </div>
            <div class="comment-text">{{ props.row.content }}</div>
            <div class="comment-option">
              <label>发布</label
              ><span
                class="switch-btn"
                :class="{ 'switch-select': props.row.hFlag == 0 }"
                @click="hiddenComment(props.$index)"
                ><span class="switch-slider"></span
              ></span>
              <el-button
                type="danger"
                @click="deleteComment(props.row.id, props.row.title)"
                >删除</el-button
              >
            </div>
          </section>
        </template>
      </el-table-column>
      <el-table-column prop="time" label="日期" width="200" />
      <el-table-column prop="aType" label="类别" width="200" />
      <el-table-column prop="title" label="标题" />
    </el-table>
  </main>
  <footer>
    <div class="demo-pagination-block">
      <el-pagination
        layout="total,  prev, pager, next, jumper"
        :total="sum"
        :pager-count="5"
        :page-size="10"
        @current-change="getCommentList(page)"
        @prev-click="getCommentList(page)"
        @next-click="getCommentList(page)"
      />
    </div>
  </footer>
</template>

<script>
import { ref } from "@vue/reactivity";
import { useStore } from "vuex";
import { computed, onMounted } from "@vue/runtime-core";
import { DeleteComments } from "@/api/apis/article";
export default {
  setup() {
    const store = useStore();
    const sum = computed(() => store.state.Article.commentSum);
    const listData = computed(() => store.state.Article.commentList);
    const typeDate = ["全部", "技术", "杂文", "故事"];
    const typeFlag = ref(0);
    const typeSelect = (index) => {
      typeFlag.value = index;
      getSum(index);
      getCommentList(0);
    };
    const hiddenComment = (index) => {
      store.dispatch("hiddenComment", [index, listData.value[index]]);
    };
    const getSum = (aType) => {
      store.dispatch("getCommentSum", aType);
    };
    const getCommentList = (page) => {
      store.dispatch("getCommentList", { type: typeFlag.value, page });
    };
    const deleteComment = (id, title) => {
      const form = new FormData();
      form.append("id", id);
      form.append("title", title);
      DeleteComments(
        form,
        (Res) => {
          if (Res.res == 1) {
            ElMessage("删除成功");
            getCommentList(0);
          } else ElMessage("删除失败");
        },
        (err) => {
          ElMessage("删除失败");
          console.log(err);
        }
      );
    };
    onMounted(() => {
      getSum(0);
      getCommentList(0);
    });
    return {
      sum,
      listData,
      typeDate,
      typeFlag,
      typeSelect,
      hiddenComment,
      deleteComment,
      getCommentList,
    };
  },
};
</script>

<style lang="scss" scoped>
header {
  margin-bottom: 1rem;
  padding: 1rem;
  display: flex;
  align-items: center;
  position: relative;
  background-color: white;
  box-shadow: 0px 0px 1px 0px darkgray;
  overflow: hidden;
  min-width: 23.5rem;
  animation: load-bottom 1s;
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
  label {
    margin-right: 1rem;
  }
}
.el-radio-button {
  margin: 0rem 0.5rem;
}
.el-radio-button::v-deep > .el-radio-button__inner {
  border: 1px solid #409eff !important;
  border-radius: 0.2rem !important;
}
.el-select::v-deep > .el-radio-button__inner {
  background-color: #409eff;
  color: white;
}
main {
  margin-bottom: 1rem;
  padding: 1rem;
  position: relative;
  min-height: 65vh;
  min-width: 23.5rem;
  background-color: white;
  box-shadow: 0px 0px 1px 0px darkgray;
  overflow: hidden;
  animation: load-bottom 1s;
}
.comment-main {
  padding: 0.5rem;
}
.comment-text {
  margin-top: 0.5rem;
  padding: 0.5rem;
  background-color: rgba($color: gray, $alpha: 0.1);
  border-radius: 0.2rem;
  box-shadow: 0px 0px 1px 0px darkgrey;
}
.switch-btn {
  display: inline-block;
  padding: 0.15rem;
  width: 3rem;
  height: 1.5rem;
  background-color: rgba($color: gray, $alpha: 0.2);
  border-radius: 0.75rem;
  transition-duration: 0.3s;
}
.switch-slider {
  display: inline-block;
  height: 1.2rem;
  width: 1.2rem;
  border-radius: 0.6rem;
  background-color: white;
  transition-duration: 0.3s;
}
.switch-select {
  background-color: dodgerblue;
  .switch-slider {
    position: relative;
    left: 1.5rem;
  }
}
.comment-option {
  margin-top: 0.5rem;
  display: flex;
  align-items: center;
  & > span,
  & > label {
    margin-right: 0.5rem;
  }
}
footer {
  padding: 1rem;
  background-color: white;
  box-shadow: 0px 0px 1px 0px darkgray;
  animation: load-bottom 1s;
}
</style>