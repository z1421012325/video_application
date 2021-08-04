package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"video_application/server/models"
	"video_application/server/request"
	"video_application/server/response"
)

// 搜索
func Search(c *gin.Context){
	var req request.Search
	if err := response.BindResponse(c,&req);err!=nil{
		return
	}

	videosList,err :=models.GetVideo(req.Key,req.Page,req.Number)
	if err != nil {
		c.JSON(
			200,
			response.NewResponseData(
				response.DatabaseErr,
				"搜索失败",
				nil,
				errors.New("搜索失败")))
	}

	c.JSON(
		200,
		response.NewResponseData(
			response.SUCCESS_CODE,
			"SUCCESS",
			videosList,
			nil))

}
