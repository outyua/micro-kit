package server

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/config"
	"net/http"
)

type server struct {
	g *gin.Engine
}

func NewServer() http.Handler {
	debug := config.Get("debug").Bool(true)
	if !debug {
		gin.SetMode(gin.ReleaseMode)
	}
	g := gin.New()
	g.Use(gin.Logger())
	g.Use(gin.Recovery())

	g.GET("/status", statusRoute)

	addRouter(g)
	return g
}
