package response

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"video_application/server/conf"
)

type resData struct {
	Code int
	Msg string
	Error string
	Data interface{}
}



// 通用返回请求数据
func NewResponseData(code int,msg string,data interface{},err error) (res resData){
	debug := conf.Config.DebugMode
	if gin.DebugMode != debug {
		res.Code = code
		res.Msg = msg
	} else {
		res.Error = err.Error()
	}
	b,_:= json.Marshal(data)
	res.Data = string(b)
	return res
}



// 参数映射错误返回
func BindResponse(c *gin.Context,req interface{}) (err error) {
	if err = c.ShouldBind(req); err != nil{
		c.JSON(
			201,
			NewResponseData(
				BindErr,
				"参数异常",
				nil,
				err))
	}
	return
}
