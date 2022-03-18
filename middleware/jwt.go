package middleware

import (
	"github.com/gin-gonic/gin"
	"time"
	"todo_list/pkg/e"
	"todo_list/pkg/utils"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		code := e.SUCCESS
		//var data interface{}
		token := c.GetHeader("Authorization")
		if token == "" {
			code = 404
		} else {
			claims, err := utils.ParseToken(token)
			if err != nil {
				code = e.ErrorAuthCheckTokenFail //无权限
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ErrorAuthCheckTokenTimeout //token无效
			}
		}
		if code != e.SUCCESS {
			c.JSON(400, gin.H{
				"status": code,
				"msg":    "Token解析错误",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
