import { Entity, PrimaryColumn, Column } from "typeorm"

@Entity()
export class article {
    @PrimaryColumn()
    id: number

    @Column()
    title: string

    @Column()
    atype: number

    @Column()
    illustrate: string

    @Column()
    cover: string

    @Column()
    fcontent: string

    @Column()
    rcontent: string

    @Column()
    tags: string

    @Column()
    hflag: number

    @Column({ type: "datetime" })
    time: string
}