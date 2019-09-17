package main

import (
	"fmt"
	"net/http"
	"log"
	"github.com/gorilla/mux"
	"encoding/json"
)

func main(){
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/contact", Contact)
	router.HandleFunc("/movies", MovieList)
	router.HandleFunc("/movie/{id}", MovieShow)

	fmt.Println("The server is running in http://localhost:8080")
	server := http.ListenAndServe(":8080", router)
	log.Fatal(server)
}

func Index(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello World from my web server with GO")
}

func Contact(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "This is the Contact Page")
}

func MovieList(w http.ResponseWriter, r *http.Request){
	movies := Movies{
		Movie{"Limitless", 2013, "Unknown"},
		Movie{"Batman Begins", 1999, "Unknown"},
		Movie{"Fast & Furious", 1999, "Your Mum"},
	}
	//fmt.Fprintf(w, "This is the Movies List")
	json.NewEncoder(w).Encode(movies)
}

func MovieShow(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	movieId := params["id"]

	fmt.Fprintf(w, "Movie with id %s", movieId)
}