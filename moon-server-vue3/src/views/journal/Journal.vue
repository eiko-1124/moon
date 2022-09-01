<template>
  <main class="diary">
    <section>
      <h3>管理员日志</h3>
      <el-scrollbar
        class="el-scrollbar"
        v-infinite-scroll="getATimes"
        :infinite-scroll-disabled="aLoading"
      >
        <ul>
          <li
            v-for="(time, index) in aTime"
            :key="index"
            @click="loadASite(time, time.flag)"
          >
            <div class="diaty-time">
              <el-icon class="journal-icon"
                ><ArrowRight :class="{ icon: true, rotate: time.flag }"
              /></el-icon>
              <p>{{ time.date }}</p>
              <p>{{ time.text }}</p>
            </div>
            <div v-if="time.flag" class="diary-detailed">
              <p v-for="(site, index) in time.diary" :key="index">
                {{ site }}
              </p>
            </div>
          </li>
        </ul>
      </el-scrollbar>
    </section>
    <section>
      <h3>访客日志</h3>
      <el-scrollbar
        class="el-scrollbar"
        v-infinite-scroll="getUTimes"
        :infinite-scroll-disabled="uLoading"
      >
        <ul>
          <li
            v-for="(time, index) in uTime"
            :key="index"
            @click="loadUSite(time, time.flag)"
          >
            <div class="diaty-time">
              <el-icon class="journal-icon"
                ><ArrowRight :class="{ icon: true, rotate: time.flag }"
              /></el-icon>
              <p>{{ time.date }}</p>
              <p>{{ time.text }}</p>
            </div>
            <div v-if="time.flag" class="diary-detailed">
              <p v-for="(site, index) in time.diary" :key="index">
                {{ site }}
              </p>
            </div>
          </li>
        </ul>
      </el-scrollbar>
    </section>
  </main>
  <main class="message">
    <section>
      <h3>留言</h3>
      <el-scrollbar
        class="el-scrollbar"
        v-infinite-scroll="getMessage"
        :infinite-scroll-disabled="mLoading"
      >
        <div class="message-container">
          <el-card
            class="box-card"
            v-for="(message, index) in messages"
            :key="index"
          >
            <template #header>
              <div class="card-header">
                <span>{{ message.name }}</span>
                <el-button
                  class="button"
                  text
                  @click="deleteMessage(message.id)"
                  >删除</el-button
                >
              </div>
            </template>
            <div class="card-text">
              <p>{{ message.text }}</p>
            </div>
          </el-card>
        </div>
      </el-scrollbar>
    </section>
  </main>
</template>

<script>
import { computed, onMounted, ref, watch } from "@vue/runtime-core";
import { useStore } from "vuex";
export default {
  setup() {
    const store = useStore();
    const aLoading = ref(false);
    const uLoading = ref(false);
    const mLoading = ref(false);
    const aTime = computed(() => store.state.Diary.aTime);
    const aTimeSum = computed(() => store.state.Diary.aTimeSum);
    const uTime = computed(() => store.state.Diary.uTime);
    const uTimeSum = computed(() => store.state.Diary.uTimeSum);
    const messages = computed(() => store.state.Diary.messages);
    const messageSum = computed(() => store.state.Diary.messageSum);
    const getATimes = () => {
      aLoading.value = true;
      if (aTimeSum.value > aTime.value.length) {
        store.dispatch("getATime", [
          aLoading,
          { dType: 0, page: aTime.value.length },
        ]);
      }
    };
    const getUTimes = () => {
      uLoading.value = true;
      if (uTimeSum.value > uTime.value.length) {
        store.dispatch("getUTime", [
          uLoading,
          { dType: 1, page: uTime.value.length },
        ]);
      }
    };
    const getMessage = () => {
      mLoading.value = true;
      if (messageSum.value > messages.value.length) {
        store.dispatch("getMessages", [messages.value.length, mLoading]);
      }
    };
    const loadASite = (site, flag) => {
      site.flag = !flag;
      // 留个bug
      getASite(site);
    };
    const loadUSite = (site, flag) => {
      site.flag = !flag;
      // 留个bug
      getUSite(site);
    };
    const getASite = (time) => {
      if (time.sum > time.diary.length) {
        time.loading = true;
        store.dispatch("getADiary", time);
      }
    };
    const getUSite = (time) => {
      if (time.sum > time.diary.length) {
        time.loading = true;
        store.dispatch("getUDiary", time);
      }
    };
    const deleteMessage = (id) => {
      store.dispatch("deleteMessage", [
        id,
        () => {
          ElMessage("删除成功");
        },
        () => {
          ElMessage("删除失败");
        },
      ]);
    };
    watch(aTimeSum, () => {
      getATimes();
    });
    watch(uTimeSum, () => {
      getUTimes();
    });
    watch(messageSum, () => {
      getMessage();
    });
    onMounted(() => {
      store.dispatch("getASum");
      store.dispatch("getUSum");
      store.dispatch("getMessageSum");
    });
    return {
      aTime,
      aTimeSum,
      aLoading,
      uTime,
      uTimeSum,
      uLoading,
      messages,
      messageSum,
      mLoading,
      getATimes,
      loadASite,
      getASite,
      getUTimes,
      loadUSite,
      getUSite,
      getMessage,
      deleteMessage,
    };
  },
};
</script>

<style lang="scss" scoped>
.diary {
  height: 25rem;
}
main {
  display: flex;
  min-height: 23.5rem;
  & > section:nth-child(1) {
    margin-right: 1rem;
  }
}
section {
  padding: 1rem;
  position: relative;
  flex: 1;
  background-color: white;
  overflow: hidden;
  box-shadow: 0px 0px 1px 0px darkgray;
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
}
h3 {
  margin-bottom: 0.5rem;
}
ul {
  padding-left: 0.5rem;
}
.diaty-time {
  display: flex;
  align-items: center;
  padding: 0.5rem 0rem;
  p {
    white-space: nowrap;
    overflow: hidden;
  }
  & > p:nth-child(2) {
    width: 8rem;
  }
  & > p:nth-child(3) {
    text-overflow: ellipsis;
    font-size: 0.8rem;
    opacity: 0.7;
  }
  &:hover {
    cursor: pointer;
    & > p:nth-child(2),
    .journal-icon {
      color: darkcyan;
    }
  }
}
.journal-icon {
  margin-right: 0.2rem;
}
.message {
  margin-top: 1rem;
  animation: load-bottom 1s;
  & > section {
    margin-right: 0rem !important;
  }
}
.el-scrollbar {
  background-color: white;
}
.icon {
  transition-duration: 0.5s;
}
.rotate {
  transform: rotateZ(90deg);
}
.diary-detailed {
  max-height: 100%;
  animation: loadSite 0.5s;
  p {
    padding: 0.2rem;
    font-size: 0.9rem;
    color: darkcyan;
  }
}
.message-container {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 1rem;
}
.box-card {
  animation: site-load 1s;
}
.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}
.card-text {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 6rem;
  p {
    word-break: break-all;
    text-align: center;
  }
}
@keyframes loadSite {
  0% {
    transform: scale(0);
  }
  100% {
    transform: scale(1);
  }
}

@media screen and (max-width: 960px) {
  .message-container {
    grid-template-columns: repeat(3, 1fr);
  }
}

@media screen and (max-width: 780px) {
  main {
    min-width: 35rem;
  }
  .message-container {
    grid-template-columns: repeat(2, 1fr);
  }
}

@media screen and (max-width: 560px) {
  main {
    min-width: 23.5rem;
  }
  .diary {
    flex-direction: column;
    & > section:nth-child(1) {
      margin-right: 0rem;
      margin-bottom: 1rem;
    }
  }
  .message-container {
    grid-template-columns: repeat(1, 1fr);
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