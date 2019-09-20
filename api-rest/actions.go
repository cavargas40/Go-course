package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var collection = getMongoSession().DB("GoCourse").C("movies")

var movies = Movies{
	Movie{"Limitless", 2013, "Unknown"},
	Movie{"Batman Begins", 1999, "Unknown"},
	Movie{"Fast & Furious", 1999, "Your Mum"},
}

func getMongoSession() *mgo.Session {
	session, err := mgo.Dial("mongodb://localhost")

	if err != nil{
		panic(err)
	}

	return session
}

func responseMovie(w http.ResponseWriter, status int, results Movie) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(results)
}

func responseMovies(w http.ResponseWriter, status int, results []Movie) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(results)
}


func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World from my web server with GO")
}

func MoviesList(w http.ResponseWriter, r *http.Request) {
	var results []Movie
	err := collection.Find(nil).Sort("-_id").All(&results)

	if err != nil {
		log.Fatal(err)
	}else{
		fmt.Println("Results", results)
	}

	responseMovies(w, 200, results)
}

func MovieShow(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movieId := params["id"]

	if !bson.IsObjectIdHex(movieId) {
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(movieId)

	results := Movie{}
	err := collection.FindId(oid).One(&results)

	if err != nil {
		w.WriteHeader(404)
		return
	} else {
		responseMovie(w, 200, results)
	}

}

func Contact(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "This is the Contact Page")
}

func MovieAdd(w http.ResponseWriter, r *http.Request){
	decoder := json.NewDecoder(r.Body)

	var movieData Movie
	err := decoder.Decode(&movieData)

	if err != nil {
		panic(err)
	}

	defer r.Body.Close()

	err = collection.Insert(movieData)

	if err != nil {
		w.WriteHeader(500)
		return
	}
	
	responseMovie(w, 200, movieData)
}

func MovieUpdate(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movieId := params["id"]

	if !bson.IsObjectIdHex(movieId) {
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(movieId)
	decoder := json.NewDecoder(r.Body)

	var movieData Movie
	err := decoder.Decode(&movieData)
	if err != nil {
		panic(err)
		w.WriteHeader(500)
		return
	} 

	defer r.Body.Close()

	document := bson.M{"_id": oid}
	change := bson.M{"$set": movieData}
	err = collection.Update(document, change)

	if err != nil {
		panic(err)
		w.WriteHeader(500)
		return
	} 
	
	responseMovie(w, 200 , movieData)
}

type Message struct {
	Status string `json:"status"`
	Message string `json:"message"`
}


//binding of object to methods
func (this *Message) setStatus(data string){
	this.Status = data
}

func (this *Message) setMessage(data string){
	this.Message = data
}


func MovieRemove(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movieId := params["id"]

	if !bson.IsObjectIdHex(movieId) {
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(movieId)

	err := collection.RemoveId(oid)

	if err != nil {
		w.WriteHeader(404)
		return
	} 

	//results := Message{"success", "The movie with ID "+ movieId +" has been deleted correctly"}
	message := new(Message)
	message.setStatus("success")
	message.setMessage("The movie with ID "+ movieId +" has been deleted correctly")

	results := message

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(results)
}

