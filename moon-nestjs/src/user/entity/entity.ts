import { info } from "../../entity/info.entity"
import { signs } from "../../entity/signs.entity"
import { flink } from "../..//entity/flink.entity"

export interface infoRes {
    res: number,
    info: info
}

export interface signRes {
    res: number,
    signs: signs[]
}

export interface aboutRes {
    res: number
    text: string
}

export interface flinkRes {
    res: number
    flinks: flink[]
}