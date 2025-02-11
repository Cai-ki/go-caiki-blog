package auth

import (
	"net/http"
	"strings"

	"github.com/Cai-ki/go-caiki-blog/utils"
	"github.com/gin-gonic/gin"
)

func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			utils.RespondWithJSON(c, http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == "" {
			utils.RespondWithJSON(c, http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}

		claims, err := Jwt.ParseToken(tokenString)
		if err != nil {
			utils.RespondWithJSON(c, http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}

		c.Set("username", claims.Username)
		c.Next()
	}
}
