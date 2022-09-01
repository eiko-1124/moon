package moodrouter

import (
	"database/sql"
	"moon-manage-serve/module/mood"

	"github.com/labstack/echo"
)

func UseRouter(e *echo.Group, Db *sql.DB) {
	mood.SetDb(Db)
	cRouter := e.Group("/mood")
	cRouter.POST("/moodSubmit", mood.SubmitMood)
	cRouter.GET("/getMoodList", mood.GetMoods)
	cRouter.GET("/getMoodSum", mood.GetMoodSum)
	cRouter.GET("/getLinks", mood.GetLinks)
	cRouter.GET("/getComments", mood.GetComments)
	cRouter.POST("/deleteComment", mood.DeleteComment)
	cRouter.POST("/deleteMood", mood.DeleteMood)
	cRouter.POST("/noteSubmit", mood.SetNote)
	cRouter.GET("/getOldNote", mood.GetOldNote)
	cRouter.GET("/getNewNotes", mood.GetNewNotes)
	cRouter.GET("/getNoteSum", mood.GetNoteSum)
	cRouter.POST("/deleteNote", mood.DeleteNote)
}
