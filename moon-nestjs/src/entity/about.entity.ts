import { Entity, PrimaryColumn, Column } from "typeorm"

@Entity()
export class about {
    @PrimaryColumn()
    id: number

    @Column()
    rtext: string

    @Column()
    ftext: string
}