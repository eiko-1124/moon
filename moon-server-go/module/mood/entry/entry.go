package entry

import sqlEntry "moon-manage-serve/sql/moodDto/entry"

type Res struct {
	Res int `json:"res"`
}

type MoodRes struct {
	Res   int             `json:"res"`
	Moods []sqlEntry.Mood `json:"moods"`
}

type MoodSumRes struct {
	Res int `json:"res"`
	Sum int `json:"sum"`
}

type LinkRes struct {
	Res   int      `json:"res"`
	Links []string `json:"links"`
}

type CommentRes struct {
	Res      int                 `json:"res"`
	Comments []sqlEntry.MComment `json:"comments"`
}

type OldNoteRes struct {
	Res  int            `json:"res"`
	Note *sqlEntry.Note `json:"note"`
}

type NewNoteRes struct {
	Res   int             `json:"res"`
	Notes []sqlEntry.Note `json:"notes"`
}

type NoteSumRes struct {
	Res int `json:"res"`
	Sum int `json:"sum"`
}
