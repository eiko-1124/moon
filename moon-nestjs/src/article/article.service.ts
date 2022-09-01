import { Injectable } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { LessThan, MoreThan, Repository } from 'typeorm';
import { article } from '../entity/article.entity';
import { tags } from '../entity/tags.entity';
import { comment } from '../entity/comment.entity';
import { diary } from '../entity/diary.entity';
import { articleListRes, articleRes, articleSite, commentForm, commentRes, cTagRes, nearRes, recommend, recommendRes, submitRes, sumRes, tagsRes, titleRes, typeRes } from './entity/entity'

@Injectable()
export class ArticleService {
    constructor(
        @InjectRepository(article)
        private articleRepository: Repository<article>,

        @InjectRepository(tags)
        private tagsRepository: Repository<tags>,

        @InjectRepository(comment)
        private commentRepository: Repository<comment>,

        @InjectRepository(diary)
        private diaryRepository: Repository<diary>,
    ) { }

    async getArticleSum(): Promise<sumRes> {
        const res: sumRes = {
            res: 0,
            sum: 0
        }
        try {
            const count = await this.articleRepository.count({ where: { hflag: 0 } })
            res.res = 1
            res.sum = count
        } catch (e) {
            console.log(e)
        }
        return res
    }

    async getTpyeSum(): Promise<sumRes> {
        const res: sumRes = {
            res: 0,
            sum: 0
        }
        try {
            const count = await this.articleRepository.createQueryBuilder().select('COUNT(DISTINCT article.atype)', "sum").getRawOne()
            res.res = 1
            res.sum = count.sum
        } catch (e) {
            console.log(e)
        }
        return res
    }

    async getTagSum(): Promise<sumRes> {
        const res: sumRes = {
            res: 0,
            sum: 0
        }
        try {
            const count = await this.tagsRepository.createQueryBuilder().select('COUNT(DISTINCT tags.tag)', "sum").getRawOne()
            res.res = 1
            res.sum = count.sum
        } catch (e) {
            console.log(e)
        }
        return res
    }

    async getTags(): Promise<tagsRes> {
        const res: tagsRes = {
            res: 0,
            tags: []
        }
        try {
            const tags = await this.tagsRepository.createQueryBuilder().select('DISTINCT tags.tag').getRawMany()
            let tagArr: string[] = []
            for (let tag of tags) {
                tagArr.push(tag.tag)
            }
            res.res = 1
            res.tags = tagArr
        } catch (e) {
            console.log(e)
        }
        return res
    }

    async getArticleList(page: number): Promise<articleListRes> {
        const res: articleListRes = {
            res: 0,
            list: []
        }
        try {
            const list: articleSite[] = await this.articleRepository.createQueryBuilder()
                .leftJoinAndSelect(comment, "comment", "article.title=comment.title")
                .select(`article.id as id,article.title as title,article.atype as atype,article.illustrate as illustrate,article.cover as cover,article.time as time,article.tags as tags,count(comment.id) as comment`)
                .where("article.hflag=0")
                .groupBy("article.id")
                .limit(6)
                .offset(page * 6)
                .orderBy("article.id", "DESC")
                .getRawMany()
            res.list = list
            res.res = 1
        } catch (e) {
            console.log(e)
        }
        return res
    }

    async getRecommend(): Promise<recommendRes> {
        const res: recommendRes = {
            res: 0,
            list: []
        }
        try {
            const list: recommend[] = await this.articleRepository.find({ select: { id: true, title: true, cover: true, time: true }, where: { hflag: 0 }, order: { id: "DESC" }, skip: 0, take: 5 })
            res.list = list
            res.res = 1
        } catch (e) {
            console.log(e)
        }
        return res
    }

    async getTitle(index: number): Promise<titleRes> {
        const res: titleRes = {
            res: 0,
            titles: []
        }
        try {
            const titles = await this.articleRepository.find({ select: { id: true, title: true, time: true }, order: { id: "DESC" }, skip: index, take: 10 })
            res.titles = titles
            res.res = 1
        } catch (e) {
            console.log(e)
        }
        return res
    }

    async getType(): Promise<typeRes> {
        const res: typeRes = {
            res: 0,
            types: []
        }
        try {
            const types = await this.articleRepository.createQueryBuilder().select('DISTINCT article.atype').where('article.hflag=0').getRawMany()
            types.forEach(site => {
                res.types.push(site.atype)
            })
            res.res = 1
        } catch (e) {
            console.log(e)
        }
        return res
    }

    async getClassifySum(type: number, tag: string): Promise<sumRes> {
        const res: sumRes = {
            res: 0,
            sum: 0
        }
        try {
            let row
            if (type != 0) {
                row = await this.articleRepository.createQueryBuilder()
                    .leftJoinAndSelect(tags, "tags", "tags.id=article.id")
                    .select(`count(article.id) as count`)
                    .where(`article.hflag=0 and article.atype=${type} and tags.tag="${tag}"`)
                    .getRawOne()
            }
            else {
                row = await this.articleRepository.createQueryBuilder()
                    .leftJoinAndSelect(tags, "tags", "tags.id=article.id")
                    .select(`count(article.id) as count`)
                    .where(`article.hflag=0 and tags.tag="${tag}"`)
                    .getRawOne()
            }
            res.sum = row.count
            res.res = 1
        } catch (e) {
            console.log(e)
        }
        return res
    }

    async getClassifyTags(type: number): Promise<cTagRes> {
        const res: cTagRes = {
            res: 0,
            tags: []
        }
        try {
            let rows: article[]
            if (type != 0) {
                rows = await this.articleRepository.find({ select: { tags: true }, where: { atype: type, hflag: 0 } })
            } else {
                rows = await this.articleRepository.find({ select: { tags: true }, where: { hflag: 0 } })
            }
            let tags: string[] = []
            for (let row of rows) {
                let tagArr = row.tags.split("#").filter((i) => i && i.trim())
                tagArr.forEach(site => {
                    if (tags.indexOf(site) == -1) tags.push(site)
                })
            }
            res.tags = tags
            res.res = 1
        } catch (e) {
            console.log(e)
        }
        return res
    }

    async getClassifyArticle(type: number, tag: string, index: number): Promise<articleListRes> {
        const res: articleListRes = {
            res: 0,
            list: []
        }
        if (index < 0) return res
        try {
            let rows: articleSite[]
            if (type != 0) {
                rows = await this.articleRepository.createQueryBuilder()
                    .leftJoinAndSelect(tags, "tags", "tags.id=article.id")
                    .leftJoinAndSelect(comment, "comment", "article.title=comment.title")
                    .select(`article.id as id,article.title as title,article.atype as atype,article.illustrate as illustrate,article.cover as cover,article.time as time,article.tags as tags,count(comment.id) as comment`)
                    .where(`article.hflag=0 and article.atype=${type} and tags.tag="${tag}"`)
                    .groupBy("article.id")
                    .limit(6)
                    .offset(index * 6)
                    .orderBy("article.id", "DESC")
                    .getRawMany()
            } else {
                rows = await this.articleRepository.createQueryBuilder()
                    .leftJoinAndSelect(tags, "tags", "tags.id=article.id")
                    .leftJoinAndSelect(comment, "comment", "article.title=comment.title")
                    .select(`article.id as id,article.title as title,article.atype as atype,article.illustrate as illustrate,article.cover as cover,article.time as time,article.tags as tags,count(comment.id) as comment`)
                    .where(`article.hflag=0 and tags.tag="${tag}"`)
                    .groupBy("article.id")
                    .limit(6)
                    .offset(index * 6)
                    .orderBy("article.id", "DESC")
                    .getRawMany()
            }
            res.list = rows
            res.res = 1
        } catch (e) {
            console.log(e)
        }
        return res
    }

    async getArticle(id: number): Promise<articleRes> {
        const res: articleRes = {
            res: 0,
            article: null,
            comment: 0
        }
        try {
            const article: article = await this.articleRepository.findOne({ where: { id } })
            const sum: number = await this.commentRepository.count({ where: { title: article.title } })
            res.article = article
            res.comment = sum
            res.res = 1
        } catch (e) {
            console.log(e)
        }
        return res
    }

    async getNear(id: number): Promise<nearRes> {
        const res: nearRes = {
            res: 0,
            near: {
                nextId: 0,
                preId: 0,
                nextTitle: '',
                preTitle: ''
            }
        }
        try {
            const next = await this.articleRepository.findOne({ select: { id: true, title: true }, where: { id: MoreThan(id), hflag: 0 } })
            const pre = await this.articleRepository.findOne({ select: { id: true, title: true }, where: { id: LessThan(id), hflag: 0 } })
            if (next != null) {
                res.near.nextId = next.id
                res.near.nextTitle = next.title
            }
            if (pre != null) {
                res.near.preId = pre.id
                res.near.preTitle = pre.title
            }
            res.res = 1
        } catch (e) {
            console.log(e)
        }
        return res
    }

    async getComments(title: string, index: number): Promise<commentRes> {
        const res: commentRes = {
            res: 0,
            comments: []
        }
        try {
            const comments = await this.commentRepository.find({ where: { title, hflag: 1 }, take: 10, skip: index })
            res.comments = comments
            res.res = 1
        } catch (e) {
            console.log(e)
        }
        return res
    }

    async setComment(comment: commentForm): Promise<submitRes> {
        const res: submitRes = {
            res: 0
        }
        try {
            await this.commentRepository.insert({ ...comment })
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
}
