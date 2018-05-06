package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/satori/go.uuid"
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
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(body, &list)
	if err != nil {
		panic(err)
	}
	list.ID = uuid.Must(uuid.NewV4()).String()
	list.Date = time.Now().Unix()
	SaveListToDb(list)
}

//UpdateList : update the information of a list on the db
func UpdateList(w http.ResponseWriter, r *http.Request) {

}

//DeleteList : delete the list selected from the database
func DeleteList(w http.ResponseWriter, r *http.Request) {

}
