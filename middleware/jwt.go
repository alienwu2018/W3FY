package middleware

import (
	"github.com/gin-gonic/gin"
	"strings"
	"time"
	"w3fy/pkg/e"
	"w3fy/pkg/util"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var msg string
		var data interface{}

		code = e.OK
		token := c.GetHeader("Authorization")
		if token == "" {
			code = e.UNAUTHORIZED
		} else {
			if !strings.HasPrefix(token, "Bearer ") {
				code = e.UNAUTHORIZED
			} else {
				token = token[len("Bearer "):]
				claims, err := util.ParseToken(token)
				if err != nil {
					code = e.UNAUTHORIZED
					msg = "Invalid token"
				} else if time.Now().Unix() > claims.ExpiresAt {
					code = e.UNAUTHORIZED
					msg = "Invalid token"
				}
				c.Set("AuthData", claims)
			}
		}

		if code != e.OK {
			if msg == "" {
				msg = e.GetMsg(code)
			}
			c.JSON(code, gin.H{
				"code": code,
				"msg":  msg,
				"data": data,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
