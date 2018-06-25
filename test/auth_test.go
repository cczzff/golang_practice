package test

import (
	"testing"
	"fmt"
	"google.golang.org/grpc"
	"golang_practice/core"
	"golang.org/x/net/context"
)

func TestAuthClient(t *testing.T) {
	conn, err := grpc.Dial("localhost:8088", grpc.WithInsecure())
	if err != nil {
		fmt.Println("Dial error: ", err)
		return
	}

	cli := core_auth.NewAuthServiceClient(conn)

	req := &core_auth.RegisterReq{}
	res, err := cli.Register(context.Background(), req)
	fmt.Println(res, err)
}
