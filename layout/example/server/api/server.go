package api

import (
	"context"
	"fmt"
	"github.com/cro4k/doc/docer"
	"github.com/cro4k/doc/export/markdown"
	"github.com/cro4k/gms/layout/example/config"
	"github.com/cro4k/gms/layout/public/global"
	"github.com/cro4k/micro/runner"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

type server struct {
	srv *http.Server
	doc *docer.Engine
}

func (s *server) Run() error {
	e := gin.New()

	s.api(e)
	s.document(e)
	s.srv = &http.Server{
		Addr:    fmt.Sprintf(":%d", config.APIPort),
		Handler: e,
	}
	logrus.Infoln("api server listen on", s.srv.Addr)
	return s.srv.ListenAndServe()
}

func (s *server) Shutdown(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}

func (s *server) document(e *gin.Engine) {
	if global.C().Produce() {
		return
	}
	var dir = "doc"
	_ = os.RemoveAll(dir)
	err := markdown.Export(dir, s.doc.Docs().Decode().Group(), true)
	if err != nil {
		logrus.Error(err)
	}
	e.StaticFS("/doc", http.Dir(dir))
}

func NewServer() runner.Runner {
	return new(server)
}
