package user

import (
	"api/db"
	"api/model"
	"common/auth"
	"common/resp"
	"common/web"
	"github.com/gin-gonic/gin"
)

// Info
// @Summary      用户信息
// @Description  获取当前登录用户的详细信息
// @Tags         用户
// @Security Jwt
// @Produce      json
// @Accept       json
// @Success      200  {object}  resp.Response{code=int,data=model.User,msg=string}
// @Failure      500  {object}  resp.Response{code=int,data=any,msg=string}
// @Router       /api/user/info [get]
func Info(c *gin.Context) {
	userId := c.MustGet(auth.UserId).(string)
	user := model.User{Id: userId}
	db.DB.Find(&user)
	resp.SUCCESS(c, user)
}

// UpdateUserInfo
// @Summary      更新用户信息
// @Description  更新当前登录用户的详细信息
// @Tags         用户
// @Security Jwt
// @Produce      json
// @Accept       json
// @Param        user  body  model.User true  "用户信息"
// @Success      200  {object}  resp.Response{code=int,data=any,msg=string}
// @Failure      500  {object}  resp.Response{code=int,data=any,msg=string}
// @Router       /api/user/info [put]
func UpdateUserInfo(c *gin.Context) {
	userId := c.MustGet(auth.UserId).(string)
	user := web.BindJSON[*model.User](c)
	user.Id = userId
	db.DB.Save(&user)
	resp.SUCCESS(c, nil)
}
