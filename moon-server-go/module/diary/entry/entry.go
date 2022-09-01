package entry

import (
	sqlEntry "moon-manage-serve/sql/diaryDto/entry"
)

type DiaryTimeRes struct {
	Res    int                 `json:"res"`
	Diarys []sqlEntry.DiaryOne `json:"diarys"`
}

type DiaryAllRes struct {
	Res    int      `json:"res"`
	Diarys []string `json:"diarys"`
}

type SumRes struct {
	Res int `json:"res"`
	Sum int `json:"sum"`
}

type MessageRes struct {
	Res      int                `json:"res"`
	Messages []sqlEntry.Message `json:"messages"`
}

type Res struct {
	Res int `json:"res"`
}

type NearRes struct {
	Res  int               `json:"res"`
	Near []sqlEntry.Active `json:"near"`
}

type DiaryRes struct {
	Res   int              `json:"res"`
	Diary []sqlEntry.Diary `json:"diary"`
}
