package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	uuid "github.com/satori/go.uuid"
)

var db *sql.DB
var err error
var listTable = "List"
var listItemTable = "ListItems"

func initDb(name string) {
	fmt.Println("Attempting to Initialize MySQL Database")
	db, err = sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/")

	listColumns := "(id varchar(40), owner varchar(32), title varchar(32), date int(11))"
	listItemColumns := "(id varchar(32), list varchar(32), data varchar(1000))"

	fmt.Printf("    - Connected to DB ")
	if err != nil {
		fmt.Println("[FAILED]")
		panic(err.Error())
	} else {
		fmt.Println("[SUCCESS]")
	}

	//Create db and required tables that dont exist
	createDb(name)
	createTable(listTable, listColumns)
	createTable(listItemTable, listItemColumns)

	fmt.Println("MySQL Database fully initialized")
}

func createDb(name string) {
	fmt.Printf("    - Creating and use '" + name + "' database. ")

	//Create given Database if it doesnt exist
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + name)
	if err != nil {
		fmt.Println("[FAILED]")
		panic(err)
	} else {
		fmt.Println("[SUCCESS]")
	}

	//Select (use) given database
	_, err = db.Exec("USE " + name)
	if err != nil {
		panic(err)
	}
}

func createTable(name string, columns string) {
	fmt.Printf("    - Creating " + name + " tables if they dont exist ")
	//Create the necessary tables for openlist
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS " + name + " " + columns)
	if err != nil {
		fmt.Println("[FAILED]")
		panic(err)
	} else {
		fmt.Println("[SUCCESS]")
	}

}

//FormatListForDb will take a List and format a string for the db.Prepare
func FormatListForDb(l List) string {
	return ""
}

//SaveListInDb will save the provided list to the db
func SaveListInDb(list List) {
	list.ID = uuid.Must(uuid.NewV4()).String()
	list.Date = time.Now().Unix()
	fmt.Printf("Inserting '" + list.Title + "' into List table ")
	fmt.Printf(list.insertSQL())
	_, err = db.Exec(list.insertSQL())
	if err != nil {
		fmt.Println("[FAILED]")
		panic(err)
	} else {
		fmt.Println("[SUCCESS]")
	}
}
