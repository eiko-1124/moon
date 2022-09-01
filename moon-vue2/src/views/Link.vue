<template>
  <div>
    <header :style="{ 'background-image': 'url(' + background + ')' }"></header>
    <div class="link-container">
      <h3>友联</h3>
      <main>
        <article>
          <div class="link-box">
            <h3>Link</h3>
            <div
              class="link-site"
              v-for="(flink, index) in flinks"
              :key="index"
            >
              <label>{{ flink.name + " : " }}</label>
              <a :href="flink.flink" target="_blank">{{ flink.flink }}</a>
            </div>
          </div>
        </article>
        <aside>
          <div class="home-aside">
            <section>
              <tag></tag>
            </section>
            <section><recent></recent></section>
          </div>
        </aside>
      </main>
    </div>
  </div>
</template>

<script>
import Tag from "@/components/Tag.vue";
import Recent from "@/components/Recent.vue";
export default {
  name: "Link",
  components: {
    Tag,
    Recent,
  },
  data() {
    return {
      flinks: [],
    };
  },
  computed: {
    background() {
      return this.$route.meta.background;
    },
  },
  mounted() {
    this.getLinks();
  },
  methods: {
    getLinks() {
      this.$axios
        .get("/user/flinks")
        .then((Res) => {
          if (Res.res == 1) this.flinks = Res.flinks;
        })
        .catch((Err) => {
          console.log(Err);
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
.link-container,
.link-box {
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
.link-container {
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
.link-box {
  position: relative;
  padding: 0.5rem;
  margin-bottom: 1rem;
  height: calc(100% - 1rem);
  border-radius: 0.5rem;
  box-shadow: rgba(9, 30, 66, 0.25) 0px 4px 8px -2px,
    rgba(9, 30, 66, 0.08) 0px 0px 0px 1px;
  background-image: linear-gradient(
    135deg,
    rgba(255, 255, 255, 1),
    rgba(255, 255, 255, 0.7)
  );
  z-index: 2;
  h3 {
    justify-content: left;
  }
}
.link-site {
  margin: 0rem 1rem 1rem;
  label {
    letter-spacing: 0.1rem;
    color: darkslategray;
  }
  a {
    color: gray;
    text-decoration: none;
    &:hover {
      color: teal;
      text-decoration: underline teal;
    }
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