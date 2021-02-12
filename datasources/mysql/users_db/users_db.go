package users_db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var (
	Client *sql.DB
)

func init() {
	dataSource := "armin:Armin@godiz3011@tcp(127.0.0.1:3306)/users_db"
	var err error
	Client, err = sql.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	if err := Client.Ping(); err != nil {
		panic(err)
	}
	log.Print("dataBase successfully configured")
}
