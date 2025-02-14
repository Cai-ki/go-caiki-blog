package comment

import (
	"github.com/Cai-ki/go-caiki-blog/models"
	"github.com/Cai-ki/go-caiki-blog/pkg/storage"
)

type commentService interface {
	ListComments(post *models.Posts, comments *[]models.Comments) (err error)
	CreateComment(comment *models.Comments) (err error)
}

type commentServiceImpl struct {
}

var _ commentService = (*commentServiceImpl)(nil)

var Service = commentServiceImpl{}

func (commentServiceImpl) ListComments(post *models.Posts, comments *[]models.Comments) (err error) {
	db := storage.DB.GetDB()
	if err = db.Preload("User").Model(&models.Comments{}).
		Where("post_id =?", post.ID).
		Find(comments).Error; err != nil {
		return err
	}
	return nil
}

func (commentServiceImpl) CreateComment(comment *models.Comments) (err error) {
	db := storage.DB.GetDB()
	if err = db.Create(comment).Error; err != nil {
		return err
	}
	return nil
}
