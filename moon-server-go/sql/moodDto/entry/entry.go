package entry

type Mood struct {
	Id       int        `json:"id"`
	Text     string     `json:"text"`
	Pic      string     `json:"pic"`
	Time     string     `json:"time"`
	Like     int        `json:"like"`
	Comment  int        `json:"comment"`
	Links    []string   `json:"links"`
	Comments []MComment `json:"comments"`
}

type MComment struct {
	Cid     int    `json:"cid"`
	Name    string `json:"name"`
	Cname   string `json:"cname"`
	Content string `json:"content"`
	Cflag   int    `json:"cFlag"`
	Time    string `json:"time"`
}

type Note struct {
	Id   int    `json:"id"`
	Text string `json:"text"`
	Time string `json:"time"`
}
