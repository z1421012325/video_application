package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"video_application/server/models"
	"video_application/server/request"
	"video_application/server/response"
	"video_application/server/tools"
)

// 观看视频 or 点赞
func LookVideo(c *gin.Context) {
	var req request.LookVideo
	if err := response.BindResponse(c,&req);err!=nil{
		return
	}

	_,uid := GetRequestTokenANDUid(c)
	var link models.ViewVideo
	tools.ReqMapStruct(&req,&link)
	if link.LookOrLike(uid) != nil{
		c.JSON(
			401,
			response.NewResponseData(
				response.DatabaseErr,
				"操作失败!",
				nil,
				errors.New("点赞或者查询视频失败")))
		return
	}

	if req.LikeVideo != 0 {
		c.JSON(
			200,
			response.NewResponseData(
				response.SUCCESS_CODE,
				"SUCCESS",
				nil,
				nil))
	}else {
		video := models.ResultVideo(req.Vid)
		c.JSON(
			200,
			response.NewResponseData(
				response.SUCCESS_CODE,
				"SUCCESS",
				video,
				nil))
	}
}

// 评论
func UserComment(c *gin.Context){
	var req request.Comment
	if err := response.BindResponse(c,&req);err!=nil{
		return
	}

	_,uid := GetRequestTokenANDUid(c)
	var comment models.Comments





}



// 点赞评论

// 相似视频
