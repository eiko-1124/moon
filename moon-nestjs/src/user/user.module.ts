import { Module } from '@nestjs/common';
import { TypeOrmModule } from "@nestjs/typeorm"
import { UserService } from './user.service';
import { UserController } from './user.controller';
import { info } from "../entity/info.entity"
import { signs } from "../entity/signs.entity"
import { about } from "../entity/about.entity"
import { flink } from '../entity/flink.entity';

@Module({
  imports: [TypeOrmModule.forFeature([info, signs, about, flink])],
  controllers: [UserController],
  providers: [UserService]
})
export class UserModule { }
