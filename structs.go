package main

import (
	"fmt"
	"strconv"
)

//List struct
type List struct {
	ID    string `json:"id"`
	Owner string `json:"owner"`
	Title string `json:"title"`
	Date  int64  `json:"date"`
	Items []Task `json:"items"`
}

func (l List) insertSQL() string {
	return "INSERT INTO " + listTable + " VALUES ('" + l.ID + "','" + l.Owner + "','" + l.Title + "','" + strconv.Itoa(int(l.Date)) + "')"
}

func (l List) updateSQL() string {
	return "UPDATE " + listTable + " SET owner='" + l.Owner + "',title='" + l.Title + "',date='" + strconv.Itoa(int(l.Date)) + "' WHERE id='" + l.ID + "'"
}

//User struct
type User struct {
	ID        string `json:"id"`
	FirstName string `json:"first"`
	LastName  string `json:"last"`
	Email     string `json:"email"`
	Auth      string `json:"auth"`
	Token     string `json:"token"`
}

func (u User) insertSQL() {
	fmt.Printf("    - Inserting user '" + u.ID + "' into database. ")
	_, err = db.Exec(fmt.Sprintf("INSERT INTO %s VALUES (%s,%s,%s,%s,%s,%s)", userTable, u.ID, u.FirstName, u.LastName, u.Email, u.Auth, u.Token))
	didError(err)
}

func (u User) updateSQL() {
	fmt.Printf("    - Updating user '" + u.ID + "' in database. ")
	_, err = db.Exec(fmt.Sprintf("UPDATE %s SET first='%s' last='%s' email='%s' auth='%s' token='%s' WHERE id='%s'", userTable, u.FirstName, u.LastName, u.Email, u.Auth, u.Token, u.ID))
	didError(err)
}

//Task struct
type Task struct {
	ID          string   `json:"id"`
	Title       string   `json:"list"`
	Description string   `json:"description"`
	Created     string   `json:"created"`
	Due         string   `json:"due"`
	Owners      []User   `json:"owners"`
	Section     []string `json:"section"`
}

func (t Task) insertSQL() {
	fmt.Printf("    - Inserting task '" + t.ID + "' into database. ")
	_, err = db.Exec(fmt.Sprintf("INSERT INTO %s VALUES (%s,%s,%s,%s,%s,%s,%s)", taskTable, t.ID, t.Title, t.Description, t.Created, t.Due, t.Owners, t.Section))
	didError(err)
	for i := 0; i < len(t.Owners); i++ {
		fmt.Printf("    - Inserting task/owner relationship into database. ")
		_, err = db.Exec(fmt.Sprintf("INSERT INTO %s VALUES (%s,%s)", userTaskTable, t.ID, t.Owners[i].ID))
		didError(err)
	}
}

func (t Task) updateSQL() {
	fmt.Printf("    - Updating task '" + t.ID + "' in database. ")
	_, err = db.Exec(fmt.Sprintf("UPDATE %s SET title='%s' description='%s' created='%s' due='%s' WHERE id='%s'", taskTable, t.Title, t.Description, t.Created, t.Due, t.ID))
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

//TODO: Eventually replace hardcoded logic with struct to init db and tables
// type TableDec struct {
// 	name
// }
