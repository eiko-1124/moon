<template>
  <div
    class="message"
    :style="{ 'background-image': 'url(' + background + ')' }"
  >
    <div class="message-xl">
      <section class="message-type">
        <div class="type-select" @click="typeOpen">
          <button v-show="typeFlag">&#xf019f;</button
          ><button v-show="!typeFlag">&#xe61d;</button>
        </div>
        <div
          class="type-unselect"
          :class="{ close: selectType }"
          @click="typeChoose"
        >
          <button v-show="!typeFlag">&#xf019f;</button
          ><button v-show="typeFlag">&#xe61d;</button>
        </div>
      </section>
      <section>
        <input v-model="name" class="in-name" placeholder="name" />
      </section>
      <section>
        <input v-model="text" class="in-text" type="text" placeholder="text" />
      </section>
      <section>
        <button class="message-submit" @click="submit">发送</button>
      </section>
    </div>
    <div class="message-sl">
      <div>
        <section>
          <input
            v-model="text"
            class="in-text"
            type="text"
            placeholder="text"
          />
        </section>
      </div>
      <div>
        <section class="message-type">
          <div class="type-select" @click="typeOpen">
            <button v-show="typeFlag">&#xf019f;</button
            ><button v-show="!typeFlag">&#xe61d;</button>
          </div>
          <div
            class="type-unselect"
            :class="{ close: selectType }"
            @click="typeChoose"
          >
            <button v-show="!typeFlag">&#xf019f;</button
            ><button v-show="typeFlag">&#xe61d;</button>
          </div>
        </section>
        <section>
          <input v-model="name" class="in-name" placeholder="name" />
        </section>
        <section>
          <button class="message-submit" @click="submit">发送</button>
        </section>
      </div>
    </div>
    <div class="damu">
      <vueDanmaku ref="danmaku" v-model="List" use-slot style="height: 100%">
        <template v-slot:dm="{ danmu }">
          <span v-if="danmu.mtype == 1"
            >&#xf019f;
            <p>{{ danmu.name + " : " }}</p>
            <p>{{ danmu.text }}</p>
            <p>{{ danmu.days }}</p></span
          >
          <span v-if="danmu.mtype == 2" :index="index"
            >&#xe61d;
            <p>{{ danmu.name + " : " }}</p>
            <p>{{ danmu.text }}</p>
            <p>{{ danmu.days }}</p></span
          >
        </template>
      </vueDanmaku>
    </div>
  </div>
</template>

<script>
import { formatdays } from "@/utils/format";
import vueDanmaku from "vue-danmaku";
export default {
  name: "Message",
  components: { vueDanmaku },
  data() {
    return {
      selectType: true,
      typeFlag: false,
      count: 0,
      List: [],
      text: "",
      name: "",
    };
  },
  computed: {
    background() {
      return this.$route.meta.background;
    },
  },
  mounted() {
    this.getMessage();
  },
  methods: {
    typeOpen() {
      this.selectType = !this.selectType;
    },
    typeChoose() {
      this.selectType = !this.selectType;
      this.typeFlag = !this.typeFlag;
    },
    getMessage() {
      this.$axios
        .get("/note/messages", { index: this.count, num: 200 })
        .then((Res) => {
          if (Res.res == 1) {
            formatdays(Res.messages);
            this.List = Res.messages;
          }
        })
        .catch((Err) => {
          console.log(Err);
        });
    },
    submit() {
      if (this.name == "" || this.text == "") return;
      const form = {
        name: this.name,
        text: this.text,
        mtype: this.typeFlag ? 1 : 2,
        time: new Date(),
      };
      this.$axios
        .post("/note/setMessage", form)
        .then((Res) => {
          if (Res.res == 1) {
            this.$refs.danmaku.add(form);
          }
        })
        .catch((Err) => {
          console.log(Err);
        });
    },
  },
};
</script>

<style lang="scss" scoped>
.message {
  height: 100vh;
  width: 100vw;
  background-size: cover;
  background-position: center center;
  overflow: hidden;
}
.message-xl {
  position: relative;
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100%;
  width: 100%;
}
.message-sl {
  position: relative;
  display: none;
  justify-content: center;
  align-items: center;
  flex-direction: column;
  height: 100%;
  width: 100%;
  & > div {
    margin: 0.5rem;
    display: flex;
  }
}
section {
  margin-right: 0.5rem;
  height: 3rem;
  input {
    padding: 1rem;
    height: 3rem;
    outline: none;
    border-radius: 0.4rem;
    font-size: 1rem;
    color: greenyellow;
    border: 1px solid lightblue;
    background-color: transparent;
  }
  .in-text {
    width: 20rem;
  }
  .in-name {
    width: 6rem;
  }
}
.type-select,
.type-unselect {
  margin-bottom: 0.5rem;
  height: 3rem;
  width: 3rem;
  overflow: hidden;
  button {
    height: 3rem;
    width: 3rem;
    color: greenyellow;
    font-size: 2rem;
    border-radius: 0.4rem;
    border: 1px solid lightblue;
    background-color: transparent;
  }
}
.type-unselect {
  transition-duration: 0.3s;
}
.message-submit {
  height: 3rem;
  width: 5rem;
  color: lightblue;
  border-radius: 0.4rem;
  border: 1px solid lightblue;
  font-size: 1rem;
  letter-spacing: 0.2rem;
  background-color: transparent;
}
.close {
  height: 0rem;
}
.damu {
  position: absolute;
  top: 4rem;
  width: 100%;
  height: 10rem;
}
span {
  display: flex;
  align-items: center;
  flex-wrap: nowrap;
  white-space: nowrap;
  color: greenyellow;
  font-size: 2rem;
  height: 2.5rem;
  text-shadow: 0.2rem 0rem 0.8rem lightblue;
  p {
    margin-left: 0.5rem;
    font-size: 1rem;
    color: lightblue;
    text-shadow: 0px 0px 1px greenyellow, 0px 0px 1px greenyellow;
  }
}
@keyframes move {
  0% {
    left: 100vw;
  }
  100% {
    left: -100vw;
  }
}

@media screen and (max-width: 759px) {
  .message-xl {
    display: none;
  }
  .message-sl {
    display: flex;
  }
}
</style>