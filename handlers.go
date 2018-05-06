package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//GetLists : Get all lists
func GetLists(w http.ResponseWriter, r *http.Request) {

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

}

//DeleteList : delete the list selected from the database
func DeleteList(w http.ResponseWriter, r *http.Request) {

}
