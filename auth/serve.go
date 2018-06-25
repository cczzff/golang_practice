package auth

import (
	"golang.org/x/net/context"
	"golang_practice/core/auth"
	"net"
	"fmt"
	"google.golang.org/grpc"
)

type AuthServer struct {
	db DBModel
	authModel AuthModel
}

func (a *AuthServer) Register(ctx context.Context, request *core_auth.RegisterReq) (account *core_auth.Account, err error) {
	account = &core_auth.Account{
		Id:        "1",
		Username:  "cc",
		Password:  "ccc",
		CreatedAt: nil,
	}
	err = nil

	return

}

func Run(){

	//起服务
	lis, err := net.Listen("tcp", ":"+ "8765")
	if err != nil {
		fmt.Println("listen to tcp error", err)
	}
	s := grpc.NewServer()

	core_auth.RegisterAuthServiceServer(s, &AuthServer{})

	s.Serve(lis)

}



