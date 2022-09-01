import { Injectable } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { Repository } from 'typeorm';
import { info } from "../entity/info.entity"
import { signs } from "../entity/signs.entity"
import { about } from '../entity/about.entity';
import { flink } from '../entity/flink.entity';
import { aboutRes, flinkRes, infoRes, signRes } from "./entity/entity"

@Injectable()
export class UserService {
    constructor(
        @InjectRepository(info)
        private infoRepository: Repository<info>,

        @InjectRepository(signs)
        private signRepository: Repository<signs>,

        @InjectRepository(about)
        private aboutRepository: Repository<about>,

        @InjectRepository(flink)
        private flinkRepository: Repository<flink>
    ) { }

    async getInfo(): Promise<infoRes> {
        const res: infoRes = {
            res: 0,
            info: null
        }
        try {
            const info: info = await this.infoRepository.findOne({ where: { id: 1 } })
            res.res = 1
            res.info = info
        } catch (e) {
            console.log(e)
        }
        return res
    }

    async getSigns(): Promise<signRes> {
        const res: signRes = {
            res: 0,
            signs: []
        }
        try {
            const signs: signs[] = await this.signRepository.find()
            res.res = 1
            res.signs = signs
        } catch (e) {
            console.log(e)
        }
        return res
    }

    async getAbout(): Promise<aboutRes> {
        const res: aboutRes = {
            res: 0,
            text: ""
        }
        try {
            const about: about = await this.aboutRepository.findOne({ where: { id: 1 } })
            res.res = 1
            res.text = about.ftext
        } catch (e) {
            console.log(e)
        }
        return res
    }

    async getLinks(): Promise<flinkRes> {
        const res: flinkRes = {
            res: 0,
            flinks: []
        }
        try {
            const flinks: flink[] = await this.flinkRepository.find()
            res.flinks = flinks
            res.res = 1
        } catch (e) {
            console.log(e)
        }
        return res
    }
}
