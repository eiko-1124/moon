<template>
  <div>
    <header :style="{ 'background-image': 'url(' + background + ')' }"></header>
    <div class="classify-container">
      <h3>分类</h3>
      <main>
        <article>
          <div class="classify-select">
            <div>
              <label>分类:</label>
              <div>
                <button
                  v-for="index in types"
                  :key="index"
                  :class="{ 'select-bnt': typeOption[index].type == type }"
                  @click="typeSelect(typeOption[index].type)"
                >
                  {{ typeOption[index].text }}
                </button>
              </div>
            </div>
            <div>
              <label>标签:</label>
              <div>
                <button
                  v-for="(site, index) in tags"
                  :key="index"
                  :class="{ 'select-bnt': tag == site }"
                  @click="tagSelect(site)"
                >
                  {{ site }}
                </button>
              </div>
            </div>
          </div>
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
            <section>
              <tag></tag>
            </section>
            <section>
              <recent></recent>
            </section>
          </div>
        </aside>
      </main>
    </div>
  </div>
</template>

<script>
import Tag from "@/components/Tag.vue";
import Recent from "@/components/Recent.vue";
import Card from "@/components/Card.vue";
import { formatTime, formatTpye, formatTags } from "@/utils/format.js";
export default {
  name: "Classify",
  components: {
    Tag,
    Recent,
    Card,
  },
  data() {
    return {
      typeOption: [
        {
          type: 0,
          text: "全部",
        },
        {
          type: 1,
          text: "技术",
        },
        {
          type: 2,
          text: "杂文",
        },
        {
          type: 3,
          text: "故事",
        },
      ],
      types: [],
      type: 0,
      tags: [],
      tag: 0,
      pageNow: 1,
      pageSum: 1,
      articleList: [],
      closeFlag: false,
    };
  },
  computed: {
    background() {
      return this.$route.meta.background;
    },
  },
  mounted() {
    this.init();
  },
  beforeRouteUpdate(to, from, next) {
    next();
    this.init();
  },
  methods: {
    async init() {
      this.closeFlag = false;
      this.articleList = [];
      await this.getType();
      this.getRouteType();
      await this.getTag();
      this.getRouteTag();
      await this.getArticleSum();
      this.getRoutePage();
      this.getArticle();
    },
    getType() {
      return new Promise((resolve, reject) => {
        this.$axios
          .get("/article/type")
          .then((Res) => {
            if (Res.res == 1) {
              this.types = Res.types;
              this.types.unshift(0);
              resolve();
            } else reject();
          })
          .catch((Err) => {
            console.log(Err);
            reject();
          });
      });
    },
    getRouteType() {
      const type = +this.$route.params.type;
      if (this.types.indexOf(type) == -1) this.$router.push("/Classify/0/0/1");
      else {
        this.type = type;
        this.getTag();
      }
    },
    getTag() {
      return new Promise((resolve, reject) => {
        this.$axios
          .get("/article/classifyTags", { type: this.type })
          .then((Res) => {
            if (Res.res == 1) {
              this.tags = Res.tags;
              resolve();
            } else reject();
          })
          .catch((Err) => {
            console.log(Err);
            reject();
          });
      });
    },
    getRouteTag() {
      const tag = this.$route.params.tag;
      if (tag == "0") {
        this.tag = this.tags[0];
        return;
      }
      if (this.tags.indexOf(tag) == -1)
        this.$router.push({
          name: "Classify",
          params: {
            type: this.$route.params.type,
            tag: this.tags[0],
            page: 1,
          },
        });
      else {
        this.tag = tag;
      }
    },
    getArticleSum() {
      return new Promise((resolve, reject) => {
        this.$axios
          .get("/article/classifySum", {
            type: this.type,
            tag: this.tag,
          })
          .then((Res) => {
            if (Res.res == 1) {
              this.pageSum = Math.ceil(Res.sum / 6);
              resolve();
            } else reject();
          })
          .catch((Err) => {
            console.log(Err);
            resolve();
          });
      });
    },
    getRoutePage() {
      const page = +this.$route.params.page;
      if (page <= 0 || page > this.pageSum) {
        this.$router.push({
          name: "Classify",
          params: {
            type: this.$route.params.type,
            tag: this.$route.params.tag,
            page: 1,
          },
        });
      } else this.pageNow = page;
    },
    getArticle() {
      this.$axios
        .get("/article/classifyArticle", {
          type: this.type,
          tag: this.tag,
          index: this.pageNow - 1,
        })
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
    changePage(page) {
      this.pageNow = page;
      this.closeFlag = true;
      setTimeout(() => {
        this.$router.push({
          name: "Classify",
          params: {
            type: this.type,
            tag: this.tag,
            page: page,
          },
        });
      }, 100);
    },
    formatList(list) {
      formatTime(list);
      formatTpye(list);
      formatTags(list);
    },
    typeSelect(type) {
      this.closeFlag = true;
      setTimeout(() => {
        this.$router.push({
          name: "Classify",
          params: {
            type: type,
            tag: 0,
            page: 1,
          },
        });
      }, 100);
    },
    tagSelect(tag) {
      this.closeFlag = true;
      setTimeout(() => {
        this.$router.push({
          name: "Classify",
          params: {
            type: this.type,
            tag: tag,
            page: 1,
          },
        });
      }, 100);
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
.classify-container {
  animation: load-down 1s;
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
  position: relative;
  margin: 1rem;
  flex: 3;
  max-width: 50rem;
  z-index: 2;
}
aside {
  position: relative;
  margin: 1rem;
  flex: 1;
  max-width: 25rem;
  z-index: 2;
  section {
    margin-bottom: 1rem;
  }
}
.home-aside {
  position: sticky;
  top: 1rem;
}
.classify-select {
  padding: 0.5rem;
  border-radius: 0.5rem;
  box-shadow: rgba(9, 30, 66, 0.25) 0px 4px 8px -2px,
    rgba(9, 30, 66, 0.08) 0px 0px 0px 1px;
  background-image: linear-gradient(
    135deg,
    rgba(255, 255, 255, 1),
    rgba(255, 255, 255, 0.7)
  );
  & > div {
    margin: 0.5rem;
    display: flex;
    label {
      margin-right: 0.5rem;
      color: #f06292;
      white-space: nowrap;
      line-height: 2rem;
    }
    button {
      margin: 0.2rem 0.2rem;
      padding: 0.2rem 0.5rem;
      background-color: transparent;
      border: 1px solid #f06292;
      color: #f06292;
      border-radius: 0.2rem;
      transition-duration: 0.1s;
      &:hover {
        background-color: rgba($color: #f06292, $alpha: 0.15);
      }
    }
  }
}
.select-bnt {
  background-color: #f06292 !important;
  color: white !important;
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