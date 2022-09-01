<template>
  <div
    :class="{ 'top-nav': true, 'top-switch': topSwitch, 'top-bg': topBg }"
    ref="topNav"
  >
    <header>
      <h2 v-for="(font, index) in title" :key="index">{{ font }}</h2>
    </header>
    <nav>
      <button>搜索</button>
      <button>
        <router-link to="/">首页</router-link>
      </button>
      <button><router-link to="/Classify/0/0/1">分类</router-link></button>
      <button><router-link to="/PickOn">归档</router-link></button>
      <button><router-link to="/Record">记录</router-link></button>
      <button><router-link to="/Link">友联</router-link></button>
      <button><router-link to="/About">关于</router-link></button>
      <button><router-link to="/Message">留言</router-link></button>
    </nav>
  </div>
</template>

<script>
export default {
  name: "TopNav",
  data() {
    return {
      topSwitch: false,
      topBg: false,
      scrollTop: 0,
      title: "Mn+O2",
    };
  },
  mounted() {
    this.addScrollEvent();
  },
  methods: {
    addScrollEvent() {
      const view = document.querySelector(".view");
      view.addEventListener("scroll", this.scrollEvent);
    },
    scrollEvent() {
      const view = document.querySelector(".view");
      const top = view.scrollTop;
      if (top > this.scrollTop && this.scrollTop > 100) this.topSwitch = true;
      else this.topSwitch = false;
      if (top > 10) this.topBg = true;
      else this.topBg = false;
      this.scrollTop = top;
    },
  },
  beforeDestroy() {
    window.removeEventListener("scroll", this.scrollEvent);
  },
  computed: {
    load() {
      return this.$store.state.loadFlag;
    },
  },
  watch: {
    load() {
      if (!this.load) {
        this.$nextTick(() => {
          setTimeout(() => {
            this.addScrollEvent();
          }, 200);
        });
      }
    },
  },
};
</script>

<style lang="scss" scoped>
.top-nav {
  padding: 0.5rem;
  position: fixed;
  top: 0rem;
  width: 100%;
  height: 3.5rem;
  transition-duration: 0.3s;
  z-index: 9999;
}
.top-switch {
  top: -4rem;
}
.top-bg {
  background-color: rgba($color: white, $alpha: 0.9);
  box-shadow: 0px 0px 1px 0px gray;
  nav {
    button {
      color: black;
      &:hover {
        background-color: rgba(211, 211, 211, 0.3);
      }
    }
  }
}
header {
  display: flex;
  margin: 0.25rem 2rem;
  width: 8rem;
  cursor: default;
  h2 {
    margin: 0rem 0.05rem;
    color: lightblue;
  }
}
nav {
  margin-top: -2.5rem;
  margin-right: 2rem;
  height: 3rem;
  float: right;
  button {
    margin: 0.75rem 0.4rem;
    height: 1.5rem;
    width: 4rem;
    border: none;
    border-radius: 0.1rem;
    background-color: transparent;
    color: whitesmoke;
    font-weight: 500;
    line-height: 1.5rem;
    transition-duration: 0.3s;
    &:hover {
      background-color: rgba($color: white, $alpha: 0.3);
    }
    a {
      color: inherit;
      text-decoration: none;
    }
  }
}
@media screen and (max-width: 759px) {
  .top-nav {
    display: none;
  }
}
</style>