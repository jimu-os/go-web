package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/alipay"
	"github.com/go-pay/xlog"
	"github.com/jimu-server/logger"
	"webServer"
)

func main() {
	webServer.Engine.POST("/notify", func(context *gin.Context) {
		var notifyReq gopay.BodyMap
		var err error
		notifyReq, err = alipay.ParseNotifyToBodyMap(context.Request) // c.Request 是 gin 框架的写法
		if err != nil {
			xlog.Error(err)
			return
		}
		logger.Info("notifyReq: %+v", notifyReq)
	})
	webServer.Engine.Run(":8080") // 监听并在 0.0.0.0:8080 上启动服务
}
