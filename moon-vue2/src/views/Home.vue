<template>
  <div class="home">
    <cover class="home-cover" :sentence="info.sentence"></cover>
    <div class="home-main">
      <h3 class="title">home</h3>
      <main>
        <article>
          <sign></sign>
          <card
            :pageNow="pageNow"
            :pageSum="pageSum"
            :articleList="articleList"
            :closeFlag="closeFlag"
            v-on:changePage="changePage"
          ></card>
        </article>
        <aside>
          <div class="home-aside">
            <user :info="info"></user>
            <tag></tag>
          </div>
        </aside>
      </main>
    </div>
  </div>
</template>

<script>
import Cover from "@/components/Cover.vue";
import Card from "@/components/Card.vue";
import Sign from "@/components/Sign.vue";
import User from "@/components/User.vue";
import Tag from "@/components/Tag.vue";
import { formatTime, formatTpye, formatTags } from "@/utils/format.js";
export default {
  name: "Home",
  components: { Cover, Card, Sign, User, Tag },
  data() {
    return {
      info: {},
      pageNow: 0,
      pageSum: 1,
      articleList: [],
      closeFlag: false,
    };
  },
  mounted() {
    this.getInfo();
    this.getRoutePage();
    this.getPageSum();
  },
  beforeRouteUpdate(to, from, next) {
    next();
    this.articleList = [];
    this.closeFlag = false;
    this.getRoutePage();
    this.getPageSum();
  },
  methods: {
    getInfo() {
      this.$axios
        .get("/user/info")
        .then((Res) => {
          if (Res.res == 1) {
            this.info = Res.info;
          }
        })
        .catch((Err) => {
          console.log(Err);
        });
    },
    getRoutePage() {
      if (this.$route.params.page != null) {
        this.pageNow = +this.$route.params.page;
      }
    },
    getPageSum() {
      this.$axios
        .get("article/articleSum")
        .then((Res) => {
          if (Res.res == 1) {
            this.pageSum = Math.ceil(Res.sum / 6);
            this.pageSum = this.pageSum > 0 ? this.pageSum : 1;
            if (this.pageNow > this.pageSum) {
              this.pageNow = 1;
              this.$router.push("/Home/1");
            }
          }
          this.getArticleList();
        })
        .catch((Err) => {
          console.log(Err);
        });
    },
    getArticleList() {
      this.$axios
        .get("/article/articleList", { page: this.pageNow - 1 })
        .then((Res) => {
          if (Res.res == 1) {
            this.formatList(Res.list);
            this.articleList = Res.list;
          }
        })
        .catch((Err) => {
          console.log(Err);
        });
    },
    formatList(list) {
      formatTime(list);
      formatTpye(list);
      formatTags(list);
    },
    changePage(page) {
      this.pageNow = page;
      this.closeFlag = true;
      setTimeout(() => {
        this.$router.push({
          name: "Home",
          params: {
            page: page,
          },
        });
      }, 200);
    },
  },
};

Card;
</script>

<style lang="scss" scoped>
.title {
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
}
.home-cover {
  animation: load-up 1s;
}
.home-main {
  animation: load-down 1s;
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
