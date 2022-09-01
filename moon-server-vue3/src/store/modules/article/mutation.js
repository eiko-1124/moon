export const mutations = {
    editArticle(state, [key, value]) {
        state.article[key] = value
    },
    setCover(state, value) {
        state.article.cover = value
    },
    setAllTags(state, value) {
        value.unshift("全部")
        state.tags = value
    },
    setArticleSum(state, value) {
        state.articleSum = value
    },
    setArticleList(state, value) {
        if (value == null) {
            state.articleList = []
        } else {
            for (let site of value) {
                if (site.aType == 1) site.aType = "技术"
                if (site.aType == 2) site.aType = "杂文"
                if (site.aType == 3) site.aType = "故事"
                let tagArr = site.tags.split("#").filter((i) => i && i.trim())
                site.tagArr = tagArr
            }
            state.articleList = value
        }
    },
    setArticleStatus(state, [index, hFlag]) {
        state.articleList[index].hFlag = hFlag
    },
    initArticle(state, article) {
        state.article = article
    },
    setCommentSum(state, value) {
        state.commentSum = value
    },
    setCommentList(state, value) {
        if (value == null) {
            state.commentList = []
        } else {
            for (let site of value) {
                if (site.aType == 1) site.aType = "技术"
                if (site.aType == 2) site.aType = "杂文"
                if (site.aType == 3) site.aType = "故事"
            }
            state.commentList = value
        }
    },
    setCommentStatus(state, [index, hFlag]) {
        state.commentList[index].hFlag = hFlag
    },
    setTagsDetails(state, tags) {
        state.tagsDetails = tags
    }
}