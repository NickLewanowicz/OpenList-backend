package main

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	uuid "github.com/satori/go.uuid"
)

var db *sql.DB
var err error
var dbName = "openlist"
var listTable = "List"
var listItemTable = "ListItems"

func initDb() {
	attempts := 1
	for err != nil || attempts < 2 {
		fmt.Println("Attempting to Initialize MySQL Database [" + strconv.Itoa(attempts) + "/50]")
		db, err = sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/")
		_, err = db.Query("SHOW DATABASES")
		if attempts > 50 {
			break
		}
		attempts++
		time.Sleep(time.Second)
	}

	listColumns := "(id varchar(40), owner varchar(32), title varchar(32), date int(11))"
	listItemColumns := "(id varchar(32), list varchar(32), data varchar(1000))"

	fmt.Printf("    - Connected to DB ")
	if err != nil {
		fmt.Println("[FAILED]")
		panic(err)
	} else {
		fmt.Println("[SUCCESS]")
	}

	//Create db and required tables that dont exist
	createDb(dbName)
	createTable(listTable, listColumns)
	createTable(listItemTable, listItemColumns)

	fmt.Printf("MySQL Database fully initialized\n\n")
}

func createDb(name string) {
	fmt.Printf("    - Creating '" + name + "' database. ")

	//Create given Database if it doesnt exist
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + name)
	if err != nil {
		fmt.Println("[FAILED]")
		panic(err)
	} else {
		fmt.Println("[SUCCESS]")
	}

	//Select (use) given database
	fmt.Printf("    - Use database '" + name + "'. ")
	_, err = db.Exec("USE " + name)
	if err != nil {
		fmt.Println("[FAILED]")
		panic(err)
	} else {
		fmt.Println("[SUCCESS]")
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

//GetListsInDb will get all lists in db (eventually of a particular owner)
func GetListsInDb(owner string) []List {
	var lists []List
	fmt.Printf("Fetching all lists of '" + owner + "' from " + listTable)

	results, err := db.Query("SELECT id, owner, title, date FROM " + listTable)
	if err != nil {
		fmt.Println("[FAILED]")
		panic(err)
	} else {
		fmt.Println("[SUCCESS]")
	}
	for results.Next() {
		var list List

		err = results.Scan(&list.ID, &list.Owner, &list.Title, &list.Date)
		if err != nil {
			panic(err.Error())
		}

		lists = append(lists, list)
	}
	return lists
}

//GetListInDb will get a list by list id
func GetListInDb(id string) List {
	var list List
	fmt.Printf("Fetching list with id '" + id + "' from " + listTable)

	result, err := db.Query("SELECT id, owner, title, date FROM " + listTable + " WHERE id='" + id + "'")
	if err != nil {
		fmt.Println("[FAILED]")
		panic(err)
	} else {
		fmt.Println("[SUCCESS]")
	}
	result.Next()

	err = result.Scan(&list.ID, &list.Owner, &list.Title, &list.Date)
	if err != nil {
		panic(err.Error())
	}

	return list
}

//SaveListInDb will save the provided list to the db
func SaveListInDb(list List) {
	list.ID = uuid.Must(uuid.NewV4(), err).String()
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

//UpdateListInDb will take list and update it in db
func UpdateListInDb(list List) {
	list.Date = time.Now().Unix()
	fmt.Printf("Updating '" + list.Title + "' into List table ")
	fmt.Printf(list.updateSQL())
	_, err = db.Exec(list.updateSQL())
	if err != nil {
		fmt.Println("[FAILED]")
		panic(err)
	} else {
		fmt.Println("[SUCCESS]")
	}
}
