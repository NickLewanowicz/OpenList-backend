package main

import (
	"net/http"
	//"github.com/graph-gophers/graphql-go"
)

//Resolver is a type for functions which return query data
type Resolver struct{}

//GraphQl is an endpoint to handle all graphql queries
func GraphQl(w http.ResponseWriter, r *http.Request) {

}

//GraphiQl is the enpoint for exploration of the query schema
func GraphiQl(w http.ResponseWriter, r *http.Request) {
	w.Write(page)
}
