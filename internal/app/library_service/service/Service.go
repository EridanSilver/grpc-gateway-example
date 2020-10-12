package service

import (
	"context"
	"github.com/erizaver/grpc-gateway-example/pkg/pb/library"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) GetBook(ctx context.Context, req *library.GetBookRequest) (*library.GetBookResponse, error) {
	return &library.GetBookResponse{
		Data: "Thats a new book you asked for",
	}, nil
}
