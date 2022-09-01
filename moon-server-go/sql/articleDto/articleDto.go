package articledto

import (
	"database/sql"
	"moon-manage-serve/sql/articleDto/entry"
	"strings"
	"time"
)

var db *sql.DB

func SetDb(Db *sql.DB) {
	db = Db
}

func UpdateArticle(article entry.SqlArticle) error {
	str := "UPDATE article SET title=?,atype=?,illustrate=?,cover=?,fcontent=?,rcontent=?,tags=? WHERE id=?"
	qry, err := db.Prepare(str)
	if err != nil {
		return err
	}
	defer qry.Close()
	_, err = qry.Exec(article.Title, article.Atpye, article.Illustrate, article.Cover, article.Fcontent, article.Rcontent, article.Tags, article.Id)
	if err != nil {
		return err
	}
	err = DeleteTags(article.Id)
	if err != nil {
		return err
	}
	err = InsetTags(article.Id, article.Tags)
	if err != nil {
		return err
	}
	return nil
}

func QueryMaxId() (int, error) {
	var id int
	str := "SELECT COUNT(article.id) FROM article"
	qry, err := db.Prepare(str)
	if err != nil {
		return 0, err
	}
	defer qry.Close()
	row, err := qry.Query()
	if err != nil {
		return 0, err
	}
	defer row.Close()
	if row.Next() {
		err := row.Scan(&id)
		if err != nil {
			return 0, err
		}
	}
	return id + 1, nil
}

func InsetArticle(article entry.SqlArticle) error {
	str := "insert into article (id,title,atype,illustrate,tags,cover,time,fcontent,rcontent,hflag) value(?,?,?,?,?,?,?,?,?,?)"
	qry, err := db.Prepare(str)
	if err != nil {
		return err
	}
	defer qry.Close()
	_, err = qry.Exec(article.Id, article.Title, article.Atpye, article.Illustrate, article.Tags, article.Cover, article.Atime, article.Fcontent, article.Rcontent, 0)
	if err != nil {
		return err
	}
	err = InsetTags(article.Id, article.Tags)
	if err != nil {
		return err
	}
	return nil
}

func DeleteTags(id int) error {
	str := "DELETE FROM tags WHERE id=?"
	qry, err := db.Prepare(str)
	if err != nil {
		return err
	}
	defer qry.Close()
	_, err = qry.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

func InsetTags(id int, tags string) error {
	sqlStr := "insert into tags(id, tag) values"
	tagArray := strings.Split(tags, "#")
	vals := []interface{}{}
	for index, tag := range tagArray {
		if index == len(tagArray)-1 {
			sqlStr += "(?, ?)"
		} else {
			sqlStr += "(?, ?), "
		}
		vals = append(vals, id, tag)
	}
	qry, err := db.Prepare(sqlStr)
	if err != nil {
		return err
	}
	defer qry.Close()
	_, err = qry.Exec(vals...)
	if err != nil {
		return err
	}
	return nil
}

func QueryTags() ([]string, error) {
	str := "SELECT DISTINCT tags.tag FROM tags"
	qry, err := db.Prepare(str)
	if err != nil {
		return nil, err
	}
	defer qry.Close()
	rows, err := qry.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var tags []string
	for rows.Next() {
		var tag string
		rows.Scan(&tag)
		tags = append(tags, tag)
	}
	return tags, nil
}

func QueryArticleSum(date string, aType string, tag string) (int, error) {
	var str string = "SELECT COUNT(DISTINCT article.id) FROM tags,article WHERE article.id = tags.id"
	if date != "0" {
		str = str + " AND date_sub(curdate(), interval " + date + " day) <= date(article.time)"
	}
	if tag != "0" {
		str = str + " AND tags.tag = " + tag
	}
	if aType != "0" {
		str = str + " AND article.atype=" + aType
	}
	qry, err := db.Prepare(str)
	if err != nil {
		return 0, err
	}
	defer qry.Close()
	row, err := qry.Query()
	if err != nil {
		return 0, err
	}
	defer row.Close()
	var sum int
	if row.Next() {
		row.Scan(&sum)
	}
	return sum, nil
}

func QueryArticleList(date string, aType string, tag string, page string) ([]entry.ArticleSite, error) {
	var str string = "SELECT DISTINCT article.id,article.title,article.atype,article.cover,article.tags,article.illustrate,article.time,article.hflag FROM tags,article WHERE article.id = tags.id"
	if date != "0" {
		str = str + " AND date_sub(curdate(), interval " + date + " day) <= date(article.time)"
	}
	if tag != "0" {
		str = str + " AND tags.tag = \"" + tag + "\""
	}
	if aType != "0" {
		str = str + " AND article.atype=" + aType
	}
	str = str + " ORDER BY id DESC" + " LIMIT " + page + ", 10"
	qry, err := db.Prepare(str)
	if err != nil {
		return nil, err
	}
	defer qry.Close()
	rows, err := qry.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []entry.ArticleSite
	for rows.Next() {
		var site entry.ArticleSite
		rows.Scan(&site.Id, &site.Title, &site.Atpye, &site.Cover, &site.Tags, &site.Illustrate, &site.Atime, &site.Hflag)
		list = append(list, site)
	}
	return list, nil
}

func UpdateArticleStatus(id string, hFlag int) error {
	str := "UPDATE article SET hflag=? WHERE id=?"
	qry, err := db.Prepare(str)
	if err != nil {
		return err
	}
	defer qry.Close()
	_, err = qry.Exec(hFlag, id)
	if err != nil {
		return err
	}
	return nil
}

func QueryMaxMoodId() (int, error) {
	var id int
	str := "SELECT MAX(mood.id) FROM mood"
	qry, err := db.Prepare(str)
	if err != nil {
		return 0, err
	}
	defer qry.Close()
	row, err := qry.Query()
	if err != nil {
		return 0, err
	}
	defer row.Close()
	if row.Next() {
		err := row.Scan(&id)
		if err != nil {
			return 0, err
		}
	}
	return id + 1, nil
}

func UpdateMood(id int, text string, pic string) error {
	str := "insert into mood(id, text,pic,time) values(?,?,?,?)"
	qry, err := db.Prepare(str)
	if err != nil {
		return err
	}
	defer qry.Close()
	date := time.Now().Unix()
	_, err = qry.Exec(id, text, pic, date)
	if err != nil {
		return err
	}
	return nil
}

func QueryArticle(id string) (*entry.SqlArticle, error) {
	str := "SELECT id,title,atype,illustrate,cover,fcontent,rcontent,tags from article where id=?"
	qry, err := db.Prepare(str)
	if err != nil {
		return nil, err
	}
	row, err := qry.Query(id)
	if err != nil {
		return nil, err
	}
	article := new(entry.SqlArticle)
	if row.Next() {
		err = row.Scan(&article.Id, &article.Title, &article.Atpye, &article.Illustrate, &article.Cover, &article.Fcontent, &article.Rcontent, &article.Tags)
		if err != nil {
			return nil, err
		}
	}
	return article, nil
}

func QueryCommentSum(aType string) (int, error) {
	var str string = "SELECT COUNT(comment.id) FROM comment"
	if aType != "0" {
		str = str + ",article WHERE article.atype=" + aType + " AND article.title=comment.title"
	}
	qry, err := db.Prepare(str)
	if err != nil {
		return 0, err
	}
	defer qry.Close()
	row, err := qry.Query()
	if err != nil {
		return 0, err
	}
	defer row.Close()
	var sum int
	if row.Next() {
		row.Scan(&sum)
	}
	return sum, nil
}

func QueryCommentList(aType string, page string) ([]entry.CommentSite, error) {
	var str string = "SELECT article.atype,comment.id,comment.title,comment.time,comment.content,comment.email,comment.name,comment.cname,comment.hflag,comment.cflag FROM comment,article WHERE article.title=comment.title"
	if aType != "0" {
		str = str + " AND article.atype=" + aType
	}
	str = str + " ORDER BY id DESC" + " LIMIT " + page + ", 10"
	qry, err := db.Prepare(str)
	if err != nil {
		return nil, err
	}
	defer qry.Close()
	rows, err := qry.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []entry.CommentSite
	for rows.Next() {
		var site entry.CommentSite
		rows.Scan(&site.Atype, &site.Id, &site.Title, &site.Atime, &site.Content, &site.Email, &site.Name, &site.Cname, &site.Hflag, &site.Cflag)
		list = append(list, site)
	}
	return list, nil
}

func UpdateCommentStatus(id string, title string, hFlag int) error {
	str := "UPDATE comment SET hflag=? WHERE id=? and title=?"
	qry, err := db.Prepare(str)
	if err != nil {
		return err
	}
	defer qry.Close()
	_, err = qry.Exec(hFlag, id, title)
	if err != nil {
		return err
	}
	return nil
}

func DeleteComment(id string, title string) error {
	str := "DELETE FROM comment WHERE (comment.id=? AND comment.title=?) OR (comment.title=? AND comment.cid=?)"
	qry, err := db.Prepare(str)
	if err != nil {
		return err
	}
	defer qry.Close()
	_, err = qry.Exec(id, title, title, id)
	if err != nil {
		return err
	}
	return nil
}

func DeleteArticle(id string) error {
	str := "DELETE article FROM article WHERE article.id=?"
	qry, err := db.Prepare(str)
	if err != nil {
		return err
	}
	defer qry.Close()
	_, err = qry.Exec(id)
	if err != nil {
		return err
	}
	str = "ALTER TABLE article DROP id"
	qry, err = db.Prepare(str)
	if err != nil {
		return err
	}
	defer qry.Close()
	_, err = qry.Exec()
	if err != nil {
		return err
	}
	str = "ALTER TABLE article ADD id INT NOT NULL PRIMARY KEY AUTO_INCREMENT FIRST"
	qry, err = db.Prepare(str)
	if err != nil {
		return err
	}
	defer qry.Close()
	_, err = qry.Exec()
	if err != nil {
		return err
	}
	err = DeleteAComment(id)
	if err != nil {
		return err
	}
	return nil
}

func DeleteATags(id string) error {
	str := "DELETE FROM tags WHERE id=?"
	qry, err := db.Prepare(str)
	if err != nil {
		return err
	}
	defer qry.Close()
	_, err = qry.Exec(id)
	if err != nil {
		return err
	}
	str = "UPDATE tags SET tags.id=tags.id-1 WHERE tags.id>?"
	qry, err = db.Prepare(str)
	if err != nil {
		return err
	}
	defer qry.Close()
	_, err = qry.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

func DeleteAComment(title string) error {
	str := "DELETE FROM comment WHERE title=?"
	qry, err := db.Prepare(str)
	if err != nil {
		return err
	}
	defer qry.Close()
	_, err = qry.Exec(title)
	if err != nil {
		return err
	}
	return nil
}

func QueryTagsDetails() ([]entry.Tag, error) {
	str := "select tag,count(*) as count from tags group by tag;"
	qry, err := db.Prepare(str)
	if err != nil {
		return nil, err
	}
	defer qry.Close()
	var tags []entry.Tag
	rows, err := qry.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var tag entry.Tag
		rows.Scan(&tag.Tag, &tag.Num)
		tags = append(tags, tag)
	}
	return tags, nil
}

func QueryMaxTags() ([]entry.Tag, error) {
	str := "select tag,count(tag) as number from tags group by tag order by number desc limit 3"
	qry, err := db.Prepare(str)
	if err != nil {
		return nil, err
	}
	defer qry.Close()
	var tags []entry.Tag
	rows, err := qry.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var tag entry.Tag
		rows.Scan(&tag.Tag, &tag.Num)
		tags = append(tags, tag)
	}
	return tags, nil
}
