package articlerouter

import (
	"database/sql"
	"moon-manage-serve/module/article"

	"github.com/labstack/echo"
)

func UseRouter(router *echo.Group, db *sql.DB) {
	article.SetDb(db)
	cRouter := router.Group("/article")
	cRouter.POST("/submit", article.SubmitArticle)
	cRouter.GET("/getArticle", article.GetArticle)
	cRouter.GET("/getTags", article.GetTags)
	cRouter.GET("/getTagsDetails", article.GetTagsDetails)
	cRouter.GET("/getArticleSum", article.GetArticleSum)
	cRouter.GET("/getArticleList", article.GetArticleList)
	cRouter.POST("/setArticleStatus", article.SetArticleStatus)
	cRouter.POST("/deleteArticle", article.DeleteArticle)
	cRouter.POST("/updateMood", article.UpdateMood)
	cRouter.GET("/getCommentSum", article.GetCommentSum)
	cRouter.GET("/getCommentList", article.GetCommentList)
	cRouter.POST("/setCommentStatus", article.SetCommentStatus)
	cRouter.POST("/deleteComment", article.DeleteComment)
	cRouter.GET("/getMaxTags", article.GetMaxTags)
}
