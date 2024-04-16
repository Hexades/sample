package sqlite

import (
	"log"
	"reflect"

	"github.com/hexades/samples/gorilla-sqlite-sample/models"
	"github.com/hexades/sqlite"
)

func NewRepository() {
	sqlite.NewRepository()
	sqlite.SendEvent(sqlite.NewEvent("sample_sqlite.db", sqlite.BasicOpenFunc))

}

func GetMember(member *models.Member) error {
	evt := sqlite.NewEvent(member, sqlite.ReadFirstFunc)
	sqlite.SendEvent(evt)
	response := evt.Receive()
	log.Println(reflect.TypeOf(response.Value))
	return response.Err
}

func Insert(member *models.Member) error {
	evt := sqlite.NewEvent(member, sqlite.BasicInsertFunc)
	log.Println("Sending insert: :", evt)
	sqlite.SendEvent(evt)
	response := evt.Receive()
	return response.Err
}
