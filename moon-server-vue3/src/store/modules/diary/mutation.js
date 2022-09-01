
export const mutations = {
    setATime(state, value) {
        if (value != null) {
            for (let site of value) {
                site.flag = false
                site.loading = false
                site.diary = []
                state.aTime.push(site)
            }
        }
    },
    setASum(state, value) {
        state.aTimeSum = value
    },
    setADiary(state, [time, value]) {
        if (value != null) {
            for (let site of state.aTime) {
                if (site.date == time) {
                    site.diary = [...site.diary, ...value]
                    break
                }
            }
        }
    },
    setUTime(state, value) {
        if (value != null) {
            for (let site of value) {
                site.flag = false
                site.loading = false
                site.diary = []
                state.uTime.push(site)
            }
        }
    },
    setUSum(state, value) {
        state.uTimeSum = value
    },
    setUDiary(state, [time, value]) {
        if (value != null) {
            for (let site of state.uTime) {
                if (site.date == time) {
                    site.diary = [...site.diary, ...value]
                    break
                }
            }
        }
    },
    setMessageSum(state, value) {
        state.messageSum = value
    },
    setMessages(state, messages) {
        if (messages != null) {
            state.messages = [...state.messages, ...messages]
        }
    },
    deleteMessage(state, id) {
        for (let [index, message] of state.messages.entries()) {
            if (message.id == id) {
                state.messages.splice(index, 1)
                state.moodSum--
                for (let i = 0; i < index; i++) {
                    state.messages[i].id--
                }
                break
            }
        }
    }
}