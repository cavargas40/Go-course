package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"log"
)

var movies = Movies{
	Movie{"Limitless", 2013, "Unknown"},
	Movie{"Batman Begins", 1999, "Unknown"},
	Movie{"Fast & Furious", 1999, "Your Mum"},
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World from my web server with GO")
}

func MoviesList(w http.ResponseWriter, r *http.Request) {

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

func MovieAdd(w http.ResponseWriter, r *http.Request){
	decoder := json.NewDecoder(r.Body)

	var movieData Movie
	err := decoder.Decode(&movieData)

	if(err != nil){
		panic(err)
	}

	defer r.Body.Close()

	log.Println(movieData)
	movies = append(movies, movieData)
}
