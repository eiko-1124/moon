<template>
  <div>
    <header :style="{ 'background-image': 'url(' + background + ')' }"></header>
    <div class="pick-container">
      <h3>归档</h3>
      <main>
        <article>
          <div class="pick-box">
            <h3>Pick On</h3>
            <div class="pick-site-head">
              <p>目前一中{{ sum }}篇文章，继续加油</p>
            </div>
            <div
              v-for="(title, index) in list"
              :key="index"
              class="pick-site"
              ref="title"
              @click="goArticle(id)"
            >
              <p>{{ title.time }}</p>
              <p>{{ title.title }}</p>
            </div>
            <div class="pick-load" v-if="loadFlag"><label>&#xe629;</label></div>
          </div>
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
import { formatDate } from "@/utils/format";
export default {
  name: "PickOn",
  components: {
    Recent,
    Tag,
  },
  data() {
    return {
      sum: 0,
      list: [],
      loadFlag: false,
    };
  },
  computed: {
    background() {
      return this.$route.meta.background;
    },
  },
  mounted() {
    this.getSum();
    this.getTitles();
  },
  methods: {
    getSum() {
      this.$axios.get("/article/articleSum").then((Res) => {
        if (Res.res == 1) {
          this.sum = Res.sum;
        }
      });
    },
    getTitles() {
      return new Promise((resolve, reject) => {
        this.loadFlag = true;
        this.$axios
          .get("/article/title", { index: this.list.length })
          .then((Res) => {
            this.loadFlag = false;
            if (Res.res == 1) {
              formatDate(Res.titles);
              this.addObserve(Res.titles);
              resolve();
            }
          })
          .catch((Err) => {
            console.log(Err);
            reject();
          });
      });
    },
    addObserve(list) {
      const timer = setInterval(() => {
        if (list.length == 0) {
          clearInterval(timer);
          this.$nextTick(() => {
            const target = this.$refs.title[this.$refs.title.length - 1];
            let options = {
              root: null,
              rootMargin: "0px",
              threshold: 1,
            };
            let observer = new IntersectionObserver((entries) => {
              entries.forEach((entry) => {
                if (entry.isIntersecting) {
                  observer.unobserve(target);
                  if (this.sum > this.list.length) this.getTitles();
                }
              });
            }, options);
            observer.observe(target);
          });
        } else {
          this.list.push(list.shift());
        }
      }, 200);
    },
    goArticle(id) {
      this.$router.push({
        name: "Article",
        params: {
          id: id,
        },
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
.pick-container,
.pick-box {
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
.pick-container {
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
.pick-box {
  position: relative;
  padding: 0.5rem;
  border-radius: 0.5rem;
  box-shadow: rgba(9, 30, 66, 0.25) 0px 4px 8px -2px,
    rgba(9, 30, 66, 0.08) 0px 0px 0px 1px;
  background-image: linear-gradient(
    135deg,
    rgba(255, 255, 255, 1),
    rgba(255, 255, 255, 0.7)
  );
  z-index: 2;
  & > h3 {
    justify-content: left;
  }
}
.pick-site-head {
  position: relative;
  margin: 0rem 1rem;
  padding: 0rem 1rem;
  padding-left: 5rem;
  &::before {
    content: "";
    display: block;
    position: absolute;
    top: 0rem;
    left: 2.45rem;
    height: 100%;
    width: 0.1rem;
    background-color: darkcyan;
    border-top-left-radius: 0.1rem;
    border-top-right-radius: 0.1rem;
  }
  &::after {
    content: "";
    display: block;
    position: absolute;
    top: 50%;
    left: 2rem;
    height: 1rem;
    width: 1rem;
    background-color: darkcyan;
    border-radius: 0.5rem;
    transform: translateY(-50%);
  }
}
.pick-site {
  position: relative;
  margin: 0rem 1rem;
  padding: 0.5rem 1rem;
  padding-left: 5rem;
  animation: site-load 0.3s;
  p {
    font-size: 0.9rem;
    opacity: 0.9;
  }
  &::before {
    content: "";
    display: block;
    position: absolute;
    top: 0rem;
    left: 2.45rem;
    height: 100%;
    width: 0.1rem;
    background-color: darkcyan;
    border-top-left-radius: 0.1rem;
    border-top-right-radius: 0.1rem;
  }
  &::after {
    content: "";
    display: block;
    position: absolute;
    top: 50%;
    left: 2.1rem;
    height: 0.8rem;
    width: 0.8rem;
    background-color: white;
    box-shadow: 0px 0px 1px 1px darkcyan;
    border-radius: 0.4rem;
    transform: translateY(-50%);
  }
  &:hover {
    & > p:nth-child(2) {
      color: darkcyan;
      cursor: pointer;
    }
  }
}
.pick-site:nth-last-child(1) {
  margin-bottom: 2.5rem;
  &::before {
    content: "";
    display: block;
    position: absolute;
    top: 0rem;
    left: 2.45rem;
    height: 80%;
    width: 0.1rem;
    background-color: darkcyan;
    border-top-left-radius: 0.1rem;
    border-top-right-radius: 0.1rem;
  }
}
.pick-load {
  padding: 1rem;
  text-align: center;
  color: darkcyan;
  overflow: hidden;
  label {
    display: block;
    animation: load 1.5s infinite;
    font-size: 1.2rem;
  }
}

@keyframes load {
  from {
    transform: rotateZ(0deg);
  }
  to {
    transform: rotateZ(360deg);
  }
}

@keyframes site-load {
  0% {
    max-height: 0rem;
    opacity: 0;
    padding: 0rem 1rem 0rem 5rem;
  }
  100% {
    max-height: 10rem;
    opacity: 1;
    padding: 0.5rem 1rem;
    padding-left: 5rem;
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