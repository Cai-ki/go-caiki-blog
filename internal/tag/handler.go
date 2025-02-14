package tag

import (
	"net/http"

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
	err := Service.ListTags(&tags)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to list tags")
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

func CreateTagHandler(c *gin.Context) {
	name := c.Param("name")
	tag := models.Tags{
		Name: name,
	}
	err := Service.CreateTag(&tag)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to create tag")
		return
	}

	utils.RespondWithJSON(c, http.StatusCreated, tagInfo{
		ID:   tag.ID,
		Name: tag.Name,
	})
}
