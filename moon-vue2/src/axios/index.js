import axios from "axios";
import qs from 'qs'

const config = {
    baseURL: 'http://localhost:3000/',
    timeout: 20000,
}

const _axios = axios.create(config)

export default {
    get: (url, params) => {
        return new Promise((resolve, reject) => {
            _axios.get(url, { params }).then(res => {
                resolve(res.data)
            }).catch(err => {
                reject(err)
            })
        })
    },
    post: (url, params) => {
        return new Promise((resolve, reject) => {
            _axios.post(url, qs.stringify(params)).then(res => {
                resolve(res.data)
            }).catch(err => {
                reject(err)
            })
        })
    }
}