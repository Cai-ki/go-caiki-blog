package handler

import (
	"net/http"
	"strconv"

	"github.com/Cai-ki/go-caiki-blog/models"
	"github.com/Cai-ki/go-caiki-blog/pkg/cgin"
	"github.com/Cai-ki/go-caiki-blog/utils"
)

type commentInfo struct {
	Username  string `json:"username"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
}

func ListCommentsHandler(c *cgin.Context) {
	postIDStr := c.Param("id")
	postID, err := strconv.Atoi(postIDStr)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid post ID")
		return
	}

	post := models.Posts{}
	post.ID = uint(postID)

	if err = PostService.GetPost(&post); err != nil {
		utils.RespondWithError(c, http.StatusNotFound, "Post not found")
		return
	}

	comments := []models.Comments{}

	if err = CommentService.ListComments(&post, &comments); err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to list comments")
		return
	}

	response := []commentInfo{}

	for _, comment := range comments {
		response = append(response, commentInfo{
			Username:  comment.User.Username,
			Content:   comment.Content,
			CreatedAt: comment.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	utils.RespondWithJSON(c, http.StatusOK, response)
}

func CreateCommentHandler(c *cgin.Context) {
	var request struct {
		Content string `json:"content"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request body")
		return
	}

	postIDStr := c.Param("id")
	postID, err := strconv.Atoi(postIDStr)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid post ID")
		return
	}

	userID := c.GetUint("user_id")
	post := models.Posts{}
	post.ID = uint(postID)

	if err = PostService.GetPost(&post); err != nil {
		utils.RespondWithError(c, http.StatusNotFound, "Post not found")
		return
	}

	comment := models.Comments{
		PostID:  uint(postID),
		UserID:  userID,
		Content: request.Content,
	}

	if err = CommentService.CreateComment(&comment); err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to create comment")
		return
	}

	response := struct {
		CreatedAt string `json:"created_at"`
	}{
		CreatedAt: comment.CreatedAt.Format("2006-01-02 15:04:05"),
	}

	utils.RespondWithJSON(c, http.StatusCreated, response)
}
