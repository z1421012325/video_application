package mid

import (
	"errors"
	"github.com/gin-gonic/gin"
	"video_application/server/cache"
	"video_application/server/response"
	"video_application/server/tools"
)

func VerifyUser() gin.HandlerFunc {

	return func(c *gin.Context) {

		StrToken := tools.GetRequestToken(c)	// 得到请求中的token
		uid := tools.GetRequestUid(c)
		redisToken := cache.RedisGet(string(uid))
		if StrToken != redisToken {
			respondWithError(
				201,
				response.NewResponseData(
					response.ValueCode,
					"账号异常",
					nil,
					errors.New("账号token与redis缓存不同")),
				c)
			return
		}
		c.Next()
	}
}

func respondWithError(code int, msg interface{}, c *gin.Context) {
	c.JSON(code, msg)
	c.Abort()
}
