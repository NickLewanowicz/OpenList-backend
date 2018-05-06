package main

import (
	"encoding/json"
	"fmt"
	"io"
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
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 999999))
	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &list); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	fmt.Println(body, list)
}

//UpdateList : update the information of a list on the db
func UpdateList(w http.ResponseWriter, r *http.Request) {

}

//DeleteList : delete the list selected from the database
func DeleteList(w http.ResponseWriter, r *http.Request) {

}
