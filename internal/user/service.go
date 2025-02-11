package user

import (
	"crypto/sha256"

	"github.com/Cai-ki/go-caiki-blog/models"
	"github.com/Cai-ki/go-caiki-blog/pkg/storage"
)

type UserService interface {
	Register(username, email, password string) (models.Users, error)
	// 返回 JWT Token
	Login(email, password string) (string, error)
	GetUserByEmail(email string) (models.Users, error)
	GetUserByName(username string) (models.Users, error)
}

type userServiceImpl struct {
}

var _ UserService = (*userServiceImpl)(nil)

var userService = userServiceImpl{}

func (userServiceImpl) Register(username, email, password string) (user models.Users, err error) {
	data := []byte(password)
	hash := sha256.Sum256(data)

	userInfo := models.Users{
		Username: username,
		Email:    email,
		Password: string(hash[:]),
	}

	db := storage.DB.GetDB()
	err = db.Create(&userInfo).Error
	if err != nil {
		return models.Users{}, err
	}

	db.Model(&models.Users{}).Where("username = ?", username).First(&user)

	return user, nil
}

func (userServiceImpl) Login(email, password string) (string, error) { return "", nil }

func (userServiceImpl) GetUserByEmail(email string) (models.Users, error) {
	return models.Users{}, nil
}

func (userServiceImpl) GetUserByName(username string) (models.Users, error) {
	return models.Users{}, nil
}
