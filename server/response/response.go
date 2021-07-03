package response

import (
	"github.com/gin-gonic/gin"
	"video_application/server/conf"
)

type ResData struct {
	Code int
	Msg string
	Error string
	Data interface{}
}


func NewResponse(code int,msg string,data interface{},err error) (res ResData ){

	debug := conf.Config.DebugMode
	if gin.DebugMode != debug {
		res.Data = data
		res.Code = code
		res.Msg = msg
	} else {
		res.Error = err.Error()
	}
	return res
}




func BindResponse(c *gin.Context,req interface{}) (err error) {
	if err = c.ShouldBind(req); err != nil{
		c.JSON(
			201,
			NewResponse(
				BindErr,
				"参数异常",
				nil,
				err))
	}
	return
}
