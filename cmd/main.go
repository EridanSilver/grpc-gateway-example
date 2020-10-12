package main

import (
	"flag"

	libService "github.com/erizaver/grpc-gateway-example/internal/app/library_service/service"
	userService "github.com/erizaver/grpc-gateway-example/internal/app/user_service/service"
	"github.com/erizaver/grpc-gateway-example/pkg/pb/library"
	"github.com/erizaver/grpc-gateway-example/pkg/pb/user"
	"google.golang.org/grpc"
	"net/http"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	gsrv := grpc.NewServer()

	userSrv := userService.NewService()
	err := user.RegisterUserServiceHandlerServer(ctx, mux, userSrv)
	user.RegisterUserServiceServer(gsrv, userSrv)

	libSrv := libService.NewService()
	err = library.RegisterLibServiceHandlerServer(ctx, mux, libSrv)
	library.RegisterLibServiceServer(gsrv, libSrv)

	if err != nil {
		return err
	}

	return http.ListenAndServe(":8080", mux)
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
