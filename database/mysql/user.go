package user

import (
	"database/sql"
	"fmt"

	// mysql driver to connect to database
	_ "github.com/go-sql-driver/mysql"
)

var (
	// UserDB db returned after successful connection
	UserDB   *sql.DB
	userName = "root"
	password = "mypassword"
	host     = "127.0.0.1"
	port     = "3306"
	dbName   = "book-store-user"
)

func init() {
	var err error
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", userName, password, host, port, dbName)
	UserDB, err = sql.Open("mysql", dataSource)
	if err != nil {
		panic(err.Error())
	}
	if err := UserDB.Ping(); err != nil {
		panic(err.Error())
	}
	fmt.Println("data base successfully connected")
}
