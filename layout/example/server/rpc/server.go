package rpc

import (
	"context"
	"fmt"
	"github.com/cro4k/gms/layout/example/config"
	"github.com/cro4k/gms/layout/example/server/rpc/service"
	"github.com/cro4k/gms/layout/public/etcd"
	"github.com/cro4k/gms/layout/public/rpc/naming"
	"github.com/cro4k/gms/layout/public/rpc/rpcmessage"
	"github.com/cro4k/micro/registry"
	"github.com/cro4k/micro/runner"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

type server struct {
	srv *grpc.Server
}

func (s *server) Run() error {
	s.srv = grpc.NewServer()
	rpcmessage.RegisterExampleUserServiceServer(s.srv, new(service.ExampleUserService))
	if err := s.register("user"); err != nil {
		return err
	}
	listen := fmt.Sprintf("127.0.0.1:%d", config.RPCPort)
	addr, err := net.ResolveTCPAddr("tcp", listen)
	if err != nil {
		return err
	}
	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return err
	}
	logrus.Infoln("api server listen on", listen)
	return s.srv.Serve(listener)
}

func (s *server) register(service ...string) error {
	for _, v := range service {
		serviceName := fmt.Sprintf("%s/%s", naming.ServiceExample, v)
		serviceAddr := fmt.Sprintf("%s:%d", config.RPCHost, config.RPCPort)
		_, err := registry.Register(etcd.CLI(), serviceName, serviceAddr)
		if err != nil {
			return err
		}
		logrus.Infof("rpc service %s registered on %s", serviceName, serviceAddr)
	}
	return nil
}

func (s *server) Shutdown(ctx context.Context) error {
	s.srv.GracefulStop()
	return nil
}

func NewServer() runner.Runner {
	return new(server)
}
