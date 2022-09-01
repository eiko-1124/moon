<template>
  <div class="load">
    <div class="load-box">
      <span class="icon-1">
        <span>
          <img v-show="moodFlag" src="../assets/icon/icon-happy.png" />
          <img v-show="!moodFlag" src="../assets/icon/icon-smile.png" />
        </span>
      </span>
      <span class="icon-2">
        <span>
          <img v-show="moodFlag" src="../assets/icon/icon-sad.png" />
          <img v-show="!moodFlag" src="../assets/icon/icon-cry.png" />
        </span>
      </span>
    </div>
  </div>
</template>

<script>
export default {
  name: "Load",
  data() {
    return {
      moodFlag: true,
      timer: null,
    };
  },
  mounted() {
    setTimeout(() => {
      this.moodFlag = !this.moodFlag;
      this.timer = setInterval(() => {
        this.moodFlag = !this.moodFlag;
      }, 2000);
    }, 700);
  },
  beforeDestroy() {
    clearInterval(this.timer);
  },
};
</script>

<style lang="scss" scoped>
.load {
  position: fixed;
  top: 0rem;
  left: 0rem;
  width: 100vw;
  height: 100vh;
  background-color: rgba($color: white, $alpha: 0.4);
  z-index: 10000;
}
.load-box {
  position: relative;
  top: 50%;
  left: 50%;
  width: 18rem;
  height: 3rem;
  transform: translate(-50%, -50%);
}
span {
  position: absolute;
  display: inline-block;
  width: 5rem;
  height: 3rem;
  border-radius: 0.5rem;
  box-shadow: 0px 0px 1px 1px white;
  img {
    margin: 0.3rem 1.3rem;
    height: 80%;
  }
}
.icon-1 {
  left: 0rem;
  background-color: rgba(185, 205, 246, 0.6);
  animation: icon1-back 2s linear infinite;
  span {
    left: 1rem;
    background-color: rgba(185, 205, 246, 1);
    animation: icon1-pre 2s linear infinite;
  }
}
.icon-2 {
  right: 0rem;
  background-color: rgba(185, 205, 246, 0.6);
  animation: icon2-back 2s linear infinite;
  span {
    right: 1rem;
    background-color: rgba(185, 205, 246, 1);
    animation: icon2-pre 2s linear infinite;
  }
}

@keyframes icon1-back {
  0% {
    left: 0rem;
  }
  35% {
    left: 3rem;
  }
  50% {
    left: 3rem;
  }
  85% {
    left: 0rem;
  }
  100% {
    left: 0rem;
  }
}
@keyframes icon1-pre {
  0% {
    left: 1rem;
  }
  35% {
    left: 1rem;
  }
  50% {
    left: -1rem;
  }
  85% {
    left: -1rem;
  }
  100% {
    left: 1rem;
  }
}
@keyframes icon2-back {
  0% {
    right: 0rem;
  }
  35% {
    right: 3rem;
  }
  50% {
    right: 3rem;
  }
  85% {
    right: 0rem;
  }
  100% {
    right: 0rem;
  }
}
@keyframes icon2-pre {
  0% {
    right: 1rem;
  }
  35% {
    right: 1rem;
  }
  50% {
    right: -1rem;
  }
  85% {
    right: -1rem;
  }
  100% {
    right: 1rem;
  }
}
</style>