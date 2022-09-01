package diaryrouter

import (
	"database/sql"
	"moon-manage-serve/module/diary"

	"github.com/labstack/echo"
)

func UseRouter(router *echo.Group, db *sql.DB) {
	diary.SetDb(db)
	cRouter := router.Group("/diary")
	cRouter.GET("/getDiaryTime", diary.GetDiaryTime)
	cRouter.GET("/getDiaryAll", diary.GetDiaryAll)
	cRouter.GET("/getDiarySum", diary.GetDiarySum)
	cRouter.GET("/getMessageSum", diary.GetMessageSum)
	cRouter.GET("/getMessage", diary.GetMessage)
	cRouter.POST("/deleteMessage", diary.DeleteMessage)
	cRouter.GET("/getDiaryNear", diary.GetNearDiary)
	cRouter.GET("/getDiary", diary.GetDiary)
}
