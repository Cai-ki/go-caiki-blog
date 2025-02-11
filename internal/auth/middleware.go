package auth

import (
	"net/http"
	"strings"

	"github.com/Cai-ki/go-caiki-blog/utils"
	"github.com/gin-gonic/gin"
)

func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 注意写前端时，请求头Authorization的格式为Bearer tokenString
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

		ok, err := Jwt.ValidateClaimsExists(claims)

		if err != nil {
			utils.RespondWithJSON(c, http.StatusInternalServerError, "Internal Server Error")
		}

		if !ok {
			utils.RespondWithJSON(c, http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}

		c.Set("username", claims.Username)
		c.Set("email", claims.Email)
		c.Next()
	}
}
