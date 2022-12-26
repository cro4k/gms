package main

import (
	"context"
	"github.com/cro4k/gms/layout/example/internal/dao"
	"github.com/cro4k/gms/layout/example/server/api"
	"github.com/cro4k/gms/layout/example/server/rpc"
	_ "github.com/cro4k/gms/layout/public/logs"
	"github.com/cro4k/micro/runner"
	"github.com/sirupsen/logrus"
)

func main() {
	dao.Migrate()
	runner.Join(
		api.NewServer(),
		rpc.NewServer(),
	)
	runner.Run(func(err error) { logrus.Error(err) })
	runner.WaitSignal()
	runner.Shutdown(context.Background(), func(e error) {})
}
