import { Module } from '@nestjs/common';
import { TypeOrmModule } from '@nestjs/typeorm';
import { UserModule } from "./user/user.module"
import { ArticleModule } from './article/article.module';
import { NoteModule } from './note/note.module';

@Module({
  imports: [
    UserModule,
    ArticleModule,
    NoteModule,
    TypeOrmModule.forRoot({
      type: 'mysql',
      host: '120.79.73.206',
      port: 3306,
      username: "dmeiko",
      password: "#Eiko1124",
      database: "dmeiko_DB",
      entities: ["dist/entity/*.entity{.ts,.js}"],
      synchronize: false
    })
  ],
})
export class AppModule { }
