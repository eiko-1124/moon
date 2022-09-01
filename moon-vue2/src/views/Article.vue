<template>
  <div class="article">
    <header :style="{ backgroundImage: 'url(' + article.cover + ')' }">
      <h1 style="color: white">{{ article.title }}</h1>
      <div style="margin: 0.2rem; font-size: 0.9rem">
        <label style="color: white"
          ><label style="margin-right: 0.5rem">&#xe616;</label
          >{{ article.time }}</label
        >
        <span class="splice-line"></span>
        <label style="color: white"
          ><label style="margin-right: 0.5rem">&#xe695;</label
          >{{ comment }}</label
        >
      </div>
      <div style="margin: 0.2rem; font-size: 0.9rem" class="article-tags">
        <label v-for="(tag, index) in article.tagArr" :key="index">
          <label style="color: white"
            ><label style="margin-right: 0.5rem">&#xe87d;</label
            >{{ tag }}</label
          >
          <span class="splice-line"></span>
        </label>
      </div>
    </header>
    <div class="article-body">
      <h3>文章</h3>
      <main>
        <article>
          <section class="article-text">
            <h3>Article</h3>
            <p v-html="article.fcontent"></p>
          </section>
          <section class="article-recommend">
            <div class="recommend-left">
              <p>Previous Post<img src="../assets/icon/right-arrow.png" /></p>
              <p v-if="near.preId != 0">{{ near.preTitle }}</p>
              <label style="opacity: 0.8" v-if="near.preId == 0"
                >没有更多了</label
              >
            </div>
            <div class="recommend-right" style="text-align: right">
              <p>Next Post<img src="../assets/icon/left-arrow.png" /></p>
              <p v-if="near.nextId != 0">{{ near.nextTitle }}</p>
              <label style="opacity: 0.8" v-if="near.nextId == 0"
                >没有更多了</label
              >
            </div>
          </section>
          <section class="article-comment">
            <h3>评论</h3>
            <comment :sum="comment" :title="article.title"></comment>
          </section>
        </article>
        <aside>
          <div style="position: sticky; top: 1rem">
            <recent></recent>
            <amessage></amessage>
          </div>
        </aside>
      </main>
    </div>
  </div>
</template>

<script>
import Comment from "@/components/Comment.vue";
import Recent from "@/components/Recent.vue";
import Amessage from "@/components/Amessage.vue";
import axios from "@/axios";
import { formatTime, formatTpye, formatTags } from "@/utils/format.js";
export default {
  name: "Article",
  components: { Comment, Recent, Amessage },
  data() {
    return {
      article: {
        title: "",
      },
      comment: 0,
      comments: [],
      near: {},
    };
  },
  mounted() {
    document.body.scrollTo(0, 0);
    this.getNear();
  },
  beforeRouteEnter(to, from, next) {
    axios
      .get("/article/article", { id: to.params.id })
      .then((Res) => {
        if (Res.res == 1) {
          const image = new Image();
          image.src = Res.article.cover;
          image.onload = () => {
            next((vm) => {
              formatTime([Res.article]);
              formatTpye([Res.article]);
              formatTags([Res.article]);
              vm.article = Res.article;
              vm.comment = Res.comment;
              vm.$store.commit("setLoad", false);
              vm.$store.commit("setNav");
            });
          };
        }
      })
      .catch((Err) => {
        console.log(Err);
      });
  },
  methods: {
    getNear() {
      this.$axios
        .get("/article/near", { id: this.article.id })
        .then((Res) => {
          if (Res.res == 1) this.near = Res.near;
        })
        .catch((Err) => {
          console.log(Err);
        });
    },
  },
};
</script>

<style lang="scss" scoped>
.article {
  width: 100vw;
}
header {
  position: relative;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 25rem;
  background-size: cover;
  background-position: center center;
  animation: load-up 1s;
}
.article-body {
  position: relative;
  background-image: linear-gradient(
    180deg,
    rgba(255, 255, 255, 1),
    rgba(135, 206, 235, 0.4) 350px
  );
  animation: load-down 1s;
  & > h3 {
    margin-top: 1rem;
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

main {
  display: flex;
  justify-content: center;
  width: 100%;
}
article {
  position: relative;
  margin: 1rem;
  flex: 3;
  max-width: 50rem;
  z-index: 2;
  background-image: linear-gradient(
    135deg,
    rgba(255, 255, 255, 1),
    rgba(255, 255, 255, 0.65)
  );
  box-shadow: rgba(9, 30, 66, 0.25) 0px 4px 8px -2px,
    rgba(9, 30, 66, 0.08) 0px 0px 0px 1px;
  border-radius: 0.5rem;
  section {
    & > h3 {
      height: 4rem;
      display: flex;
      align-items: center;
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
}
.article-text,
.article-recommend {
  position: relative;
  padding: 1rem;
  &::after {
    display: block;
    content: "";
    position: absolute;
    bottom: 0rem;
    height: 1px;
    width: calc(100% - 2rem);
    background-image: radial-gradient(
      rgba(135, 206, 250, 0.8),
      rgba(135, 206, 250, 0.2),
      rgba(135, 206, 250, 0)
    );
  }
}
.article-text {
  p {
    padding: 0rem 1rem 1rem;
  }
}
.article-text > p ::v-deep hr {
  margin: 0.5rem 0rem;
  border-left: none;
  border-right: none;
  border-bottom: none;
}
.article-recommend {
  p {
    display: flex;
    align-items: center;
    transition-duration: 0.5s;
  }
  img {
    height: 1.5rem;
    transition-duration: 0.5s;
  }
}
.recommend-left {
  padding-left: 1rem;
  position: relative;
  display: inline-block;
  width: 50%;
  p {
    flex-direction: row;
  }
  &::after {
    display: block;
    content: "";
    position: absolute;
    right: 0rem;
    bottom: 0rem;
    height: 100%;
    width: 1px;
    background-image: radial-gradient(
      rgba(135, 206, 250, 1),
      rgba(135, 206, 250, 0)
    );
  }
}
.recommend-right {
  padding-right: 1rem;
  display: inline-block;
  width: 50%;
  p {
    flex-direction: row-reverse;
  }
}
.recommend-left,
.recommend-right {
  &:hover {
    img {
      margin: 0rem 1rem;
    }
    & > p:nth-child(2) {
      color: darkcyan;
      cursor: pointer;
    }
  }
}
.article-comment {
  position: relative;
  padding: 1rem;
}
aside {
  position: relative;
  margin: 1rem;
  flex: 1;
  max-width: 25rem;
  z-index: 2;
}
.splice-line {
  margin: 0rem 0.8rem;
  border-left: 1px solid white;
}
.article-tags > label:nth-last-child(1) {
  label {
    cursor: pointer;
  }
  span {
    display: none;
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