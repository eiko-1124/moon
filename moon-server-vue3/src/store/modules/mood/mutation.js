export const mutations = {
    setMoodList(state, value) {
        state.moods = value
    },
    setMoodSum(state, value) {
        state.moodSum = value
    },
    loadMood(state, value) {
        state.moods = [...state.moods, ...value]
    },
    submitMood(state, value) {
        state.moodSum++
        state.moods.unshift(value[0])
    },
    loadComment(state, [value, id]) {
        for (let mood of state.moods) {
            if (mood.id == id) {
                mood.comments = [...mood.comments, ...value]
            }
        }
    },
    deleteComment(state, { id, cid }) {
        for (let mood of state.moods) {
            if (mood.id == id) {
                for (let [index, comment] of mood.comments.entries()) {
                    if (comment.cid == cid) {
                        mood.comments.splice(index, 1)
                        mood.comment--
                        for (let i = 0; i < index; i++) {
                            mood.comments[i].cid--
                        }
                        break
                    }
                }
            }
        }
    },
    deleteMood(state, id) {
        for (let [index, mood] of state.moods.entries()) {
            if (mood.id == id) {
                state.moods.splice(index, 1)
                state.moodSum--
                for (let i = 0; i < index; i++) {
                    state.moods[i].id--
                }
                break
            }
        }
    },
    setOldNote(state, note) {
        state.oldNote = note
    }
    ,
    setNotes(state, notes) {
        if (notes != null) {
            state.notes = [...state.notes, ...notes]
        }
    },
    setNoteSum(state, sum) {
        state.noteSum = sum
    },
    deleteNote(state, id) {
        for (let [index, note] of state.notes.entries()) {
            if (note.id == id) {
                state.notes.splice(index, 1)
                state.noteSum--
                for (let i = 0; i < index; i++) {
                    state.notes[i].id--
                }
                break
            }
        }
    },
}