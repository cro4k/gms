package api

import (
	"github.com/cro4k/doc/docer"
	"github.com/cro4k/gms/layout/example/server/api/controller"
	"github.com/cro4k/gms/layout/public/apiutil/middle"
	"github.com/gin-gonic/gin"
)

func (s *server) api(e *gin.Engine) {
	e.Use(middle.UUID)
	e.Use(middle.Logger)
	e.Use(middle.Token)
	
	router := docer.Wrap(e)
	router.POST("/api/hello", controller.SayHello)
	s.doc = router
}
