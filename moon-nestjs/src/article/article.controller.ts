import { Body, Controller, Get, Post, Query } from '@nestjs/common';
import { ArticleService } from './article.service';
import { articleListRes, articleRes, commentForm, commentRes, cTagRes, nearRes, recommendRes, submitRes, sumRes, tagsRes, titleRes, typeRes } from './entity/entity';

@Controller('article')
export class ArticleController {
  constructor(private readonly articleService: ArticleService) { }

  @Get('articleSum')
  getArticleSum(): Promise<sumRes> {
    return this.articleService.getArticleSum()
  }

  @Get('typeSum')
  getTypeSum(): Promise<sumRes> {
    return this.articleService.getTpyeSum()
  }

  @Get('tagSum')
  getTagSum(): Promise<sumRes> {
    return this.articleService.getTagSum()
  }

  @Get('tags')
  getTags(): Promise<tagsRes> {
    return this.articleService.getTags()
  }

  @Get('articleList')
  getArticleList(@Query('page') page): Promise<articleListRes> {
    return this.articleService.getArticleList(page)
  }

  @Get('recommend')
  getRecommend(): Promise<recommendRes> {
    return this.articleService.getRecommend()
  }

  @Get('title')
  getArticleTitle(@Query('index') index: number): Promise<titleRes> {
    return this.articleService.getTitle(index)
  }

  @Get('classifySum')
  getClassifySum(@Query('type') type: number, @Query('tag') tag: string): Promise<sumRes> {
    return this.articleService.getClassifySum(type, tag)
  }

  @Get('classifyTags')
  getClassifyTags(@Query('type') type: number): Promise<cTagRes> {
    return this.articleService.getClassifyTags(type)
  }

  @Get('type')
  getType(): Promise<typeRes> {
    return this.articleService.getType()
  }

  @Get('classifyArticle')
  getClassifyArticle(@Query('type') type: number, @Query('index') index: number, @Query('tag') tag: string): Promise<articleListRes> {
    return this.articleService.getClassifyArticle(type, tag, index)
  }

  @Get('article')
  getArticle(@Query('id') id: number): Promise<articleRes> {
    return this.articleService.getArticle(id)
  }

  @Get('near')
  getNear(@Query('id') id: number): Promise<nearRes> {
    return this.articleService.getNear(id)
  }

  @Get('comments')
  getCommnets(@Query('title') title: string, @Query('index') index: number): Promise<commentRes> {
    return this.articleService.getComments(title, index)
  }

  @Post('submit')
  setComment(@Body() comment: commentForm): Promise<submitRes> {
    return this.articleService.setComment(comment)
  }
}
