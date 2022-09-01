import { Module } from '@nestjs/common';
import { TypeOrmModule } from "@nestjs/typeorm"
import { ArticleService } from './article.service';
import { ArticleController } from './article.controller';
import { article } from "../entity/article.entity"
import { tags } from '../entity/tags.entity';
import { comment } from '../entity/comment.entity'
import { diary } from '../entity/diary.entity';

@Module({
  imports: [TypeOrmModule.forFeature([article, tags, comment, diary])],
  controllers: [ArticleController],
  providers: [ArticleService]
})
export class ArticleModule { }
