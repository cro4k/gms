package middle

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"time"
)

func Logger(ctx *gin.Context) {
	start := time.Now()
	ctx.Next()
	info := fmt.Sprintf("| %3d | %-10v | %-16s |%6s | %s",
		ctx.Writer.Status(),
		time.Since(start),
		ctx.RemoteIP(),
		ctx.Request.Method,
		ctx.Request.RequestURI,
	)
	logrus.WithContext(ctx).Info(info)
}
