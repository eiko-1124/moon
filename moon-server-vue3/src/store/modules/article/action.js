import { getAllTag, getArticleSum, getArticleList, setArticleStatus, GetArticle, GetCommentSum, GetCommentList, SetCommentStatus, GetTagsDetails } from '@/api/apis/article'
export const actions = {
    getAllTags(content) {
        getAllTag((Res) => {
            if (Res.res == 1) {
                content.commit('setAllTags', Res.tags)
            }
        }, error => {
            console.log(error)
        })
    },
    getArticleSum(content, data) {
        getArticleSum(data, Res => {
            if (Res.res == 1) {
                content.commit('setArticleSum', Res.sum)
            }
        }, Err => {
            console.log(Err)
        })
    },
    getArticleList(content, data) {
        getArticleList(data, Res => {
            if (Res.res == 1) {
                content.commit('setArticleList', Res.list)
            }
        }, Err => {
            console.log(Err)
        })
    },
    hiddenArticle(content, [index, hFlag, id]) {
        const form = new FormData()
        form.append('id', id)
        form.append('hFlag', hFlag)
        setArticleStatus(form, Res => {
            if (Res.res == 1) content.commit('setArticleStatus', [index, hFlag])
        }, Err => {
            console.log(Err)
        })
    },
    getArticle(content, id) {
        GetArticle(id, (Res) => {
            if (Res.res == 1) {

                content.commit("initArticle", Res.article)
            }
        }, (Err) => {
            console.log(Err)
        })
    },
    getCommentSum(content, aType) {
        GetCommentSum(aType, Res => {
            if (Res.res == 1) {
                content.commit('setCommentSum', Res.sum)
            }
        }, Err => {
            console.log(Err)
        })
    },
    getCommentList(content, data) {
        GetCommentList(data, Res => {
            if (Res.res == 1) {
                content.commit('setCommentList', Res.list)
            }
        }, Err => {
            console.log(Err)
        })
    },
    hiddenComment(content, [index, date]) {
        const form = new FormData()
        form.append('id', date.id)
        form.append('title', date.title)
        form.append('hFlag', (date.hFlag + 1) % 2)
        SetCommentStatus(form, Res => {
            if (Res.res == 1) content.commit('setCommentStatus', [index, (date.hFlag + 1) % 2])
        }, Err => {
            console.log(Err)
        })
    },
    getTagsDetails(content) {
        GetTagsDetails(Res => {
            if (Res.res == 1) content.commit("setTagsDetails", Res.tags)
        }, Err => {
            console.log(Err)
        })
    }
}