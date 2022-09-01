import { GetMoodList, GetMoodSum, GetComments, DeleteComment, DeleteMood, GetOldNote, GetNotes, GetNoteSum, DeleteNote } from "@/api/apis/mood"
export const actions = {
    getMoodList(content) {
        GetMoodList(0, Res => {
            if (Res.res == 1) {
                for (let mood of Res.moods) {
                    if (mood.links == null) {
                        mood.links = []
                    }
                    if (mood.comments == null) {
                        mood.comments = []
                    }
                }
                content.commit('setMoodList', Res.moods)
            }
        }, Err => {
            console.log(Err)
        })
    },
    submitMood(content) {
        GetMoodList(0, Res => {
            if (Res.res == 1) {
                for (let mood of Res.moods) {
                    if (mood.links == null) {
                        mood.links = []
                    }
                    if (mood.comments == null) {
                        mood.comments = []
                    }
                }
                content.commit('submitMood', Res.moods)
            }
        }, Err => {
            console.log(Err)
        })
    },
    getMoodSum(content) {
        GetMoodSum(Res => {
            if (Res.res == 1) {
                content.commit('setMoodSum', Res.sum)
            }
        }, Err => {
            console.log(Err)
        })
    },
    loadMood(content, num) {
        GetMoodList(num, Res => {
            if (Res.res == 1 && Res.moods != null) {
                for (let mood of Res.moods) {
                    if (mood.links == null) {
                        mood.links = []
                    }
                    if (mood.comments == null) {
                        mood.comments = []
                    }
                }
                content.commit('loadMood', Res.moods)
            }
        }, Err => {
            console.log(Err)
        })
    },
    loadComments(content, date) {
        GetComments(date, Res => {
            if (Res.res == 1 && Res.comments != null) {
                content.commit('loadComment', [Res.comments, date.id])
            }
        }, Err => {
            console.log(Err)
        })
    },
    deleteComment(content, [date, success, error]) {
        const form = new FormData()
        form.append('id', date.id)
        form.append('cid', date.cid)
        DeleteComment(form, Res => {
            if (Res.res == 1) {
                success()
                content.commit("deleteComment", date)
            } else error()
        }, Err => {
            console.log(Err)
            error()
        })
    },
    deleteMood(content, [id, success, error]) {
        const form = new FormData()
        form.append('id', id)
        DeleteMood(form, Res => {
            if (Res.res == 1) {
                success()
                content.commit("deleteMood", id)
            } else error()
        }, Err => {
            console.log(Err)
            error()
        })
    },
    getOldNote(content) {
        GetOldNote(Res => {
            if (Res.res == 1) {
                content.commit("setOldNote", Res.note)
            }
        }, Err => {
            console.log(Err)
        })
    },
    getNotes(content, loading) {
        GetNotes(content.state.notes.length, Res => {
            if (Res.res == 1) {
                content.commit("setNotes", Res.notes)
                loading.value = false
            }
        }, Err => {
            console.log(Err)
        })
    },
    getNoteSum(content) {
        GetNoteSum(Res => {
            if (Res.res == 1) {
                content.commit("setNoteSum", Res.sum)
            }
        }, Err => {
            console.log(Err)
        })
    },
    deleteNote(content, [id, success, error]) {
        const form = new FormData()
        form.append("id", id)
        DeleteNote(form, Res => {
            if (Res.res == 1) {
                success()
                content.commit("deleteNote", id)
            } else error()
        }, Err => {
            console.log(Err)
            error()
        })
    }
}