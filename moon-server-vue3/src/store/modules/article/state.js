export const state = () => {
    return {
        article: {
            id: 0,
            title: '',
            aType: 0,
            tag: '',
            illustrate: '',
            cover: '',
            fContent: '',
            rContent: ''
        },
        tags: [],
        articleSum: 0,
        articleList: [],
        commentSum: 0,
        commentList: [],
        tagsDetails: []
    }
}