package main

import (
	"context"
	"fmt"
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

//GetUser will get user with id
func (r *Resolver) GetUser(ctx context.Context, args struct{ Email string }) (*User, error) {
	return getUser(ctx, string(args.Email))
}

func getUser(ctx context.Context, Email string) (*User, error) {
	var user User
	fmt.Printf("Fetching user with ID '" + Email + "' from " + userTable)

	result, err := db.Query(fmt.Sprintf("SELECT id, first, last, email FROM %s where id='%s'", userTable, Email))
	didError(err)
	result.Next()

	err = result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email)
	didError(err)
	return &user, nil
}
