import axios from "../axios/axios.js"

export const MoodSubmit = (form, success, error) => {
    axios.post("./api/mood/moodSubmit", form).then(Res => {
        success(Res.data)
    }).catch(Err => {
        error(Err)
    })
}

export const GetMoodList = (num, success, error) => {
    axios.get("./api/mood/getMoodList", {
        params: {
            num
        }
    }).then(Res => {
        success(Res.data)
    }).catch(Err => {
        error(Err)
    })
}

export const GetMoodSum = (success, error) => {
    axios.get("./api/mood/getMoodSum").then(Res => {
        success(Res.data)
    }).catch(Err => {
        error(Err)
    })
}

export const GetComments = (date, success, error) => {
    axios.get("./api/mood/getComments", {
        params: {
            'id': date.id,
            'page': date.page
        }
    }).then(Res => {
        success(Res.data)
    }).catch(Err => {
        error(Err)
    })
}

export const DeleteComment = (form, success, error) => {
    axios.post("./api/mood/deleteComment", form).then(Res => {
        success(Res.data)
    }).catch(Err => {
        error(Err)
    })
}

export const DeleteMood = (form, success, error) => {
    axios.post("./api/mood/deleteMood", form).then(Res => {
        success(Res.data)
    }).catch(Err => {
        error(Err)
    })
}

export const SetNote = (form, success, error) => {
    axios.post("./api/mood/noteSubmit", form).then(Res => {
        success(Res.data)
    }).catch(Err => {
        error(Err)
    })
}

export const GetOldNote = (success, error) => {
    axios.get("./api/mood/getOldNote").then(Res => {
        success(Res.data)
    }).catch(Err => {
        error(Err)
    })
}

export const GetNotes = (page, success, error) => {
    axios.get("./api/mood/getNewNotes", {
        params: {
            page
        }
    }).then(Res => {
        success(Res.data)
    }).catch(Err => {
        error(Err)
    })
}

export const GetNoteSum = (success, error) => {
    axios.get("./api/mood/getNoteSum").then(Res => {
        success(Res.data)
    }).catch(Err => {
        error(Err)
    })
}

export const DeleteNote = (form, success, error) => {
    axios.post("./api/mood/deleteNote", form).then(Res => {
        success(Res.data)
    }).catch(Err => {
        error(Err)
    })
}