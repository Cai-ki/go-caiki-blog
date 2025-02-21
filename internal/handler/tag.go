package handler

import (
	"net/http"
	"strconv"

	"github.com/Cai-ki/go-caiki-blog/models"
	"github.com/Cai-ki/go-caiki-blog/utils"
	"github.com/gin-gonic/gin"
)

type tagInfo struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func ListTagsHandler(c *gin.Context) {
	tags := []models.Tags{}
	err := TagService.ListTags(&tags)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to list tags")
		return
	}

	tagInfos := []tagInfo{}
	for _, tag := range tags {
		tagInfos = append(tagInfos, tagInfo{
			ID:   tag.ID,
			Name: tag.Name,
		})
	}

	utils.RespondWithJSON(c, http.StatusOK, tagInfos)
}

func ConnectTagsHandler(c *gin.Context) {
	var req struct {
		PostID uint     `json:"post_id"`
		Tags   []string `json:"tags"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request payload")
		return
	}

	post := models.Posts{}
	post.ID = req.PostID
	err := postService.GetPost(&post)
	if err != nil {
		utils.RespondWithError(c, http.StatusNotFound, "Post not found")
		return
	}

	tags := []models.Tags{}
	for _, name := range req.Tags {
		tags = append(tags, models.Tags{
			Name: name,
		})
	}

	err = TagService.ConnectTags(&post, &tags)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to connect tags")
		return
	}

	utils.RespondWithError(c, http.StatusCreated, "Tags connected successfully")
}

func ListPostTagsHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid post ID")
		return
	}
	post := models.Posts{}
	post.ID = uint(id)
	tags := []models.Tags{}
	err = TagService.ListPostTags(&post, &tags)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to list post tags")
		return
	}

	tagInfos := []tagInfo{}
	for _, tag := range tags {
		tagInfos = append(tagInfos, tagInfo{
			ID:   tag.ID,
			Name: tag.Name,
		})
	}

	utils.RespondWithJSON(c, http.StatusOK, tagInfos)
}
