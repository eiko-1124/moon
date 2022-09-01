import { Entity, PrimaryColumn, Column } from "typeorm"

@Entity()
export class diary {
    @PrimaryColumn()
    id: number

    @PrimaryColumn()
    dtype: number

    @Column()
    text: string

    @Column({ type: "datetime" })
    time: string
}