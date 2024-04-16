package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/hexades/samples/gorilla-sqlite-sample/models"
	"github.com/stretchr/testify/assert"
)

func TestSampleSuite(t *testing.T) {
	go main()
	_ = os.Remove("sample_sqlite.db")
	//wait for setup complete...
	time.Sleep(8 * time.Second)
	resp, err := http.Get("http://localhost:8080/ping")
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	log.Println(resp.StatusCode)
	assert.NotEqual(t, resp.StatusCode, 404)

	member := &models.Member{
		MemberId: "0",
		First:    "Fool",
		Last:     "Hardy",
	}
	var buf bytes.Buffer
	err = json.NewEncoder(&buf).Encode(member)
	if err != nil {
		log.Fatal(err)
	}

	resp, err = http.Post("http://localhost:8080/member", "application/json", &buf)
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	//
	///
	//curl -X POST http://localhost:8080/member  -H "Content-Type: application/json" -d {"member_id":"0","first":"Fool","last":"Hardy"}

	resp, err = http.Get("http://localhost:8080/member/0")
	assert.Nil(t, err)
	assert.NotNil(t, resp)

	m := &models.Member{}
	json.NewDecoder(resp.Body).Decode(m)

	assert.Equal(t, member, m)
	log.Println("Round trip insert and get of: ", m)

	_, _ = http.Get("http://localhost:8080/shutdown")

}
