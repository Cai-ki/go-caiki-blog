package post

import (
	"net/http"

	"github.com/Cai-ki/go-caiki-blog/utils"
	"github.com/gin-gonic/gin"
)

type createPostRequestInfo struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	// Tags    []string `json:"tags"`
}

type createPostResponseInfo struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	UserID    uint   `json:"user_id"`
	CreatedAt string `json:"created_at"`
}

func CreatePostHandler(c *gin.Context) {
	var req createPostRequestInfo
	if err := c.BindJSON(&req); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request body")
		return
	}

	userID := c.GetUint("user_id")

	post, err := Service.CreatePost(userID, req.Title, req.Content)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to create post")
		return
	}

	res := createPostResponseInfo{
		ID:        post.ID,
		Title:     post.Title,
		UserID:    post.UserID,
		CreatedAt: post.CreatedAt.Format("2006-01-02 15:04:05"),
	}

	utils.RespondWithJSON(c, http.StatusCreated, res)
}

type listPostsRequestInfo struct {
	Page  int `form:"page"`
	Limit int `form:"limit"`
}

type listPostsResponseInfo struct {
	Posts []postInfo `json:"posts"`
}

type postInfo struct {
	ID     uint   `json:"id"`
	Title  string `json:"title"`
	Author struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
	} `json:"author"`
	CreatedAt string `json:"created_at"`
}

func ListPostsHandler(c *gin.Context) {
	var req listPostsRequestInfo
	if err := c.BindQuery(&req); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request query")
		return
	}

	posts, err := Service.ListPosts(req.Page, req.Limit)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to list posts")
		return
	}

	res := listPostsResponseInfo{
		Posts: make([]postInfo, len(posts)),
	}

	for i, post := range posts {
		res.Posts[i] = postInfo{
			ID:    post.ID,
			Title: post.Title,
			Author: struct {
				ID   uint   `json:"id"`
				Name string `json:"name"`
			}{post.UserID, post.User.Username},
			CreatedAt: post.CreatedAt.Format("2006-01-02 15:04:05"),
		}
	}

	utils.RespondWithJSON(c, http.StatusOK, res)
}

func GetPostHandler(c *gin.Context) {}

func UpdatePostHandler(c *gin.Context) {}

func DeletePostHandler(c *gin.Context) {}
