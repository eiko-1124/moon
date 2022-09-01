import { Tickets, ChatLineSquare, CollectionTag, CoffeeCup } from "@element-plus/icons-vue";
export default [
    {
        title: '文章统计',
        text: '篇文章',
        key: "article",
        icon: Tickets,
        color: 'lightseagreen',
        link: [{
            title: '写文章',
            path: '/EditBlog/0'
        }, {
            title: '读文章',
            path: '/BlogList'
        }]
    }, {
        title: '评论统计',
        text: '条评论',
        key: "comment",
        icon: ChatLineSquare,
        color: 'dodgerblue',
        link: [{
            title: '管理评论',
            path: '/Comment'
        }, {
            title: '查看评论',
            path: '/Comment'
        }]
    }, {
        title: '标签统计',
        text: '个标签',
        key: "tag",
        icon: CollectionTag,
        color: 'lightcoral',
        link: [{
            title: '管理标签',
            path: '/Tag'
        }, {
            title: '查看标签',
            path: '/Tag'
        }]
    }, {
        title: '记录统计',
        text: '个记录',
        key: "mood",
        icon: CoffeeCup,
        color: 'darkorange',
        link: [{
            title: '发布记录',
            path: '/EditMood'
        }, {
            title: '管理记录',
            path: '/Note'
        }]
    }
]