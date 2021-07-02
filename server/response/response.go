package response

import (
	"github.com/gin-gonic/gin"
	"video_application/server/configure"
)

type ResData struct {
	Code int
	Msg string
	Error string
	Data interface{}
}


func NewResponse(code int,msg string,data interface{},err error) (res ResData ){

	debug := configure.Config.DebugMode
	if gin.DebugMode != debug {
		res.Data = data
		res.Code = code
		res.Msg = msg
	} else {
		res.Error = err.Error()
	}
	return res
}