import { article } from "../../entity/article.entity"
import { comment } from "../../entity/comment.entity"

export interface sumRes {
    res: number,
    sum: number
}

export interface tagsRes {
    res: number,
    tags: string[]
}

export interface articleListRes {
    res: number,
    list: articleSite[]
}

export interface articleSite {
    id: number,
    title: string,
    atype: number,
    illustrate: string,
    cover: string,
    time: string,
    tags: string
    comment: number
}

export interface recommendRes {
    res: number,
    list: recommend[]
}

export interface recommend {
    id: number,
    title: string,
    time: string,
    cover: string
}

export interface titleRes {
    res: number,
    titles: articleTitle[]
}

interface articleTitle {
    id: number,
    title: string,
    time: string
}

export interface cTagRes {
    res: number,
    tags: string[]
}

export interface typeRes {
    res: number,
    types: number[]
}

export interface articleRes {
    res: number,
    article: article,
    comment: number
}

export interface nearRes {
    res: number
    near: near
}

export interface near {
    nextId: number
    nextTitle: string
    preId: number
    preTitle: string
}

export interface commentRes {
    res: number
    comments: comment[]
}

export interface commentForm {
    title: string
    id: number
    cid: number
    time: string
    content: string
    name: string
    cname: string
    cflag: number
    hflag: number
    email: string
}

export interface submitRes {
    res: number
}