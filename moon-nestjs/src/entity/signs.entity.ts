import { Entity, PrimaryColumn, Column } from "typeorm"

@Entity()
export class signs {
    @PrimaryColumn()
    id: number

    @Column()
    content: string
}