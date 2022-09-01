import { GetDiaryTime, GetDiarySum, GetDiaryAll, GetMessageSum, GetMessages, DeleteMessage } from "@/api/apis/diary"
export const actions = {
    getATime(content, [loading, date]) {
        GetDiaryTime(date, Res => {
            if (Res.res == 1) {
                loading.value = false
                content.commit("setATime", Res.diarys)
            }
        }, Err => {
            console.log(Err)
            loading.value = false
        })
    },
    getASum(content) {
        GetDiarySum(0, Res => {
            if (Res.res == 1) {
                content.commit("setASum", Res.sum)
            }
        }, Err => {
            console.log(Err)
        })
    },
    getADiary(content, time) {
        const data = {
            page: time.diary.length,
            type: 0,
            date: time.date
        }
        GetDiaryAll(data, Res => {
            if (Res.res == 1) {
                time.loading = false
                content.commit("setADiary", [data.date, Res.diarys])
            }
        }, Err => {
            console.log(Err)
            time.loading = false
        })
    },
    getUTime(content, [loading, date]) {
        GetDiaryTime(date, Res => {
            if (Res.res == 1) {
                loading.value = false
                content.commit("setUTime", Res.diarys)
            }
        }, Err => {
            console.log(Err)
            loading.value = false
        })
    },
    getUSum(content) {
        GetDiarySum(1, Res => {
            if (Res.res == 1) {
                content.commit("setUSum", Res.sum)
            }
        }, Err => {
            console.log(Err)
        })
    },
    getUDiary(content, time) {
        const data = {
            page: time.diary.length,
            type: 1,
            date: time.date
        }
        GetDiaryAll(data, Res => {
            if (Res.res == 1) {
                time.loading = false
                content.commit("setUDiary", [data.date, Res.diarys])
            }
        }, Err => {
            console.log(Err)
            time.loading = false
        })
    },
    getMessageSum(content) {
        GetMessageSum(Res => {
            if (Res.res == 1) {
                content.commit("setMessageSum", Res.sum)
            }
        }, Err => {
            console.log(Err)
        })
    },
    getMessages(content, [page, loading]) {
        GetMessages(page, Res => {
            if (Res.res == 1) {
                loading.value = false
                content.commit("setMessages", Res.messages)
            }
        }, Err => {
            console.log(Err)
        })
    },
    deleteMessage(content, [id, success, error]) {
        const form = new FormData()
        form.append('id', id)
        DeleteMessage(form, Res => {
            if (Res.res == 1) {
                success()
                content.commit("deleteMessage", id)
            } else error()
        }, Err => {
            console.log(Err)
            error()
        })
    },
}