package auth

import (
	"golang.org/x/net/context"
	"golang_practice/core/auth"
	"net"
	"fmt"
	"google.golang.org/grpc"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"flag"
	"net/http"
	"golang_practice/util/protoutil"
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

func (s *AuthServer) Register(ctx context.Context, request *core_auth.RegisterReq) (account *core_auth.Account, err error) {
	account, err = s.db.addUser(request.Username, request.Password)
	return

}

func Run() {

	//起服务
	lis, err := net.Listen("tcp", ":"+"8209")
	if err != nil {
		fmt.Println("listen to tcp error", err)
	}
	s := grpc.NewServer()
	authServer, err := NewAuthServer()
	if err != nil {
		fmt.Println("new auth server error: ", err)
	}

	core_auth.RegisterAuthServiceServer(s, authServer)

	go s.Serve(lis)


	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard, &protoutil.JSONPb{OrigName: true}))
	opts := []grpc.DialOption{grpc.WithInsecure()}

	var (
		echoEndpoint = flag.String("echo_endpoint", "localhost:8209", "endpoint of YourService")
	)


	err = core_auth.RegisterAuthServiceHandlerFromEndpoint(ctx, mux, *echoEndpoint, opts)
	if err != nil {
		fmt.Println("gateway error!: ", err)
	}

	err = http.ListenAndServe(":8080", mux)

	if err!= nil {
		fmt.Println("ListenAndServe error!: ", err)
	}


}
