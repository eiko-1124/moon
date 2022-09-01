package entry

import "time"

type SqlArticle struct {
	Id         int       `json:"id"`
	Title      string    `json:"title"`
	Atpye      int       `json:"aType"`
	Illustrate string    `json:"illustrate"`
	Tags       string    `json:"tag"`
	Atime      time.Time `json:"time"`
	Cover      string    `json:"cover"`
	Fcontent   string    `json:"fContent"`
	Rcontent   string    `json:"rContent"`
}

type ArticleSite struct {
	Id         int    `json:"id"`
	Title      string `json:"title"`
	Atpye      int    `json:"aType"`
	Illustrate string `json:"illustrate"`
	Tags       string `json:"tags"`
	Atime      string `json:"time"`
	Cover      string `json:"cover"`
	Hflag      int    `json:"hFlag"`
}

type CommentSite struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Atime   string `json:"time"`
	Atype   string `json:"aType"`
	Content string `json:"content"`
	Name    string `json:"name"`
	Cname   string `json:"cName"`
	Hflag   int    `json:"hFlag"`
	Cflag   int    `json:"cFlag"`
	Email   string `json:"email"`
}

type Tag struct {
	Tag string `json:"tag"`
	Num int    `json:"num"`
}
