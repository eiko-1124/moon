package entry

// User
type User struct {
	Id   string `json:"id" form:"id" query:"id"`
	Pswd string `json:"pswd" form:"pswd" query:"pswd"`
}

type LoginRes struct {
	Token string `json:"token"`
	Res   string `json:"res"`
}
