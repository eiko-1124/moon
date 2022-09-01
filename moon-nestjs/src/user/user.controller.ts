import { Controller, Get } from '@nestjs/common';
import { UserService } from './user.service';
import { aboutRes, flinkRes, infoRes, signRes } from "./entity/entity"

@Controller('user')
export class UserController {
  constructor(private readonly userService: UserService) { }

  @Get("info")
  getInfo(): Promise<infoRes> {
    return this.userService.getInfo()
  }

  @Get("signs")
  getSigns(): Promise<signRes> {
    return this.userService.getSigns()
  }

  @Get("about")
  getAbout(): Promise<aboutRes> {
    return this.userService.getAbout()
  }

  @Get("flinks")
  getLinks(): Promise<flinkRes> {
    return this.userService.getLinks()
  }
}
