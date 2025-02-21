package tag

import (
	"errors"

	"github.com/Cai-ki/go-caiki-blog/models"
	"github.com/Cai-ki/go-caiki-blog/pkg/storage"
	"gorm.io/gorm"
)

type TagService interface {
	ListTags(tags *[]models.Tags) error
	ConnectTags(post *models.Posts, tags *[]models.Tags) error
}

type tagServiceImpl struct {
}

var _ TagService = (*tagServiceImpl)(nil)

var Service = tagServiceImpl{}

func (tagServiceImpl) ListTags(tags *[]models.Tags) error {
	db := storage.DB.GetDB()

	if err := db.Find(tags).Error; err != nil {
		return err
	}
	return nil
}

func (tagServiceImpl) ConnectTags(post *models.Posts, tags *[]models.Tags) error {
	db := storage.DB.GetDB()

	if err := Service.CreateTags(tags); err != nil {
		return err
	}

	if err := db.Model(post).Association("Tags").Replace(tags); err != nil {
		return err
	}
	return nil
}

func (tagServiceImpl) CreateTags(tags *[]models.Tags) error {
	db := storage.DB.GetDB()

	for i := range *tags {
		tag := &(*tags)[i]
		var existingTag models.Tags
		if err := db.Where("name = ?", tag.Name).First(&existingTag).Error; err == nil {
			tag.ID = existingTag.ID
		} else if errors.Is(err, gorm.ErrRecordNotFound) {
			if err := db.Create(tag).Error; err != nil {
				return err
			}
		} else {
			return err
		}
	}
	return nil
}

func (tagServiceImpl) ListPostTags(post *models.Posts, tags *[]models.Tags) error {
	db := storage.DB.GetDB()

	if err := db.Model(post).Association("Tags").Find(tags); err != nil {
		return err
	}
	return nil
}
