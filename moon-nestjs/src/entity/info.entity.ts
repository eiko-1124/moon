import { Entity, PrimaryColumn, Column } from "typeorm"

@Entity()
export class info {
    @PrimaryColumn()
    id: number

    @Column()
    name: string

    @Column()
    notice: string

    @Column()
    illustate: string

    @Column()
    sentence: string

    @Column()
    qq: string

    @Column()
    github: string
}