package main

import "time"

//List struct
type List struct {
	ID    string     `json:"id,omitempty"`
	Owner string     `json:"owner,omitempty"`
	Title string     `json:"title,omitempty"`
	Date  time.Time  `json:"date"`
	Items []ListItem `json:"items"`
}

//ListItem struct
type ListItem struct {
	ID   string `json:"id,omitempty"`
	List string `json:"list,omitempty"`
	Data string `json:"data,omitempty"`
}
