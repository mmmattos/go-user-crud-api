package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
)

const (
	hosts      = "localhost:27017"
	database   = "db"
	username   = ""
	password   = ""
	collection = "users"
)

type MongoStore struct {
	session *mgo.Session
}

var mongoStore = MongoStore{}

func main() {

	fmt.Println("Starting Users API on port 8080...!!")
	//Create MongoDB session
	session := initialiseMongo()
	mongoStore.session = session

	r := mux.NewRouter()

	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/", getRoot).Methods(http.MethodGet)
	api.HandleFunc("/users", get).Methods(http.MethodGet)
	api.HandleFunc("/users", post).Methods(http.MethodPost)
	api.HandleFunc("/users/{id}", put).Methods(http.MethodPut)
	api.HandleFunc("/users/{id}", delete).Methods(http.MethodDelete)

	log.Fatal(http.ListenAndServe(":8080", r))
}

func initialiseMongo() (session *mgo.Session) {

	info := &mgo.DialInfo{
		Addrs:    []string{hosts},
		Timeout:  60 * time.Second,
		Database: database,
		Username: username,
		Password: password,
	}

	session, err := mgo.DialWithInfo(info)
	if err != nil {
		panic(err)
	}
	return
}
