package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"video_application/server/cache"
	"video_application/server/models"
	"video_application/server/request"
	"video_application/server/response"
	"video_application/server/tools"
)

// 上传视频
func UpVideo(c *gin.Context) {
	var req request.UpVideo
	if err := response.BindResponse(c,&req);err!=nil{
		return
	}

	// 根据业务需求 其中video 是需要第三方来支持存储,如aliyun oos
	/* 假设-- 前端上传视频时,先发送一个请求,含有数据 video标题,时长,分辨率,用户id... 等
		后端某个handler接受,生成aliyun oos凭证,返回能与aliyun oos 进行get(看上传进度),post(上传),del等操作的url
		前端直接将视频传输到oos中,等待完毕则进行该handle请求
	 */
	// 当然还有其他业务,比如video标题,video简介,video视频封面... 关键词检测
	// 由第三方或者自己写轮子
	// 该项目实际算是个demo 实际需求根据业务变化

	token,uid := GetRequestTokenANDUid(c)
	if req.Uid != uid && token == cache.CacheGet(string(req.Uid))  {
		c.JSON(
			401,
			response.NewResponseData(
				response.IdentityErr,
				"账号异常,无法上传视频,请重新登录!",
				nil,
				errors.New("uid和token错误")))
		return
	}

	req.CheckVideoType()

	var video models.Video
	tools.ReqMapStruct(&req,&video)
	if video.NewVideo(req.Uid) != nil {
		c.JSON(
			401,
			response.NewResponseData(
				response.DatabaseErr,
				"上传视频失败,请稍后再试!",
				nil,
				errors.New("数据库创建异常")))
		return
	}

	c.JSON(
		200,
		response.NewResponseData(
			response.SUCCESS_CODE,
			"上传视频成功!",
			nil,
			nil))
	return
}


// 删除视频
func DleVideo(c *gin.Context) {
	var req request.DelVideo
	if err := response.BindResponse(c,&req);err!=nil{
		return
	}

	token,uid := GetRequestTokenANDUid(c)
	if req.Uid != uid && token == cache.CacheGet(string(req.Uid))  {
		c.JSON(
			401,
			response.NewResponseData(
				response.IdentityErr,
				"账号异常,无法删除视频,请重新尝试!",
				nil,
				errors.New("uid和token错误")))
		return
	}

	var video models.Video
	tools.ReqMapStruct(&req,&video)
	if video.DelVideo() != nil {
		c.JSON(
			401,
			response.NewResponseData(
				response.DatabaseErr,
				"删除视频失败,请稍后再试!",
				nil,
				errors.New("视频删除异常")))
		return
	}

	c.JSON(
		200,
		response.NewResponseData(
			response.SUCCESS_CODE,
			"视频删除成功!",
			nil,
			nil))
	return
}

