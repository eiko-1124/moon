<template>
  <div class="recent">
    <h3>近期推荐</h3>
    <ul>
      <li
        v-for="(recommend, index) in list"
        :key="index"
        @click="goArticle(recommend.id)"
      >
        <section class="recent-cover">
          <div
            class="recent-cover-image"
            :style="{ 'background-image': 'url(' + recommend.cover + ')' }"
          ></div>
        </section>
        <section class="recent-text">
          <h4>{{ recommend.title }}</h4>
          <label><label>&#xe60d;</label>{{ recommend.time }}</label>
        </section>
      </li>
    </ul>
  </div>
</template>

<script>
import { formatDate } from "@/utils/format";
export default {
  name: "Recent",
  data() {
    return {
      list: [],
    };
  },
  mounted() {
    this.getRecomment();
  },
  methods: {
    getRecomment() {
      this.$axios
        .get("/article/recommend")
        .then((Res) => {
          if (Res.res == 1) {
            formatDate(Res.list);
            this.list = Res.list;
          }
        })
        .catch((Err) => {
          console.log(Err);
        });
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
.recent {
  background-image: linear-gradient(
    to bottom,
    rgba(255, 255, 255, 1),
    rgba(255, 255, 255, 0.6)
  );
  box-shadow: rgba(9, 30, 66, 0.25) 0px 4px 8px -2px,
    rgba(9, 30, 66, 0.08) 0px 0px 0px 1px;
  border-radius: 0.5rem;
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
ul {
  padding: 0rem 1rem 0.5rem;
  list-style: none;
  li {
    margin-bottom: 1rem;
    display: flex;
    height: 3rem;
    &:hover {
      .recent-cover-image {
        height: 120%;
        width: 120%;
      }
      h4 {
        color: darkcyan;
      }
    }
  }
}
.recent-cover {
  width: 5rem;
  overflow: hidden;
  border-radius: 0.2rem;
}
.recent-cover-image {
  position: relative;
  top: 50%;
  left: 50%;
  width: 100%;
  height: 100%;
  background-size: cover;
  background-position: center center;
  transform: translate(-50%, -50%);
  transition-duration: 0.5s;
}
.recent-text {
  padding-left: 1rem;
  h4 {
    padding-bottom: 0.4rem;
    height: 1.8rem;
    font-weight: 400;
    line-height: 1.4rem;
    color: #4c6280;
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
  }
  & > label {
    display: flex;
    align-items: center;
    color: #4c6280;
    & > label {
      margin-right: 0.2rem;
    }
  }
}

@media screen and (max-width: 759px) {
  .recent {
    display: none;
  }
}
</style>