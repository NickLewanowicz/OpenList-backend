package main

import (
	"net/http"
	//"github.com/graph-gophers/graphql-go"
)

//GraphQl is an endpoint to handle all graphql queries
func GraphQl(w http.ResponseWriter, r *http.Request) {
}

//GraphiQl is the enpoint for exploration of the query schema
func GraphiQl(w http.ResponseWriter, r *http.Request) {
	w.Write(page)
}
