import { Entity, PrimaryColumn, Column } from "typeorm"

@Entity()
export class comment {
    @PrimaryColumn()
    id: number

    @PrimaryColumn()
    title: string

    @Column()
    content: string

    @Column()
    name: string

    @Column()
    cname: string

    @Column()
    cflag: number

    @Column()
    hflag: number

    @Column()
    email: string

    @Column()
    cid: number

    @Column({ type: "datetime" })
    time: string
}