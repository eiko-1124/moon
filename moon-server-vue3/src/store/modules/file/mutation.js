export const mutations = {
    waitingUpload(state, file) {
        state.uploading.push(file)
    },
    setPercentage(state, [id, value]) {
        for (let file of state.uploading) {
            if (file.id == id) {
                file.percentage = value
            }
        }
    },
    overUpload(state, [id, file]) {
        for (let i = 0; i < state.uploading.length; i++) {
            if (state.uploading[i].id == id) {
                state.uploading.splice(i, 1)
            }
        }
        state.hadUpload.unshift(file)
    },
    failUpload(state, id) {
        for (let i = 0; i < state.uploading.length; i++) {
            if (state.uploading[i].id == id) {
                state.uploading.splice(i, 1)
            }
        }
    },
    setAllFiles(state, files) {
        if (files != null) state.allUpload = [...state.allUpload, ...files]
    },
    setArticleCoverSum(state, sum) {
        state.articleCoverSum = sum
    },
    setArticleCovers(state, covers) {
        if (covers != null) state.articleCovers = [...state.articleCovers, ...covers]
    },
    setMoodImageSum(state, sum) {
        state.moodImageSum = sum
    },
    setMoodImages(state, covers) {
        if (covers != null) state.moodImages = [...state.moodImages, ...covers]
    },
    setOtherFileSum(state, sum) {
        state.otherFileSum = sum
    },
    setOtherFiles(state, links) {
        if (links != null) state.otherFiles = [...state.otherFiles, ...links]
    },
}