package auth

import (
	"golang.org/x/net/context"
	"golang_practice/core/auth"
	"net"
	"fmt"
	"google.golang.org/grpc"
	"time"
	"github.com/golang/protobuf/ptypes"
)

type AuthServer struct {
	db        *DBModel
	authModel AuthModel
}

func NewAuthServer() (s *AuthServer, err error) {
	s = &AuthServer{}
	s.authModel = AuthModel{}
	s.db, err = NewDBModel("mysql", "root:123456@(localhost:3306)/gm3?charset=utf8")
	return
}

func (a *AuthServer) Register(ctx context.Context, request *core_auth.RegisterReq) (account *core_auth.Account, err error) {
	createAt, err := ptypes.TimestampProto(time.Now())
	if err != nil {
		return
	}



	account = &core_auth.Account{
		Username:  "cc",
		Password:  "ccc",
		CreatedAt: createAt,
	}

	return

}

func Run() {

	//起服务
	lis, err := net.Listen("tcp", ":"+"8088")
	if err != nil {
		fmt.Println("listen to tcp error", err)
	}
	s := grpc.NewServer()
	authServer, err :=  NewAuthServer()
	if err != nil{
		fmt.Println("new auth server error: ", err)
	}

	core_auth.RegisterAuthServiceServer(s, authServer)

	s.Serve(lis)

}
