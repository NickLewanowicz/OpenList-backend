package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func initDb() {
	var err error
	db, err = sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/openlistdb")

	if err != nil {
		fmt.Println("Failed to connect to DB")
		panic(err.Error())
	}
	defer db.Close()

	fmt.Println("Successfully connected to DB")
}

//SaveListToDb will save the provided list to the db
func SaveListToDb(l List) {
}
