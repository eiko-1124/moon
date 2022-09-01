<template>
  <nav>
    <section>
      <label>时间</label>
      <div>
        <el-radio-group class="el-radio-group">
          <el-radio-button
            :class="{ 'el-radio-button': true, 'el-select': timeFlag == index }"
            v-for="(sTime, index) in timeData"
            :key="index"
            :label="sTime"
            @change="timeSelect(index)"
          />
        </el-radio-group>
      </div>
    </section>
    <section>
      <label>分类</label>
      <div>
        <el-radio-group class="el-radio-group">
          <el-radio-button
            :class="{ 'el-radio-button': true, 'el-select': typeFlag == index }"
            v-for="(sType, index) in typeData"
            :key="index"
            :label="sType"
            @change="typeSelect(index)"
          />
        </el-radio-group>
      </div>
    </section>
    <section>
      <label>标签</label>
      <div>
        <el-radio-group class="el-radio-group">
          <el-radio-button
            v-for="(tag, index) in tagData"
            :key="index"
            :class="{ 'el-radio-button': true, 'el-select': tagFlag == index }"
            :label="tag"
            @change="tagSelect(index)"
          />
        </el-radio-group>
      </div>
    </section>
  </nav>
  <article>
    <el-table :data="listData" :border="true" style="width: 100%; height: auto">
      <el-table-column type="expand">
        <template v-slot="props">
          <div>
            <section class="expand-container">
              <label>标签:</label>
              <el-tag
                class="el-tag"
                type="success"
                v-for="(tag, index) in props.row.tagArr"
                :key="index"
                >{{ tag }}</el-tag
              >
            </section>
            <section class="expand-container">
              <label>描述:</label>
              {{ props.row.illustrate }}
            </section>
            <section class="expand-container expand-option-container">
              <label>发布</label
              ><span
                class="switch-btn"
                :class="{ 'switch-select': props.row.hFlag == 0 }"
                @click="
                  hiddenArticle(props.$index, props.row.id, props.row.hFlag)
                "
                ><span class="switch-slider"></span
              ></span>
              <span class="expend-option">
                <el-button type="primary" @click="edit(props.row.id)"
                  >编辑</el-button
                >
                <el-button type="danger" @click="deleteArticle(props.row.id)"
                  >删除</el-button
                >
              </span>
            </section>
          </div>
        </template>
      </el-table-column>
      <el-table-column prop="id" label="编号" width="100" />
      <el-table-column prop="aType" label="类别" width="100" />
      <el-table-column prop="title" label="标题" />
      <el-table-column prop="time" label="日期" width="180" />
    </el-table>
  </article>
  <footer>
    <div class="demo-pagination-block">
      <el-pagination
        layout="total,  prev, pager, next, jumper"
        :total="sum"
        :pager-count="5"
        :page-size="10"
        @current-change="getArticleList(page)"
        @prev-click="getArticleList(page)"
        @next-click="getArticleList(page)"
      />
    </div>
  </footer>
</template>

<script>
import { ref } from "@vue/reactivity";
import { useStore } from "vuex";
import { computed, onMounted } from "@vue/runtime-core";
import { useRouter } from "vue-router";
import { DeleteArticle } from "@/api/apis/article";
export default {
  setup() {
    const store = useStore();
    const router = useRouter();
    store.dispatch("getAllTags");
    const timeFlag = ref(0);
    const typeFlag = ref(0);
    const tagFlag = ref(0);
    let sum = computed(() => store.state.Article.articleSum);
    const timeData = [
      "全部",
      "最近一年",
      "最近三个月",
      "最近一个月",
      "最近一周",
    ];
    const typeData = ["全部", "技术", "杂文", "故事"];
    const tagData = computed(() => store.state.Article.tags);
    const listData = computed(() => store.state.Article.articleList);
    const timeSelect = (index) => {
      timeFlag.value = index;
      getArticleSum();
      getArticleList(0);
    };
    const typeSelect = (index) => {
      typeFlag.value = index;
      getArticleSum();
      getArticleList(0);
    };
    const tagSelect = (index) => {
      tagFlag.value = index;
      getArticleSum();
      getArticleList(0);
    };
    const getArticleSum = () => {
      const data = {
        time: timeFlag.value,
        type: typeFlag.value,
        tag: tagFlag.value == 0 ? 0 : store.state.Article.tags[tagFlag.value],
      };
      store.dispatch("getArticleSum", data);
    };
    const getArticleList = (page) => {
      const data = {
        time: timeFlag.value,
        type: typeFlag.value,
        tag: tagFlag.value == 0 ? 0 : store.state.Article.tags[tagFlag.value],
        page: page,
      };
      store.dispatch("getArticleList", data);
    };
    const hiddenArticle = (index, id, hFlag) => {
      const cFlag = hFlag == 0 ? 1 : 0;
      store.dispatch("hiddenArticle", [index, cFlag, id]);
    };
    const edit = (index) => {
      store.commit("setActiveName", 1);
      store.commit("setActiveBtn", 0);
      router.push("/EditBlog/" + index);
    };
    const deleteArticle = (id) => {
      const form = new FormData();
      form.append("id", id);
      DeleteArticle(
        form,
        (Res) => {
          if (Res.res == 1) {
            ElMessage("删除成功");
            getArticleList(0);
          } else ElMessage("删除失败");
        },
        (Err) => {
          console.log(Err);
          ElMessage("删除失败");
        }
      );
    };
    onMounted(() => {
      getArticleSum();
      getArticleList(0);
    });
    return {
      listData,
      timeData,
      typeData,
      tagData,
      timeFlag,
      tagFlag,
      typeFlag,
      sum,
      timeSelect,
      typeSelect,
      tagSelect,
      hiddenArticle,
      edit,
      deleteArticle,
    };
  },
};
</script>

<style lang="scss" scoped>
nav,
article,
footer {
  margin-bottom: 1rem;
  padding: 1rem;
  background-color: white;
  box-shadow: 0px 0px 1px 0px darkgray;
  animation: load-bottom 1s;
}
article {
  min-height: 29.5rem;
}
nav {
  section {
    padding: 1rem 1rem 0rem;
    display: flex;
    & > label {
      width: 3rem;
    }
    & > div {
      flex: 1;
    }
  }
}
.el-radio-group {
  padding: 0rem 1rem;
}
.el-radio-button {
  margin: 0rem 0.5rem 0.5rem;
}
.el-radio-button::v-deep > .el-radio-button__inner {
  border: 1px solid #409eff !important;
  border-radius: 0.2rem !important;
}
nav {
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
.expand-container {
  padding: 0.5rem 1rem;
  & > label {
    margin-right: 0.5rem;
  }
}
.expand-option-container {
  display: flex;
  align-items: center;
}
.expend-option {
  position: absolute;
  right: 1rem;
}
.el-select::v-deep > .el-radio-button__inner {
  background-color: #409eff;
  color: white;
}
.el-tag {
  margin-right: 0.5rem;
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
</style>