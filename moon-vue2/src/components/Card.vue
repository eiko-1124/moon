<template>
  <div class="card" ref="card">
    <section
      v-for="(article, index) in list"
      :key="index"
      :class="{ close: closeFlag }"
      @click="goArticle(article.id)"
    >
      <div class="card-cover">
        <div
          class="cover"
          :style="{ 'background-image': 'url(' + article.cover + ')' }"
        ></div>
      </div>
      <div class="card-main">
        <div class="card-info">
          <label class="card-type" @click="goClassify(article.atype)"
            ><label>&#xe659;</label>{{ article.type }}</label
          >
          <h4>{{ article.title }}</h4>
          <p>
            {{ article.illustrate }}
          </p>
        </div>
        <div class="card-tag">
          <div class="card-tag-container">
            <label
              v-for="(tag, index) in article.tagArr"
              :key="index"
              @click.stop="goTag(tag)"
              ><label>&#xe87d;</label>{{ tag }}</label
            >
          </div>
          <div class="card-tag-time">
            <label><label>&#xe616;</label>{{ article.time }}</label>
          </div>
          <div class="card-tag-comment">
            <label><label>&#xe695;</label>{{ article.comment }}</label>
          </div>
          <img
            class="card-crow"
            src="../assets/image/crow-1.png"
            v-if="index % 2 == 1"
          />
          <img
            class="card-crow"
            src="../assets/image/crow-2.png"
            v-if="index % 2 == 0"
          />
        </div>
      </div>
    </section>
    <p
      :class="{ close: closeFlag }"
      class="card-note"
      v-if="pageSum == pageNow && flag"
    >
      没有更多啦
    </p>
    <footer
      class="page-container"
      v-if="pageSum > 1 && flag"
      :class="{ close: closeFlag }"
    >
      <span @click="changePage(1)">首页</span>
      <span @click="changePage(pageNow - 1)" :class="{ prohibit: pageNow == 1 }"
        >上一页</span
      >
      <div class="page-bnt">
        <div
          style="position: relative; transition-duration: 0.5s"
          ref="pageSlider"
        >
          <span
            :class="{ 'page-now': index == pageNow }"
            v-for="index in pageSum"
            :key="index"
            @click="changePage(index)"
            >{{ index }}</span
          >
        </div>
      </div>
      <span
        @click="changePage(pageNow + 1)"
        :class="{ prohibit: pageNow == pageSum }"
        >下一页</span
      >
    </footer>
  </div>
</template>

<script>
export default {
  name: "Card",
  props: ["pageNow", "pageSum", "articleList", "closeFlag"],
  data() {
    return {
      flag: false,
      list: [],
    };
  },
  methods: {
    changePage(page) {
      this.$emit("changePage", page);
    },
    goArticle(id) {
      this.$router.push({
        name: "Article",
        params: {
          id: id,
        },
      });
    },
    goTag(tag) {
      this.$router.push({
        name: "Classify",
        params: {
          type: 0,
          tag: tag,
          page: 1,
        },
      });
    },
    goClassify(type) {
      this.$router.push({
        name: "Classify",
        params: {
          type: type,
          tag: 0,
          page: 1,
        },
      });
    },
  },
  watch: {
    pageNow() {
      if (!this.$refs.pageSlider) return;
      if (this.pageSum <= 5) return;
      if (this.pageNow <= 2) {
        this.$refs.pageSlider.style.left = "0rem";
      }
      if (this.pageNow > 2 && this.pageSum - this.pageNow > 2) {
        this.$refs.pageSlider.style.left = `-${(this.pageNow - 3) * 3.5}rem`;
      }
      if (this.pageSum - this.pageNow <= 2) {
        this.$refs.pageSlider.style.left = `-${(this.pageSum - 5) * 3.5}rem`;
      }
    },
    articleList() {
      this.list = [];
      this.flag = false;
      if (this.articleList.length != 0) {
        let count = 0;
        const timer = setInterval(() => {
          if (count == this.articleList.length) {
            clearInterval(timer);
            this.flag = true;
          } else {
            this.list.push(this.articleList[count]);
            count++;
          }
        }, 200);
      }
    },
  },
};
</script>

<style lang="scss" scoped>
section {
  position: relative;
  margin: 1rem 0rem;
  display: flex;
  height: 16rem;
  border-radius: 0.5rem;
  box-shadow: rgba(9, 30, 66, 0.25) 0px 4px 8px -2px,
    rgba(9, 30, 66, 0.08) 0px 0px 0px 1px;
  background-image: linear-gradient(
    135deg,
    rgba(255, 255, 255, 1),
    rgba(255, 255, 255, 0.7)
  );
  z-index: 2;
  overflow: hidden;
  animation: load 0.5s;
}
.card-cover {
  margin: 1rem;
  flex: 3;
  overflow: hidden;
  border-radius: 0.2rem;
}
.cover {
  position: relative;
  top: 50%;
  left: 50%;
  height: 100%;
  width: 100%;
  background: rgba($color: #000000, $alpha: 0.1);
  transform: translateX(-50%) translateY(-50%);
  background-size: cover;
  background-position: center center;
  transition-duration: 0.5s;
}
section:hover {
  .cover {
    height: 110%;
    width: 110%;
  }
}
.card-main {
  display: grid;
  grid-template-rows: 5fr 3fr;
  padding: 1rem;
  flex: 4;
}
.card-info {
  padding-bottom: 0.5rem;
  .card-type {
    display: flex;
    align-items: center;
    color: darkcyan;
    letter-spacing: 0.1rem;
    transition-duration: 0.5s;
    label {
      height: 1.5rem;
      font-size: 1.5rem;
      line-height: 1.5rem;
    }
    &:hover {
      letter-spacing: 0.5rem;
      cursor: pointer;
    }
  }
  h4 {
    margin-top: 0.2rem;
    color: cadetblue;
    letter-spacing: 0.1rem;
    font-weight: 400;
    transition-duration: 0.5s;
    &:hover {
      cursor: pointer;
    }
  }
  p {
    margin-top: 0.8rem;
    display: -webkit-box;
    -webkit-box-orient: vertical;
    -webkit-line-clamp: 3;
    overflow: hidden;
    font-size: 0.9rem;
    color: rgba($color: gray, $alpha: 0.98);
  }
}
.card-tag {
  border-top: 1px solid darkgray;
}
.card-tag-container {
  margin-top: 0.5rem;
  display: flex;
  flex-wrap: wrap;
  & > label {
    display: flex;
    align-items: center;
    margin-right: 0.4rem;
    font-size: 0.9rem;
    color: lightseagreen;
    cursor: pointer;
    & > label {
      margin-right: 0.2rem;
      cursor: pointer;
    }
  }
}
.card-tag-time,
.card-tag-comment {
  margin-top: 0.5rem;
  display: flex;
  flex-wrap: wrap;
  label {
    display: flex;
    align-items: center;
    margin-right: 0.4rem;
    font-size: 1rem;
    color: rgba($color: gray, $alpha: 0.98);
    img {
      margin-right: 0.1rem;
      height: 1rem;
    }
  }
}
.card-crow {
  position: absolute;
  height: 8rem;
  right: 0rem;
  bottom: -2rem;
}
.page-container {
  position: relative;
  padding: 1rem 0rem;
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 2;
  .page-bnt {
    padding: 1px;
    max-width: calc(17.5rem + 2px);
    height: calc(2.25rem + 2px);
    white-space: nowrap;
    overflow: hidden;
    span {
      width: 3rem;
      height: 2.25rem;
    }
  }
  span {
    display: inline-block;
    margin-right: 0.5rem;
    padding: 0.5rem 1rem;
    background-color: white;
    border-radius: 0.2rem;
    box-shadow: rgba(9, 30, 66, 0.25) 0px 4px 8px -2px,
      rgba(9, 30, 66, 0.08) 0px 0px 0px 1px;
    cursor: pointer;
    color: cadetblue;
    text-align: center;
  }
  .page-now {
    background-color: cadetblue;
    color: white;
  }
  .prohibit {
    cursor: not-allowed;
  }
}
.card-note {
  margin: 1rem 0rem;
  padding: 1rem;
  text-align: center;
  box-shadow: rgba(9, 30, 66, 0.25) 0px 4px 8px -2px,
    rgba(9, 30, 66, 0.08) 0px 0px 0px 1px;
  background-image: linear-gradient(
    135deg,
    rgba(255, 255, 255, 1),
    rgba(255, 255, 255, 0.7)
  );
  color: gray;
  border-radius: 0.5rem;
  animation: load 0.5s;
}
.close {
  margin: 0rem !important;
  padding: 0rem !important;
  height: 0rem !important;
  overflow: hidden;
  transition-duration: 0.1s;
}
@keyframes load {
  0% {
    opacity: 0;
    transform: translateY(100%);
  }
  100% {
    opacity: 1;
    transform: translateY(0%);
  }
}

@media screen and (max-width: 759px) {
  section {
    display: block;
    height: auto;
  }
  .card-cover {
    height: 8rem;
  }
  .card-main {
    display: block;
  }
  .card-info {
    height: 6rem;
  }
  .card-type {
    p {
      white-space: nowrap;
      overflow: hidden;
      text-overflow: ellipsis;
    }
  }
  .page-bnt {
    display: none;
  }
}
</style>