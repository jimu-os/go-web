package user

import (
	"api/config"
	"common/auth"
	"common/web"
)

func init() {
	api := web.Engine.Group("/api/user", auth.Authorization(config.Evn.Auth.AccessSecret))
	api.GET("/info", Info)
}
