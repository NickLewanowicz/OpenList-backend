package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	//"github.com/nicklewanowicz/kickit-backend/db"
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	structs "github.com/nicklewanowicz/kickit-backend/structs"
)

func getSchema(path string) (string, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func main() {
	//Init Router
	fmt.Println("Server starting up...")
	s, err := getSchema("./schema.graphql")
	didError(err)

	schema := graphql.MustParseSchema(s, &structs.Resolver{})
	http.Handle("/query", &relay.Handler{Schema: schema})
	r := NewRouter()
	structs.InitDb()
	log.Fatal(http.ListenAndServe(":8000", r))

	log.Fatal(http.ListenAndServe(":8080", nil))
}

//didError will signal an error
func didError(err error) {
	if err != nil {
		fmt.Println("[FAILED]")
		panic(err)
	} else {
		fmt.Println("[SUCCESS]")
	}
}

var page = []byte(`
	<!DOCTYPE html>
	<html>
		<head>
			<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/graphiql/0.10.2/graphiql.css" />
			<script src="https://cdnjs.cloudflare.com/ajax/libs/fetch/1.1.0/fetch.min.js"></script>
			<script src="https://cdnjs.cloudflare.com/ajax/libs/react/15.5.4/react.min.js"></script>
			<script src="https://cdnjs.cloudflare.com/ajax/libs/react/15.5.4/react-dom.min.js"></script>
			<script src="https://cdnjs.cloudflare.com/ajax/libs/graphiql/0.10.2/graphiql.js"></script>
		</head>
		<body style="width: 100%; height: 100%; margin: 0; overflow: hidden;">
			<div id="graphiql" style="height: 100vh;">Loading...</div>
			<script>
				function graphQLFetcher(graphQLParams) {
					return fetch("/query", {
						method: "post",
						headers: {
							'Accept': 'application/json',
							'Content-Type': 'application/json',
							'Authorization': 'Bearer f66aa7b3-add4-4fa2-88c8-534b2905c6ca'
						},
						body: JSON.stringify(graphQLParams),
						credentials: "include",
					}).then(function (response) {
						return response.text();
					}).then(function (responseBody) {
						try {
							return JSON.parse(responseBody);
						} catch (error) {
							return responseBody;
						}
					});
				}
				ReactDOM.render(
					React.createElement(GraphiQL, {fetcher: graphQLFetcher}),
					document.getElementById("graphiql")
				);
			</script>
		</body>
	</html>
	`)
