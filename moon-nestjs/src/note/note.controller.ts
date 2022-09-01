import { Body, Controller, Get, Post, Query } from '@nestjs/common';
import { comment, commentRes, imageRes, likeRes, moodRes, submitRes, noteRes, messageRes, messageForm } from './entity/entity';
import { NoteService } from './note.service';

@Controller('note')
export class NoteController {
  constructor(private readonly noteService: NoteService) { }

  @Get('moods')
  getMoods(@Query('index') index): Promise<moodRes> {
    return this.noteService.getMoods(index)
  }

  @Get('comment')
  getComment(@Query('index') index: number, @Query('id') id: number): Promise<commentRes> {
    return this.noteService.getComment(id, index)
  }

  @Post('submit')
  submitCommnet(@Body() comment: comment): Promise<submitRes> {
    return this.noteService.submit(comment)
  }

  @Get('image')
  getImage(@Query('pic') pic: string): Promise<imageRes> {
    return this.noteService.getImage(pic)
  }

  @Get('like')
  setLike(@Query('id') id: number, @Query('like') like: number): Promise<likeRes> {
    return this.noteService.setLike(id, like)
  }

  @Get('notes')
  getNotes(): Promise<noteRes> {
    return this.noteService.getNotes()
  }

  @Get('messages')
  getMessages(@Query('index') index: number, @Query('num') num: number): Promise<messageRes> {
    return this.noteService.getMessages(index, num)
  }

  @Post('setMessage')
  setMessage(@Body() form: messageForm): Promise<submitRes> {
    return this.noteService.setMessage(form)
  }
}
