<template>
  <div class="tag">
    <h3>标签</h3>
    <div class="tag-container">
      <label v-for="(tag, index) in tags" :key="index" @click="goTag(tag)">{{
        tag
      }}</label>
    </div>
    <footer>
      <span v-for="n in 5" :key="n">
        <img src="../assets/icon/leaf.png" />
      </span>
    </footer>
  </div>
</template>

<script>
export default {
  name: "Tag",
  data() {
    return {
      tags: [],
    };
  },
  mounted() {
    this.getTags();
  },
  methods: {
    getTags() {
      this.$axios
        .get("/article/tags")
        .then((Res) => {
          if (Res.res == 1) {
            this.tags = Res.tags;
          }
        })
        .catch((Err) => {
          console.log(Err);
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
  },
};
</script>

<style lang="scss" scoped>
.tag {
  background-image: linear-gradient(
    to bottom,
    rgba(255, 255, 255, 1),
    rgba(255, 255, 255, 0.7)
  );
  border-radius: 0.5rem;
  box-shadow: rgba(9, 30, 66, 0.25) 0px 4px 8px -2px,
    rgba(9, 30, 66, 0.08) 0px 0px 0px 1px;
}
h3 {
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
.tag-container {
  padding: 0rem 1rem 0.2rem;
  display: flex;
  justify-content: center;
  align-items: center;
  flex-wrap: wrap;
  label {
    margin: 0.2rem;
    padding: 0.2rem 0.5rem;
    border: 1px solid darkcyan;
    border-radius: 2rem;
    white-space: nowrap;
    text-overflow: ellipsis;
    color: darkcyan;
    opacity: 0.7;
    cursor: pointer;
    &:hover {
      background-color: rgba($color: darkcyan, $alpha: 1);
      color: white;
      transition-duration: 0.5s;
    }
  }
}
footer {
  position: relative;
  height: 3.5rem;
  span {
    position: absolute;
    display: block;
    height: 1.2rem;
    width: 1.2rem;
    img {
      height: 100%;
    }
  }
  & > span:nth-child(1) {
    right: 1.65rem;
    animation: shine 4s infinite;
  }
  & > span:nth-child(2) {
    top: 0.65rem;
    right: 2.6rem;
    transform: rotateZ(-75deg);
    animation: shine 4s infinite 0.8s;
  }
  & > span:nth-child(3) {
    top: 0.65rem;
    right: 0.7rem;
    transform: rotateZ(75deg);
    animation: shine 4s infinite 0.2s;
  }
  & > span:nth-child(4) {
    top: 1.7rem;
    right: 2.2rem;
    transform: rotateZ(-150deg);
    animation: shine 4s infinite 0.6s;
  }
  & > span:nth-child(5) {
    top: 1.7rem;
    right: 1.1rem;
    transform: rotateZ(150deg);
    animation: shine 4s infinite 0.4s;
  }
}

@keyframes shine {
  0% {
    opacity: 0.4;
  }
  50% {
    opacity: 1;
  }
  100% {
    opacity: 0.4;
  }
}
@media screen and (max-width: 759px) {
  .tag {
    display: none;
  }
}
</style>