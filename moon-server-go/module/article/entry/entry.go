package entry

import (
	sqlEntry "moon-manage-serve/sql/articleDto/entry"
)

type ArticleForm struct {
	Id         int    `form:"id"`
	Title      string `form:"title"`
	Atype      int    `form:"aType"`
	Tags       string `form:"tag"`
	Illustrate string `form:"illustrate"`
	Cover      string `form:"cover"`
	Fcontent   string `form:"fContent"`
	Rcontent   string `form:"rContent"`
}

type Res struct {
	Res int `json:"res"`
}

type TagsRes struct {
	Res  int      `json:"res"`
	Tags []string `json:"tags"`
}

type SumRes struct {
	Res int `json:"res"`
	Sum int `json:"sum"`
}

type ListRes struct {
	Res  int                    `json:"res"`
	List []sqlEntry.ArticleSite `json:"list"`
}

type ArticleRes struct {
	Res     int                  `json:"res"`
	Article *sqlEntry.SqlArticle `json:"article"`
}

type CListRes struct {
	Res  int                    `json:"res"`
	List []sqlEntry.CommentSite `json:"list"`
}

type TagsDetailsRes struct {
	Res  int            `json:"res"`
	Tags []sqlEntry.Tag `json:"tags"`
}
