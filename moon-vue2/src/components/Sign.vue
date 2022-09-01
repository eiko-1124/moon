<template>
  <div class="sign">
    <ul>
      <li v-for="(sign, index) in signs" :key="index">{{ sign.content }}</li>
      <li v-if="signs.length != 0">{{ signs[0].content }}</li>
    </ul>
  </div>
</template>

<script>
export default {
  name: "Sign",
  data() {
    return {
      signs: [],
    };
  },
  mounted() {
    this.getSigns();
  },
  methods: {
    getSigns() {
      this.$axios
        .get("/user/signs")
        .then((Res) => {
          if (Res.res == 1) {
            this.signs = Res.signs;
          }
        })
        .catch((Err) => {
          console.log(Err);
        });
    },
  },
};
</script>

<style lang="scss" scoped>
.sign {
  position: relative;
  height: 2.5rem;
  box-shadow: rgba(9, 30, 66, 0.25) 0px 4px 8px -2px,
    rgba(9, 30, 66, 0.08) 0px 0px 0px 1px;
  background-color: rgba(255, 255, 255, 0.9);
  border-radius: 0.3rem;
  z-index: 2;
  overflow: hidden;
}
ul {
  position: relative;
  list-style: none;
  animation: sRotate 40s infinite linear;
  li {
    height: 2.5rem;
    padding: 0.5rem 1rem;
    line-height: 1.5rem;
    text-align: center;
    text-overflow: clip;
    color: lightslategray;
  }
}

@keyframes sRotate {
  0% {
    top: 0rem;
  }
  24% {
    top: 0rem;
  }
  25% {
    top: -2.5rem;
  }
  49% {
    top: -2.5rem;
  }
  50% {
    top: -5rem;
  }
  74% {
    top: -5rem;
  }
  75% {
    top: -7.5rem;
  }
  99% {
    top: -7.5rem;
  }
  100% {
    top: -10rem;
  }
}
</style>