import { Entity, PrimaryColumn, Column } from "typeorm"

@Entity()
export class flink {
    @PrimaryColumn()
    id: number

    @Column()
    name: string

    @Column()
    flink: string
}