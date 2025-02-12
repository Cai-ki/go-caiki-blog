package post

import (
	"github.com/Cai-ki/go-caiki-blog/models"
	"github.com/Cai-ki/go-caiki-blog/pkg/storage"
)

type PostService interface {
	CreatePost(user_id uint, title string, content string) (res models.Posts, err error)
	ListPosts(page int, limit int) (res []models.Posts, err error)
}

type postServiceImpl struct {
}

var _ PostService = (*postServiceImpl)(nil)

var Service = postServiceImpl{}

func (postServiceImpl) CreatePost(user_id uint, title string, content string) (res models.Posts, err error) {
	post := models.Posts{UserID: user_id, Title: title, Content: content}

	db := storage.DB.GetDB()
	if err = db.Create(&post).Error; err != nil {
		return models.Posts{}, err
	}

	db.Model(&models.Posts{}).Where("user_id = ?", user_id).First(&res)

	if res.ID == 0 {
		return models.Posts{}, err
	}

	return res, nil
}

func (postServiceImpl) ListPosts(page int, limit int) (res []models.Posts, err error) {
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}

	db := storage.DB.GetDB()
	db = db.Preload("User").Model(&models.Posts{})
	if err = db.Offset((page - 1) * limit).Limit(limit).Find(&res).Error; err != nil {
		return []models.Posts{}, err
	}

	return res, nil
}
