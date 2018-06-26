package test

import (
	"testing"
	"fmt"
	"google.golang.org/grpc"
	"golang.org/x/net/context"
	"golang_practice/core/auth"
)

func TestAuthClient(t *testing.T) {
	conn, err := grpc.Dial("localhost:8209", grpc.WithInsecure())
	if err != nil {
		fmt.Println("Dial error: ", err)
		return
	}

	cli := core_auth.NewAuthServiceClient(conn)

	req := &core_auth.RegisterReq{
		Username: "ll",
		Password: "hahaha",
	}
	res, err := cli.Register(context.Background(), req)
	fmt.Println(res, err)
}
