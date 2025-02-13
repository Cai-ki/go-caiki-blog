package routes

import (
	"github.com/Cai-ki/go-caiki-blog/internal/auth"
	"github.com/Cai-ki/go-caiki-blog/internal/comment"
	"github.com/Cai-ki/go-caiki-blog/internal/post"
	users "github.com/Cai-ki/go-caiki-blog/internal/user"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/api/users", users.RegisterHandler)
	r.POST("/api/auth/login", users.LoginHandler)
	r.GET("/api/users/:username", users.GetUserHandler)

	r.GET("/api/posts", post.ListPostsHandler)
	r.GET("/api/posts/:id", post.GetPostHandler)

	authPosts := r.Group("/api/auth/posts")
	authPosts.Use(auth.JwtMiddleware())
	authPosts.POST("/", post.CreatePostHandler)
	authPosts.PUT("/:id", post.UpdatePostHandler)
	authPosts.DELETE("/:id", post.DeletePostHandler)

	r.GET("/api/comments/:id", comment.ListCommentsHandler)
	authComments := r.Group("/api/auth/comments")
	authComments.Use(auth.JwtMiddleware())
	authComments.POST("/:id", comment.CreateCommentHandler)

}
