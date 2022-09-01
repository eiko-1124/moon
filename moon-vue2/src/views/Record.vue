<template>
  <div>
    <header :style="{ 'background-image': 'url(' + background + ')' }"></header>
    <div class="record-container">
      <h3>记录</h3>
      <main>
        <article>
          <div v-for="(mood, index) in moods" :key="index" ref="moods">
            <div class="record-box">
              <p style="padding: 2rem 1rem 0rem; text-align: center">
                {{ mood.text }}
              </p>
              <div class="record-show">
                <div
                  v-for="(image, index) in mood.images"
                  :key="index"
                  style="
                    padding-bottom: 60%;
                    height: 0rem;
                    overflow: hidden;
                    background-size: cover;
                    background-position: center center;
                  "
                  :style="{ backgroundImage: 'url(' + image + ')' }"
                ></div>
              </div>
              <h3>{{ mood.time }}</h3>
              <div class="record-user">
                <div class="record-head">
                  <img
                    src="http://localhost:3000/static/portrait/mood-head.jpg"
                  />
                </div>
                <label>Mn+O2</label>
                <div class="record-option">
                  <label :class="{ 'record-like': mood.likeFlag }"
                    >&#xe8ab;{{ mood.mlike }}</label
                  >
                  <label>&#xe8b4;{{ mood.comment }}</label>
                  <label>&#xf0141;</label>
                  <div class="record-button">
                    <label v-if="!mood.likeFlag" @click="setLike(mood)"
                      >点赞</label
                    >
                    <label v-if="mood.likeFlag" @click="deleteLike(mood)"
                      >取消点赞</label
                    >
                    <label v-if="!mood.commentFlag" @click="openComment(index)"
                      >评论</label
                    >
                    <label v-if="mood.commentFlag" @click="openComment(index)"
                      >收起评论</label
                    >
                  </div>
                </div>
              </div>
            </div>
            <div class="record-comment" v-if="mood.commentFlag">
              <div
                v-for="(comment, mindex) in mood.comments"
                :key="mindex"
                @click="setCname(comment.name, index)"
              >
                <label class="comment-name">{{ comment.name }}</label
                ><label class="comment-name" v-if="comment.cflag == 1">{{
                  " 回复 " + comment.cname
                }}</label
                ><label style="word-break: break-all">{{
                  " : " + comment.content
                }}</label>
                <label
                  style="margin-left: 1rem; font-size: 0.8rem; color: gray"
                  >{{ comment.days }}</label
                >
              </div>
              <div
                class="record-comment-more"
                v-if="mood.loadFlag && mood.comments.length < mood.comment"
              >
                <button v-if="!mood.loadBnt" @click="moreComment(index)">
                  更多
                </button>
                <button class="load-bnt" v-if="mood.loadBnt">&#xe629;</button>
              </div>
              <div class="record-input" v-if="mood.inputFlag">
                <div>
                  <input ref="name" placeholder="昵称" /><button
                    @click="submit(index)"
                  >
                    评论
                  </button>
                  <label
                    ref="cname"
                    v-if="mood.cname != ''"
                    style="
                      margin: 0rem 1rem;
                      font-size: 0.8rem;
                      color: darkcyan;
                    "
                    >{{ "@" + mood.cname
                    }}<button
                      style="
                        font-size: 0.8rem;
                        color: darkcyan;
                        background-color: transparent;
                        width: 2rem;
                      "
                      @click="deleteCname(index)"
                    >
                      取消
                    </button></label
                  >
                </div>
                <textarea ref="content" placeholder="评论"></textarea>
              </div>
            </div>
          </div>
        </article>
        <aside>
          <div class="home-aside">
            <section><plan></plan></section>
            <section><recent></recent></section>
          </div>
        </aside>
      </main>
    </div>
  </div>
</template>

<script>
import Recent from "@/components/Recent.vue";
import Plan from "@/components/Plan.vue";
import { formatDate, formatdays } from "@/utils/format";
import { updateLike, formatLikes } from "@/utils/localLikes";
export default {
  name: "Record",
  components: {
    Plan,
    Recent,
  },
  data() {
    return {
      moods: [],
      randerMoods: [],
    };
  },
  computed: {
    background() {
      return this.$route.meta.background;
    },
  },
  mounted() {
    this.addObserver();
  },
  methods: {
    getMoods() {
      return new Promise((resolve, reject) => {
        this.$axios
          .get("/note/moods", { index: this.moods.length })
          .then((Res) => {
            if (Res.res == 1) {
              this.format(Res.moods);
              this.randerMoods = Res.moods;
              resolve();
            }
          })
          .catch((Err) => {
            console.log(Err);
            reject();
          });
      });
    },
    addObserver() {
      this.getMoods().then(() => {
        if (this.randerMoods.length != 0) {
          this.moods.push(this.randerMoods.shift());
          this.$nextTick(() => {
            const moods = this.$refs.moods;
            this.moodObserver(moods[moods.length - 1]);
          });
        }
      });
    },
    moodObserver(target) {
      let options = {
        root: null,
        rootMargin: "0px",
        threshold: 1,
      };
      let observer = new IntersectionObserver((entries) => {
        entries.forEach((entry) => {
          if (entry.isIntersecting) {
            observer.unobserve(target);
            const moodNotes = this.$refs.moods;
            this.getImage(moodNotes.length - 1);
            if (this.randerMoods.length != 0) {
              this.moods.push(this.randerMoods.shift());
              this.$nextTick(() => {
                this.moodObserver(moodNotes[moodNotes.length - 1]);
              });
            } else this.addObserver();
          }
        });
      }, options);
      observer.observe(target);
    },
    format(list) {
      formatDate(list);
      formatLikes(list);
      list.forEach((site) => {
        site.images = [];
        site.comments = [];
        site.randerComments = [];
        site.commentFlag = false;
        site.loadFlag = false;
        site.loadBnt = false;
        site.inputFlag = false;
        site.cname = "";
      });
    },
    getImage(index) {
      if (this.moods[index].pic != null) {
        this.$axios
          .get("/note/image", { pic: this.moods[index].pic })
          .then((Res) => {
            if (Res.res == 1) {
              this.moods[index].images = Res.paths;
            }
          })
          .catch((Err) => {
            console.log(Err);
          });
      }
    },
    async openComment(index) {
      const target = this.moods[index];
      target.commentFlag = !target.commentFlag;
      if (target.commentFlag) {
        if (target.randerComments.length == 0) await this.getComment(index);
        const timer = setInterval(() => {
          if (target.randerComments.length == 0) {
            target.loadFlag = true;
            setTimeout(() => {
              target.inputFlag = true;
            }, 200);
            clearInterval(timer);
          } else {
            target.comments.push(target.randerComments.shift());
          }
        }, 200);
      } else {
        target.randerComments = target.comments;
        target.comments = [];
        target.loadFlag = false;
        target.inputFlag = false;
      }
    },
    async moreComment(index) {
      const target = this.moods[index];
      target.loadBnt = true;
      await this.getComment(index);
      target.loadFlag = false;
      const timer = setInterval(() => {
        if (target.randerComments.length == 0) {
          target.loadFlag = true;
          target.loadBnt = false;
          clearInterval(timer);
        } else {
          target.comments.push(target.randerComments.shift());
        }
      }, 200);
    },
    getComment(index) {
      return new Promise((resolve, reject) => {
        const num =
          this.moods[index].comments.length +
          this.moods[index].randerComments.length;
        if (num < this.moods[index].comment) {
          this.$axios
            .get("/note/comment", { index: num, id: this.moods[index].id })
            .then((Res) => {
              if (Res.res == 1) {
                formatdays(Res.comments);
                this.moods[index].randerComments = Res.comments;
                resolve();
              } else reject();
            })
            .catch((Err) => {
              console.log(Err);
              reject();
            });
        } else {
          resolve();
        }
      });
    },
    setCname(name, index) {
      this.moods[index].cname = name;
    },
    deleteCname(index) {
      this.moods[index].cname = "";
    },
    submit(index) {
      if (
        this.$refs.name[index].value == "" ||
        this.$refs.content[index].value == ""
      )
        return;
      const comment = {
        name: this.$refs.name[index].value,
        content: this.$refs.content[index].value,
        cflag: this.moods[index].cname.length == 0 ? 0 : 1,
        cname: this.moods[index].cname,
        time: new Date(),
      };
      this.$axios
        .post("/note/submit", {
          id: this.moods[index].id,
          name: comment.name,
          content: comment.content,
          cflag: comment.cflag,
          cname: comment.cname,
          time: comment.time,
        })
        .then((Res) => {
          if (Res.res == 1) {
            this.$refs.name[index].value = "";
            this.$refs.content[index].value = "";
            if (this.moods[index].comments.length == this.moods[index].comment)
              this.moods[index].comments.push(comment);
            this.moods[index].comment++;
            formatdays(this.moods[index].comments);
          }
        })
        .catch((Err) => {
          console.log(Err);
        });
    },
    setLike(mood) {
      updateLike(mood.id);
      mood.likeFlag = true;
      mood.mlike++;
      this.updateLike(mood.id, mood.mlike);
    },
    deleteLike(mood) {
      updateLike(mood.id);
      mood.likeFlag = false;
      mood.mlike--;
      this.updateLike(mood.id, mood.mlike);
    },
    updateLike(id, like) {
      this.$axios
        .get("/note/like", { id, like })
        .then((Res) => {
          if (Res.res != 1) console.log("点赞失败");
        })
        .catch((Err) => {
          console.log(Err);
        });
    },
  },
};
</script>

<style lang="scss" scoped>
header {
  position: relative;
  display: flex;
  justify-content: center;
  align-items: center;
  height: 75vh;
  background-size: cover;
  background-position: center center;
  animation: load-up 1s;
}
.record-container,
.record-box {
  & > h3 {
    height: 4rem;
    display: flex;
    align-items: center;
    justify-content: center;
    padding: 1rem;
    font-weight: 500;
    color: darkcyan;
    &::before,
    &::after {
      content: "";
      margin: 0.5rem;
      height: 1px;
      width: 4rem;
      background-color: darkcyan;
    }
  }
}
.record-container {
  animation: load-down 1s;
}
main {
  display: flex;
  justify-content: center;
  width: 100%;
  background-image: linear-gradient(
    180deg,
    white,
    rgba(135, 206, 235, 0.4) 350px
  );
}
article {
  margin: 1rem;
  flex: 3;
  max-width: 50rem;
}
aside {
  position: relative;
  margin: 1rem;
  flex: 1;
  max-width: 25rem;
  z-index: 2;
}
.home-aside {
  position: sticky;
  top: 1rem;
  section {
    margin-bottom: 1rem;
  }
}
.record-box {
  position: relative;
  margin-bottom: 1rem;
  border-radius: 0.5rem;
  box-shadow: rgba(9, 30, 66, 0.25) 0px 4px 8px -2px,
    rgba(9, 30, 66, 0.08) 0px 0px 0px 1px;
  background-image: linear-gradient(
    135deg,
    rgba(255, 255, 255, 1),
    rgba(255, 255, 255, 0.7)
  );
  z-index: 2;
  animation: mood-load 1s;
  h3 {
    padding: 1rem 1rem 0rem;
    justify-content: right;
    height: 3rem;
  }
}
.record-user {
  position: relative;
  padding: 0rem 1rem 1rem;
  display: flex;
  align-items: center;
  color: darkcyan;
}
.record-head {
  margin-right: 0.5rem;
  height: 3rem;
  width: 3rem;
  overflow: hidden;
  border-radius: 1.5rem;
  img {
    height: 100%;
  }
}
.record-option {
  position: absolute;
  right: 1rem;
  bottom: 1rem;
  width: 10rem;
  text-align: right;
  overflow: hidden;
  label {
    margin-right: 0.5rem;
    display: inline-block;
    height: 2rem;
    line-height: 2rem;
    vertical-align: bottom;
    color: gray;
  }
  & > label:nth-child(3) {
    font-size: 1.5rem;
    color: darkcyan;
  }
  &:hover {
    .record-button {
      transform: translateX(0%);
    }
  }
}
.record-like {
  color: red !important;
}
.record-button {
  position: absolute;
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  top: 0rem;
  height: 100%;
  width: 100%;
  background-color: cadetblue;
  border-radius: 0.1rem;
  transition-duration: 0.3s;
  transform: translateX(100%);
  label {
    margin: 0rem;
    color: white;
    text-align: center;
    cursor: pointer;
  }
  & > label:nth-child(1) {
    border-right: 1px solid white;
  }
}
.record-show {
  padding: 1rem 1rem 0rem;
  display: grid;
  grid-template: auto / repeat(3, 1fr);
  gap: 0.5rem;
}
.record-comment {
  margin-bottom: 1rem;
  border-radius: 0.2rem;
  box-shadow: rgba(9, 30, 66, 0.25) 0px 4px 8px -2px,
    rgba(9, 30, 66, 0.08) 0px 0px 0px 1px;
  overflow: hidden;
  text-decoration: 0.1s;
  & > div {
    padding: 0.5rem;
    background-image: linear-gradient(
      135deg,
      rgba(255, 255, 255, 1),
      rgba(255, 255, 255, 0.7)
    );
    overflow: hidden;
    animation: comment-load 0.2s linear;
  }
}
.comment-name {
  color: darkcyan;
}
.record-input {
  & > div {
    display: flex;
    align-items: center;
  }
  input,
  textarea {
    outline: none;
    color: darkcyan;
    border: 1px solid darkcyan;
  }
  input {
    padding: 0rem 0.2rem;
    margin: 0.2rem;
    margin-left: 0rem;
    height: 1.8rem;
    width: 8rem;
  }
  button {
    height: 1.8rem;
    width: 4rem;
    background-color: darkcyan;
    color: white;
    border: none;
    cursor: pointer;
  }
  textarea {
    padding: 0.2rem;
    width: 100%;
    height: 5rem;
    resize: none;
    letter-spacing: 0.05rem;
    &::-webkit-scrollbar {
      width: 0.3rem;
    }
    &::-webkit-scrollbar-thumb {
      background-color: lightgray;
    }
  }
}
.record-comment-more {
  button {
    background-color: transparent;
    border: none;
    color: darkcyan;
  }
  & > button:nth-child(1) {
    cursor: pointer;
  }
  .load-bnt {
    animation: load-bnt 1s linear infinite;
  }
}

@keyframes mood-load {
  0% {
    opacity: 0;
    transform: translateY(100%);
  }
  100% {
    opacity: 1;
    transform: translateY(0%);
  }
}

@keyframes comment-load {
  0% {
    opacity: 0;
    padding: 0rem;
    max-height: 0rem;
  }
  100% {
    opacity: 1;
    padding: 0.5rem;
    max-height: 10rem;
  }
}

@keyframes load-bnt {
  0% {
    transform: rotateZ(0deg);
  }
  100% {
    transform: rotateZ(360deg);
  }
}

@media screen and (max-width: 1080px) {
  article {
    flex: 2;
  }
  aside {
    flex: 1;
  }
}

@media screen and (max-width: 759px) {
  aside {
    display: none;
  }
}
</style>