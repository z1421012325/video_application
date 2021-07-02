package service




import (
	"github.com/gin-gonic/gin"
)


// 启动app
func RunService(addr string) (err error) {

	//if !configure.Config.DebugMode {
	//	gin.SetMode(gin.ReleaseMode)
	//}

	server := gin.Default()

	// 加点插件 todo
	server.Use()
	err = server.Run(addr)
	return err
}