import axios from "../axios/axios.js"

export const getInfo = (success, error) => {
    axios.get('/api/admin/getInfo',
    ).then((Response) => {
        success(Response.data)
    }).catch((Error => {
        error(Error.response.data)
    }))
}

export const getSigns = (success, error) => {
    axios.get('/api/admin/getSigns',
    ).then((Response) => {
        success(Response.data)
    }).catch((Error => {
        error(Error.response.data)
    }))
}

export const getLinks = (success, error) => {
    axios.get('/api/admin/getLinks',
    ).then((Response) => {
        success(Response.data)
    }).catch((Error => {
        error(Error.response.data)
    }))
}

export const getArticleSum = (success, error) => {
    axios.get('/api/admin/getArticleSum',
    ).then((Response) => {
        success(Response.data)
    }).catch((Error => {
        error(Error)
    }))
}

export const getMoodSum = (success, error) => {
    axios.get('/api/admin/getMoodSum',
    ).then((Response) => {
        success(Response.data)
    }).catch((Error => {
        error(Error)
    }))
}

export const getLinkSum = (success, error) => {
    axios.get('/api/admin/getLinkSum',
    ).then((Response) => {
        success(Response.data)
    }).catch((Error => {
        error(Error)
    }))
}

export const getTagSum = (success, error) => {
    axios.get('/api/admin/getTagSum',
    ).then((Response) => {
        success(Response.data)
    }).catch((Error => {
        error(Error)
    }))
}

export const getCommentSum = (success, error) => {
    axios.get('/api/admin/getCommentSum',
    ).then((Response) => {
        success(Response.data)
    }).catch((Error => {
        error(Error)
    }))
}

export const updateLink = (form, success, error) => {
    axios.post('api/admin/setLink', form).then(Response => {
        success(Response.data)
    }).catch(Error => {
        error(Error)
    })
}

export const updateSign = (form, success, error) => {
    axios.post('api/admin/setSign', form).then(Response => {
        success(Response.data)
    }).catch(Error => {
        error(Error)
    })
}

export const deleteLink = (form, success, error) => {
    axios.post('api/admin/deleteLink', form).then(Response => {
        success(Response.data)
    }).catch(Error => {
        error(Error)
    })
}

export const updateInfo = (form, success, error) => {
    axios.post('api/admin/updateInfo', form).then(Response => {
        success(Response.data)
    }).catch(Error => {
        error(Error)
    })
}

export const UpdateAbout = (form, success, error) => {
    axios.post('api/admin/updateAbout', form).then(Response => {
        success(Response.data)
    }).catch(Error => {
        error(Error)
    })
}

export const GetAbout = (success, error) => {
    axios.get('/api/admin/getAbout',
    ).then((Response) => {
        success(Response.data)
    }).catch((Error => {
        error(Error)
    }))
}

export const GetTags = (success, error) => {
    axios.get('/api/article/getMaxTags',
    ).then((Response) => {
        success(Response.data)
    }).catch((Error => {
        error(Error)
    }))
}

export const GetNear = (dType, success, error) => {
    axios.get('/api/diary/getDiaryNear', {
        params: {
            type: dType
        }
    }
    ).then((Response) => {
        success(Response.data)
    }).catch((Error => {
        error(Error)
    }))
}

export const GetDiary = (dType, success, error) => {
    axios.get('/api/diary/getDiary', {
        params: {
            type: dType
        }
    }
    ).then((Response) => {
        success(Response.data)
    }).catch((Error => {
        error(Error)
    }))
}