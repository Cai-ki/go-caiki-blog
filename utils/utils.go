package utils

import "github.com/gin-gonic/gin"

func RespondWithError(c *gin.Context, code int, message string) {
	c.JSON(code, Error{Code: code, Message: message})
}

func RespondWithJSON(c *gin.Context, code int, data interface{}) {
	c.JSON(code, data)
}
