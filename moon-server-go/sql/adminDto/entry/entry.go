package entry

type Info struct {
	Name      string `json:"name"`
	Notice    string `json:"notice"`
	Illustate string `json:"illustate"`
	Sentence  string `json:"sentence"`
	QQ        string `json:"qq"`
	Github    string `json:"github"`
}

type Sign struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
}

type FLink struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Flink string `json:"flink"`
}

type About struct {
	Rander string `json:"rander"`
	Format string `json:"format"`
}
