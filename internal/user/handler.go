package user

import (
	"net/http"
	"time"

	"github.com/Cai-ki/go-caiki-blog/utils"
	"github.com/gin-gonic/gin"
)

type registerRequestInfo struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type registerResponseInfo struct {
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	// AvatarURL string `json:"avatar_url"`
}

func RegisterHandler(c *gin.Context) {
	var req registerRequestInfo

	if err := c.BindJSON(&req); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, err)
		return
	}

	user, err := userService.Register(req.Username, req.Email, req.Password)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err)
		return
	}

	userInfo := registerResponseInfo{Username: user.Username, Email: user.Email, CreatedAt: user.CreatedAt}
	utils.RespondWithJSON(c, http.StatusCreated, userInfo)
}

type loginRequestInfo struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type loginResponseInfo struct {
	Token string `json:"token"`
}

func LoginHandler(c *gin.Context) {
	var req loginRequestInfo

	if err := c.BindJSON(&req); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, err)
		return
	}

	token, err := userService.Login(req.Email, req.Password)
	if err != nil {
		utils.RespondWithError(c, http.StatusUnauthorized, err)
		return
	}

	userInfo := loginResponseInfo{Token: token}
	utils.RespondWithJSON(c, http.StatusOK, userInfo)
}

type getUserResponseInfo struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	// CreatedAt time.Time `json:"created_at"`
	// AvatarURL string `json:"avatar_url"`
}

func GetUserHandler(c *gin.Context) {
	username := c.Param("username")

	user, err := userService.GetUserByName(username)
	if err != nil {
		utils.RespondWithError(c, http.StatusNotFound, err)
		return
	}

	userInfo := getUserResponseInfo{Username: user.Username, Email: user.Email}
	utils.RespondWithJSON(c, http.StatusOK, userInfo)
}
