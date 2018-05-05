package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func initDb(name string) {
	fmt.Println("Attempting to Initialize MySQL Database")
	db, err = sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/")
	createDb(name)

	fmt.Printf("    - Connected to DB ")
	if err != nil {
		fmt.Println("[FAILED]")
		panic(err.Error())
	}
	fmt.Println("[SUCCESS]")
	defer db.Close()

}

func createDb(name string) {
	fmt.Printf("    - Creating and use '" + name + "' database. ")
	_, err = db.Exec("CREATE DATABASE  IF NOT EXISTS " + name)
	if err != nil {
		fmt.Println("[FAILED]")
		panic(err)
	}
	fmt.Println("[SUCCESS]")

	_, err = db.Exec("USE " + name)
	if err != nil {
		panic(err)
	}
}

func createTable(name string) {
	_, err = db.Exec("CREATE TABLE " + name + " ( id integer, data varchar(32) )")
	if err != nil {
		panic(err)
	}
}

// //SaveListToDb will save the provided list to the db
// func SaveListToDb(l List) {
// 	stmtSave, err := db.Prepare("INSERT INTO list VALUES (:id, :owner, :title, :date)")
// }
