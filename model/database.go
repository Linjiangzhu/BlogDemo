package model

import (
	"database/sql"
	"fmt"

	// not get golint complain
	_ "github.com/go-sql-driver/mysql"
)

// InitDB return a sql database pointer
func InitDB() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/dbname")
	if err != nil {
		panic(err.Error())
	}
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("connected")
	}
	return db
}
