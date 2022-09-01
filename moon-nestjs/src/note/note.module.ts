import { Module } from '@nestjs/common';
import { NoteService } from './note.service';
import { NoteController } from './note.controller';
import { TypeOrmModule } from '@nestjs/typeorm';
import { mood } from '../entity/mood.entity';
import { pcomment } from '../entity/pcomment.entity';
import { mImage } from '../entity/mImage.entity';
import { note } from '../entity/note.entity';
import { message } from '../entity/message.entity';
import { diary } from '../entity/diary.entity';

@Module({
  imports: [TypeOrmModule.forFeature([mood, pcomment, mImage, note, message, diary])],
  controllers: [NoteController],
  providers: [NoteService]
})
export class NoteModule { }
