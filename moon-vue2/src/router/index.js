import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'
import { myStore } from "@/store"

Vue.use(VueRouter)

const routes = [
  {
    path: '/Home/:page',
    name: 'Home',
    component: Home,
    meta: {
      background: "http://localhost:3000/static/background/home.png"
    }
  },
  {
    path: '/Article/:id',
    name: 'Article',
    component: () => import(/* webpackChunkName: "article" */ '../views/Article.vue')
  },
  {
    path: '/Message',
    name: 'Message',
    component: () => import(/* webpackChunkName: "message" */ '../views/Message.vue'),
    meta: {
      background: "http://localhost:3000/static/background/message.png"
    }
  },
  {
    path: '/About',
    name: 'About',
    component: () => import(/* webpackChunkName: "about" */ '../views/About.vue'),
    meta: {
      background: "http://localhost:3000/static/background/about.png"
    }
  },
  {
    path: '/Link',
    name: 'Link',
    component: () => import(/* webpackChunkName: "link" */ '../views/Link.vue'),
    meta: {
      background: "http://localhost:3000/static/background/link.png"
    }
  },
  {
    path: '/Record',
    name: 'Record',
    component: () => import(/* webpackChunkName: "link" */ '../views/Record.vue'),
    meta: {
      background: "http://localhost:3000/static/background/record.png"
    }
  },
  {
    path: '/PickOn',
    name: 'PickOn',
    component: () => import(/* webpackChunkName: "link" */ '../views/PickOn.vue'),
    meta: {
      background: "http://localhost:3000/static/background/pickon.png"
    }
  },
  {
    path: '/Classify/:type/:tag/:page',
    name: 'Classify',
    component: () => import(/* webpackChunkName: "link" */ '../views/Classify.vue'),
    meta: {
      background: "http://localhost:3000/static/background/classify.png"
    }
  }
]


const originalPush = VueRouter.prototype.push

VueRouter.prototype.push = function push(location) {
  return originalPush.call(this, location).catch(err => err)
}

const router = new VueRouter({
  routes
})

const store = myStore()

router.beforeEach((to, form, next) => {
  if (to.path == '/') router.push('/Home/1')
  store.commit("setLoad", true)
  if (to.name != 'Article') {
    const image = new Image();
    image.src = to.meta.background;
    image.onload = () => {
      next();
      store.commit("setLoad", false)
      setTimeout(() => {
        if (!store.state.navFlag)
          store.commit("setNav", true)
      }, 200)
    }
  } else {
    next()
  }
})

export default router
