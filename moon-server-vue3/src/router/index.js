import { createRouter, createWebHashHistory } from 'vue-router'
import Moon from '@/views/moon/Moon'
import Login from '@/views/login/Login'
import { changeRouter } from '@/api/apis/login'
import { myStore } from '@/store'

const routes = [
  {
    path: '/',
    name: 'Moon',
    component: Moon,
    children: [
      {
        path: '/Dashboard',
        name: 'Dashboard',
        component: () => import(/* webpackChunkName: "Dashboard" */ '../views/base/Dashboard.vue'),
        meta: { name: "仪表盘" }
      }, {
        path: '/UserConf',
        name: 'UserConf',
        component: () => import(/* webpackChunkName: "UserConf" */ '../views/base/UserConf.vue'),
        meta: { name: "基本配置" }
      }, {
        path: '/BlogList',
        name: 'BlogList',
        component: () => import(/* webpackChunkName: "BlogList" */ '../views/blog/BlogList.vue'),
        meta: { name: "管理文章" }
      }, {
        path: '/EditBlog/:id',
        name: 'EditBlog',
        component: () => import(/* webpackChunkName: "EditBlog" */ '../views/blog/EditBlog.vue'),
        meta: { name: "添加文章" }
      }, {
        path: '/EditMood',
        name: 'EditMood',
        component: () => import(/* webpackChunkName: "EditMood" */ '../views/mood/EditMood.vue'),
        meta: { name: "发布记录" }
      }, {
        path: '/Note',
        name: 'Note',
        component: () => import(/* webpackChunkName: "Note" */ '../views/mood/Note.vue'),
        meta: { name: "笔记安排" }
      }, {
        path: '/Comment',
        name: 'Comment',
        component: () => import(/* webpackChunkName: "Comment" */ '../views/comment/Comment.vue'),
        meta: { name: "我的评论" }
      }, {
        path: '/Tag',
        name: 'Tag',
        component: () => import(/* webpackChunkName: "Comment" */ '../views/tag/Tag.vue'),
        meta: { name: "我的标签" }
      }, {
        path: '/Recovery',
        name: 'Recovery',
        component: () => import(/* webpackChunkName: "Recovery" */ '../views/recovery/Recovery.vue'),
        meta: { name: "文件回收" }
      }, {
        path: '/About',
        name: 'About',
        component: () => import(/* webpackChunkName: "About" */ '../views/journal/About.vue'),
        meta: { name: "关于本站" }
      }, {
        path: '/Journal',
        name: 'Journal',
        component: () => import(/* webpackChunkName: "Journal" */ '../views/journal/Journal.vue'),
        meta: { name: "我的日志" }
      },
    ]
  },
  {
    path: '/Login',
    name: 'Login',
    component: Login
  }
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

router.beforeEach(async (to, form, next) => {
  if (to.name != "Login") {
    let hasToken = false
    const cookies = document.cookie.split(';');
    for (let i = 0; i < cookies.length; i++) {
      const cookie = cookies[i].trim();
      if (cookie.indexOf("token=") == 0) {
        hasToken = true
        break
      }
    }
    if (!hasToken) {
      next({
        name: 'Login'
      })
    }
    else {
      const res = await changeRouter()
      if (!res) {
        next({
          name: 'Login'
        })
      } else {
        if (to.name == 'Moon') {
          next({
            name: 'Dashboard'
          })
        }
        else next()
      }
    }
  }
  else next()
})

const store = myStore()

router.afterEach((to, from) => {
  let cache = null
  if (from.name != undefined) cache = { path: from.path, name: from.meta.name }
  store.commit('setViewCache', [to.meta.name, cache])
})

export default router
