package controller

import (
	"github.com/cro4k/doc/docer"
	"github.com/cro4k/ginx"
	"github.com/gin-gonic/gin"
)

type SayHelloRequest struct {
	ginx.Empty
	Name string `json:"name" doc:"must"`
}

type SayHelloResponse struct {
	Message string `json:"message"`
}

func sayHello(c *gin.Context) {
	ctx, err := ginx.With[SayHelloRequest](c).Bind()
	if err != nil {
		ctx.Logger().Error(err)
		ctx.FailError(err)
		return
	}
	ctx.OK(&SayHelloResponse{
		Message: "Hello " + ctx.Body.Name + "!",
	})
}

var SayHello = &docer.Doc{
	Name:    "Login",
	Req:     &SayHelloRequest{},
	Rsp:     &SayHelloResponse{},
	Header:  docer.JSONHeader,
	Handler: sayHello,
}
