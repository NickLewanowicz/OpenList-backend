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
var dbName = "kickit"

var listTable = "List"
var listItemTable = "ListItem"
var userTable = "User"
var projTable = "Project"
var sectTable = "Section"
var taskTable = "Task"
var userTaskTable = "UserTask"
var userProjTable = "UserProject"
var taskSectTable = "TaskSection"

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

	//Users have __id__, first, last, email, auth (type), token (string)
	userCol := "(id varchar(40), first varchar(12), last varchar(12), email varchar(32), auth varchar(40), token varchar(40), PRIMARY KEY (id))"

	//Projects have __id__, title, created
	projCol := "(id varchar(40), title varchar(32), created int(11), PRIMARY KEY (id))"

	//Sections have __id__, title, **projectId**, position
	sectCol := "(id varchar(40), title varchar(32), projectId varchar(40), position int(2), PRIMARY KEY (id), FOREIGN KEY (projectId) REFERENCES Project(id))"

	//Tasks have __id__, title, desc, created, due
	taskCol := "(id varchar(40), title varchar(32), description varchar(1000), created int(11), due int(11), PRIMARY KEY (id))"

	userProjCol := "(userId varchar(40), projId varchar(40), FOREIGN KEY (userId) REFERENCES User(id), FOREIGN KEY (projId) REFERENCES Project(id))"
	userTaskCol := "(userId varchar(40), taskId varchar(40), FOREIGN KEY (userId) REFERENCES User(id), FOREIGN KEY (taskId) REFERENCES Task(id))"
	sectTaskCol := "(sectId varchar(40), taskId varchar(40), FOREIGN KEY (sectId) REFERENCES Section(id), FOREIGN KEY (taskId) REFERENCES Task(id))"

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
	createTable(userTable, userCol)
	createTable(projTable, projCol)
	createTable(sectTable, sectCol)
	createTable(taskTable, taskCol)
	createTable(userTaskTable, userTaskCol)
	createTable(userProjTable, userProjCol)
	createTable(taskSectTable, sectTaskCol)

	fmt.Printf("MySQL Database fully initialized\n\n")
}

func createDb(name string) {
	fmt.Printf("    - Creating '" + name + "' database. ")

	//Create given Database if it doesnt exist
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + name)
	didError(err)

	//Select (use) given database
	fmt.Printf("    - Use database '" + name + "'. ")
	_, err = db.Exec("USE " + name)
	didError(err)
}

func createTable(name string, columns string) {
	fmt.Printf("    - Creating " + name + " tables if they dont exist ")
	//Create the necessary tables for kickit
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS " + name + " " + columns)
	didError(err)

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
	didError(err)
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
	didError(err)
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
	didError(err)
}

//UpdateListInDb will take list and update it in db
func UpdateListInDb(list List) {
	list.Date = time.Now().Unix()
	fmt.Printf("Updating '" + list.Title + "' into List table ")
	fmt.Printf(list.updateSQL())
	_, err = db.Exec(list.updateSQL())
	didError(err)
}
