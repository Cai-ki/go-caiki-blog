package routes

import (
	users "github.com/Cai-ki/go-caiki-blog/internal/user"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/api/users", users.RegisterHandler)
	r.POST("/api/auth/login", users.LoginHandler)
	r.GET("/api/users/:username", users.GetUserHandler)
}
