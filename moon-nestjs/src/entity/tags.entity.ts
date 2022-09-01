import { Entity, PrimaryColumn, Column } from "typeorm"

@Entity()
export class tags {
    @PrimaryColumn()
    id: number

    @Column()
    tag: string
}