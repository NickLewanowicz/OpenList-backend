package main

import (
	"fmt"
	// "github.com/leesper/couchdb-golang"
	// "encoding/json"
	// "strconv"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//List Struct
type List struct {
	ID    string `json:"id,omitempty"`
	Owner string `json:"owner,omitempty"`
	Title string `json:"title,omitempty"`
	Date  string `json:"date,omitempty"`
}

//ListItem Struct
type ListItem struct {
	ID   string `json:"id,omitempty"`
	List string `json:"list,omitempty"`
	Data string `json:"data,omitempty"`
}

func main() {
	//Init Router
	fmt.Println("Server starting up...")
	r := mux.NewRouter()

	//Route Handlers
	r.HandleFunc("/v1/api/list", getLists).Methods("GET")
	r.HandleFunc("/v1/api/list/{id}", getList).Methods("GET")
	r.HandleFunc("/v1/api/list/{id}", createList).Methods("PUT")
	r.HandleFunc("/v1/api/list/{id}", updateList).Methods("GET")
	r.HandleFunc("/v1/api/list/{id}", deleteList).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))

}
