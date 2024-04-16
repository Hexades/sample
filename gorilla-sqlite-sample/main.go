package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hexades/gorilla"

	"github.com/hexades/samples/gorilla-sqlite-sample/models"
	orm "github.com/hexades/samples/gorilla-sqlite-sample/sqlite-orm"
)

func main() {

	gorilla.NewServer()

	gorilla.SendEvent(gorilla.NewEvent(gorilla.ServerStart("localhost:8080", 15, 15)))
	gorilla.SendEvent(gorilla.NewEvent(gorilla.HandlerFunc("/member/{id}", GetMemberHandler)))
	gorilla.SendEvent(gorilla.NewEvent(gorilla.HandlerFunc("/member", InsertMemberHandler)))
	gorilla.SendEvent(gorilla.NewEvent(gorilla.HandlerFunc("/shutdown", ShutdownHandler)))
	gorilla.SendEvent(gorilla.NewEvent(gorilla.HandlerFunc("/ping", gorilla.PingHandler)))

	orm.NewRepository()

	//TODO add signals for alive and quit.
	doKeepAlive()

}

var keepAlive = true

func doKeepAlive() {
	for keepAlive == true {
	}
	log.Println("Shutdown")

}

var GetMemberHandler = func(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	m := &models.Member{MemberId: id}
	err := orm.GetMember(m)
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
	m := &models.Member{}
	json.NewDecoder(r.Body).Decode(m)
	log.Println("Decoded: ", m)
	err := orm.Insert(m)
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
