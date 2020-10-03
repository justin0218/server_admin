package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server_admin/pkg/jwt"
	"server_admin/pkg/resp"
)

func VerifyToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("TOKEN")
		uid, err := jwt.VerifyToken(token)
		if err != nil {
			resp.RespCode(c, http.StatusUnauthorized, "未授权")
			c.Abort()
			return
		}
		c.Set("uid", uid)
		c.Next()
	}
}
