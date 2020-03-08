package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/logger"
)

func Login(c *gin.Context) {

	username := c.GetString("username")

	logger.Info(username)

	c.JSON(200, gin.H{
		"code": 20000,
		"data": gin.H{
			"token": 99999,
		},
	})
}

func Info(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 20000,
		"data": gin.H{
			"name":   "outyua",
			"avatar": "https://media.ifanrusercontent.com/tavatar/28/d3/28d39bd75da9ef227188eae4fd73777cc8b30bfc.jpg",
		},
	})
}

func Logout(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 20000,
		"data": gin.H{
			"status": "ok",
		},
	})
}
