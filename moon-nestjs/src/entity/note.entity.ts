import { Entity, PrimaryColumn, Column } from "typeorm"

@Entity()
export class note {
    @PrimaryColumn()
    id: number

    @Column()
    text: string

    @Column({ type: "date" })
    time: string
}