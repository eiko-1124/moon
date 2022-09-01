<template>
  <div class="cover" :style="{ 'background-image': 'url(' + background + ')' }">
    <main>
      <aside ref="aside">
        <ul>
          <li v-for="(list, index) in notes" :key="index">
            <div
              class="note"
              v-for="(note, index) in list"
              :key="index"
              :style="{ opacity: note }"
            ></div>
          </li>
        </ul>
      </aside>
      <article>
        <h4>
          {{ sentence }}
        </h4>
      </article>
    </main>
    <footer ref="footer">
      <canvas ref="myCanvas"></canvas>
    </footer>
  </div>
</template>

<script>
export default {
  name: "Cover",
  props: ["sentence"],
  data() {
    return {
      coordinates: [],
      footerTimer: null,
      notes: [],
      noteTimer: null,
    };
  },
  mounted() {
    this.init();
  },
  computed: {
    background() {
      return this.$route.meta.background;
    },
  },
  methods: {
    init() {
      this.randCoordinate();
      this.footerTimer = setTimeout(this.paint, 50);
      this.createNotes();
      this.noteTimer = setTimeout(this.resetNotes, 30);
    },
    paint() {
      const height = this.$refs.footer.offsetHeight;
      const width = this.$refs.footer.offsetWidth;
      const canvas = this.$refs.myCanvas;
      canvas.width = width;
      canvas.height = height;
      if (canvas.getContext) {
        this.resetCoordinate();
        let context = canvas.getContext("2d");
        const linearGradient = context.createLinearGradient(0, 0, 0, height);
        linearGradient.addColorStop(0, "rgba(255,255,255,0.6)");
        linearGradient.addColorStop(1, "rgba(255,255,255,1)");
        context.fillStyle = linearGradient;
        context.beginPath();
        context.moveTo(0, height * (1 / 4 + this.coordinates[0].yAxis / 400));
        context.bezierCurveTo(
          width * (1 / 10 + this.coordinates[1].xAxis / 200),
          height * (this.coordinates[1].yAxis / 200),
          width * (1 / 2 + this.coordinates[2].xAxis / 200),
          height * (this.coordinates[2].yAxis / 200),
          width,
          height * (1 / 4 + this.coordinates[3].yAxis / 400)
        );
        context.lineTo(width, height);
        context.lineTo(0, height);
        context.lineTo(0, height * (1 / 4 + this.coordinates[0].yAxis / 400));
        context.fill();
      }
    },
    randCoordinate() {
      for (let i = 0; i < 4; i++) {
        const spot = {
          xAxis: Math.floor(60 * Math.random()),
          yAxis: Math.floor(100 * Math.random()),
          xTurn: Math.random() * 2 > 1,
          yTurn: Math.random() * 2 > 1,
        };
        this.coordinates.push(spot);
      }
    },
    resetCoordinate() {
      for (let spot of this.coordinates) {
        if (spot.xAxis >= 60) spot.xTurn = false;
        if (spot.xAxis <= 0) spot.xTurn = true;
        if (spot.xTurn) spot.xAxis += 1;
        else spot.xAxis -= 1;
        if (spot.yAxis >= 100) spot.yTurn = false;
        if (spot.yAxis <= 0) spot.yTurn = true;
        if (spot.yTurn) spot.yAxis += 1;
        else spot.yAxis -= 1;
      }
    },
    createNotes() {
      if (!this.$refs.aside) return;
      for (let i = 0; i < 10; i++) {
        const num = Math.floor(Math.random() * 9);
        this.notes.push([]);
        for (let j = 0; j < num; j++) {
          const opacity = 0.5 + Math.ceil(Math.random() * 5) / 10;
          this.notes[i].push(opacity);
        }
      }
    },
    resetNotes() {
      if (!this.$refs.aside) return;
      const col = Math.floor(Math.random() * 10);
      const flag = Math.random() >= 0.5 ? true : false;
      if (this.notes[col].length == 0) {
        this.notes[col].push(0.5 + Math.ceil(Math.random() * 5) / 10);
        return;
      }
      if (this.notes[col].length == 8) {
        this.notes[col].pop();
        return;
      }
      if (flag) {
        this.notes[col].push(0.5 + Math.ceil(Math.random() * 5) / 10);
        return;
      } else {
        this.notes[col].pop();
        return;
      }
    },
  },
};
</script>

<style lang="scss" scoped>
.cover {
  position: relative;
  height: 100vh;
  background-position: center center;
  background-size: cover;
}
main {
  margin: 0rem;
  position: relative;
  display: flex;
  justify-content: center;
  align-items: center;
  width: 100%;
  height: calc(100% - 8rem);
  z-index: 2;
}
aside {
  margin: 0.5rem;
  width: 12rem;
  height: 6rem;
  background-color: rgba($color: lightgray, $alpha: 0.7);
  border-radius: 0.2rem;
  ul {
    display: flex;
    list-style: none;
    padding: 1rem;
    height: 100%;
    li {
      display: flex;
      flex-direction: column-reverse;
      padding: 0.1rem;
      width: 1rem;
      height: 100%;
    }
  }
}
.note {
  margin: 0.1rem 0rem;
  height: 0.3rem;
  background-color: white;
  border-radius: 0.1rem;
}
article {
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 1rem;
  height: 6rem;
  width: 50rem;
  background-color: rgba($color: lightgray, $alpha: 0.7);
  border-radius: 0.2rem;
  h4 {
    text-align: center;
    font-weight: 500;
    letter-spacing: 0.05rem;
  }
}
footer {
  width: 100%;
  height: 8rem;
  overflow: hidden;
}

@media screen and (max-width: 759px) {
  aside {
    display: none;
  }
}
</style>