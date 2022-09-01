import { createStore } from 'vuex'
import { User } from './modules/user'
import { Article } from './modules/article'
import { File } from './modules/file'
import { Mood } from './modules/mood'
import { Diary } from './modules/diary'

const store = createStore({
  modules: {
    User,
    Article,
    File,
    Mood,
    Diary
  }
})

export const myStore = () => store

export default store
