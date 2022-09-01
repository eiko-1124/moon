<template>
  <div class="patel">
    <span
      v-for="patel in patels"
      :key="patel.id"
      :style="{
        top: patel.top + 'px',
        right: patel.right + 'px',
        height: patel.height + 'px',
      }"
    >
      <img src="../assets/image/patel.png" />
    </span>
  </div>
</template>

<script>
export default {
  name: "Patel",
  data() {
    return {
      patels: [],
    };
  },
  mounted() {
    setInterval(this.fallDown, 400);
    setTimeout(() => {
      setInterval(this.recoveryPatel, 400);
    }, 10000);
  },
  methods: {
    createPatel() {
      const flag = Math.random() > 0.3;
      let top, right;
      if (flag) {
        const width = document.body.clientWidth;
        top = -100;
        right = Math.floor(Math.random() * (width + 100)) - 100;
      } else {
        const height = document.body.clientHeight;
        right = -100;
        top = Math.floor(Math.random() * (height + 100)) - 100;
      }
      const height = Math.floor(Math.random() * 50) + 40;
      return {
        id: Symbol(),
        top,
        right,
        height,
      };
    },
    fallDown() {
      const patel = this.createPatel();
      this.patels.push(patel);
      this.$nextTick(() => {
        const height = document.body.clientHeight;
        patel.right = patel.right - (patel.top - height) * 1.25;
        patel.top = height;
      });
    },
    recoveryPatel() {
      this.patels.shift();
    },
  },
};
</script>

<style lang="scss" scoped>
.patel {
  top: 0rem;
  width: 0rem;
  height: 0rem;
  span {
    position: absolute;
    display: block;
    transition-duration: 10s;
    transition-timing-function: linear;
    z-index: 1;
    animation: Rotate 5s infinite linear;
    img {
      height: 100%;
      animation: flip 16s infinite linear;
    }
  }
}

@keyframes Rotate {
  0% {
    transform: rotateZ(0deg);
  }
  100% {
    transform: rotateZ(360deg);
  }
}
@keyframes flip {
  0% {
    transform: rotateY(0deg);
  }
  100% {
    transform: rotateY(360deg);
  }
}
</style>