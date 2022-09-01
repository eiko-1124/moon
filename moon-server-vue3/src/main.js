import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import eventBus from 'vue3-eventbus'
import * as ElementPlusIconsVue from '@element-plus/icons-vue'
import '../src/assets/css/main.css'

const app = createApp(App)
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
    app.component(key, component)
}
app.use(eventBus).use(store).use(router).mount('#app')