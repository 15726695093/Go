package service

import (
	userv1 "clock-in/api/user/v1"
	"context"
	"log"
	"testing"

	"google.golang.org/grpc"
)

var ctx = context.TODO()
var client = newUserClient("127.0.0.1:9001")

func TestGetUserById(t *testing.T) {
	reply, err := client.GetUserById(ctx, &userv1.GetUserByIdRequest{
		Id: []int64{1, 2},
	})
	if err != nil {
		t.Error(err.Error())
	} else {
		log.Printf("%v", reply.User)
	}
}

func TestGetUserByName(t *testing.T) {
	reply, err := client.GetUserByName(ctx, &userv1.GetUserByNameRequest{
		Name: "malney1",
	})
	if err != nil {
		t.Error(err.Error())
	} else {
		log.Printf("%v", reply.User)
	}
}

func TestSearchUserByName(t *testing.T) {
	reply, err := client.SearchUserByName(ctx, &userv1.SearchUserByNameRequest{
		Name: "y",
	})
	if err != nil {
		t.Error(err.Error())
	} else {
		log.Printf("%v", reply.User)
	}
}

func TestRemoveUser(t *testing.T) {
	reply, err := client.RemoveUser(ctx, &userv1.RemoveUserRequest{
		Id: []int64{1, 2},
	})
	if err != nil {
		t.Error(err.Error())
	}
	log.Printf("%v", reply)
}

func TestSaveUser(t *testing.T) {
	reply, err := client.SaveUser(ctx, &userv1.SaveUserRequest{
		User: &userv1.User{
			Name:     "test4",
			Password: "0000",
			Phone:    "123456",
		},
	})
	if err != nil {
		t.Error(err.Error())
	} else {
		log.Printf("%v", reply.User)
	}
}

func newUserClient(addr string) userv1.UserServiceClient {
	conn, err := grpc.Dial(
		addr,
		grpc.WithInsecure(),
	)
	if err != nil {
		panic(err.Error())
	}
	return userv1.NewUserServiceClient(conn)
}
