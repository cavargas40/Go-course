package main

import (
	"fmt"
	"net/http"
	"log"
	"github.com/gorilla/mux"
)

func main(){
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/contact", Contact)

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
