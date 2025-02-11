package utils

import "github.com/gin-gonic/gin"

func RespondWithError(c *gin.Context, code int, err error) {
	c.JSON(code, err.Error())
}

func RespondWithJSON(c *gin.Context, code int, data interface{}) {
	c.JSON(code, data)
}
