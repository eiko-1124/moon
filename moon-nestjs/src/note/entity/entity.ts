import { message } from "src/entity/message.entity"
import { mood } from "../../entity/mood.entity"
import { pcomment } from "../../entity/pcomment.entity"

export interface moodRes {
    res: number
    moods: mood[]
}

export interface commentRes {
    res: number
    comments: pcomment[]
}

export interface comment {
    id: number
    name: string
    cname: string
    content: string
    cflag: number
    time: string
}

export interface submitRes {
    res: number
}

export interface imageRes {
    res: number
    paths: string[]
}

export interface likeRes {
    res: number
}

export interface noteRes {
    res: number
    notes: noteSite[]
}

interface noteSite {
    time: string
    text: string
}

export interface messageRes {
    res: number
    messages: message[]
}

export interface messageForm {
    name: string
    text: string
    mtype: number
    time: string
}