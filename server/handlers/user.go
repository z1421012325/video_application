package handlers

import (
	"video_application/server/request"
	"video_application/server/response"

	"github.com/gin-gonic/gin"
)

// 登录
func UserLogin(c *gin.Context){
	var req request.UserLogin
	if err := c.ShouldBind(&req); err != nil{
		c.JSON(201,response.NewResponse(response.LoginCode,"登录参数异常",nil,err))
		return
	}


}

// 注册
func UserRegister(){}

// 退出
func UserOutlogin(){}

// 修改用户资料

// 修改密码




