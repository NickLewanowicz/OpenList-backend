package main

import "time"

//List struct
type List struct {
	ID    string     `json:"id"`
	Owner string     `json:"owner"`
	Title string     `json:"title"`
	Date  time.Time  `json:"date"`
	Items []ListItem `json:"items"`
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
