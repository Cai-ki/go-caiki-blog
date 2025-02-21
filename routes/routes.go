package routes

import (
	"github.com/Cai-ki/go-caiki-blog/internal/handler"
	"github.com/Cai-ki/go-caiki-blog/internal/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/api/users", handler.RegisterHandler)
	r.POST("/api/auth/login", handler.LoginHandler)
	r.GET("/api/users/:username", handler.GetUserHandler)

	r.GET("/api/posts", handler.ListPostsHandler)
	r.GET("/api/posts/:id", handler.GetPostHandler)

	authPosts := r.Group("/api/auth/posts")
	authPosts.Use(middleware.JwtMiddleware())
	authPosts.POST("/", handler.CreatePostHandler)
	authPosts.PUT("/:id", handler.UpdatePostHandler)
	authPosts.DELETE("/:id", handler.DeletePostHandler)

	r.GET("/api/comments/:id", handler.ListCommentsHandler)
	authComments := r.Group("/api/auth/comments")
	authComments.Use(middleware.JwtMiddleware())
	authComments.POST("/:id", handler.CreateCommentHandler)

	r.GET("/api/tags", handler.ListTagsHandler)

	r.GET("/test/tags/:id", handler.ListPostTagsHandler)
	r.POST("/test/tags/connect", handler.ConnectTagsHandler)
}
