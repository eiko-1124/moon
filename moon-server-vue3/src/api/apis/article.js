import axios from "../axios/axios.js"

export const submitArticle = (form, success, error) => {
    axios.post("./api/article/submit", form).then(Res => {
        success(Res.data)
    }).catch(Err => {
        error(Err)
    })
}

export const getAllTag = (success, error) => {
    axios.get("./api/article/getTags").then(Res => {
        success(Res.data)
    }).catch(Err => {
        error(Err)
    })
}

export const GetTagsDetails = (success, error) => {
    axios.get("./api/article/getTagsDetails").then(Res => {
        success(Res.data)
    }).catch(Err => {
        error(Err)
    })
}

export const getArticleSum = (data, success, error) => {
    axios.get("./api/article/getArticleSum", {
        params: {
            type: data.type,
            time: data.time,
            tag: data.tag
        }
    }).then(Res => {
        success(Res.data)
    }).catch(Err => {
        error(Err)
    })
}

export const getArticleList = (data, success, error) => {
    axios.get("./api/article/getArticleList", {
        params: {
            type: data.type,
            time: data.time,
            tag: data.tag,
            page: data.page
        }
    }).then(Res => {
        success(Res.data)
    }).catch(Err => {
        error(Err)
    })
}

export const setArticleStatus = (form, success, error) => {
    axios.post("./api/article/setArticleStatus", form).then(Res => {
        success(Res.data)
    }).catch(Err => {
        error(Err)
    })
}

export const DeleteArticle = (form, success, error) => {
    axios.post("./api/article/deleteArticle", form).then(Res => {
        success(Res.data)
    }).catch(Err => {
        error(Err)
    })
}

export const UpdateMood = (form, success, error) => {
    axios.post("./api/article/updateMood", form).then(Res => {
        success(Res.data)
    }).catch(Err => {
        error(Err)
    })
}

export const GetArticle = (id, success, error) => {
    axios.get("./api/article/getArticle", {
        params: {
            id
        }
    }).then(Res => {
        success(Res.data)
    }).catch(Err => {
        error(Err)
    })
}

export const GetCommentSum = (aType, success, error) => {
    axios.get("./api/article/getCommentSum", {
        params: {
            type: aType
        }
    }).then(Res => {
        success(Res.data)
    }).catch(Err => {
        error(Err)
    })
}

export const GetCommentList = (data, success, error) => {
    axios.get("./api/article/getCommentList", {
        params: {
            type: data.type,
            page: data.page
        }
    }).then(Res => {
        success(Res.data)
    }).catch(Err => {
        error(Err)
    })
}

export const SetCommentStatus = (form, success, error) => {
    axios.post("./api/article/setCommentStatus", form).then(Res => {
        success(Res.data)
    }).catch(Err => {
        error(Err)
    })
}

export const DeleteComments = (form, success, error) => {
    axios.post("./api/article/deleteComment", form).then(Res => {
        success(Res.data)
    }).catch(Err => {
        error(Err)
    })
}