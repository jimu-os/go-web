package main

import (
	"api/config"
	"api/docs"
	"api/logs"
	"common/web"
	"fmt"
	"go.uber.org/zap"

	_ "api/controller/pub"
	_ "api/controller/user"
	_ "api/role"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	docs.SwaggerInfo.Title = "web-kit"
	docs.SwaggerInfo.Description = "api 文档"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
}

// @securityDefinitions.apikey Jwt
// @in header
// @name Authorization
func main() {
	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", config.Evn.Host, config.Evn.Port),
		Handler: web.Engine,
	}

	go func() {
		logs.Log.Info("server start docs url http://localhost:8080/swagger/index.html")
		if err := server.ListenAndServe(); err != nil {
			panic(err.Error())
		}
	}()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-signals:
		if err := zap.L().Sync(); err != nil {
			logs.Log.Error("sync zap log error", zap.Error(err))
		}
		if err := server.Close(); err != nil {
			logs.Log.Error("close server error", zap.Error(err))
		}
		logs.Log.Info("server shutdown")
	}

}

func ApiInfoInit() {
	getwd, _ := os.Getwd()
	logs.Log.Info(getwd)
}
