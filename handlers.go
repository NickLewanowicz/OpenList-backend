package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//GetLists : Get all lists
func GetLists(w http.ResponseWriter, r *http.Request) {
	var lists = GetListsInDb("")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(lists); err != nil {
		panic(err)
	}
	return
}

//GetList : Get single list based on ID
func GetList(w http.ResponseWriter, r *http.Request) {

}

//CreateList : Create a list on the DB
func CreateList(w http.ResponseWriter, r *http.Request) {
	var list List
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(body, &list)
	if err != nil {
		panic(err)
	}
	SaveListInDb(list)
}

//UpdateList : update the information of a list on the db
func UpdateList(w http.ResponseWriter, r *http.Request) {
	//Eventually use the item in the URL rather then just body
	var list List
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(body, &list)
	if err != nil {
		panic(err)
	}
	UpdateListInDb(list)

}

//DeleteList : delete the list selected from the database
func DeleteList(w http.ResponseWriter, r *http.Request) {

}
