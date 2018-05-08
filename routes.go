package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

//Route Struct
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

//Routes Struct
type Routes []Route

//NewRouter Function
func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {

		var handler http.Handler

		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

var routes = Routes{
	Route{
		"GetLists",
		"GET",
		"/api/v1/lists",
		GetLists,
	},
	Route{
		"GetList",
		"GET",
		"/api/v1/lists/{id}",
		GetList,
	},
	Route{
		"CreateList",
		"POST",
		"/api/v1/lists",
		CreateList,
	},
	Route{
		"UpdateList",
		"PUT",
		"/api/v1/lists/{id}",
		UpdateList,
	},
	Route{
		"DeleteList",
		"DELETE",
		"/api/v1/lists/{id}",
		DeleteList,
	},
	Route{
		"GraphQl",
		"POST",
		"/api/graphql",
		GraphQl,
	},
	Route{
		"GraphQl",
		"GET",
		"/api/graphql",
		GraphQl,
	},
}
