package middleware

import (
	"api-gateway/util"
	"api-gateway/util/e"
	"time"

	"github.com/gin-gonic/gin"
)

// JWT token验证中间件
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}
		code = e.SUCCESS
		token := c.GetHeader("Authorization")
		if token == "" {
			code = e.ErrorAuthNotFound
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = e.ErrorAuthCheckTokenFail
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ErrorAuthCheckTokenTimeout
			}
		}
		if code != e.SUCCESS {
			c.JSON(200, gin.H{
				"code": code,
				"msg":  e.GetMsg(uint(code)),
				"data": data,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
