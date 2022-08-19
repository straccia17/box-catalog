package services

import (
	"github.com/gin-gonic/gin"
)

func RetrieveUserInfo(c *gin.Context) (string, string) {
	id := c.GetString("UserID")
	email := c.GetString("Email")
	return id, email
}
