package routes

import (
	"github.com/Cai-ki/go-caiki-blog/internal/auth"
	"github.com/Cai-ki/go-caiki-blog/internal/post"
	users "github.com/Cai-ki/go-caiki-blog/internal/user"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/api/users", users.RegisterHandler)
	r.POST("/api/auth/login", users.LoginHandler)
	r.GET("/api/users/:username", users.GetUserHandler)

	authPosts := r.Group("/api/auth/posts")
	authPosts.Use(auth.JwtMiddleware())
	authPosts.POST("/", post.CreatePostHandler)
	authPosts.PUT("/:id", post.UpdatePostHandler)
	authPosts.DELETE("/:id", post.DeletePostHandler)

	r.GET("/api/posts", post.ListPostsHandler)
	r.GET("/api/posts/:id", post.GetPostHandler)
}
