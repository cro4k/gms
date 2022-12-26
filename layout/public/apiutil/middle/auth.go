package middle

import (
	"context"
	"github.com/cro4k/gms/layout/public/etcd"
	"github.com/cro4k/gms/layout/public/rpc/naming"
	"github.com/cro4k/gms/layout/public/rpc/rpcmessage"
	"github.com/cro4k/micro/discovery"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
	"time"
)

func Token(ctx *gin.Context) {
	tokenStr := strings.TrimPrefix(ctx.GetHeader("Authorization"), "Bearer ")
	if tokenStr == "" {
		tokenStr, _ = ctx.Cookie("token")
	}
	if tokenStr == "" {
		return
	}
	conn, err := discovery.Discover(etcd.CLI(), naming.ServiceAuthToken)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		logrus.WithContext(ctx).Error(err)
		return
	}
	auth := rpcmessage.NewAuthServiceClient(conn)
	timeout, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	token, err := auth.TokenVerify(timeout, &rpcmessage.TokenVerifyRequest{Token: tokenStr})
	//if err != nil {
	//	ctx.AbortWithStatus(http.StatusInternalServerError)
	//	logrus.WithContext(ctx).Error(err)
	//	return
	//}
	ctx.Set("token_error", err)
	ctx.Set("uid", token.Id)
}

func Auth(ctx *gin.Context) {
	err, _ := ctx.Get("token_error")
	id := ctx.GetString("uid")
	if err != nil || id == "" {
		ctx.AbortWithStatus(http.StatusForbidden)
	}
}
