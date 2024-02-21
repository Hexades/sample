package main

import (
	"encoding/json"
	"log"
	"net/http"
	"reflect"

	"github.com/hexades/sqlite"

	"github.com/gorilla/mux"
	"github.com/hexades/gorilla"
)

func main() {
	sqlite.NewRepository()
	sqlite.SendEvent(sqlite.NewEvent("sample_sqlite.db", sqlite.BasicOpenFunc))

	gorilla.NewServer()

	gorilla.SendEvent(gorilla.NewEvent(gorilla.ServerStart("localhost:8080", 15, 15)))
	gorilla.SendEvent(gorilla.NewEvent(gorilla.HandlerFunc("/member/{id}", GetMemberHandler)))
	gorilla.SendEvent(gorilla.NewEvent(gorilla.HandlerFunc("/member", InsertMemberHandler)))
	gorilla.SendEvent(gorilla.NewEvent(gorilla.HandlerFunc("/shutdown", ShutdownHandler)))
	gorilla.SendEvent(gorilla.NewEvent(gorilla.HandlerFunc("/ping", gorilla.PingHandler)))

	go doKeepAlive()

}

func GetMember(member *Member) error {
	evt := sqlite.NewEvent(member, sqlite.ReadFirstFunc)
	sqlite.SendEvent(evt)
	response := evt.Receive()
	log.Println(reflect.TypeOf(response.Value))
	return response.Err
}

func Insert(member *Member) error {
	evt := sqlite.NewEvent(member, sqlite.BasicInsertFunc)
	log.Println("Sending insert: :", evt)
	sqlite.SendEvent(evt)
	response := evt.Receive()
	return response.Err
}

var keepAlive = true

func doKeepAlive() {
	for keepAlive == true {
	}
	log.Println("Shutdown")

}

var GetMemberHandler = func(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	m := &Member{MemberId: id}
	err := GetMember(m)
	if err == nil {
		bytes, _ := json.Marshal(m)
		w.Write(bytes)
		//w.WriteHeader(http.StatusOK)
	} else {
		//w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
	}
}
var InsertMemberHandler = func(w http.ResponseWriter, r *http.Request) {
	m := &Member{}
	json.NewDecoder(r.Body).Decode(m)
	log.Println("Decoded: ", m)
	err := Insert(m)
	if err == nil {
		bytes, _ := json.Marshal(m)
		w.Write(bytes)
		//w.WriteHeader(http.StatusNotFound)
	} else {
		//w.WriteHeader(http.StatusOK)
		w.Write([]byte(err.Error()))
	}
}

var ShutdownHandler = func(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	rsp := "Shutdown command"
	log.Println(rsp)
	w.Write([]byte(rsp))
	keepAlive = false
}
