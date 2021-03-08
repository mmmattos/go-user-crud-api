package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "ROOT was called"}`))
}

func get(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET USERS!")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	col := mongoStore.session.DB(database).C(collection)
	var results []User

	err := col.Find(nil).All(&results)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	jsonString, err := json.Marshal(results)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//Set content-type http header
	w.Header().Set("content-type", "application/json")

	//Send back data as response
	w.Write(jsonString)
}

func post(w http.ResponseWriter, r *http.Request) {
	fmt.Println("POST USER!")

	col := mongoStore.session.DB(database).C(collection)

	//Retrieve body from http request
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		panic(err)
	}

	//Save data into User struct
	var _user User
	err = json.Unmarshal(b, &_user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	_user.Password = getHash([]byte(password))
	_user.ID = bson.NewObjectId()

	//Insert User into MongoDB
	err = col.Insert(_user)
	if err != nil {
		panic(err)
	}

	//Convert User struct into json
	jsonString, err := json.Marshal(_user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//Set content-type http header
	w.Header().Set("content-type", "application/json")

	//Send back data as response
	w.Write(jsonString)

}

func put(w http.ResponseWriter, r *http.Request) {

	col := mongoStore.session.DB(database).C(collection)

	//Retrieve body from http request
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		panic(err)
	}

	//Save data into User struct
	var _user User
	err = json.Unmarshal(b, &_user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	vars := mux.Vars(r)
	var idToUpdate string = vars["id"]

	if idToUpdate == "" {
		http.Error(w, err.Error(), 500)
		return
	}

	var updateMap map[string]string
	updateMap = make(map[string]string)
	if _user.Age != "" {
		updateMap["aga"] = string(_user.Age)
	}
	if _user.Address != "" {
		updateMap["address"] = _user.Address
	}
	if _user.Password != "" {
		updateMap["password"] = getHash([]byte(_user.Password))
	}

	// // Embed body in bson.M
	// finalbody, err := json.Marshal(b)
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// var finalbodymap map[string]interface{}
	// if err = json.Unmarshal(finalbody, &finalbodymap); err != nil {
	// 	log.Println(err)
	// }

	filter := bson.M{"_id": bson.ObjectIdHex(idToUpdate)}
	change := bson.M{"$set": updateMap}
	err = col.Update(filter, change)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//Set content-type http header
	w.Header().Set("content-type", "application/json")

	//Send back OK
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Uaer updated"}`))
}

func delete(w http.ResponseWriter, r *http.Request) {

	col := mongoStore.session.DB(database).C(collection)

	vars := mux.Vars(r)
	idToDelete := vars["id"]

	if idToDelete == "" {
		http.Error(w, "No ID provided to delete!", 400)
		return
	}

	err := col.Remove(bson.M{"_id": bson.ObjectIdHex(idToDelete)})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//Set content-type http header
	w.Header().Set("content-type", "application/json")

	//Send back OK
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Uaer removed"}`))
}
