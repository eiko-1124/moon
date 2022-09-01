package entry

import (
	"mime/multipart"
	sqlEntry "moon-manage-serve/sql/fileDto/entry"
)

type File struct {
	Name string `json:"fileName"`
	Link string `json:"link"`
}

type UpdateRes struct {
	Res  string `json:"res"`
	File File   `json:"file"`
}

type FileQueue struct {
	FileName string
	Order    int
	Queue    map[int]*multipart.FileHeader
}

type PicRes struct {
	Res int    `json:"res"`
	Id  string `json:"id"`
}

type SumRes struct {
	Res int `json:"res"`
	Sum int `json:"sum"`
}

type AcoverRes struct {
	Res    int              `json:"res"`
	Covers []sqlEntry.Cover `json:"covers"`
}

type ImageRes struct {
	Res    int              `json:"res"`
	Images []sqlEntry.Image `json:"images"`
}

type OtherFilesRes struct {
	Res   int                  `json:"res"`
	Files []sqlEntry.OtherFile `json:"files"`
}
