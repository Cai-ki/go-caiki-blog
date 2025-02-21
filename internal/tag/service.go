package tag

import (
	"errors"

	"github.com/Cai-ki/go-caiki-blog/models"
	"github.com/Cai-ki/go-caiki-blog/pkg/storage"
	"gorm.io/gorm"
)

type TagService interface {
	ListTags(tags *[]models.Tags) (err error)
	ConnectTags(post *models.Posts, tags *[]models.Tags) (err error)
}

type tagServiceImpl struct {
}

var _ TagService = (*tagServiceImpl)(nil)

var Service = tagServiceImpl{}

func (tagServiceImpl) ListTags(tags *[]models.Tags) (err error) {
	db := storage.DB.GetDB()

	if err = db.Find(tags).Error; err != nil {
		return
	}
	return
}

func (tagServiceImpl) ConnectTags(post *models.Posts, tags *[]models.Tags) (err error) {
	db := storage.DB.GetDB()

	if err = Service.CreateTags(tags); err != nil {
		return
	}

	if err = db.Model(post).Association("Tags").Replace(tags); err != nil {
		return
	}
	return
}

func (tagServiceImpl) CreateTags(tags *[]models.Tags) (err error) {
	db := storage.DB.GetDB()

	for i := range *tags {
		tag := &(*tags)[i]
		var existingTag models.Tags
		if err = db.Where("name = ?", tag.Name).First(&existingTag).Error; err == nil {
			tag.ID = existingTag.ID
		} else if errors.Is(err, gorm.ErrRecordNotFound) {
			if err = db.Create(tag).Error; err != nil {
				return
			}
		} else {
			return
		}
	}
	return
}

func (tagServiceImpl) ListPostTags(post *models.Posts, tags *[]models.Tags) (err error) {
	db := storage.DB.GetDB()

	if err = db.Model(post).Association("Tags").Find(tags); err != nil {
		return
	}
	return
}
