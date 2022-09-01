import axios from "../axios/axios.js"

export const GetDiaryTime = (date, success, error) => {
    axios.get("./api/diary/getDiaryTime", {
        params: {
            type: date.type,
            page: date.page
        }
    }).then(Res => {
        success(Res.data)
    }).catch(Err => {
        error(Err)
    })
}

export const GetDiarySum = (dType, success, error) => {
    axios.get("./api/diary/getDiarySum", {
        params: {
            type: dType
        }
    }).then(Res => {
        success(Res.data)
    }).catch(Err => {
        error(Err)
    })
}

export const GetDiaryAll = (data, success, error) => {
    axios.get("./api/diary/getDiaryAll", {
        params: {
            type: data.type,
            page: data.page,
            date: data.date
        }
    }).then(Res => {
        success(Res.data)
    }).catch(Err => {
        error(Err)
    })
}

export const GetMessageSum = (success, error) => {
    axios.get("./api/diary/getMessageSum"
    ).then(Res => {
        success(Res.data)
    }).catch(Err => {
        error(Err)
    })
}

export const GetMessages = (page, success, error) => {
    axios.get("./api/diary/getMessage", {
        params: {
            page
        }
    }).then(Res => {
        success(Res.data)
    }).catch(Err => {
        error(Err)
    })
}

export const DeleteMessage = (form, success, error) => {
    axios.post("./api/diary/deleteMessage", form).then(Res => {
        success(Res.data)
    }).catch(Err => {
        error(Err)
    })
}