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

	listColumns := "(id varchar(32), owner varchar(32), title varchar(32), date int(11))"

	fmt.Printf("    - Connected to DB ")
	if err != nil {
		fmt.Println("[FAILED]")
		panic(err.Error())
	}
	fmt.Println("[SUCCESS]")
	createDb(name)
	createTable("list", listColumns)
	defer db.Close()

}

func createDb(name string) {
	fmt.Printf("    - Creating and use '" + name + "' database. ")

	//Create given Database if it doesnt exist
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + name)
	if err != nil {
		fmt.Println("[FAILED]")
		panic(err)
	}
	fmt.Println("[SUCCESS]")

	//Select (use) given database
	_, err = db.Exec("USE " + name)
	if err != nil {
		panic(err)
	}
}

func createTable(name string, columns string) {
	fmt.Printf("    - Creating " + name + "tables if they dont exist ")
	//Create the necessary tables for openlist
	_, err = db.Exec("CREATE TABLE " + name + " " + columns)
	if err != nil {
		fmt.Println("[FAILED]")
		panic(err)
	}
	fmt.Println("[SUCCESS]")

}

// //SaveListToDb will save the provided list to the db
// func SaveListToDb(l List) {
// 	stmtSave, err := db.Prepare("INSERT INTO list VALUES (:id, :owner, :title, :date)")
// }
