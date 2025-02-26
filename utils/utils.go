package utils

import "github.com/Cai-ki/go-caiki-blog/pkg/cgin"

func RespondWithError(c *cgin.Context, code int, message string) {
	c.JSON(code, Error{Code: code, Message: message})
}

func RespondWithJSON(c *cgin.Context, code int, data interface{}) {
	c.JSON(code, data)
}
