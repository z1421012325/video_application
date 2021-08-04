package handlers

import (
	"errors"

	"video_application/server/cache"
	"video_application/server/request"
	"video_application/server/response"
	"video_application/server/models"
	"video_application/server/tools"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/contrib/sessions"
)

// 登录
func UserLogin(c *gin.Context){
	var req request.UserLogin
	if err := response.BindResponse(c,&req);err!=nil{
		return
	}

	// 密码验证
	var user models.User
	tools.ReqMapStruct(&req,&user)
	user.GetUserPassword()
	if !tools.CheckPassword(user.PassWord,req.Password){
		c.JSON(
			201,
			response.NewResponseData(
				response.UserPassword,
				"账号密码错误",
				nil,
				nil))
		return
	}

	// 返回token和session
	token,ok := tools.NewToken(&user)
	if !ok {
		c.JSON(
			201,
			response.NewResponseData(
				response.TokenErr,
				"登录异常",
				nil,
				errors.New("tokne加密异常")))
		return
	}

	session := sessions.Default(c)
	session.Set("u_id", user.Uid)
	session.Set("token", token)
	_ = session.Save()
	_ = cache.CacheSet(string(user.Uid),token,86400 * 30)

	c.JSON(
		200,
		response.NewResponseData(
			response.SUCCESS_CODE,
			"登录成功",
			nil,
			nil))
}

// 注册
func UserRegister(c *gin.Context){
	var req request.UserRegister
	if err := response.BindResponse(c,&req);err!=nil{
		return
	}

	// 查询账号是否存在
	var user models.User
	tools.ReqMapStruct(&req,&user)
	user.GetUserName()
	if user.NickName != "" {
		c.JSON(
			201,
			response.NewResponseData(
				response.RegisterErr,
				"账号已存在",
				nil,
				errors.New("账号已存在")))
		return
	}

	// 注册账号
	user.PassWord = tools.EnCryptionPassword(user.PassWord)
	err := user.Registry()
	if err != nil {
		c.JSON(
			401,
			response.NewResponseData(
				response.RegisterErr,
				"账号注册异常",
				nil,
				err))
		return
	}

	user.GetUserName()
	token,_ := tools.NewToken(&user)
	setUidToken(c,string(user.Uid),token)

	c.JSON(
		200,
		response.NewResponseData(
			response.SUCCESS_CODE,
			"注册成功",
			nil,
			nil))
}

// 退出
func UserOutlogin(c *gin.Context){
	delUidToken(c)
	c.JSON(
		200,
		response.NewResponseData(
			response.SUCCESS_CODE,
			"登出成功",
			nil,
			nil))
}



// 修改用户资料
func ModifyUserData(c *gin.Context){
	var req request.ModifyUserData
	if err := response.BindResponse(c,&req);err!=nil{
		return
	}

	_,uid := GetRequestTokenANDUid(c)
	if req.Uid != uid {
		c.JSON(
			201,
			response.NewResponseData(
				response.IdentityErr,
				"账号异常,无法修改,请重新登录!",
				nil,
				errors.New("请求uid与session中uid不同")))
		return
	}

	var user models.User
	tools.ReqMapStruct(&req,&user)
	err := user.ModifyUserData()
	if err != nil {
		c.JSON(
			201,
			response.NewResponseData(
				response.ModifyErr,
				"修改用户资料异常",
				nil,
				err))
		return
	}
	c.JSON(
		200,
		response.NewResponseData(
			response.SUCCESS_CODE,
			"修改用户资料成功",
			nil,
			nil))
}

// 修改密码
func ModifyUserPassword(c *gin.Context){
	var req request.ModifyUserPassword
	if err := response.BindResponse(c,&req);err!=nil{
		return
	}

	token,uid := GetRequestTokenANDUid(c)
	if req.Uid != uid && token == cache.CacheGet(string(req.Uid)) &&
		req.NewPassword != req.OldPassword {
		c.JSON(
			201,
			response.NewResponseData(
				response.IdentityErr,
				"账号异常,无法修改,请重新登录!",
				nil,
				errors.New("uid和token错误")))
		return
	}

	var user models.User
	tools.ReqMapStruct(&req,&user)
	user.GetToUidUser()
	if !tools.CheckPassword(tools.EnCryptionPassword(user.PassWord),req.NewPassword){
		c.JSON(
			201,
			response.NewResponseData(
				response.UserPassword,
				"密码错误",
				nil,
				errors.New("新老密码错误")))
		return
	}

	user.PassWord = tools.EnCryptionPassword(req.NewPassword)
	if err := user.ModifyUserPassword(); err != nil {
		c.JSON(
			201,
			response.NewResponseData(
				response.ModifyErr,
				"账号异常,无法修改,稍后再试!",
				nil,
				errors.New("数据库修改密码错误")))
		return
	}
	// success
	c.JSON(
		200,
		response.NewResponseData(
			response.SUCCESS_CODE,
			"修改密码成功,请重新登录!",
			nil,
			nil))

	delUidToken(c)
}






// 从请求中找到token,uid
func GetRequestTokenANDUid(c *gin.Context) (token string,uid int64) {
	session := sessions.Default(c)
	token,uid = session.Get("token").(string) , session.Get("u_id").(int64)
	return
}


func delUidToken(c *gin.Context) {
	session := sessions.Default(c)
	_,uid := GetRequestTokenANDUid(c)
	session.Delete("u_id")
	session.Delete("token")
	_ = session.Save()

	_ = cache.CacheDel(string(uid))
}

func setUidToken(c *gin.Context,uid,token string) {
	session := sessions.Default(c)
	session.Set("u_id", uid)
	session.Set("token", token)
	_ = session.Save()

	_ = cache.CacheSet(uid,token,86400 * 30)
}


