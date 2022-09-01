package diary

import (
	"database/sql"
	"encoding/json"
	"moon-manage-serve/module/diary/entry"
	diarydto "moon-manage-serve/sql/diaryDto"
	"net/http"

	"github.com/labstack/echo"
)

func SetDb(db *sql.DB) {
	diarydto.SetDb(db)
}

func GetDiaryTime(ctx echo.Context) error {
	res := new(entry.DiaryTimeRes)
	dType := ctx.QueryParam("type")
	page := ctx.QueryParam("page")
	diarys, err := diarydto.QueryDiaryTime(dType, page)
	if err != nil {
		res.Res = 0
		res.Diarys = nil
	} else {
		res.Res = 1
		res.Diarys = diarys
	}
	data, _ := json.Marshal(res)
	return ctx.String(http.StatusOK, string(data))
}

func GetDiaryAll(ctx echo.Context) error {
	res := new(entry.DiaryAllRes)
	dType := ctx.QueryParam("type")
	page := ctx.QueryParam("page")
	date := ctx.QueryParam("date")
	diarys, err := diarydto.QueryDiaryAll(dType, date, page)
	if err != nil {
		res.Res = 0
		res.Diarys = nil
	} else {
		res.Res = 1
		res.Diarys = diarys
	}
	data, _ := json.Marshal(res)
	return ctx.String(http.StatusOK, string(data))
}

func GetDiarySum(ctx echo.Context) error {
	res := new(entry.SumRes)
	dType := ctx.QueryParam("type")
	sum, err := diarydto.QueryDiarySum(dType)
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

func GetMessage(ctx echo.Context) error {
	res := new(entry.MessageRes)
	page := ctx.QueryParam("page")
	messages, err := diarydto.QueryMessage(page)
	if err != nil {
		res.Res = 0
		res.Messages = nil
	} else {
		res.Res = 1
		res.Messages = messages
	}
	data, _ := json.Marshal(res)
	return ctx.String(http.StatusOK, string(data))
}

func GetMessageSum(ctx echo.Context) error {
	res := new(entry.SumRes)
	sum, err := diarydto.QueryMessageSum()
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

func DeleteMessage(ctx echo.Context) error {
	res := new(entry.Res)
	id := ctx.FormValue("id")
	err := diarydto.DeleteMessage(id)
	if err != nil {
		res.Res = 0
	} else {
		res.Res = 1
	}
	data, _ := json.Marshal(res)
	diarydto.InsertDiary("你删除了一条留言", 0)
	return ctx.String(http.StatusOK, string(data))
}

func GetNearDiary(ctx echo.Context) error {
	res := new(entry.NearRes)
	dType := ctx.QueryParam("type")
	near, err := diarydto.QueryNearDiary(dType)
	if err != nil {
		res.Res = 0
		res.Near = nil
	} else {
		res.Res = 1
		res.Near = near
	}
	data, _ := json.Marshal(res)
	return ctx.String(http.StatusOK, string(data))
}

func GetDiary(ctx echo.Context) error {
	res := new(entry.DiaryRes)
	dType := ctx.QueryParam("type")
	diary, err := diarydto.QueryDiary(dType)
	if err != nil {
		res.Res = 0
		res.Diary = nil
	} else {
		res.Res = 1
		res.Diary = diary
	}
	data, _ := json.Marshal(res)
	return ctx.String(http.StatusOK, string(data))
}
