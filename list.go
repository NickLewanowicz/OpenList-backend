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
	return "INSERT INTO " + listTable + "VALUES ('" + l.ID + "','" + l.Owner + "','" + l.Title + "','" + strconv.Itoa(int(l.Date)) + "')"
}

func (l List) updateSQL() string {
	return "UPDATE " + listTable + " SET list id='" + l.ID + "',owner='" + l.Owner + "',title='" + l.Title + "',date='" + strconv.Itoa(int(l.Date)) + "'"
}

//ListItem struct
type ListItem struct {
	ID   string `json:"id"`
	List string `json:"list"`
	Data string `json:"data"`
}

//TODO: Eventually replace hardcoded logic with struct to init db and tables
// type TableDec struct {
// 	name
// }
