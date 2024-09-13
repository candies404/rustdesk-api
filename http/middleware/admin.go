package middleware

import (
	"Gwen/http/response"
	"Gwen/service"
	"github.com/gin-gonic/gin"
)

// AdminAuth 后台权限验证中间件
func AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		//测试先关闭
		token := c.GetHeader("api-token")
		if token == "" {
			response.Fail(c, 403, "请先登录")
			c.Abort()
			return
		}
		user := service.AllService.UserService.InfoByAccessToken(token)
		if user.Id == 0 {
			response.Fail(c, 403, "请先登录")
			c.Abort()
			return
		}

		c.Set("curUser", user)
		c.Set("token", token)

		c.Next()
	}
}