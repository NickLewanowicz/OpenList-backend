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
			Handler(route.HandlerFunc)
	}

	return router
}

var routes = Routes{
	Route{
		"GetLists",
		"GET",
		"/v1/api/lists",
		GetLists,
	},
	Route{
		"GetList",
		"GET",
		"/v1/api/lists/{id}",
		GetList,
	},
	Route{
		"CreateList",
		"POST",
		"/v1/api/lists/{id}",
		CreateList,
	},
	Route{
		"UpdateList",
		"PUT",
		"/v1/api/lists/{id}",
		UpdateList,
	},
	Route{
		"DeleteList",
		"DELETE",
		"/v1/api/lists/{id}",
		DeleteList,
	},
}
