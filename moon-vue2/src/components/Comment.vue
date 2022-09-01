<template>
  <div>
    <div class="comment-show">
      <p class="nobady" v-if="commentSum == 0">空荡荡~~</p>
      <div
        class="comment-site"
        v-for="(comment, index) in comments"
        :key="index"
        @click="selectUser(comment.id, comment.name)"
      >
        <label style="color: darkcyan">{{ comment.name }}</label>
        <label style="color: darkcyan" v-if="comment.cflag">{{
          " 回复 " + comment.cname
        }}</label>
        <label>{{ " : " + comment.content }}</label>
        <label style="margin-left: 0.5rem; font-size: 0.8rem; opacity: 0.9">{{
          comment.days
        }}</label>
      </div>
      <button
        v-if="comments.length != commentSum"
        class="comment-load"
        @click="getComment"
      >
        加载
      </button>
    </div>
    <p style="margin: 0.5rem 1rem; color: #f06292" v-if="cflag">
      @{{ cname
      }}<label @click="delUser" style="margin-left: 0.5rem; font-size: 0.8rem"
        >取消</label
      >
    </p>
    <div class="comment-input">
      <input v-model="name" placeholder="昵称" />
      <input v-model="email" placeholder="邮箱" />
      <img src="../assets/icon/user.png" />
      <img src="../assets/icon/email.png" />
      <div
        ref="content"
        class="comment-input-text"
        contenteditable="true"
      ></div>
      <img src="../assets/image/fire-keeper.png" />
      <button @click="submit">发言</button>
    </div>
  </div>
</template>

<script>
import { formatdays } from "@/utils/format";
export default {
  name: "Comment",
  props: ["sum", "title"],
  data() {
    return {
      commentSum: 0,
      comments: [],
      name: "",
      email: "",
      cflag: false,
      cname: "",
      cid: 0,
    };
  },
  mounted() {},
  methods: {
    getComment() {
      this.$axios
        .get("/article/comments", {
          title: this.title,
          index: this.comments.length,
        })
        .then((Res) => {
          if (Res.res == 1) {
            formatdays(Res.comments);
            this.comments = Res.comments;
          }
        })
        .catch((Err) => {
          console.log(Err);
        });
    },
    submit() {
      if (
        this.name == "" ||
        this.email == "" ||
        this.$refs.content.innerHTML == ""
      )
        return;
      const form = {
        title: this.title,
        id: this.comments.length + 1,
        time: new Date(),
        content: this.$refs.content.innerHTML,
        name: this.name,
        email: this.email,
        hflag: 1,
        cflag: this.cflag ? 1 : 0,
        cid: this.cid,
        cname: this.cname,
      };
      this.$axios
        .post("/article/submit", form)
        .then((Res) => {
          if (Res.res == 1) {
            if (this.commentSum == this.comments.length)
              this.comments.push(form);
            this.commentSum++;
          }
        })
        .catch((Err) => {
          console.log(Err);
        });
    },
    selectUser(id, name) {
      this.cflag = true;
      this.cname = name;
      this.cid = id;
    },
    delUser() {
      this.cflag = false;
      this.cname = "";
      this.cid = 0;
    },
  },
  watch: {
    sum() {
      this.commentSum = this.sum;
    },
    title() {
      this.getComment();
    },
  },
};
</script>

<style lang="scss" scoped>
.comment-show {
  position: relative;
  margin: 0rem 1rem 1rem;
  padding: 0.5rem 1rem;
  background-color: rgba($color: lightblue, $alpha: 0.3);
  box-shadow: 0px 0px 1px 0px lightblue;
  border-radius: 0.5rem;
  overflow: hidden;
  .nobady {
    padding: 2rem;
    text-align: center;
    color: #f06292;
  }
}
.comment-load {
  position: relative;
  left: -1rem;
  width: calc(100% + 2rem);
  height: 1rem;
  background-color: transparent;
  border: none;
  color: #f06292;
  cursor: pointer;
}
.comment-site {
  padding: 0.2rem 0rem;
  cursor: pointer;
}
.comment-input {
  position: relative;
  padding: 0rem 1rem 1rem;
  & > input {
    padding-left: 2.5rem;
    padding-right: 0.5rem;
    height: 2.5rem;
    line-height: 1.5rem;
    width: calc(50% - 1rem);
    outline: none;
    border: 1px solid dodgerblue;
    border-radius: 0.2rem;
  }
  & > input:nth-child(1) {
    margin-right: 2rem;
  }
  & > img {
    position: absolute;
    height: 1.5rem;
  }
  & > img:nth-child(3) {
    top: 0.5rem;
    left: 1.5rem;
  }
  & > img:nth-child(4) {
    top: 0.5rem;
    left: calc(50% + 1.5rem);
  }
  & > img:nth-last-child(2) {
    height: 10rem;
    top: 5rem;
    right: 1rem;
  }
  button {
    margin-top: 1rem;
    width: 100%;
    height: 2.5rem;
    border-radius: 0.2rem;
    background-color: darkcyan;
    border: none;
    color: white;
    font-size: 1rem;
    letter-spacing: 0.5rem;
    cursor: pointer;
  }
}
.comment-input-text {
  margin-top: 1rem;
  padding: 1rem 10rem 1rem 1rem;
  height: 12rem;
  border: 1px solid dodgerblue;
  border-radius: 0.2rem;
  outline: none;
  line-height: 1.5rem;
  font-size: 1rem;
  letter-spacing: 0.05rem;
  overflow: auto;

  &::-webkit-scrollbar {
    width: 5px;
  }

  &::-webkit-scrollbar-button {
    display: none;
  }

  &::-webkit-scrollbar-thumb {
    background-color: lightskyblue;
    border-radius: 0.2rem;
  }
}
</style>