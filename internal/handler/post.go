package handler

import (
	"net/http"
	"strconv"

	"github.com/Cai-ki/go-caiki-blog/models"
	"github.com/Cai-ki/go-caiki-blog/pkg/cgin"
	"github.com/Cai-ki/go-caiki-blog/utils"
)

func CreatePostHandler(c *cgin.Context) {
	var req struct {
		Title   string   `json:"title"`
		Content string   `json:"content"`
		Tags    []string `json:"tags"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request body")
		return
	}

	userID := c.GetUint("user_id")
	post := models.Posts{
		Title:   req.Title,
		Content: req.Content,
		UserID:  userID,
	}

	if err := PostService.CreatePost(&post); err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to create post")
		return
	}

	tags := []models.Tags{}
	for _, name := range req.Tags {
		tags = append(tags, models.Tags{
			Name: name,
		})
	}

	if err := TagService.ConnectTags(&post, &tags); err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to connect tags")
		return
	}

	res := struct {
		ID        uint   `json:"id"`
		Title     string `json:"title"`
		UserID    uint   `json:"user_id"`
		CreatedAt string `json:"created_at"`
	}{
		ID:        post.ID,
		Title:     post.Title,
		UserID:    post.UserID,
		CreatedAt: post.CreatedAt.Format("2006-01-02 15:04:05"),
	}

	utils.RespondWithJSON(c, http.StatusCreated, res)
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

type postDetailInfo struct {
	postInfo
	Content string `json:"content"`
}

func ListPostsHandler(c *cgin.Context) {
	var req struct {
		Page  int `form:"page"`
		Limit int `form:"limit"`
	}

	pageStr := c.Query("page")
	limitStr := c.Query("limit")

	var err error

	if pageStr != "" {
		req.Page, err = strconv.Atoi(pageStr)
		if err != nil {
			utils.RespondWithError(c, http.StatusBadRequest, "Invalid page")
			return
		}
	}

	if limitStr != "" {
		req.Limit, err = strconv.Atoi(limitStr)
		if err != nil {
			utils.RespondWithError(c, http.StatusBadRequest, "Invalid limit")
			return
		}
	}

	// if err := c.ShouldBindQuery(&req); err != nil {
	// 	utils.RespondWithError(c, http.StatusBadRequest, "Invalid request query")
	// 	return
	// }

	posts := []models.Posts{}
	if err := PostService.ListPosts(&posts, req.Page, req.Limit); err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to list posts")
		return
	}

	res := make([]postInfo, len(posts))

	for i, post := range posts {
		res[i] = postInfo{
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

func GetPostHandler(c *cgin.Context) {
	postIDStr := c.Param("id")
	postID, err := strconv.Atoi(postIDStr)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid post ID")
		return
	}

	post := models.Posts{}
	post.ID = uint(postID)

	if err := PostService.GetPost(&post); err != nil {
		utils.RespondWithError(c, http.StatusNotFound, "Post not found")
		return
	}

	res := postDetailInfo{
		postInfo: postInfo{
			ID:    post.ID,
			Title: post.Title,
			Author: struct {
				ID   uint   `json:"id"`
				Name string `json:"name"`
			}{post.UserID, post.User.Username},
			CreatedAt: post.CreatedAt.Format("2006-01-02 15:04:05"),
		},
		Content: post.Content,
	}

	utils.RespondWithJSON(c, http.StatusOK, res)
}

func UpdatePostHandler(c *cgin.Context) {
	postIDStr := c.Param("id")
	postID, err := strconv.Atoi(postIDStr)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid post ID")
		return
	}
	post := models.Posts{}
	post.ID = uint(postID)

	if err := PostService.GetPost(&post); err != nil {
		utils.RespondWithError(c, http.StatusNotFound, "Post not found")
		return
	}

	if post.UserID != c.GetUint("user_id") {
		utils.RespondWithError(c, http.StatusForbidden, "Forbidden")
		return
	}

	var req struct {
		Title   string   `json:"title"`
		Content string   `json:"content"`
		Tags    []string `json:"tags"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request body")
		return
	}

	post.Title = req.Title
	post.Content = req.Content

	if err := PostService.UpdatePost(&post); err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to update post")
		return
	}

	tags := []models.Tags{}
	for _, name := range req.Tags {
		tags = append(tags, models.Tags{
			Name: name,
		})
	}

	if err := TagService.ConnectTags(&post, &tags); err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to connect tags")
		return
	}

	if err := PostService.GetPost(&post); err != nil {
		utils.RespondWithError(c, http.StatusNotFound, "Post not found")
		return
	}

	res := postDetailInfo{
		postInfo: postInfo{
			ID:    post.ID,
			Title: post.Title,
			Author: struct {
				ID   uint   `json:"id"`
				Name string `json:"name"`
			}{post.UserID, post.User.Username},
			CreatedAt: post.CreatedAt.Format("2006-01-02 15:04:05"),
		},
		Content: post.Content,
	}

	utils.RespondWithJSON(c, http.StatusOK, res)
}

func DeletePostHandler(c *cgin.Context) {
	postIDStr := c.Param("id")
	postID, err := strconv.Atoi(postIDStr)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid post ID")
		return
	}

	post := models.Posts{}
	post.ID = uint(postID)

	if err := PostService.GetPost(&post); err != nil {
		utils.RespondWithError(c, http.StatusNotFound, "Post not found")
		return
	}

	if post.UserID != c.GetUint("user_id") {
		utils.RespondWithError(c, http.StatusForbidden, "Forbidden")
		return
	}

	if err := PostService.DeletePost(&post); err != nil {
		utils.RespondWithError(c, http.StatusNotFound, "Post not found")
		return
	}

	utils.RespondWithJSON(c, http.StatusNoContent, nil)
}
