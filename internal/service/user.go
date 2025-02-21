package service

import (
	"github.com/Cai-ki/go-caiki-blog/models"
	"github.com/Cai-ki/go-caiki-blog/pkg/jwt"
	"github.com/Cai-ki/go-caiki-blog/pkg/storage"
)

var Jwt = jwt.Jwt

type userService interface {
	Register(user *models.Users) (err error)
	Login(user *models.Users) (token string, err error)
	GetUser(user *models.Users) (err error)
}

type userServiceImpl struct {
}

var _ userService = (*userServiceImpl)(nil)

var UserService = userServiceImpl{}

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

	jwtToken, err := Jwt.GenerateToken(user.ID, user.Username, user.Email)
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
