package filerouter

import (
	"database/sql"
	"moon-manage-serve/module/file"

	"github.com/labstack/echo"
)

func UseRouter(e *echo.Group, Db *sql.DB) {
	file.SetDb(Db)
	cRouter := e.Group("/file")
	cRouter.POST("/minorFile", file.UploadMinorFile)
	cRouter.POST("/filePiece", file.UploadFilePiece)
	cRouter.GET("/getAllFiles", file.GetAllFiles)
	cRouter.POST("/moodPicture", file.UpdateMoodPicture)
	cRouter.GET("/getArticleCoverSum", file.GetArticleCoverSum)
	cRouter.GET("/getArticleCovers", file.GetArticleCovers)
	cRouter.GET("/getMoodImageSum", file.GetMoodImageSum)
	cRouter.GET("/getMoodImages", file.GetMoodImages)
	cRouter.GET("/getOtherFileSum", file.GetOtherFileSum)
	cRouter.GET("/getOtherFiles", file.GetOtherFiles)
}
