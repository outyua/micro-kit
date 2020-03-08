package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/logger"
	authSrv "micro-kit/apps/srv/admin/api/auth"
	"micro-kit/global/reg"
)

func Login(c *gin.Context) {

	srvAuthMeta := reg.NewSrvAdmin()
	authClient := authSrv.NewAuthService(srvAuthMeta.GetName(), client.DefaultClient)

	data, _ := c.GetRawData()
	var req struct {
		Username string
		Password string
	}
	json.Unmarshal(data, &req)

	logger.Error("post data ---->>", req)

	loginRes, err := authClient.Login(c, &authSrv.LoginReq{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil || loginRes == nil {
		c.JSON(200, gin.H{
			"code": 50000,
			"data": nil,
			"msg":  err,
		})
		return
	}

	c.JSON(200, gin.H{
		"code": 20000,
		"data": gin.H{
			"token": loginRes.Token,
		},
	})
}

func Info(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 20000,
		"data": gin.H{
			"name":   "outyua",
			"avatar": "https://media.ifanrusercontent.com/tavatar/28/d3/28d39bd75da9ef227188eae4fd73777cc8b30bfc.jpg",
			"roles":  []string{"admin", "editor"},
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
