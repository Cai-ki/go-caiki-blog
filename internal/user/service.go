package user

import "github.com/Cai-ki/go-caiki-blog/models"

type UserService interface {
	Register(username, email, password string) (*models.User, error)
	// 返回 JWT Token
	Login(email, password string) (string, error)
	GetUserByEmail(email string) (*models.User, error)
	GetUserByName(username string) (*models.User, error)
}

type userServiceImpl struct {
}

var _ UserService = (*userServiceImpl)(nil)

var userService = userServiceImpl{}

func (us userServiceImpl) Register(username, email, password string) (*models.User, error) {
	return &models.User{}, nil
}

func (us userServiceImpl) Login(email, password string) (string, error) { return "", nil }

func (us userServiceImpl) GetUserByEmail(email string) (*models.User, error) {
	return &models.User{}, nil
}

func (us userServiceImpl) GetUserByName(username string) (*models.User, error) {
	return &models.User{}, nil
}
