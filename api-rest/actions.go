package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World from my web server with GO")
}

func MovieList(w http.ResponseWriter, r *http.Request) {
	movies := Movies{
		Movie{"Limitless", 2013, "Unknown"},
		Movie{"Batman Begins", 1999, "Unknown"},
		Movie{"Fast & Furious", 1999, "Your Mum"},
	}
	//fmt.Fprintf(w, "This is the Movies List")
	json.NewEncoder(w).Encode(movies)
}

func MovieShow(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movieId := params["id"]

	fmt.Fprintf(w, "Movie with id %s", movieId)
}

func Contact(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "This is the Contact Page")
}
