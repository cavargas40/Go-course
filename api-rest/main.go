package main

import (
	"fmt"
	"net/http"
	"log"
)

func main(){
	router := CustomRouter()

	fmt.Println("The server is running in http://localhost:8080")
	server := http.ListenAndServe(":8080", router)
	log.Fatal(server)
}