import { minorFile, getAllFiles, GetArticleCoverSum, GetArticleCovers, GetMoodImageSum, GetMoodImages, GetOtherFileSum, GetOtherFiles } from '@/api/apis/file'
export const actions = {
    uploadFile(content, file) {
        const id = Symbol()
        const fileName = file.name
        let percentage = 0
        let per = {
            set Value(value) {
                content.commit('setPercentage', [id, value])
            }
        }
        content.commit('waitingUpload', ({ id, fileName, percentage }))
        const form = new FormData()
        form.append('file', file)
        form.append('fileName', fileName)
        minorFile(form, per, (Res) => {
            if (Res.res == "1") {
                content.commit('overUpload', [id, {
                    fileName: Res.file.fileName,
                    link: Res.file.link
                }])
            }
            else {
                content.commit('failUpload', id)
            }
        }, (Err) => {
            console.log(Err)
            content.commit('failUpload', id)
        })
    },
    getAllFiles(content, page) {
        getAllFiles(page, (data) => {
            content.commit('setAllFiles', data)
        }, (Err) => {
            console.log(Err)
        })
    },
    getArticleCoverSum(content) {
        GetArticleCoverSum(Res => {
            if (Res.res == 1) {
                content.commit('setArticleCoverSum', Res.sum)
            }
        }, Err => {
            console.log(Err)
        })
    },
    getArticleCovers(content, [page, loading]) {
        GetArticleCovers(page, Res => {
            if (Res.res == 1) {
                loading.value = false
                content.commit('setArticleCovers', Res.covers)
            }
        }, Err => {
            console.log(Err)
        })
    },
    getMoodImageSum(content) {
        GetMoodImageSum(Res => {
            if (Res.res == 1) {
                content.commit('setMoodImageSum', Res.sum)
            }
        }, Err => {
            console.log(Err)
        })
    },
    getMoodImages(content, [page, loading]) {
        GetMoodImages(page, Res => {
            if (Res.res == 1) {
                loading.value = false
                content.commit('setMoodImages', Res.images)
            }
        }, Err => {
            console.log(Err)
        })
    },
    getOtherFileSum(content) {
        GetOtherFileSum(Res => {
            if (Res.res == 1) {
                content.commit('setOtherFileSum', Res.sum)
            }
        }, Err => {
            console.log(Err)
        })
    },
    getOtherFiles(content, [page, loading]) {
        GetOtherFiles(page, Res => {
            if (Res.res == 1) {
                loading.value = false
                content.commit('setOtherFiles', Res.files)
            }
        }, Err => {
            console.log(Err)
        })
    },
}