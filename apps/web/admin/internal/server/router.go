package server

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/config"
	"micro-kit/apps/web/admin/internal/controller"
	"net/http"
)

func statusRoute(c *gin.Context) {
	name := config.Get("name").String("default")
	c.JSON(200, gin.H{
		"status": "ok",
		"name":   name,
	})
}

func addRouter(g *gin.Engine) {

	// add route here

	g.StaticFile("/", "web/index.html")
	g.StaticFS("/static", http.Dir("web/static"))

	api := g.Group("/api")

	userApi := api.Group("/user")
	userApi.POST("/login", controller.Login)
	userApi.GET("/info", controller.Info)
	userApi.POST("/logout", controller.Logout)
}
