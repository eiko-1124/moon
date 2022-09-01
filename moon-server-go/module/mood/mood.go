package mood

import (
	"database/sql"
	"encoding/json"
	"moon-manage-serve/module/mood/entry"
	diarydto "moon-manage-serve/sql/diaryDto"
	moodDto "moon-manage-serve/sql/moodDto"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func SetDb(db *sql.DB) {
	moodDto.SetDb(db)
}

func SubmitMood(ctx echo.Context) error {
	res := new(entry.Res)
	text := ctx.FormValue("text")
	pic := ctx.FormValue("pic")
	id, err := moodDto.QueryMoodId()
	if err != nil {
		res.Res = 0
	} else {
		err = moodDto.InsertMood(id+1, text, pic)
		if err != nil {
			res.Res = 0
		} else {
			res.Res = 1
		}
	}
	diarydto.InsertDiary("你发布了一条记录", 0)
	data, _ := json.Marshal(res)
	return ctx.String(http.StatusOK, string(data))
}

func GetMoods(ctx echo.Context) error {
	res := new(entry.MoodRes)
	num := ctx.QueryParam("num")
	moods, err := moodDto.QueryMoodList(num)
	if err != nil {
		res.Res = 0
		res.Moods = nil
	} else {
		res.Res = 1
		res.Moods = moods
	}
	data, _ := json.Marshal(res)
	return ctx.String(http.StatusOK, string(data))
}

func GetMoodSum(ctx echo.Context) error {
	res := new(entry.MoodSumRes)
	sum, err := moodDto.QueryMoodSum()
	if err != nil {
		res.Res = 0
		res.Sum = 0
	} else {
		res.Res = 1
		res.Sum = sum
	}
	data, _ := json.Marshal(res)
	return ctx.String(http.StatusOK, string(data))
}

func GetLinks(ctx echo.Context) error {
	res := new(entry.LinkRes)
	picStr := ctx.QueryParam("pic")
	links, err := moodDto.QueryLink(picStr)
	if err != nil {
		res.Res = 0
		res.Links = nil
	} else {
		res.Res = 1
		res.Links = links
	}
	data, _ := json.Marshal(res)
	return ctx.String(http.StatusOK, string(data))
}

func GetComments(ctx echo.Context) error {
	res := new(entry.CommentRes)
	id := ctx.QueryParam("id")
	page := ctx.QueryParam("page")
	cPage, _ := strconv.Atoi(page)
	comments, err := moodDto.QueryComments(id, cPage, 10)
	if err != nil {
		res.Res = 0
		res.Comments = nil
	} else {
		res.Res = 1
		res.Comments = comments
	}
	data, _ := json.Marshal(res)
	return ctx.String(http.StatusOK, string(data))
}

func DeleteComment(ctx echo.Context) error {
	res := new(entry.Res)
	id := ctx.FormValue("id")
	cid := ctx.FormValue("cid")
	err := moodDto.DeleteComment(id, cid)
	if err != nil {
		res.Res = 0
	} else {
		res.Res = 1
	}
	diarydto.InsertDiary("你删除了记录"+id+"的评论"+cid, 0)
	data, _ := json.Marshal(res)
	return ctx.String(http.StatusOK, string(data))
}

func DeleteMood(ctx echo.Context) error {
	res := new(entry.Res)
	id := ctx.FormValue("id")
	err := moodDto.DeleteMood(id)
	if err != nil {
		res.Res = 0
	} else {
		res.Res = 1
	}
	diarydto.InsertDiary("你删除了记录"+id, 0)
	data, _ := json.Marshal(res)
	return ctx.String(http.StatusOK, string(data))
}

func SetNote(ctx echo.Context) error {
	res := new(entry.Res)
	date := ctx.FormValue("date")
	text := ctx.FormValue("text")
	id, err := moodDto.QueryMaxNote()
	if err != nil {
		res.Res = 0
	}
	err = moodDto.InsertNote(id+1, date, text)
	if err != nil {
		res.Res = 0
	} else {
		res.Res = 1
	}
	diarydto.InsertDiary("你添加了一项安排记录", 0)
	data, _ := json.Marshal(res)
	return ctx.String(http.StatusOK, string(data))
}

func GetOldNote(ctx echo.Context) error {
	res := new(entry.OldNoteRes)
	note, err := moodDto.QueryOldNote()
	if err != nil {
		res.Res = 0
		res.Note = nil
	} else {
		res.Res = 1
		res.Note = note
	}
	data, _ := json.Marshal(res)
	return ctx.String(http.StatusOK, string(data))
}

func GetNewNotes(ctx echo.Context) error {
	res := new(entry.NewNoteRes)
	page := ctx.QueryParam("page")
	notes, err := moodDto.QueryNewNote(page)
	if err != nil {
		res.Res = 0
		res.Notes = nil
	} else {
		res.Res = 1
		res.Notes = notes
	}
	data, _ := json.Marshal(res)
	return ctx.String(http.StatusOK, string(data))
}

func GetNoteSum(ctx echo.Context) error {
	res := new(entry.NoteSumRes)
	sum, err := moodDto.QueryNoteSum()
	if err != nil {
		res.Res = 0
		res.Sum = 0
	} else {
		res.Res = 1
		res.Sum = sum
	}
	data, _ := json.Marshal(res)
	return ctx.String(http.StatusOK, string(data))
}

func DeleteNote(ctx echo.Context) error {
	res := new(entry.Res)
	id := ctx.FormValue("id")
	err := moodDto.DeleteNote(id)
	if err != nil {
		res.Res = 0
	} else {
		res.Res = 1
	}
	diarydto.InsertDiary("你移除了一项安排", 0)
	data, _ := json.Marshal(res)
	return ctx.String(http.StatusOK, string(data))
}
