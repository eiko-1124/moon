import { Bicycle, ChatLineSquare, CollectionTag, FolderOpened, Grid, Notebook, Tickets } from "@element-plus/icons-vue";

export const menu = [
    {
        mainNav: '基本配置',
        icon: Grid,
        subNav: [
            {
                title: '仪表盘',
                path: '/Dashboard'
            }, {
                title: '基本配置',
                path: '/UserConf'
            }
        ]
    }, {
        mainNav: '博客管理',
        icon: Tickets,
        subNav: [
            {
                title: '添加文章',
                path: '/EditBlog/0'
            }, {
                title: '管理文章',
                path: '/BlogList'
            }
        ]
    }, {
        mainNav: '记录笔记',
        icon: Bicycle,
        subNav: [
            {
                title: '发布记录',
                path: '/EditMood'
            }, {
                title: '笔记安排',
                path: '/Note'
            }
        ]
    }, {
        mainNav: '评论管理',
        icon: ChatLineSquare,
        subNav: [
            {
                title: '我的评论',
                path: '/Comment'
            }
        ]
    }, {
        mainNav: '标签管理',
        icon: CollectionTag,
        subNav: [
            {
                title: '我的标签',
                path: '/Tag'
            }
        ]
    }, {
        mainNav: '资源回收',
        icon: FolderOpened,
        subNav: [
            {
                title: '文件回收',
                path: '/Recovery'
            }
        ]
    }, {
        mainNav: '日志信息',
        icon: Notebook,
        subNav: [
            {
                title: '我的日志',
                path: '/Journal'
            }, {
                title: '关于本站',
                path: '/About'
            }
        ]
    }
]

export const routerIndex = {
    Dashboard: [0, 0],
    UserConf: [0, 1],
    EditBlog: [1, 0],
    BlogList: [1, 1],
    EditMood: [2, 0],
    Note: [2, 1],
    Comment: [3, 0],
    Tag: [4, 0],
    Recovery: [5, 0],
    Journal: [6, 0],
    About: [6, 1]
}