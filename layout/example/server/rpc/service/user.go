package service

import (
	"context"
	"github.com/cro4k/gms/layout/public/rpc/rpcmessage"
)

type ExampleUserService struct {
	rpcmessage.UnimplementedExampleUserServiceServer
}

var _ rpcmessage.ExampleUserServiceServer = new(ExampleUserService)

func (s *ExampleUserService) SayHello(ctx context.Context, req *rpcmessage.ExampleSayHelloRequest) (*rpcmessage.ExampleSayHelloResponse, error) {
	return &rpcmessage.ExampleSayHelloResponse{Message: "Hello, " + req.User.Name + "!"}, nil
}
