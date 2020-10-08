package service

import (
	"context"
	"github.com/erizaver/grpc-gateway-example/pkg/pb/user"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) GetById(ctx context.Context, req *user.GetByIdRequest) (*user.GetByIdResponse, error) {
	return &user.GetByIdResponse{
		Data: "SomeRespData",
	}, nil
}
