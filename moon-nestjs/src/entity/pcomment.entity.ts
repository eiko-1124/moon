import { Entity, PrimaryColumn, Column } from "typeorm"

@Entity()
export class pcomment {
    @PrimaryColumn()
    id: number

    @PrimaryColumn()
    cid: number

    @Column()
    content: string

    @Column()
    name: string

    @Column()
    cname: string

    @Column()
    cflag: number

    @Column({ type: "datetime" })
    time: string
}