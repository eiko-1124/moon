<template>
  <div class="amessage">
    <h3>最新留言</h3>
    <ul>
      <li v-for="(site, index) in list" :key="index">
        <section>
          <p>
            <img :src="icon[site.image]" />
          </p>
        </section>
        <section>
          <label
            >{{ site.name }}<label>{{ site.time }}</label></label
          >
          <p>{{ site.text }}</p>
        </section>
      </li>
    </ul>
  </div>
</template>

<script>
import { formatDate } from "@/utils/format";
export default {
  name: "Amessage",
  data() {
    return {
      list: [],
      icon: [
        require("../assets/icon/dog.png"),
        require("../assets/icon/octopus.png"),
        require("../assets/icon/kola.png"),
        require("../assets/icon/cat.png"),
        require("../assets/icon/brid.png"),
      ],
    };
  },
  mounted() {
    this.getMessages();
  },
  methods: {
    getMessages() {
      this.$axios
        .get("/note/messages", { index: 0, num: 6 })
        .then((Res) => {
          if (Res.res == 1) {
            formatDate(Res.messages);
            this.formatIcon(Res.messages);
            this.list = Res.messages;
          }
        })
        .catch((Err) => {
          console.log(Err);
        });
    },
    formatIcon(list) {
      list.forEach((site) => {
        site.image = Math.floor(Math.random() * 5);
      });
    },
  },
};
</script>

<style lang="scss" scoped>
.amessage {
  margin-top: 1rem;
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
  padding-bottom: 1rem;
  list-style: none;
  li {
    padding: 0rem 1rem 1rem;
    display: flex;
    align-items: stretch;
    & > section:nth-child(1) {
      position: relative;
      margin-right: 0.5rem;
      height: 3rem;
      width: 3rem;
      font-size: 0.8rem;
      background-color: rgba(135, 206, 235, 0.15);
      border-radius: 0.2rem;
      img {
        position: absolute;
        top: 50%;
        left: 50%;
        max-height: 2rem;
        max-width: 2rem;
        transform: translate(-50%, -50%);
      }
    }
    & > section:nth-child(2) {
      flex: 1;
      padding: 0.4rem;
      min-height: 3rem;
      background-color: rgba(135, 206, 235, 0.15);
      border-radius: 0.2rem;
      & > label {
        font-size: 0.9rem;
        color: darkcyan;
        label {
          margin-left: 0.5rem;
        }
      }
    }
  }
}
@media screen and (max-width: 759px) {
  .amessage {
    display: none;
  }
}
</style>