package main

import (
	"fmt"

	"moon-manage-serve/router"

	"moon-manage-serve/sql"

	"github.com/labstack/echo"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	Db, err := sql.Connect()
	if err != nil {
		fmt.Println("DateBase Connect ERROR")
		return
	}
	e := echo.New()
	router.CreateRouter(e, Db)
	e.Logger.Fatal(e.Start(":1323"))
}
