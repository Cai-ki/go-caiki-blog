package user

import (
	"github.com/Cai-ki/go-caiki-blog/internal/auth"
	"github.com/Cai-ki/go-caiki-blog/models"
	"github.com/Cai-ki/go-caiki-blog/pkg/storage"
)

type UserService interface {
	Register(user *models.Users) (err error)
	Login(user *models.Users) (token string, err error)
	GetUser(user *models.Users) (err error)
}

type userServiceImpl struct {
}

var _ UserService = (*userServiceImpl)(nil)

var Service = userServiceImpl{}

func (userServiceImpl) Register(user *models.Users) (err error) {
	db := storage.DB.GetDB()
	if err = db.Create(user).Error; err != nil {
		return
	}
	return
}

func (userServiceImpl) Login(user *models.Users) (token string, err error) {
	db := storage.DB.GetDB()
	if err = db.Model(&models.Users{}).Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error; err != nil {
		return
	}

	jwtToken, err := auth.Jwt.GenerateToken(user.ID, user.Username, user.Email)
	if err != nil {
		return
	}

	return jwtToken, nil
}

func (userServiceImpl) GetUser(user *models.Users) (err error) {
	db := storage.DB.GetDB()
	if err = db.Model(&models.Users{}).Where("username = ?", user.Username).First(&user).Error; err != nil {
		return
	}
	return
}
