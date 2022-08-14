package service

import (
	recordv1 "clock-in/api/record/v1"
	"context"
	"testing"

	"google.golang.org/grpc"
)

var ctx = context.TODO()
var client = newUserClient("127.0.0.1:9002")

func TestClockInOnWork(t *testing.T) {
	_, err := client.ClockInOnWork(ctx, &recordv1.ClockInOnWorkRequest{
		User: 13,
	})
	if err != nil {
		t.Error(err.Error())
	}
}

func TestClockOffOnWork(t *testing.T) {
	_, err := client.ClockInOffWork(ctx, &recordv1.ClockInOffWorkRequest{
		User: 13,
	})
	if err != nil {
		t.Error(err.Error())
	}
}

func newUserClient(addr string) recordv1.RecordServiceClient {
	conn, err := grpc.Dial(
		addr,
		grpc.WithInsecure(),
	)
	if err != nil {
		panic(err.Error())
	}
	return recordv1.NewRecordServiceClient(conn)
}
