package structs

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"
	//Will rework later
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error
var dbName = "kickit"

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

func didError(err error) {
	if err != nil {
		fmt.Println("[FAILED]")
		panic(err)
	} else {
		fmt.Println("[SUCCESS]")
	}
}
