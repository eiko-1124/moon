package entry

type DiaryOne struct {
	Id   int    `json:"id"`
	Date string `json:"date"`
	Text string `json:"text"`
	Sum  int    `json:"sum"`
}

type Message struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Text string `json:"text"`
	Time string `json:"time"`
}

type Active struct {
	Date string `json:"date"`
	Num  string `json:"num"`
}

type Diary struct {
	Text string `json:"text"`
	Time string `json:"time"`
}
