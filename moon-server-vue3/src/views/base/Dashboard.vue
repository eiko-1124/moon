<template>
  <section class="base-board">
    <div
      class="board"
      v-for="(site, index) in data"
      :key="index"
      :style="`background-color: ${site.color}`"
    >
      <h3>
        {{ site.title }}
        <el-icon class="board-title-icon"
          ><component :is="site.icon"></component
        ></el-icon>
      </h3>
      <h2>{{ sum[site.key] + site.text }}</h2>
      <div class="board-link">
        <router-link
          :to="link.path"
          v-for="(link, index) in site.link"
          :key="index"
        >
          {{ link.title }}</router-link
        >
      </div>
      <span class="board-circle-1"></span>
      <span class="board-circle-2"></span>
    </div>
  </section>
  <section class="charts-container">
    <div class="charts-box" ref="pie"></div>
    <div class="charts-box" ref="category"></div>
  </section>
  <section class="log-container">
    <div class="log-box">
      <h3>近期记录</h3>
      <ul>
        <li v-for="(diary, index) in aDiary" :key="index">
          <router-link to="/">
            <el-row :gutter="10">
              <el-col :span="18"
                ><div class="el-col-text">
                  {{ index + "、" + diary.text }}
                </div></el-col
              >
              <el-col :span="6"
                ><div class="el-col-text">{{ diary.time }}</div></el-col
              >
            </el-row>
          </router-link>
        </li>
      </ul>
    </div>
    <div class="log-box">
      <h3>最新消息</h3>
      <ul>
        <li v-for="(diary, index) in uDiary" :key="index">
          <router-link to="/">
            <el-row :gutter="10">
              <el-col :span="18"
                ><div class="el-col-text">
                  {{ index + "、" + diary.text }}
                </div></el-col
              >
              <el-col :span="6"
                ><div class="el-col-text">{{ diary.time }}</div></el-col
              >
            </el-row>
          </router-link>
        </li>
      </ul>
    </div>
  </section>
</template>

<script>
import board from "@/conf/board.js";
import { ref } from "@vue/reactivity";
import {
  onActivated,
  onDeactivated,
  onMounted,
  onBeforeUnmount,
  computed,
  watch,
} from "@vue/runtime-core";
import * as echarts from "echarts";
import { useStore } from "vuex";
export default {
  setup() {
    const store = useStore();
    const sum = computed(() => store.state.User.sum);
    const nearState = computed(() => store.state.User.nearState);
    const nearDate = computed(() => store.state.User.nearDate);
    const nearUdiary = computed(() => store.state.User.nearUdiary);
    const nearAdiary = computed(() => store.state.User.nearAdiary);
    const aDiary = computed(() => store.state.User.aDiary);
    const uDiary = computed(() => store.state.User.uDiary);
    const data = board;
    const tags = computed(() => store.state.User.tags);
    const pie = ref(null);
    const category = ref(null);
    const paintPie = () => {
      const pieChart = echarts.init(pie.value);
      pieChart.setOption({
        title: {
          text: "最热门分类",
        },
        series: [
          {
            type: "pie",
            data: [...tags.value],
          },
        ],
      });
    };
    const paintCategory = () => {
      const categoryChart = echarts.init(category.value);
      categoryChart.setOption({
        tooltip: {
          trigger: "axis",
          axisPointer: { type: "cross" },
        },
        title: {
          text: "近期活跃度",
        },
        xAxis: {
          type: "category",
          name: "活跃时期",
          axisTick: {
            alignWithLabel: true,
          },
          data: [...nearDate.value],
        },
        yAxis: [
          {
            type: "value",
            name: "活跃度",
            splitNumber: 4,
            minInterval: 25,
            position: "left",
          },
        ],
        series: [
          {
            data: [...nearAdiary.value],
            type: "line",
          },
          {
            data: [...nearUdiary.value],
            type: "line",
          },
        ],
      });
    };
    const rePaint = () => {
      echarts.getInstanceByDom(category.value).resize();
      echarts.getInstanceByDom(pie.value).resize();
    };
    watch(tags, () => {
      paintPie();
    });
    watch(nearState, () => {
      paintCategory();
    });
    onMounted(() => {
      if (aDiary.value.length == 0 || uDiary.value.length == 0)
        store.dispatch("getDiary");
      store.dispatch("getSum");
      store.dispatch("getTags");
      if (nearState.value == 0) store.dispatch("getNear");
      else paintCategory();
      window.addEventListener("resize", rePaint);
    });
    onActivated(() => {
      window.addEventListener("resize", rePaint);
    });
    onDeactivated(() => {
      window.removeEventListener("resize", rePaint);
    });
    onBeforeUnmount(() => {
      window.removeEventListener("resize", rePaint);
    });
    return {
      data,
      pie,
      category,
      sum,
      nearDate,
      nearAdiary,
      nearUdiary,
      uDiary,
      aDiary,
    };
  },
};
</script>

<style lang="scss" scoped>
section {
  margin-bottom: 1rem;
  min-width: 23.5rem;
}
.base-board {
  padding: 1.5rem;
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  column-gap: 1.5rem;
  background-color: white;
  box-shadow: 0px 0px 1px 0px darkgray;
  animation: load-bottom 1s;
}
.board {
  position: relative;
  padding-bottom: 60%;
  height: 0rem;
  color: white;
  overflow: hidden;
  h2,
  h3 {
    position: absolute;
    margin: 0rem 0.5rem;
    line-height: 2rem;
    height: 2rem;
    font-weight: 500;
  }
  h2 {
    height: 2rem;
    width: 100%;
    top: calc(50% - 1rem);
    letter-spacing: 0.1rem;
    text-align: center;
    font-weight: 600;
  }
}
.board-circle-1 {
  position: absolute;
  display: block;
  top: 45%;
  right: 0%;
  width: 60%;
  padding-bottom: 60%;
  background-color: rgba($color: white, $alpha: 0.3);
  border-radius: 100%;
}
.board-circle-2 {
  position: absolute;
  display: block;
  top: 0%;
  right: 0%;
  width: 60%;
  padding-bottom: 60%;
  background-color: rgba($color: white, $alpha: 0.3);
  border-radius: 100%;
  transform: translate(30%, -30%);
}
.board-title-icon {
  position: absolute;
  margin: 0.5rem 0.25rem;
}
.board-link {
  position: absolute;
  bottom: 0.5rem;
  margin: 0rem 0.5rem;
  & > a {
    margin-right: 0.2rem;
    color: white;
    text-decoration: none;
    font-size: 0.8rem;
    transition-duration: 0.3s;
    &:hover {
      color: cyan;
    }
  }
}
.charts-container {
  position: relative;
  display: grid;
  grid-template-columns: 3fr 5fr;
  column-gap: 1rem;
  overflow: hidden;
}
.charts-box {
  max-width: 100%;
  padding: 1rem;
  height: 15rem;
  background-color: white;
  box-shadow: 0px 0px 1px 0px darkgray;
  animation: load-bottom 1s;
}
.log-container {
  display: grid;
  grid-template-columns: 1fr 1fr;
  column-gap: 1rem;
}
.log-box {
  padding: 1rem;
  height: 25rem;
  background-color: white;
  box-shadow: 0px 0px 1px 0px lightgray;
  animation: load-bottom 1s;
  h3 {
    height: 1.5rem;
  }
  ul {
    list-style: none;
    padding: 1rem;
  }
  li {
    height: 2rem;
  }
  a {
    text-decoration: none;
    color: cadetblue;
    &:hover {
      color: dodgerblue;
    }
  }
}
.el-col-text {
  white-space: nowrap;
  text-overflow: ellipsis;
  overflow: hidden;
}

@media screen and (max-width: 780px) {
  .base-board {
    grid-template-columns: repeat(2, 1fr);
    grid-template-rows: repeat(2, 1fr);
    gap: 1.5rem;
  }
  .charts-container {
    grid-template-columns: repeat(1, 1fr);
    grid-template-rows: repeat(2, 1fr);
    gap: 1.5rem;
  }
}
@media screen and (max-width: 560px) {
  .log-container {
    grid-template-columns: repeat(1, 1fr);
    grid-template-rows: repeat(2, 1fr);
    gap: 1.5rem;
  }
}
</style>
