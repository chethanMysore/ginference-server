package middlewares

import (
	"example/ginference-server/utils/token"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearerToken := c.GetHeader("Authorization")
		tokenStr := ""
		if len(strings.Split(bearerToken, " ")) == 2 {
			tokenStr = strings.Split(bearerToken, " ")[1]
		} else {
			c.String(http.StatusUnauthorized, "Unauthorized - Please provide valid bearer token")
			c.Abort()
			return
		}
		_, err := token.ValidateToken(tokenStr)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized - Please provide valid bearer token")
			c.Abort()
			return
		}
		c.Next()
	}
}
