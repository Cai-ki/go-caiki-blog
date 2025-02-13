package comment

import (
	"github.com/Cai-ki/go-caiki-blog/models"
	"github.com/Cai-ki/go-caiki-blog/pkg/storage"
)

type commentService interface {
	ListComments(postID int) ([]models.Comments, error)
	CreateComment(comment models.Comments) (models.Comments, error)
}

type commentServiceImpl struct {
}

var _ commentService = (*commentServiceImpl)(nil)

var Service = commentServiceImpl{}

func (s *commentServiceImpl) ListComments(postID int) ([]models.Comments, error) {
	db := storage.DB.GetDB()
	var comments []models.Comments
	err := db.Model(&models.Comments{}).Where("post_id =?", postID).Preload("User").Find(&comments).Error
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func (s *commentServiceImpl) CreateComment(comment models.Comments) (models.Comments, error) {
	db := storage.DB.GetDB()
	err := db.Create(&comment).Error
	if err != nil {
		return models.Comments{}, err
	}
	return comment, nil
}
