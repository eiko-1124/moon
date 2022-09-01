package adminRouter

import (
	"moon-manage-serve/module/admin"

	"github.com/labstack/echo"

	_ "github.com/go-sql-driver/mysql"

	"database/sql"
)

func UseRouter(router *echo.Group, db *sql.DB) {
	cRouter := router.Group("/admin")
	cRouter.GET("/changeRouter", admin.ChangeRouter)
	cRouter.GET("/getInfo", admin.GetInfo)
	cRouter.GET("/getSigns", admin.GetSigns)
	cRouter.GET("/getLinks", admin.GetLinks)
	cRouter.GET("/getArticleSum", admin.GetArticleSum)
	cRouter.GET("/getMoodSum", admin.GetMoodSum)
	cRouter.GET("/getLinkSum", admin.GetLinkSum)
	cRouter.GET("/getTagSum", admin.GetTagSum)
	cRouter.GET("/getCommentSum", admin.GetCommentSum)
	cRouter.POST("/setLink", admin.UpdateLink)
	cRouter.POST("/setSign", admin.UpdateSign)
	cRouter.POST("/deleteLink", admin.DeleteLink)
	cRouter.POST("/updateInfo", admin.UpdateInfo)
	cRouter.POST("/updateAbout", admin.UpdateAbout)
	cRouter.GET("/getAbout", admin.GetAbout)
}
