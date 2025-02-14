package tag

import (
	"github.com/Cai-ki/go-caiki-blog/models"
	"github.com/Cai-ki/go-caiki-blog/pkg/storage"
)

type TagService interface {
	ListTags(tags *[]models.Tags) error
	CreateTag(tag *models.Tags) error
}

type tagServiceImpl struct {
}

var _ TagService = (*tagServiceImpl)(nil)

var Service = tagServiceImpl{}

func (tagServiceImpl) ListTags(tags *[]models.Tags) error {
	db := storage.DB.GetDB()
	err := db.Find(tags).Error
	if err != nil {
		return err
	}
	return nil
}

func (tagServiceImpl) CreateTag(tag *models.Tags) error {
	db := storage.DB.GetDB()

	err := db.Create(tag).Error
	if err != nil {
		return err
	}
	return nil
}
