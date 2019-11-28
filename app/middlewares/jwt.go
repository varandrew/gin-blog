package middlewares

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/varandrew/gin-product/app/pkg/errno"
	"github.com/varandrew/gin-product/app/utils"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data interface{}

		code := errno.SUCCESS

		token := c.Query("token")
		if token != "" {
			claims, err := utils.ParseToken(token)
			if err != nil {
				code = errno.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = errno.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		} else {
			code = errno.INVALID_PARAMS
		}

		if code.Code != errno.SUCCESS.Code {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code.Code,
				"msg":  code.Message,
				"data": data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
