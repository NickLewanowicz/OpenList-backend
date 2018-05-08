package main

import "strconv"

//List struct
type List struct {
	ID    string     `json:"id"`
	Owner string     `json:"owner"`
	Title string     `json:"title"`
	Date  int64      `json:"date"`
	Items []ListItem `json:"items"`
}

func (l List) insertSQL() string {
	return "INSERT INTO " + listTable + " VALUES ('" + l.ID + "','" + l.Owner + "','" + l.Title + "','" + strconv.Itoa(int(l.Date)) + "')"
}

func (l List) updateSQL() string {
	return "UPDATE " + listTable + " SET owner='" + l.Owner + "',title='" + l.Title + "',date='" + strconv.Itoa(int(l.Date)) + "' WHERE id='" + l.ID + "'"
}

//ListItem struct
type ListItem struct {
	ID   string `json:"id"`
	List string `json:"list"`
	Data string `json:"data"`
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

func (u User) insertSQL() string {
	return "INSERT INTO " + userTable + " VALUES ('" + u.ID + "','" + u.FirstName + "','" + u.LastName + "','" + u.Email + "','" + u.Auth + "','" + u.Token + "')"
}

func (u User) updateSQL() string {
	return "UPDATE " + userTable + " SET first='" + u.FirstName + "',last='" + u.LastName + "',email='" + u.Email + "',auth='" + u.Auth + "',token='" + u.Token + "' WHERE id='" + u.ID + "'"
}

//ex
// type User {
//     id: String
//     firstName: String
//     lastName: String
//     email: String
//     projects: [Project]
//   }

//   type Project {
//     id: String
//     title: String
//     created: Int
//     owners: [User]
//     sections: [Section]
//   }

//   type Section {
//     id: String
//     title: String
//     projectId: String
//     index: int
//     tasks: [Task]
//   }

//   type Task {
//     id: String
//     title: String
//     description: String
//     created: Int
//     due: Int
//     owners: [User]
//     sectionId: String
//   }

//TODO: Eventually replace hardcoded logic with struct to init db and tables
// type TableDec struct {
// 	name
// }
