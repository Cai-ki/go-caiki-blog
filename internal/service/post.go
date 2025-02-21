package service

import (
	"github.com/Cai-ki/go-caiki-blog/models"
	"github.com/Cai-ki/go-caiki-blog/pkg/storage"
)

type postService interface {
	CreatePost(post *models.Posts) (err error)
	ListPosts(posts *[]models.Posts, page int, limit int) (err error)
	GetPost(post *models.Posts) (err error)
	DeletePost(post *models.Posts) (err error)
	UpdatePost(post *models.Posts) (err error)
}

type postServiceImpl struct {
}

var _ postService = (*postServiceImpl)(nil)

var PostService = postServiceImpl{}

func (postServiceImpl) CreatePost(post *models.Posts) (err error) {
	db := storage.DB.GetDB()
	if err = db.Create(&post).Error; err != nil {
		return
	}

	return
}

func (postServiceImpl) ListPosts(posts *[]models.Posts, page int, limit int) (err error) {
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}

	db := storage.DB.GetDB()
	db = db.Preload("User").Model(&models.Posts{})
	if err = db.Offset((page - 1) * limit).Limit(limit).Find(posts).Error; err != nil {
		return
	}

	return
}

func (postServiceImpl) GetPost(post *models.Posts) (err error) {
	db := storage.DB.GetDB()
	db = db.Preload("User").Model(&models.Posts{})

	if err = db.First(&post).Error; err != nil {
		return
	}
	return
}

func (postServiceImpl) DeletePost(post *models.Posts) (err error) {
	db := storage.DB.GetDB()

	if err = db.Delete(&post).Error; err != nil {
		return
	}
	return
}

func (postServiceImpl) UpdatePost(post *models.Posts) (err error) {
	db := storage.DB.GetDB()
	if err = db.Model(post).Updates(&post).Error; err != nil {
		return
	}
	return
}
