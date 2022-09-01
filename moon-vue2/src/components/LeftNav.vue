<template>
  <div
    class="left-nav"
    :class="{ 'open-nav': openFlag, 'top-switch': topSwitch, 'top-bg': topBg }"
    @click="changNav"
  >
    <div class="nav-body" :class="{ open: openFlag }">
      <div style="position: relative; width: 18rem">
        <div class="nav-head-box"></div>
        <div class="nav-head-portrait"></div>
        <h3 class="nav-name">Mn+O2</h3>
        <div class="nav-button">
          <button>
            <router-link to="/">首页</router-link>
          </button>
          <button><router-link to="/Classify/0/0/1">分类</router-link></button>
          <button><router-link to="/PickOn">归档</router-link></button>
          <button><router-link to="/Record">记录</router-link></button>
          <button><router-link to="/Link">友联</router-link></button>
          <button><router-link to="/About">关于</router-link></button>
          <button><router-link to="/Message">留言</router-link></button>
        </div>
      </div>
    </div>
    <div class="nav-switch" :class="{ close: openFlag }">
      <img src="../assets/icon/icon-menu.png" />
      <button>菜单</button>
    </div>
  </div>
</template>

<script>
export default {
  name: "LeftNav",
  data() {
    return {
      openFlag: false,
      topSwitch: false,
      topBg: false,
      scrollTop: 0,
    };
  },
  mounted() {
    this.addScrollEvent();
  },
  methods: {
    changNav() {
      this.openFlag = !this.openFlag;
    },
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
          this.addScrollEvent();
        });
      }
    },
  },
};
</script>

<style lang="scss" scoped>
.left-nav {
  display: none;
  position: absolute;
  height: 3rem;
  width: 100%;
  z-index: 9999;
  transition-duration: 0.5s;
}
.nav-switch {
  display: flex;
  align-items: center;
  cursor: pointer;
  img {
    margin: 0.5rem;
    height: 1.5rem;
  }
  button {
    background-color: transparent;
    border: none;
  }
  transition-duration: 0.5s;
}
.nav-body {
  display: inline-flex;
  flex-direction: column;
  height: 100vh;
  width: 0rem;
  background-image: linear-gradient(180deg, white, rgba(208, 228, 236, 0.95));
  animation: open 0.5s linear;
  transition-duration: 0.5s;
  overflow: hidden;
  overflow: auto;
  &::-webkit-scrollbar {
    display: none;
  }
}
.nav-head-box {
  height: 12rem;
  width: 100%;
  background-image: url(http://localhost:3000/static/portrait/head-background.png);
  background-position: center center;
  background-size: cover;
}
.nav-head-portrait {
  position: absolute;
  top: 9rem;
  right: 1rem;
  width: 6rem;
  height: 6rem;
  border-radius: 3rem;
  background-image: url(http://localhost:3000/static/portrait/head-portrait.png);
  background-position: center center;
  background-size: cover;
  box-shadow: rgba(9, 30, 66, 0.25) 0px 4px 8px -2px,
    rgba(9, 30, 66, 0.08) 0px 0px 0px 1px;
}
.nav-name {
  margin: 1rem 4rem;
  color: darkcyan;
}
.nav-button {
  button {
    margin: 0.2rem 10%;
    padding: 0.5rem;
    width: 80%;
    background-color: transparent;
    border: none;
    text-indent: 1rem;
    letter-spacing: 1rem;
    border-radius: 0.2rem;
    transition-duration: 0.3s;
    a {
      color: black;
      text-decoration: none;
    }
    &:hover {
      background-color: rgba($color: gray, $alpha: 0.6);
    }
  }
}
.open {
  width: 18rem;
}
.open-nav {
  width: 100vw;
  height: 100vh;
}
.close {
  opacity: 0;
}
.top-switch {
  top: -4rem;
}
.top-bg {
  background-color: rgba($color: white, $alpha: 0.9);
  box-shadow: 0px 0px 1px 0px gray;
}
@media screen and (max-width: 759px) {
  .left-nav {
    display: flex;
    align-items: flex-start;
  }
}
</style>