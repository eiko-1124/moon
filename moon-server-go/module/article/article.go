package article

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"moon-manage-serve/module/article/entry"
	articleDto "moon-manage-serve/sql/articleDto"
	sqlEntry "moon-manage-serve/sql/articleDto/entry"
	diarydto "moon-manage-serve/sql/diaryDto"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
)

func SetDb(db *sql.DB) {
	articleDto.SetDb(db)
}

func SubmitArticle(ctx echo.Context) error {
	res := new(entry.Res)
	articleForm := new(entry.ArticleForm)
	if err := ctx.Bind(articleForm); err != nil {
		res.Res = 0
		data, _ := json.Marshal(res)
		return ctx.String(http.StatusOK, string(data))
	}
	sqlArticle := new(sqlEntry.SqlArticle)
	sqlArticle.Title = articleForm.Title
	sqlArticle.Atpye = articleForm.Atype
	sqlArticle.Cover = articleForm.Cover
	sqlArticle.Tags = articleForm.Tags
	sqlArticle.Illustrate = articleForm.Illustrate
	sqlArticle.Fcontent = articleForm.Fcontent
	sqlArticle.Rcontent = articleForm.Rcontent
	if articleForm.Id == 0 {
		var err error
		sqlArticle.Atime = time.Now()
		sqlArticle.Id, err = articleDto.QueryMaxId()
		if err != nil {
			res.Res = 0
		} else {
			err = articleDto.InsetArticle(*sqlArticle)
			if err != nil {
				res.Res = 0
			} else {
				res.Res = 1
			}
		}
		data, _ := json.Marshal(res)
		return ctx.String(http.StatusOK, string(data))
	} else {
		sqlArticle.Id = articleForm.Id
		err := articleDto.UpdateArticle(*sqlArticle)
		if err != nil {
			fmt.Println(err)
			res.Res = 0
		} else {
			res.Res = 1
		}
		diarydto.InsertDiary("你添加了文章《"+articleForm.Title+"》", 0)
		data, _ := json.Marshal(res)
		return ctx.String(http.StatusOK, string(data))
	}
}

func GetTags(ctx echo.Context) error {
	res := new(entry.TagsRes)
	tags, err := articleDto.QueryTags()
	if err != nil {
		res.Res = 0
		res.Tags = nil
	} else {
		res.Res = 1
		res.Tags = tags
	}
	data, _ := json.Marshal(res)
	return ctx.String(http.StatusOK, string(data))
}

func GetArticleSum(ctx echo.Context) error {
	aType := ctx.QueryParam("type")
	date := ctx.QueryParam("time")
	tag := ctx.QueryParam("tag")
	res := new(entry.SumRes)
	sum, err := articleDto.QueryArticleSum(date, aType, tag)
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

func GetArticleList(ctx echo.Context) error {
	res := new(entry.ListRes)
	aType := ctx.QueryParam("type")
	date := ctx.QueryParam("time")
	tag := ctx.QueryParam("tag")
	page := ctx.QueryParam("page")
	cPage, _ := strconv.Atoi(page)
	var cDate string
	if date == "0" {
		cDate = "0"
	}
	if date == "1" {
		cDate = "365"
	}
	if date == "2" {
		cDate = "90"
	}
	if date == "3" {
		cDate = "30"
	}
	if date == "4" {
		cDate = "7"
	}
	list, err := articleDto.QueryArticleList(cDate, aType, tag, strconv.Itoa(cPage*10))
	if err != nil {
		res.Res = 0
		res.List = nil
		fmt.Println(err)
	} else {
		res.Res = 1
		res.List = list
	}
	data, _ := json.Marshal(res)
	return ctx.String(http.StatusOK, string(data))
}

func SetArticleStatus(ctx echo.Context) error {
	res := new(entry.Res)
	id := ctx.FormValue("id")
	hFlag := ctx.FormValue("hFlag")
	cFlag, _ := strconv.Atoi(hFlag)
	err := articleDto.UpdateArticleStatus(id, cFlag)
	if err != nil {
		res.Res = 0
	} else {
		res.Res = 1
	}
	diarydto.InsertDiary("你修改了编号"+id+"的文章发布状态", 0)
	data, _ := json.Marshal(res)
	return ctx.String(http.StatusOK, string(data))
}

func DeleteArticle(ctx echo.Context) error {
	res := new(entry.Res)
	id := ctx.FormValue("id")
	err := articleDto.DeleteArticle(id)
	if err != nil {
		fmt.Println(err)
		res.Res = 0
	} else {
		err = articleDto.DeleteATags(id)
		if err != nil {
			fmt.Println(err)
			res.Res = 0
		} else {
			res.Res = 1
		}
	}
	diarydto.InsertDiary("你删除编号为"+id+"的文章", 0)
	data, _ := json.Marshal(res)
	return ctx.String(http.StatusOK, string(data))
}

func UpdateMood(ctx echo.Context) error {
	res := new(entry.Res)
	text := ctx.FormValue("text")
	picStr := ctx.FormValue("pic")
	id, err := articleDto.QueryMaxMoodId()
	if err != nil {
		res.Res = 0
	} else {
		err = articleDto.UpdateMood(id, text, picStr)
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

func GetArticle(ctx echo.Context) error {
	res := new(entry.ArticleRes)
	id := ctx.QueryParam("id")
	var err error
	res.Article, err = articleDto.QueryArticle(id)
	if err != nil {
		fmt.Println()
		res.Res = 0
		res.Article = nil
	} else {
		res.Res = 1
	}
	data, _ := json.Marshal(res)
	return ctx.String(http.StatusOK, string(data))
}

func GetCommentSum(ctx echo.Context) error {
	aType := ctx.QueryParam("type")
	res := new(entry.SumRes)
	sum, err := articleDto.QueryCommentSum(aType)
	if err != nil {
		fmt.Println(err)
		res.Res = 0
		res.Sum = 0
	} else {
		res.Res = 1
		res.Sum = sum
	}
	data, _ := json.Marshal(res)
	return ctx.String(http.StatusOK, string(data))
}

func GetCommentList(ctx echo.Context) error {
	res := new(entry.CListRes)
	aType := ctx.QueryParam("type")
	page := ctx.QueryParam("page")
	cPage, _ := strconv.Atoi(page)
	list, err := articleDto.QueryCommentList(aType, strconv.Itoa(cPage*10))
	if err != nil {
		res.Res = 0
		res.List = nil
		fmt.Println(err)
	} else {
		res.Res = 1
		res.List = list
	}
	data, _ := json.Marshal(res)
	return ctx.String(http.StatusOK, string(data))
}

func SetCommentStatus(ctx echo.Context) error {
	res := new(entry.Res)
	id := ctx.FormValue("id")
	title := ctx.FormValue("title")
	hFlag := ctx.FormValue("hFlag")
	cFlag, _ := strconv.Atoi(hFlag)
	err := articleDto.UpdateCommentStatus(id, title, cFlag)
	if err != nil {
		res.Res = 0
	} else {
		res.Res = 1
	}
	diarydto.InsertDiary("你更改了文章《"+title+"》的第"+id+"条评论的状态", 0)
	data, _ := json.Marshal(res)
	return ctx.String(http.StatusOK, string(data))
}

func DeleteComment(ctx echo.Context) error {
	res := new(entry.Res)
	id := ctx.FormValue("id")
	title := ctx.FormValue("title")
	err := articleDto.DeleteComment(id, title)
	if err != nil {
		res.Res = 0
	} else {
		res.Res = 1
	}
	diarydto.InsertDiary("你删除了文章《"+title+"》的第"+id+"条评论", 0)
	data, _ := json.Marshal(res)
	return ctx.String(http.StatusOK, string(data))
}

func GetTagsDetails(ctx echo.Context) error {
	res := new(entry.TagsDetailsRes)
	tags, err := articleDto.QueryTagsDetails()
	if err != nil {
		res.Res = 0
		res.Tags = nil
	} else {
		res.Res = 1
		res.Tags = tags
	}
	data, _ := json.Marshal(res)
	return ctx.String(http.StatusOK, string(data))
}

func GetMaxTags(ctx echo.Context) error {
	res := new(entry.TagsDetailsRes)
	tags, err := articleDto.QueryMaxTags()
	if err != nil {
		res.Res = 0
		res.Tags = nil
	} else {
		res.Res = 1
		res.Tags = tags
	}
	data, _ := json.Marshal(res)
	return ctx.String(http.StatusOK, string(data))
}
