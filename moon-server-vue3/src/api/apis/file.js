import axios from "../axios/axios.js"

export const minorFile = (form, percentage, success, err) => {
    axios.post("/api/file/minorFile", form, {
        onUploadProgress: (progressEvent) => {
            if (progressEvent.lengthComputable) {
                percentage.Value = Math.floor(progressEvent.loaded / progressEvent.total * 100)
            }
        }
    }).then(Res => {
        success(Res.data)
    }).catch(Err => {
        err(Err)
    })
}

export const getAllFiles = (page, success, error) => {
    axios.get("/api/file/getAllFiles", {
        params: {
            page
        }
    }).then(Res => {
        success(Res.data)
    }).catch(Err => {
        error(Err)
    })
}

export const moodPicture = (form, success, error) => {
    axios.post("/api/file/moodPicture", form).then(Res => {
        success(Res.data)
    }).catch(Err => {
        error(Err)
    })
}

export const GetArticleCoverSum = (success, error) => {
    axios.get("/api/file/getArticleCoverSum").then(Res => {
        success(Res.data)
    }).catch(Err => {
        error(Err)
    })
}

export const GetArticleCovers = (page, success, error) => {
    axios.get("/api/file/getArticleCovers", {
        params: {
            page
        }
    }).then(Res => {
        success(Res.data)
    }).catch(Err => {
        error(Err)
    })
}

export const GetMoodImageSum = (success, error) => {
    axios.get("/api/file/getMoodImageSum").then(Res => {
        success(Res.data)
    }).catch(Err => {
        error(Err)
    })
}

export const GetMoodImages = (page, success, error) => {
    axios.get("/api/file/getMoodImages", {
        params: {
            page
        }
    }).then(Res => {
        success(Res.data)
    }).catch(Err => {
        error(Err)
    })
}

export const GetOtherFileSum = (success, error) => {
    axios.get("/api/file/getOtherFileSum").then(Res => {
        success(Res.data)
    }).catch(Err => {
        error(Err)
    })
}

export const GetOtherFiles = (page, success, error) => {
    axios.get("/api/file/getOtherFiles", {
        params: {
            page
        }
    }).then(Res => {
        success(Res.data)
    }).catch(Err => {
        error(Err)
    })
}

// export const bigFile = (file, percentage, success, err) => {
//     const name = file.name
//     const totalSize = file.size; // 文件总大小
//     let start = 0; // 每次上传的开始字节
//     let end = start + 5 * 1024 * 1024; // 每次上传的结尾字节
//     let chunks = []
//     while (start < totalSize) {
//         let blob = file.slice(start, end);
//         chunks.push(blob)
//         start = end;
//         end = start + 5 * 1024 * 1024;
//     }
//     Promise.all(chunks.map((chunk, id) => {
//         return spliceFile(id, name, chunk)
//     })).then(() => {
//         console.log(1)
//     }).catch(() => {
//         console.log(2)
//     })
// }


// const spliceFile = (id, name, chunk) => {
//     return new Promise((resolve, reject) => {
//         const form = new FormData()
//         form.append('id', id)
//         form.append('fileName', name)
//         form.append('file', chunk)
//         axios.post('/api/file/filePiece', form).then(Res => {
//             resolve()
//         }).catch(Err => {
//             console.log(Err)
//             reject()
//         })
//     })
// }

// const spliceFile = (id, name, chunks) => {
//     return new Promise((resolve, reject) => {
//         if (id == chunks.length) resolve()
//         else {
//             const form = new FormData()
//             form.append('id', id)
//             form.append('fileName', name)
//             form.append('file', chunks[id])
//             axios.post('/api/file/filePiece', form).then(Res => {
//                 console.log(id)
//                 spliceFile(id + 1, name, chunks).then(() => resolve())
//             }).catch(Err => {
//                 console.log(Err)
//                 reject()
//             })
//         }
//     })
// }