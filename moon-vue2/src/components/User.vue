<template>
  <div class="user">
    <section class="user-info">
      <h3>{{ info.name }}</h3>
      <div class="user-headPortrait">
        <img src="http://localhost:3000/static/portrait/head-portrait.png" />
      </div>
      <div class="user-statistics">
        <div>
          <label>文章</label>
          <p>{{ articleSum }}</p>
        </div>
        <label>|</label>
        <div>
          <label>分类</label>
          <p>{{ typeSum }}</p>
        </div>
        <label>|</label>
        <div>
          <label>标签</label>
          <p>{{ tagSum }}</p>
        </div>
      </div>
      <div class="user-option">
        <button>
          <img src="../assets/icon/github.png" />
        </button>
        <button>
          <img src="../assets/icon/qq.png" />
        </button>
        <button>加入书签</button>
      </div>
    </section>
    <section class="user-notice">
      <h3>公告</h3>
      <p>{{ info.notice }}</p>
    </section>
  </div>
</template>

<script>
export default {
  name: "User",
  props: ["info"],
  data() {
    return {
      articleSum: 0,
      tagSum: 0,
      typeSum: 0,
    };
  },
  mounted() {
    this.getSum();
  },
  methods: {
    getSum() {
      this.$axios
        .get("/article/articleSum")
        .then((Res) => {
          if (Res.res == 1) {
            this.articleSum = Res.sum;
          }
        })
        .catch((Err) => {
          console.log(Err);
        });
      this.$axios
        .get("/article/typeSum")
        .then((Res) => {
          if (Res.res == 1) {
            this.typeSum = Res.sum;
          }
        })
        .catch((Err) => {
          console.log(Err);
        });
      this.$axios
        .get("/article/tagSum")
        .then((Res) => {
          if (Res.res == 1) {
            this.tagSum = Res.sum;
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
.user-info {
  margin-bottom: 1rem;
  border-radius: 0.5rem;
  box-shadow: rgba(9, 30, 66, 0.25) 0px 4px 8px -2px,
    rgba(9, 30, 66, 0.08) 0px 0px 0px 1px;
  background-image: linear-gradient(
    to bottom,
    rgba(255, 255, 255, 1),
    rgba(255, 255, 255, 0.8)
  );
}
.user-notice {
  padding-bottom: 1rem;
  margin-bottom: 1rem;
  border-radius: 0.5rem;
  box-shadow: rgba(9, 30, 66, 0.25) 0px 4px 8px -2px,
    rgba(9, 30, 66, 0.08) 0px 0px 0px 1px;
  background-image: linear-gradient(
    to bottom,
    rgba(255, 255, 255, 1),
    rgba(255, 255, 255, 0.8)
  );
  p {
    padding: 0rem 1rem 1rem;
    color: cadetblue;
    text-align: center;
    word-break: break-all;
  }
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
.user-headPortrait {
  margin: 0rem auto;
  width: 8rem;
  height: 8rem;
  border-radius: 0.5rem;
  overflow: hidden;
  img {
    height: 100%;
  }
}
.user-statistics {
  display: grid;
  padding: 1rem 2rem;
  grid-template-columns: repeat(5, 1fr);
  color: darkcyan;
  label {
    display: inline-block;
    width: 100%;
    text-align: center;
  }
  p {
    padding: 0.5rem;
    text-align: center;
  }
}
.user-option {
  padding: 0rem 1rem 1rem;
  display: flex;
  align-items: center;
  button {
    height: 2.5rem;
    background-color: cadetblue;
    border: none;
    border-radius: 0.2rem;
    cursor: pointer;
    img {
      margin: 5%;
      height: 90%;
    }
  }
  & > button:nth-child(1),
  & > button:nth-child(2) {
    margin-right: 0.5rem;
    width: 2.5rem;
  }
  & > button:nth-child(3) {
    flex: 1;
    color: white;
  }
}

@media screen and (max-width: 759px) {
  .user {
    display: none;
  }
}
</style>