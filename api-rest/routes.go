package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{"Index", "GET", "/", Index},
	Route{"MoviesList", "GET", "/movies", MoviesList},
	Route{"MovieShow", "GET", "/movie/{id}", MovieShow},
	Route{"Contact", "GET", "/contact", Contact},
	Route{"MovieAdd", "POST", "/movie", MovieAdd},
	Route{"MovieUpdate", "PUT", "/movie/{id}", MovieUpdate},
	Route{"MovieRemove", "DELETE", "/movie/{id}", MovieRemove},
}

func CustomRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		router.
			Name(route.Name).
			Methods(route.Method).
			Path(route.Pattern).
			Handler(route.HandlerFunc)
	}

	return router
}
