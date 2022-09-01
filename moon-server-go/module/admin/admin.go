package admin

import (
	"encoding/json"
	"fmt"
	admindto "moon-manage-serve/sql/adminDto"
	diarydto "moon-manage-serve/sql/diaryDto"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func ChangeRouter(ctx echo.Context) error {
	return ctx.NoContent(http.StatusOK)
}

func GetInfo(ctx echo.Context) error {
	info, err := admindto.QueryInfo()
	if err != nil {
		fmt.Println(err.Error())
	}
	data, err := json.Marshal(info)
	if err != nil {
		fmt.Println(err.Error())
	}
	return ctx.String(http.StatusOK, string(data))
}

func GetSigns(ctx echo.Context) error {
	signs, err := admindto.QuerySigns()
	if err != nil {
		fmt.Println(err.Error())
	}
	data, err := json.Marshal(signs)
	if err != nil {
		fmt.Println(err.Error())
	}
	return ctx.String(http.StatusOK, string(data))
}

func GetLinks(ctx echo.Context) error {
	links, err := admindto.QueryLink()
	if err != nil {
		fmt.Println(err.Error())
	}
	data, err := json.Marshal(links)
	if err != nil {
		fmt.Println(err.Error())
	}
	return ctx.String(http.StatusOK, string(data))
}

func GetArticleSum(ctx echo.Context) error {
	sum, err := admindto.QueryArticleSum()
	if err != nil {
		fmt.Println(err.Error())
	}
	data, err := json.Marshal(sum)
	if err != nil {
		fmt.Println(err.Error())
	}
	return ctx.String(http.StatusOK, string(data))
}
func GetMoodSum(ctx echo.Context) error {
	sum, err := admindto.QueryMoodSum()
	if err != nil {
		fmt.Println(err.Error())
	}
	data, err := json.Marshal(sum)
	if err != nil {
		fmt.Println(err.Error())
	}
	return ctx.String(http.StatusOK, string(data))
}
func GetLinkSum(ctx echo.Context) error {
	sum, err := admindto.QueryLinkSum()
	if err != nil {
		fmt.Println(err.Error())
	}
	data, err := json.Marshal(sum)
	if err != nil {
		fmt.Println(err.Error())
	}
	return ctx.String(http.StatusOK, string(data))
}
func GetTagSum(ctx echo.Context) error {
	sum, err := admindto.QueryTagSum()
	if err != nil {
		fmt.Println(err.Error())
	}
	data, err := json.Marshal(sum)
	if err != nil {
		fmt.Println(err.Error())
	}
	return ctx.String(http.StatusOK, string(data))
}
func GetCommentSum(ctx echo.Context) error {
	sum, err := admindto.QueryCommentSum()
	if err != nil {
		fmt.Println(err.Error())
	}
	data, err := json.Marshal(sum)
	if err != nil {
		fmt.Println(err.Error())
	}
	return ctx.String(http.StatusOK, string(data))
}
func UpdateLink(ctx echo.Context) error {
	id := ctx.FormValue("id")
	uid, _ := strconv.Atoi(id)
	name := ctx.FormValue("name")
	link := ctx.FormValue("link")
	err := admindto.UpdateLink(uid, name, link)
	if err != nil {
		fmt.Println(err.Error())
	}
	res := 0
	data, err := json.Marshal(res)
	if err != nil {
		fmt.Println(err.Error())
	}
	diarydto.InsertDiary("你修改了友联“"+name+"”", 0)
	return ctx.String(http.StatusOK, string(data))
}
func UpdateSign(ctx echo.Context) error {
	id := ctx.FormValue("id")
	uid, _ := strconv.Atoi(id)
	text := ctx.FormValue("content")
	err := admindto.UpdateSign(uid, text)
	if err != nil {
		fmt.Println(err.Error())
	}
	res := 0
	data, err := json.Marshal(res)
	if err != nil {
		fmt.Println(err.Error())
	}
	diarydto.InsertDiary("你修改了签名“"+text+"”", 0)
	return ctx.String(http.StatusOK, string(data))
}
func DeleteLink(ctx echo.Context) error {
	id := ctx.FormValue("id")
	err := admindto.DeleteSign(id)
	if err != nil {
		fmt.Println(err.Error())
	}
	res := 0
	data, err := json.Marshal(res)
	if err != nil {
		fmt.Println(err.Error())
	}
	diarydto.InsertDiary("你删除了第"+id+"条友联", 0)
	return ctx.String(http.StatusOK, string(data))
}
func UpdateInfo(ctx echo.Context) error {
	key := ctx.FormValue("key")
	value := ctx.FormValue("value")
	err := admindto.UpdateInfo(key, value)
	if err != nil {
		fmt.Println(err.Error())
	}
	res := 0
	data, err := json.Marshal(res)
	if err != nil {
		fmt.Println(err.Error())
	}
	diarydto.InsertDiary("你更新了个人信息", 0)
	return ctx.String(http.StatusOK, string(data))
}
func UpdateAbout(ctx echo.Context) error {
	res := 0
	rander := ctx.FormValue("rander")
	format := ctx.FormValue("format")
	err := admindto.UpdateAbout(rander, format)
	if err != nil {
		res = 1
		fmt.Println(err.Error())
	}
	data, err := json.Marshal(res)
	if err != nil {
		fmt.Println(err.Error())
	}
	diarydto.InsertDiary("你更新了站点介绍", 0)
	return ctx.String(http.StatusOK, string(data))
}
func GetAbout(ctx echo.Context) error {
	about, err := admindto.QueryAbout()
	if err != nil {
		fmt.Println(err.Error())
	}
	data, err := json.Marshal(about)
	if err != nil {
		fmt.Println(err.Error())
	}
	return ctx.String(http.StatusOK, string(data))
}
