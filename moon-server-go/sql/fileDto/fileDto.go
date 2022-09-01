package filedto

import (
	"database/sql"
	"fmt"
	"moon-manage-serve/module/file/entry"
	sqlEntry "moon-manage-serve/sql/fileDto/entry"
)

var db *sql.DB

func SetDb(Db *sql.DB) {
	db = Db
}

func QueryFileName(name string) (bool, error) {
	str := "select * from file where name=?"
	qry, err := db.Prepare(str)
	if err != nil {
		return false, err
	}
	defer qry.Close()
	rows, err := qry.Query(name)
	if err != nil {
		return false, err
	}
	if rows.Next() {
		return true, nil
	}
	return false, nil
}

func SaveFile(name string, link string) error {
	str := "insert into  file  (name,link) value(?,?)"
	qry, err := db.Prepare(str)
	if err != nil {
		return err
	}
	defer qry.Close()
	_, err = qry.Exec(name, link)
	if err != nil {
		return err
	}
	return nil
}

func QueryAllFile(page int) ([]entry.File, error) {
	str := "select name,link from file limit ?,?"
	qry, err := db.Prepare(str)
	if err != nil {
		return nil, err
	}
	defer qry.Close()
	var files []entry.File
	rows, err := qry.Query(page*20, 20)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var file entry.File
		rows.Scan(&file.Name, &file.Link)
		files = append(files, file)
	}
	return files, nil
}

func SaveMoodImage(path string, today string) (int, error) {
	id, err := QueryImageId()
	if err != nil {
		fmt.Println(err.Error())
		return 0, nil
	}
	str := "insert into mImage (id,time,path) value(?,?,?)"
	qry, err := db.Prepare(str)
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}
	defer qry.Close()
	_, err = qry.Exec(id, today, "http://localhost:1323/"+path)
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}
	return id, nil
}

func QueryImageId() (int, error) {
	var id int
	str := "SELECT COUNT(mImage.id) FROM mImage"
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

func QueryArticleCover(page string) ([]sqlEntry.Cover, error) {
	str := "SELECT title,cover FROM article GROUP BY cover ORDER BY id DESC LIMIT ?,10"
	qry, err := db.Prepare(str)
	if err != nil {
		return nil, err
	}
	defer qry.Close()
	rows, err := qry.Query(page)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var covers []sqlEntry.Cover
	for rows.Next() {
		var cover sqlEntry.Cover
		rows.Scan(&cover.Title, &cover.Cover)
		covers = append(covers, cover)
	}
	return covers, nil
}

func QueryArticleCoverSum() (int, error) {
	str := "SELECT COUNT(DISTINCT cover) FROM article"
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

func QueryMoodImages(page string) ([]sqlEntry.Image, error) {
	str := "SELECT time,path FROM mImage GROUP BY path ORDER BY id DESC LIMIT ?,10"
	qry, err := db.Prepare(str)
	if err != nil {
		return nil, err
	}
	defer qry.Close()
	rows, err := qry.Query(page)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var images []sqlEntry.Image
	for rows.Next() {
		var image sqlEntry.Image
		rows.Scan(&image.Time, &image.Cover)
		images = append(images, image)
	}
	return images, nil
}

func QueryMoodImageSum() (int, error) {
	str := "SELECT COUNT(DISTINCT path) FROM mImage"
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

func QueryOtherFiles(page string) ([]sqlEntry.OtherFile, error) {
	str := " select file.name,file.link from file where (select count(1) as num from article where article.cover = file.link) = 0 limit ?,10"
	qry, err := db.Prepare(str)
	if err != nil {
		return nil, err
	}
	defer qry.Close()
	rows, err := qry.Query(page)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var files []sqlEntry.OtherFile
	for rows.Next() {
		var file sqlEntry.OtherFile
		rows.Scan(&file.Name, &file.Link)
		files = append(files, file)
	}
	return files, nil
}

func QueryOtherFilesSum() (int, error) {
	str := "select COUNT(file.name) from file where (select count(1) as num from article where article.cover = file.link) = 0"
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
