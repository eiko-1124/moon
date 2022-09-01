package admindto

import (
	"database/sql"
	"errors"
	"moon-manage-serve/sql/adminDto/entry"
)

var db *sql.DB

func SetDb(Db *sql.DB) {
	db = Db
}

func QueryName(id string, pswd string) (string, error) {
	str := "select name from user where id=? and pswd=?"
	qry, err := db.Prepare(str)
	if err != nil {
		return "", err
	}
	defer qry.Close()
	rows, err := qry.Query(id, pswd)
	if err != nil {
		return "", err
	}
	defer rows.Close()
	var name string
	if rows.Next() {
		err = rows.Scan(&name)
		if err != nil {
			return "", err
		}
		return name, nil
	}
	return "", errors.New("user does not exist")
}

func QueryInfo() (*entry.Info, error) {
	str := "select name,notice,illustate,sentence,qq,github from info where id=1"
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
	var info entry.Info
	if rows.Next() {
		err = rows.Scan(&info.Name, &info.Notice, &info.Illustate, &info.Sentence, &info.QQ, &info.Github)
		if err != nil {
			return nil, err
		}
		return &info, nil
	}
	return nil, errors.New("info not found")
}

func QuerySigns() ([]entry.Sign, error) {
	str := "select * from signs"
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
	var signs []entry.Sign
	for rows.Next() {
		var sign entry.Sign
		err = rows.Scan(&sign.Id, &sign.Content)
		if err != nil {
			return nil, err
		}
		signs = append(signs, sign)
	}
	return signs, nil
}

func QueryLink() ([]entry.FLink, error) {
	str := "select id,name,flink from flink"
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
	var links []entry.FLink
	for rows.Next() {
		var link entry.FLink
		err = rows.Scan(&link.Id, &link.Name, &link.Flink)
		if err != nil {
			return nil, err
		}
		links = append(links, link)
	}
	return links, nil
}

func QueryArticleSum() (int, error) {
	var str string = "SELECT COUNT(article.id) FROM article"
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
	sum := 0
	if row.Next() {
		row.Scan(&sum)
	}
	return sum, nil
}

func QueryMoodSum() (int, error) {
	var str string = "SELECT COUNT(mood.id) FROM mood"
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
	sum := 0
	if row.Next() {
		row.Scan(&sum)
	}
	return sum, nil
}

func QueryLinkSum() (int, error) {
	var str string = "SELECT COUNT(flink.id) FROM flink"
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
	sum := 0
	if row.Next() {
		row.Scan(&sum)
	}
	return sum, nil
}

func UpdateLink(id int, name string, flink string) error {
	str := "REPLACE INTO flink(id,name,flink) VALUES(?,?,?)"
	qry, err := db.Prepare(str)
	if err != nil {
		return err
	}
	defer qry.Close()
	_, err = qry.Exec(id, name, flink)
	if err != nil {
		return err
	}
	return nil
}

func UpdateSign(id int, content string) error {
	str := "REPLACE INTO signs(id,content) VALUES(?,?)"
	qry, err := db.Prepare(str)
	if err != nil {
		return err
	}
	defer qry.Close()
	_, err = qry.Exec(id, content)
	if err != nil {
		return err
	}
	return nil
}

func DeleteSign(id string) error {
	str := "DELETE FROM flink WHERE id=?"
	qry, err := db.Prepare(str)
	if err != nil {
		return err
	}
	defer qry.Close()
	_, err = qry.Exec(id)
	if err != nil {
		return err
	}
	str = "ALTER TABLE flink DROP id"
	qry, err = db.Prepare(str)
	if err != nil {
		return err
	}
	defer qry.Close()
	_, err = qry.Exec()
	if err != nil {
		return err
	}
	str = "ALTER TABLE flink ADD id INT NOT NULL PRIMARY KEY AUTO_INCREMENT FIRST"
	qry, err = db.Prepare(str)
	if err != nil {
		return err
	}
	defer qry.Close()
	_, err = qry.Exec()
	if err != nil {
		return err
	}
	return nil
}

func UpdateInfo(key string, value string) error {
	str := "UPDATE info SET " + key + "=\"" + value + "\" WHERE id=1"
	qry, err := db.Prepare(str)
	if err != nil {
		return err
	}
	defer qry.Close()
	_, err = qry.Exec()
	if err != nil {
		return err
	}
	return nil
}

func UpdateAbout(rander string, format string) error {
	str := "UPDATE about SET rtext=?,ftext=? WHERE id=1"
	qry, err := db.Prepare(str)
	if err != nil {
		return err
	}
	defer qry.Close()
	_, err = qry.Exec(rander, format)
	if err != nil {
		return err
	}
	return nil
}

func QueryAbout() (*entry.About, error) {
	str := "SELECT rtext,ftext FROM about WHERE id=1"
	qry, err := db.Prepare(str)
	if err != nil {
		return nil, err
	}
	defer qry.Close()
	row, err := qry.Query()
	if err != nil {
		return nil, err
	}
	defer row.Close()
	about := new(entry.About)
	if row.Next() {
		row.Scan(&about.Rander, &about.Format)
	}
	return about, nil
}

func QueryTagSum() (int, error) {
	var str string = "SELECT COUNT(DISTINCT tags.tag) FROM tags"
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
	sum := 0
	if row.Next() {
		row.Scan(&sum)
	}
	return sum, nil
}

func QueryCommentSum() (int, error) {
	var str string = "SELECT COUNT(*) FROM comment"
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
	sum := 0
	if row.Next() {
		row.Scan(&sum)
	}
	return sum, nil
}
