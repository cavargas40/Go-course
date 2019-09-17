package main

import (
	"fmt"
	"net/http"
	"log"
)

func main(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Hello World from my web server with GO")
	})

	fmt.Println("The server is running in http://localhost:8080")

	server := http.ListenAndServe(":8080", nil)

	log.Fatal(server)
}