package ingress

import (
	"github.com/hexades/gorilla"
	"log"
	"net/http"
	"os"
)

var keepAlive = true

func New() {
	gorilla.NewServer()
	gorilla.AddListener(new(ingress))
	gorilla.SendEvent(gorilla.NewEvent(gorilla.ServerStart("localhost:8080", 15, 15)))
	gorilla.SendEvent(gorilla.NewEvent(gorilla.HandlerFunc("/ping", gorilla.PingHandler)))
	gorilla.SendEvent(gorilla.NewEvent(gorilla.HandlerFunc("/ingress", IngressHandler)))
	gorilla.SendEvent(gorilla.NewEvent(gorilla.HandlerFunc("/shutdown", ShutdownHandler)))

	log.Println("Testing ping on loclahost:8080")
	resp, err := http.Get("http://localhost:8080/ping")
	log.Println(err)
	log.Println(resp.StatusCode)
	log.Println("Use the following in a browser or with curl")
	log.Println("localhost:8080/ping")
	log.Println("localhost:8080/ingress")
	log.Println("localhost:8080/shutdown")

	go doKeepAlive()

}

func doKeepAlive() {
	for keepAlive == true {
	}
	log.Println("Shutdown")
	os.Exit(0)

}

type ingress struct{}

// Note the OnEvent is not likely to be used...Here for testing...
func (i *ingress) OnEvent(eventChannel <-chan gorilla.Event) {
	for evt := range eventChannel {
		log.Println("Received event: ", evt)
	}
}

var IngressHandler = func(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	rsp := "Received on ingress:!"
	log.Println(rsp)
	w.Write([]byte(rsp))
}
var ShutdownHandler = func(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	rsp := "Shutdown command"
	log.Println(rsp)
	w.Write([]byte(rsp))
	keepAlive = false
}
