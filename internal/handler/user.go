package handler

import (
	"crypto/sha256"
	"net/http"
	"time"

	"github.com/Cai-ki/go-caiki-blog/internal/service"
	"github.com/Cai-ki/go-caiki-blog/models"
	"github.com/Cai-ki/go-caiki-blog/utils"
	"github.com/gin-gonic/gin"
)

var UserService = service.UserService

func RegisterHandler(c *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request body")
		return
	}

	data := []byte(req.Password)
	hash := sha256.Sum256(data)
	user := models.Users{Username: req.Username, Email: req.Email, Password: string(hash[:])}

	if err := UserService.Register(&user); err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to register user")
		return
	}

	userInfo := struct {
		Username  string    `json:"username"`
		Email     string    `json:"email"`
		CreatedAt time.Time `json:"created_at"`
		// AvatarURL string `json:"avatar_url"`
	}{Username: user.Username, Email: user.Email, CreatedAt: user.CreatedAt}
	utils.RespondWithJSON(c, http.StatusCreated, userInfo)
}

func LoginHandler(c *gin.Context) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request body")
		return
	}

	data := []byte(req.Password)
	hash := sha256.Sum256(data)
	user := models.Users{Email: req.Email, Password: string(hash[:])}

	token, err := UserService.Login(&user)
	if err != nil {
		utils.RespondWithError(c, http.StatusUnauthorized, "Invalid email or password")
		return
	}

	userInfo := struct {
		Token string `json:"token"`
	}{Token: token}

	utils.RespondWithJSON(c, http.StatusOK, userInfo)
}

func GetUserHandler(c *gin.Context) {
	username := c.Param("username")
	user := models.Users{Username: username}

	if err := UserService.GetUser(&user); err != nil {
		utils.RespondWithError(c, http.StatusNotFound, "User not found")
		return
	}

	userInfo := struct {
		Username string `json:"username"`
		// Email    string `json:"email"`
		// CreatedAt time.Time `json:"created_at"`
		// AvatarURL string `json:"avatar_url"`
	}{Username: user.Username}

	utils.RespondWithJSON(c, http.StatusOK, userInfo)
}
