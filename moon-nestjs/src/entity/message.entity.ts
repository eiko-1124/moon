import { Entity, PrimaryColumn, Column } from "typeorm"

@Entity()
export class message {
    @PrimaryColumn()
    id: number

    @Column()
    mtype: number

    @Column()
    text: string

    @Column({ type: "datetime" })
    time: string

    @Column()
    name: string
}