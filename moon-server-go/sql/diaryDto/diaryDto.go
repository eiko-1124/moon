package diarydto

import (
	"database/sql"
	"moon-manage-serve/sql/diaryDto/entry"
	"time"
)

var db *sql.DB

func SetDb(Db *sql.DB) {
	db = Db
}

func InsertDiary(text string, dType int) error {
	date := time.Now()
	id, err := QueryDiaryId(dType)
	if err != nil {
		return err
	}
	str := "INSERT INTO diary (id,dtype,text,time) VALUE(?,?,?,?)"
	qry, err := db.Prepare(str)
	if err != nil {
		return err
	}
	defer qry.Close()
	_, err = qry.Exec(id+1, dType, text, date)
	if err != nil {
		return err
	}
	return nil
}

func QueryDiaryId(dType int) (int, error) {
	str := "SELECT COUNT(id) FROM diary WHERE dtype=?"
	qry, err := db.Prepare(str)
	if err != nil {
		return 0, err
	}
	defer qry.Close()
	rows, err := qry.Query(dType)
	if err != nil {
		return 0, err
	}
	id := 0
	if rows.Next() {
		rows.Scan(&id)
	}
	return id, nil
}

func QueryDiaryTime(page string, dType string) ([]entry.DiaryOne, error) {
	str := "SELECT diary.time FROM diary WHERE diary.dtype=? GROUP BY time ORDER BY time DESC LIMIT ?,10"
	qry, err := db.Prepare(str)
	if err != nil {
		return nil, err
	}
	defer qry.Close()
	rows, err := qry.Query(dType, page)
	if err != nil {
		return nil, err
	}
	var diarys []entry.DiaryOne
	for rows.Next() {
		var diary entry.DiaryOne
		rows.Scan(&diary.Date)
		id, text, err := QueryDiaryNew(dType, diary.Date)
		if err != nil {
			return nil, err
		}
		sum, err := QueryDiarySiteSum(dType, diary.Date)
		if err != nil {
			return nil, err
		}
		diary.Sum = sum
		diary.Id = id
		diary.Text = text
		diarys = append(diarys, diary)
	}
	return diarys, nil
}

func QueryDiaryNew(dType string, date string) (int, string, error) {
	str := "SELECT diary.id,diary.text FROM diary WHERE diary.dtype=? AND diary.time=? ORDER BY diary.id DESC LIMIT 0,1"
	qry, err := db.Prepare(str)
	if err != nil {
		return 0, "", err
	}
	defer qry.Close()
	row, err := qry.Query(dType, date)
	if err != nil {
		return 0, "", err
	}
	defer row.Close()
	var id int
	var text string
	if row.Next() {
		row.Scan(&id, &text)
	}
	return id, text, nil
}

func QueryDiaryAll(dType string, date string, page string) ([]string, error) {
	str := "SELECT diary.text FROM diary WHERE diary.dtype=? AND diary.time=? ORDER BY diary.id DESC LIMIT ?,10"
	qry, err := db.Prepare(str)
	if err != nil {
		return nil, err
	}
	defer qry.Close()
	rows, err := qry.Query(dType, date, page)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var diarys []string
	for rows.Next() {
		var diary string
		rows.Scan(&diary)
		diarys = append(diarys, diary)
	}
	return diarys, nil
}

func QueryDiarySum(dType string) (int, error) {
	str := "SELECT COUNT(DISTINCT diary.time) FROM diary WHERE diary.dtype=?"
	qry, err := db.Prepare(str)
	if err != nil {
		return 0, err
	}
	defer qry.Close()
	row, err := qry.Query(dType)
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

func QueryDiarySiteSum(dType string, date string) (int, error) {
	str := "SELECT COUNT(DISTINCT diary.id) FROM diary WHERE diary.dtype=? AND diary.time=?"
	qry, err := db.Prepare(str)
	if err != nil {
		return 0, err
	}
	defer qry.Close()
	row, err := qry.Query(dType, date)
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

func QueryMessage(page string) ([]entry.Message, error) {
	str := "SELECT id,name,text,time FROM message ORDER BY id DESC LIMIT ?,10"
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
	var messages []entry.Message
	for rows.Next() {
		var message entry.Message
		rows.Scan(&message.Id, &message.Name, &message.Text, &message.Time)
		messages = append(messages, message)
	}
	return messages, nil
}

func QueryMessageSum() (int, error) {
	str := "SELECT COUNT(*) FROM message"
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

func DeleteMessage(id string) error {
	str := "DELETE FROM message WHERE id=?"
	qry, err := db.Prepare(str)
	if err != nil {
		return err
	}
	defer qry.Close()
	_, err = qry.Exec(id)
	if err != nil {
		return err
	}
	str = "UPDATE message SET message.id=message.id-1 WHERE message.id>?"
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

func QueryNearDiary(dType string) ([]entry.Active, error) {
	str := "SELECT tbl._date AS date,IFNULL(tbr.taskCount, 0) AS taskCount FROM(SELECT @s :=@s + 1 AS _index,DATE(DATE_SUB(CURRENT_DATE, INTERVAL @s DAY)) AS _date FROM information_schema. TABLES, (SELECT @s := - 1) temp WHERE @s <= 6 ORDER BY _date ) tbl LEFT JOIN ( SELECT count(*) AS taskCount, DATE(time) startdate FROM diary WHERE dtype=? GROUP BY startdate) AS tbr ON tbl._date = tbr.startdate GROUP BY tbl._date"
	qry, err := db.Prepare(str)
	if err != nil {
		return nil, err
	}
	defer qry.Close()
	rows, err := qry.Query(dType)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var actives []entry.Active
	for rows.Next() {
		var active entry.Active
		rows.Scan(&active.Date, &active.Num)
		actives = append(actives, active)
	}
	return actives, nil
}

func QueryDiary(dType string) ([]entry.Diary, error) {
	str := "SELECT text,time FROM diary WHERE dtype=? ORDER BY id DESC LIMIT 0,10"
	qry, err := db.Prepare(str)
	if err != nil {
		return nil, err
	}
	defer qry.Close()
	rows, err := qry.Query(dType)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var diarys []entry.Diary
	for rows.Next() {
		var diary entry.Diary
		rows.Scan(&diary.Text, &diary.Time)
		diarys = append(diarys, diary)
	}
	return diarys, nil
}
