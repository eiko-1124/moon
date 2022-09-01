import { Injectable } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { Repository } from 'typeorm';
import { mood } from '../entity/mood.entity';
import { pcomment } from '../entity/pcomment.entity';
import { mImage } from '../entity/mImage.entity';
import { note } from '../entity/note.entity';
import { message } from '../entity/message.entity';
import { diary } from '../entity/diary.entity';
import { comment, commentRes, imageRes, likeRes, moodRes, submitRes, noteRes, messageRes, messageForm } from './entity/entity';

@Injectable()
export class NoteService {
    constructor(
        @InjectRepository(mood)
        private moodRepository: Repository<mood>,

        @InjectRepository(pcomment)
        private commentRepository: Repository<pcomment>,

        @InjectRepository(mImage)
        private imageRepository: Repository<mImage>,

        @InjectRepository(note)
        private noteRepository: Repository<note>,

        @InjectRepository(message)
        private messageRepository: Repository<message>,

        @InjectRepository(diary)
        private diaryRepository: Repository<diary>,
    ) { }

    async getMoods(index: number): Promise<moodRes> {
        const res: moodRes = {
            res: 0,
            moods: []
        }
        try {
            const moods = await this.moodRepository.find({ order: { id: "DESC" }, skip: index, take: 6 })
            res.moods = moods
            res.res = 1
        } catch (e) {

        }
        return res
    }

    async getComment(id: number, index: number): Promise<commentRes> {
        const res: commentRes = {
            res: 0,
            comments: []
        }
        try {
            const comments: pcomment[] = await this.commentRepository.find({ where: { id: id }, order: { cid: "ASC" }, skip: index, take: 6 })
            res.comments = comments
            res.res = 1
        } catch (e) {
            console.log(e)
        }
        return res
    }

    async submit(comment: comment): Promise<submitRes> {
        const res: submitRes = {
            res: 0
        }
        try {
            const cid = await this.commentRepository.count({ where: { id: comment.id } })
            await this.commentRepository.insert({ cid: cid + 1, ...comment })
            await this.moodRepository.update({ id: comment.id }, { comment: cid })
            const id: number = await this.diaryRepository.count({ where: { dtype: 1 } })
            await this.diaryRepository.insert({
                id: id + 1,
                dtype: 1,
                text: comment.name + "留下了评论",
                time: comment.time
            })
            res.res = 1
        } catch (e) {
            console.log(e)
        }
        return res
    }

    async getImage(pic: string): Promise<imageRes> {
        const res: imageRes = {
            res: 0,
            paths: []
        }
        const images = pic.split("#").filter((i) => i && i.trim())
        try {
            let str = "SELECT path FROM mImage where id in("
            images.forEach(id => {
                str += id + ","
            });
            str = str.substring(0, str.length - 1)
            str += ")"
            const rows = await this.imageRepository.query(str)
            rows.forEach(row => {
                res.paths.push(row.path)
            })
            res.res = 1
        } catch (e) {
            console.log(e)
        }
        return res
    }

    async setLike(id: number, like: number): Promise<likeRes> {
        const res: likeRes = {
            res: 0
        }
        try {
            await this.moodRepository.update(id, { mlike: like })
            res.res = 1
        } catch (e) {
            console.log(e)
        }
        return res
    }

    async getNotes(): Promise<noteRes> {
        const res: noteRes = {
            res: 0,
            notes: []
        }
        try {
            const str: string = "SELECT text,time FROM note WHERE TO_DAYS(time) - TO_DAYS(NOW( )) >= 0"
            const rows = await this.noteRepository.query(str)
            rows.forEach(site => {
                res.notes.push(site)
            })
            res.res = 1
        } catch (e) {
            console.log(e)
        }
        return res
    }

    async getMessages(index: number, num: number): Promise<messageRes> {
        const res: messageRes = {
            res: 0,
            messages: []
        }
        try {
            const rows: message[] = await this.messageRepository.find({ order: { id: "DESC" }, take: num, skip: index })
            res.messages = rows
            res.res = 1
        } catch (e) {
            console.log(e)
        }
        return res
    }

    async setMessage(form: messageForm): Promise<submitRes> {
        const res: submitRes = {
            res: 0
        }
        try {
            const id = await this.messageRepository.count()
            await this.messageRepository.insert({ id: id + 1, ...form })
            const cid: number = await this.diaryRepository.count({ where: { dtype: 1 } })
            await this.diaryRepository.insert({
                id: cid + 1,
                dtype: 1,
                text: form.name + "留下了评论",
                time: form.time
            })
            res.res = 1
        } catch (e) {
            console.log(e)
        }
        return res
    }
}
