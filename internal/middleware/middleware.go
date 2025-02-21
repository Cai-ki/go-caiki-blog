package middleware

import (
	"net/http"
	"strings"

	"github.com/Cai-ki/go-caiki-blog/pkg/jwt"
	"github.com/Cai-ki/go-caiki-blog/utils"
	"github.com/gin-gonic/gin"
)

var Jwt = jwt.Jwt

func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 注意写前端时，请求头Authorization的格式为Bearer tokenString
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			utils.RespondWithError(c, http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == "" {
			utils.RespondWithError(c, http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}

		claims, err := Jwt.ParseToken(tokenString)
		if err != nil {
			utils.RespondWithError(c, http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}

		ok, err := Jwt.ValidateClaimsExists(claims)

		if err != nil {
			utils.RespondWithError(c, http.StatusInternalServerError, "Internal Server Error")
		}

		if !ok {
			utils.RespondWithError(c, http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}

		c.Set("user_id", claims.ID)
		c.Set("username", claims.Username)
		c.Set("email", claims.Email)
		c.Next()
	}
}
