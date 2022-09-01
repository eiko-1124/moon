package sql

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"database/sql"
)

//定义数据库连接信息
const (
	database   = "dmeiko_DB"
	dbName     = "dmeiko"
	dbAddress  = "120.79.73.206"
	dbPort     = 3306
	dbPassword = "#Eiko1124"
	charset    = "utf8"
)

func Connect() (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", dbName, dbPassword, dbAddress, dbPort, database, charset)
	Db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("mysql connect failed,detail is [%v]\n", err.Error())
		return nil, err
	}
	fmt.Println("connect success")
	return Db, nil
}
