<template>
  <main>
    <section class="note-input-container">
      <h3 class="note-input-title">日程安排</h3>
      <div class="note-input-site">
        <label>时间:</label>
        <el-date-picker
          v-model="nDate"
          type="date"
          value-format="YYYY-MM-DD"
          placeholder="Select date and time"
        />
      </div>
      <div class="note-input-site">
        <label>安排:</label>
        <el-input
          v-model="nText"
          :autosize="{ minRows: 4, maxRows: 4 }"
          type="textarea"
          placeholder="Please input"
        />
      </div>
      <img
        class="note-icon"
        src="../../assets/image/fire-keeper.png"
        @click="setNote"
      />
    </section>
    <section>
      <el-scrollbar
        class="el-scr"
        v-infinite-scroll="getNotes"
        :infinite-scroll-disabled="loading"
      >
        <div class="un-note-site">
          <label>{{ oldNote.time }}</label>
          <p>{{ oldNote.text }}</p>
        </div>
        <div class="note-site" v-for="(note, index) in notes" :key="index">
          <label>{{ note.time }}</label>
          <p>{{ note.text }}</p>
          <p><label @click="deleteNote(note.id)">取消</label></p>
        </div>
      </el-scrollbar>
    </section>
  </main>
</template>

<script>
import { ref } from "@vue/reactivity";
import { SetNote } from "@/api/apis/mood";
import { computed, onMounted, watch } from "@vue/runtime-core";
import { useStore } from "vuex";
export default {
  setup() {
    const store = useStore();
    const nDate = ref(null);
    const nText = ref("");
    const loading = ref(false);
    const sum = computed(() => store.state.Mood.noteSum);
    const oldNote = computed(() => store.state.Mood.oldNote);
    const notes = computed(() => store.state.Mood.notes);
    const setNote = () => {
      if (nDate.value != null && nText.value != null) {
        const form = new FormData();
        form.append("date", nDate.value);
        form.append("text", nText.value);
        SetNote(
          form,
          (Res) => {
            if (Res.res == 1) ElMessage("添加成功");
            else ElMessage("添加失败");
          },
          (Err) => {
            console.log(Err);
            ElMessage("添加失败");
          }
        );
      } else ElMessage("请填写完整资料");
    };
    const setNoteSum = () => {
      store.dispatch("getNoteSum");
    };
    const getNotes = () => {
      if (notes.value.length < sum.value) {
        loading.value = true;
        store.dispatch("getNotes", loading);
      }
    };
    const deleteNote = (id) => {
      store.dispatch("deleteNote", [
        id,
        () => {
          ElMessage("取消成功");
        },
        () => {
          ElMessage("取消失败");
        },
      ]);
    };
    watch(sum, () => {
      getNotes();
    });
    onMounted(() => {
      store.dispatch("getOldNote");
      setNoteSum();
    });
    return {
      nDate,
      nText,
      oldNote,
      notes,
      loading,
      setNote,
      getNotes,
      deleteNote,
    };
  },
};
</script>

<style lang="scss" scoped>
main {
  display: flex;
  min-width: 23.5rem;
  section {
    position: relative;
    padding: 1rem;
    background-color: white;
    height: 80vh;
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
      z-index: 2;
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
      z-index: 2;
    }
  }
  & > section:nth-child(1) {
    flex: 1;
  }
  & > section:nth-child(2) {
    margin-left: 1rem;
    width: 25rem;
  }
}
.note-input-container {
  padding: 5rem 1rem 10rem;
  display: flex;
  flex-direction: column;
  justify-content: center;
}
.note-input-site {
  padding-bottom: 1rem;
  display: flex;
  label {
    margin-right: 0.5rem;
    word-break: keep-all;
  }
}
.note-input-title {
  position: absolute;
  top: 1rem;
  color: darkcyan;
}
.note-icon {
  position: absolute;
  height: 10rem;
  bottom: 0rem;
  right: 0rem;
}
.el-scr {
  background-color: white;
}
.note-site,
.un-note-site {
  margin: 0rem 0.5rem;
  margin-right: 1rem;
  position: relative;
  border-left: 1px solid darkcyan;
  animation: note-load 1.5s;
  & > label {
    margin-left: 0.5rem;
    color: darkcyan;
  }
  & > p:nth-child(2) {
    margin-left: 0.5rem;
    padding: 0.5rem;
    min-height: 4rem;
    background-color: rgb(216, 222, 229);
    border-radius: 0.2rem;
  }
  & > p:nth-child(3) {
    text-align: right;
    label {
      font-size: 0.8rem;
      color: darkcyan;
      letter-spacing: 0.1rem;
      cursor: pointer;
    }
  }
  &::after {
    content: "";
    display: block;
    position: absolute;
    top: 0.3rem;
    left: -0.4rem;
    height: 0.6rem;
    width: 0.6rem;
    background-color: white;
    border: 1px solid darkcyan;
    border-radius: 0.6rem;
  }
}
.un-note-site {
  border-left: 1px solid darkgray;
  label {
    color: darkgray;
  }
  &::after {
    border: 1px solid darkgray;
  }
}
@media screen and (max-width: 780px) {
  main > section:nth-child(2) {
    width: auto;
    flex: 1;
  }
}
@media screen and (max-width: 780px) {
  main {
    flex-direction: column;
    & > section:nth-child(2) {
      margin-left: 0rem;
      margin-top: 1rem;
    }
  }
}

@keyframes note-load {
  0% {
    transform: translateX(-100%);
    opacity: 0;
  }
  50% {
    transform: translateX(-100%);
    opacity: 0;
  }
  100% {
    transform: translateX(0%);
    opacity: 1;
  }
}
</style>