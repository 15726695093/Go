package admin

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"testing"
)

var baseUrl = "http://127.0.0.1:8004"

func TestClockinOffWork(t *testing.T) {
	url := baseUrl + "/api/admin/clockin/v1/getWorkTime"
	data := map[string]interface{}{
		"user": []int64{1, 2},
		"day":  []int64{20220119, 20220118},
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
