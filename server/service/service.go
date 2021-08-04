package service




import (
	"github.com/gin-gonic/gin"
	"video_application/server/handlers"
	"video_application/server/mid"
)


// 启动app
func RunService(addr string) (err error) {

	//if !conf.Config.DebugMode {
	//	gin.SetMode(gin.ReleaseMode)
	//}

	server := gin.Default()
	// todo 全局中间件 跨域,session,限流
	// todo 根据变量设置gin mode

	// login register search ...  无需登录的handler
	server.POST("/login",handlers.UserLogin)
	server.POST("/register",handlers.UserRegister)

	// 加点插件 todo
	// 需要用户登录的handler
	server.Use(mid.VerifyUser())
	// up_video,del_video,comment ...




	err = server.Run(addr)
	return err
}