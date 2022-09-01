import { Entity, PrimaryColumn, Column } from "typeorm"

@Entity()
export class mImage {
    @PrimaryColumn()
    id: number

    @Column({ type: "datetime" })
    time: string

    @Column()
    path: string
}