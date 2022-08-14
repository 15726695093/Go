package service

import (
	v1 "clock-in/api/worktime/service/v1"
	"context"
	"log"
	"testing"

	"google.golang.org/grpc"
)

var (
	ctx    = context.TODO()
	client = newWorkServiceClient("127.0.0.1:9003")
)

func TestCreateWorkTime(t *testing.T) {
	_, err := client.CreateWorkTime(ctx, &v1.CreateWorkTimeRequest{
		User: 1,
		Record: []*v1.Record{
			{Day: 20220119, Moment: 1642644943},
			{Day: 20220119, Moment: 1642673743},
		},
	})
	if err != nil {
		t.Error(err.Error())
	}
}

func TestGetUserWorkTime(t *testing.T) {
	reply, err := client.GetUserWorkTime(ctx, &v1.GetUserWorkTimeRequest{
		User: 1,
		Day:  []int64{20220118, 20220119},
	})
	if err != nil {
		t.Error(err.Error())
	} else {
		log.Printf("%v", reply.Worktime)
	}
}

func newWorkServiceClient(addr string) v1.WorktimeServiceClient {
	conn, err := grpc.Dial(
		addr,
		grpc.WithInsecure(),
	)
	if err != nil {
		panic(err.Error())
	}
	return v1.NewWorktimeServiceClient(conn)
}
