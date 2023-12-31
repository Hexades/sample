package main

import (
	"bytes"
	"encoding/json"
	"github.com/hexades/sample/app"
	"github.com/hexades/sample/app/models"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"os"
	"testing"
)

func TestSampleSuite(t *testing.T) {
	_ = os.Remove("sample_sqlite.db")
	go app.New()

	resp, err := http.Get("http://localhost:8080/ping")
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	log.Println(resp.StatusCode)

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
