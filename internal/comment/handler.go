package comment

import (
	"net/http"
	"strconv"

	"github.com/Cai-ki/go-caiki-blog/models"
	"github.com/Cai-ki/go-caiki-blog/utils"
	"github.com/gin-gonic/gin"
)

type listCommentResponseInfo struct {
	Comments []commentInfo `json:"comments"`
}

type commentInfo struct {
	Username  string `json:"username"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
}

func ListCommentsHandler(c *gin.Context) {
	postIDStr := c.Param("id")
	postID, err := strconv.Atoi(postIDStr)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid post ID")
		return
	}

	comments, err := Service.ListComments(postID)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to list comments")
		return
	}

	response := listCommentResponseInfo{
		Comments: make([]commentInfo, len(comments)),
	}

	for i, comment := range comments {
		response.Comments[i] = commentInfo{
			Username:  comment.User.Username,
			Content:   comment.Content,
			CreatedAt: comment.CreatedAt.Format("2006-01-02 15:04:05"),
		}
	}

	utils.RespondWithJSON(c, http.StatusOK, response)
}

type createCommentRequestInfo struct {
	Content string `json:"content"`
}

type createCommentResponseInfo struct {
	CreatedAt string `json:"created_at"`
}

func CreateCommentHandler(c *gin.Context) {
	var request createCommentRequestInfo
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

	comment := models.Comments{
		PostID:  uint(postID),
		UserID:  userID,
		Content: request.Content,
	}

	comment, err = Service.CreateComment(comment)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to create comment")
		return
	}

	response := createCommentResponseInfo{
		CreatedAt: comment.CreatedAt.Format("2006-01-02 15:04:05"),
	}

	utils.RespondWithJSON(c, http.StatusCreated, response)
}
