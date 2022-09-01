package login

import (
	"encoding/json"

	"net/http"

	"time"

	jwt "github.com/dgrijalva/jwt-go"

	"github.com/labstack/echo"

	"moon-manage-serve/module/login/entry"
	admindto "moon-manage-serve/sql/adminDto"
)

func GoLogin(ctx echo.Context) error {
	user := new(entry.User)
	if err := ctx.Bind(user); err != nil {
		return ctx.String(http.StatusUnauthorized, "uncacth error")
	}
	name, err := admindto.QueryName(user.Id, user.Pswd)
	if err != nil {
		return ctx.String(http.StatusUnauthorized, "id or pswd error")
	}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = name
	claims["id"] = "2994782369@qq.com"
	claims["pswd"] = "#Eiko1124"
	claims["exp"] = time.Now().Add(time.Hour * 168).Unix()
	t, err := token.SignedString([]byte("Mn+O2"))
	if err != nil {
		return ctx.String(http.StatusUnauthorized, "uncacth error")
	}
	res := new(entry.LoginRes)
	res.Token = t
	res.Res = "login success"
	data, err := json.Marshal(res)
	if err != nil {
		return ctx.String(http.StatusUnauthorized, "uncacth error")
	}
	return ctx.String(http.StatusOK, string(data))
}
