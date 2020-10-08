package main

import (
	"flag"
	"github.com/erizaver/grpc-gateway-example/internal/app/user_service/service"
	"github.com/erizaver/grpc-gateway-example/pkg/pb/user"
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
	srv := service.NewService()

	err := user.RegisterUserServiceHandlerServer(ctx, mux, srv)
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
