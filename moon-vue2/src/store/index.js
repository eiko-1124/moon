import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

const store = new Vuex.Store({
  state: {
    navFlag: false,
    loadFlag: true
  },
  mutations: {
    setNav(state, flag) {
      state.navFlag = flag
    },
    setLoad(state, flag) {
      state.loadFlag = flag
    }
  },
  actions: {
  },
  modules: {
  }
})

export default store

export const myStore = function () {
  return store
}