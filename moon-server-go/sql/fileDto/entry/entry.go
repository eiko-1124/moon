package entry

type File struct {
	FileName string `json:"fileName"`
	Link     string `json:"link"`
}

type Cover struct {
	Title string `json:"title"`
	Cover string `json:"cover"`
}

type Image struct {
	Time  string `json:"time"`
	Cover string `json:"cover"`
}

type OtherFile struct {
	Name string `json:"name"`
	Link string `json:"link"`
}
