import { Entity, PrimaryColumn, Column } from "typeorm"

@Entity()
export class mood {
    @PrimaryColumn()
    id: number

    @Column()
    text: string

    @Column()
    pic: string

    @Column({ type: "datetime" })
    time: string

    @Column()
    mlike: number

    @Column()
    comment: number
}