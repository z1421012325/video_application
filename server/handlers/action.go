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

// 视频点击
func LookVideo(c *gin.Context) {
	var req request.LookVideo
	if err := response.BindResponse(c,&req);err!=nil{
		return
	}

	if req.LikeVideo != 0 {
		_ = cache.CacheVideoNewLook("video_look",1.0,req.Vid)
		c.JSON(
			200,
			response.NewResponseData(
				response.SUCCESS_CODE,
				"SUCCESS",
				nil,
				nil))
	}
}

// 增加评论
func UserComment(c *gin.Context){
	var req request.Comment
	if err := response.BindResponse(c,&req);err!=nil{
		return
	}

	if _,uid := GetRequestTokenANDUid(c); uid != req.FromUid{
		c.JSON(
			203,
			response.NewResponseData(
				response.IdentityErr,
				"评论异常",
				nil,
				errors.New("uid验证失败")))
	}

	// 判断回复类型
	if req.ComposeType == 0 {
		var comment models.Comments
		tools.ReqMapStruct(&req,&comment)
		if err := comment.AddComment(); err == nil{
			c.JSON(
				200,
				response.NewResponseData(
					response.SUCCESS_CODE,
					"评论成功",
					nil,
					nil))
		}
		return

	} else if req.ComposeType == 1 {
		var reply_commrnt models.ReplyComment
		tools.ReqMapStruct(&req,&reply_commrnt)
		if err :=  reply_commrnt.AddReplyComment(); err == nil{
			c.JSON(
				200,
				response.NewResponseData(
					response.SUCCESS_CODE,
					"评论成功",
					nil,
					nil))
		}
		return
	}

	c.JSON(
		201,
		response.NewResponseData(
			response.DatabaseErr,
			"评论异常",
			nil,
			errors.New("评论添加失败")))
}

// 删除评论
func DelComment(c *gin.Context){
	var req request.DelComment
	if err := response.BindResponse(c,&req);err!=nil{
		return
	}
	_,uid := GetRequestTokenANDUid(c)
	if err := models.DelComment(uid,req.Cid,req.IsReply);err!= nil{
		c.JSON(
			201,
			response.NewResponseData(
				response.DatabaseErr,
				"删除评论异常",
				nil,
				errors.New("删除评论异常")))
		return
	}
	c.JSON(
		201,
		response.NewResponseData(
			response.SUCCESS_CODE,
			"删除评论成功",
			nil,
			nil))

}

// 点赞(取消)评论,视频
func Like(c *gin.Context) {
	var req request.Like
	if err := response.BindResponse(c,&req);err!=nil{
		return
	}

	_,uid := GetRequestTokenANDUid(c)

	var like models.Likes
	tools.ReqMapStruct(&req,&like)

	if req.Status == 0 {
		if err := like.DelLike(uid); err != nil{
			c.JSON(
				201,
				response.NewResponseData(
					response.DatabaseErr,
					"取消点赞异常",
					nil,
					err))
		}
		return
	}else if req.Status == 1{
		if err := like.AddLike(uid); err != nil{
			c.JSON(
				200,
				response.NewResponseData(
					response.DatabaseErr,
					"点赞异常",
					nil,
					err))
		}
		return
	}

	c.JSON(
		200,
		response.NewResponseData(
			response.SUCCESS_CODE,
			"点赞成功",
			nil,
			nil))
}


// 相似视频
// 简单逻辑 根据 video.type  video.tag 推送类似视频
func SimilarVideo(c *gin.Context) {
	var req request.Simliar
	if err := response.BindResponse(c,&req);err!=nil{
		return
	}

	videoList := models.ReturnSimliarVideo(req.Vid)

	c.JSON(
		200,
		response.NewResponseData(
			response.SUCCESS_CODE,
			"success",
			videoList,
			nil))
}

