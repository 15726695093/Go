package interfaces

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"testing"
)

var baseUrl = "http://127.0.0.1:8000"

func TestLogin(t *testing.T) {
	url := baseUrl + "/api/clockin/v1/login"
	data := map[string]interface{}{
		"username": "yuki",
		"password": "00001",
	}
	jsondata, err := json.Marshal(data)
	if err != nil {
		t.Error(err.Error())
	}
	reply, err := http.Post(url, "application/json", bytes.NewBuffer(jsondata))
	if err != nil {
		t.Error(err.Error())
	}
	data = map[string]interface{}{}
	if json.NewDecoder(reply.Body).Decode(&data); err != nil {
		t.Error(err.Error())
	}
	log.Printf("%v", data)
}

func TestRegister(t *testing.T) {
	url := baseUrl + "/api/clockin/v1/register"
	data := map[string]interface{}{
		"name":     "yuki",
		"password": "00001",
		"phone":    "0000",
	}
	jsondata, err := json.Marshal(data)
	if err != nil {
		t.Error(err.Error())
	}
	reply, err := http.Post(url, "application/json", bytes.NewBuffer(jsondata))
	if err != nil {
		t.Error(err.Error())
	}
	data = map[string]interface{}{}
	if json.NewDecoder(reply.Body).Decode(&data); err != nil {
		t.Error(err.Error())
	}
	log.Printf("%v", data)
}

func TestClockinOnWork(t *testing.T) {
	url := baseUrl + "/api/clockin/v1/onwork"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "MXx5dWtp")
	reply, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error(err.Error())
	}
	data := map[string]interface{}{}
	if json.NewDecoder(reply.Body).Decode(&data); err != nil {
		t.Error(err.Error())
	}
	log.Printf("%v", data)
}

func TestClockinOffWork(t *testing.T) {
	url := baseUrl + "/api/clockin/v1/offwork"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "MXx5dWtp")
	reply, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error(err.Error())
	}
	data := map[string]interface{}{}
	if json.NewDecoder(reply.Body).Decode(&data); err != nil {
		t.Error(err.Error())
	}
	log.Printf("%v", data)
}
