import axios from "../axios/axios.js"

export const goLogin = (form, success, error) => {
    axios.post('/login',
        form
    ).then((Response) => {
        success(Response.data)
    }).catch((Error => {
        error(Error.response.data)
    }))
}

export const changeRouter = () => {
    const promise = new Promise((resolve, reject) => {
        axios.get('/api/admin/changeRouter').then(Response => {
            resolve(true)
        }).catch(Error => {
            if (Error.response.status == 401 || Error.response.status == 400)
                resolve(false)
            else reject(false)
        })
    })
    return promise
}
