package biz

import (
	"log"
	"testing"
	"time"
)

func TestGetDate2Int(t *testing.T) {
	log.Println(getDate2Int((time.Now())))
}
