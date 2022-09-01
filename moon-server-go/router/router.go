package router

import (
	"moon-manage-serve/module/login"
	"moon-manage-serve/router/adminRouter"
	articlerouter "moon-manage-serve/router/articleRouter"
	diaryrouter "moon-manage-serve/router/diaryRouter"
	filerouter "moon-manage-serve/router/fileRouter"
	moodrouter "moon-manage-serve/router/moodRouter"
	admindto "moon-manage-serve/sql/adminDto"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	_ "github.com/go-sql-driver/mysql"

	"database/sql"
)

func CheckToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		next(c)
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		id := claims["id"].(string)
		pswd := claims["pswd"].(string)
		name := claims["name"].(string)
		uname, err := admindto.QueryName(id, pswd)
		if err != nil || name != uname {
			return echo.NewHTTPError(http.StatusUnauthorized)
		} else {
			return nil
		}
	}
}

func setDb(Db *sql.DB) {
	admindto.SetDb(Db)
}

func CreateRouter(e *echo.Echo, Db *sql.DB) {
	setDb(Db)
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))
	e.Use(middleware.BodyLimit("20m"))
	e.POST("/login", func(ctx echo.Context) error {
		return login.GoLogin(ctx)
	})
	e.Static("/static", "assets")
	api := e.Group("/api")
	api.Use(middleware.JWT([]byte("Mn+O2")))
	api.Use(CheckToken)
	adminRouter.UseRouter(api, Db)
	filerouter.UseRouter(api, Db)
	articlerouter.UseRouter(api, Db)
	moodrouter.UseRouter(api, Db)
	diaryrouter.UseRouter(api, Db)
}
