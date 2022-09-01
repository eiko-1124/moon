<template>
  <div class="plan">
    <h3>计划</h3>
    <p
      style="padding: 1rem; text-align: center; opacity: 0.8"
      v-if="notes.length == 0"
    >
      暂时没有计划
    </p>
    <main>
      <div v-for="(note, index) in notes" :key="index">
        <span></span>
        <div>
          <p>{{ note.time }}</p>
          <p>{{ note.text }}</p>
        </div>
      </div>
    </main>
  </div>
</template>

<script>
import { formatDate } from "@/utils/format";
export default {
  name: "Plan",
  data() {
    return {
      notes: [],
    };
  },
  mounted() {
    this.getPlan();
  },
  methods: {
    getPlan() {
      this.$axios
        .get("/note/notes")
        .then((Res) => {
          if (Res.res == 1) {
            formatDate(Res.notes);
            this.notes = Res.notes;
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
.plan {
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
main {
  padding: 0rem 1rem 1rem;
  & > div {
    display: grid;
    grid-template-columns: 3rem auto;
    & > span {
      position: relative;
      height: 100%;
      width: 3rem;
      &::after {
        content: "";
        position: absolute;
        display: block;
        top: 50%;
        left: 50%;
        width: 0.8rem;
        height: 0.8rem;
        background-color: darkcyan;
        border-radius: 0.4rem;
        transform: translate(-50%, -50%);
      }
      &::before {
        content: "";
        margin: 0rem auto;
        display: block;
        height: 100%;
        width: 0.1rem;
        background-color: darkcyan;
      }
    }
  }
}

@media screen and (max-width: 759px) {
  .plan {
    display: none;
  }
}
</style>