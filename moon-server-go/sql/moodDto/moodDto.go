package moodDto

import (
	"database/sql"
	"fmt"
	"moon-manage-serve/sql/moodDto/entry"
	"strconv"
	"strings"
	"time"
)

var db *sql.DB

func SetDb(Db *sql.DB) {
	db = Db
}

func QueryMoodId() (int, error) {
	str := "SELECT COUNT(id) FROM mood"
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

func InsertMood(id int, text string, pic string) error {
	date := time.Now()
	str := "INSERT INTO mood(id,text,pic,time,mlike,comment) VALUE(?,?,?,?,0,0)"
	qry, err := db.Prepare(str)
	if err != nil {
		return err
	}
	defer qry.Close()
	_, err = qry.Exec(id, text, pic, date)
	if err != nil {
		return err
	}
	return nil
}

func QueryMoodSum() (int, error) {
	str := "SELECT COUNT(id) FROM mood"
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

func QueryMoodList(num string) ([]entry.Mood, error) {
	str := "SELECT id,text,pic,time,mlike,comment FROM mood ORDER BY id DESC LIMIT ?,5"
	qry, err := db.Prepare(str)
	if err != nil {
		return nil, err
	}
	defer qry.Close()
	rows, err := qry.Query(num)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var moods []entry.Mood
	for rows.Next() {
		var mood entry.Mood
		rows.Scan(&mood.Id, &mood.Text, &mood.Pic, &mood.Time, &mood.Like, &mood.Comment)
		mood.Links, err = QueryLink(mood.Pic)
		if err != nil {
			return nil, err
		}
		mood.Comments, err = QueryComments(strconv.Itoa(mood.Id), 0, 2)
		if err != nil {
			return nil, err
		}
		moods = append(moods, mood)
	}
	return moods, nil
}

func QueryLink(picStr string) ([]string, error) {
	if picStr == "" {
		return nil, nil
	}
	picArray := strings.Split(picStr, "#")
	str := "SELECT path FROM mImage"
	if len(picArray) > 0 {
		str += (" WHERE id=" + picArray[0])
	}
	for index, pic := range picArray {
		if index > 0 {
			str += (" OR id=" + pic)
		}
	}
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
	var links []string
	for rows.Next() {
		var link string
		rows.Scan(&link)
		links = append(links, link)
	}
	return links, nil
}

func QueryComments(id string, page int, num int) ([]entry.MComment, error) {
	str := "SELECT cid,name,cname,content,cflag,time FROM pcomment WHERE id=? ORDER BY cid DESC LIMIT ?,?"
	qry, err := db.Prepare(str)
	if err != nil {
		return nil, err
	}
	defer qry.Close()
	rows, err := qry.Query(id, page, num)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var comments []entry.MComment
	for rows.Next() {
		var comment entry.MComment
		rows.Scan(&comment.Cid, &comment.Name, &comment.Cname, &comment.Content, &comment.Cflag, &comment.Time)
		comments = append(comments, comment)
	}
	return comments, nil
}

func DeleteComment(id string, cid string) error {
	str := "DELETE FROM pcomment WHERE id=? AND cid=?"
	qry, err := db.Prepare(str)
	if err != nil {
		return err
	}
	defer qry.Close()
	_, err = qry.Exec(id, cid)
	if err != nil {
		return err
	}
	str = "UPDATE pcomment SET pcomment.cid=pcomment.cid-1 WHERE pcomment.id>? AND pcomment.id=?"
	qry, err = db.Prepare(str)
	if err != nil {
		return err
	}
	defer qry.Close()
	_, err = qry.Exec(cid, id)
	if err != nil {
		return err
	}
	str = "UPDATE mood SET mood.comment=mood.comment-1 WHERE mood.id=?"
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

func DeleteMood(id string) error {
	str := "DELETE FROM pcomment WHERE id=?"
	qry, err := db.Prepare(str)
	if err != nil {
		return err
	}
	defer qry.Close()
	_, err = qry.Exec(id)
	if err != nil {
		return err
	}
	str = "DELETE FROM mood WHERE id=?"
	qry, err = db.Prepare(str)
	if err != nil {
		return err
	}
	defer qry.Close()
	_, err = qry.Exec(id)
	if err != nil {
		return err
	}
	str = "UPDATE mood SET mood.id=mood.id-1 WHERE mood.id>?"
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

func QueryMaxNote() (int, error) {
	str := "SELECT COUNT(id) FROM note"
	qry, err := db.Prepare(str)
	if err != nil {
		return 0, err
	}
	defer qry.Close()
	row, err := qry.Query()
	if err != nil {
		return 0, err
	}
	sum := 0
	if row.Next() {
		row.Scan(&sum)
	}
	return sum, nil
}

func InsertNote(id int, date string, text string) error {
	str := "INSERT INTO note(id,time,text) VALUE(?,?,?)"
	qry, err := db.Prepare(str)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer qry.Close()
	_, err = qry.Exec(id, date, text)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func QueryOldNote() (*entry.Note, error) {
	str := "SELECT id,text,time FROM note WHERE TO_DAYS(NOW( )) - TO_DAYS(time)>= 1 ORDER BY id DESC LIMIT 0,1"
	qry, err := db.Prepare(str)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer qry.Close()
	row, err := qry.Query()
	if err != nil {
		return nil, err
	}
	defer row.Close()
	var note entry.Note
	if row.Next() {
		row.Scan(&note.Id, &note.Text, &note.Time)
	}
	return &note, nil
}
func QueryNewNote(page string) ([]entry.Note, error) {
	str := "SELECT id,text,time FROM note WHERE TO_DAYS(time) - TO_DAYS(NOW( ))>= 0 ORDER BY id ASC LIMIT ?,5"
	qry, err := db.Prepare(str)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer qry.Close()
	row, err := qry.Query(page)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	var notes []entry.Note
	for row.Next() {
		var note entry.Note
		row.Scan(&note.Id, &note.Text, &note.Time)
		notes = append(notes, note)
	}
	return notes, nil
}

func QueryNoteSum() (int, error) {
	str := "SELECT COUNT(id) FROM note WHERE TO_DAYS(time) - TO_DAYS(NOW( ))>= 0"
	qry, err := db.Prepare(str)
	if err != nil {
		fmt.Println(err)
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

func DeleteNote(id string) error {
	str := "DELETE FROM note WHERE id=?"
	qry, err := db.Prepare(str)
	if err != nil {
		return err
	}
	defer qry.Close()
	_, err = qry.Exec(id)
	if err != nil {
		return err
	}
	str = "UPDATE note SET note.id=note.id-1 WHERE note.id>?"
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
