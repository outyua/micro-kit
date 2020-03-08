package main

import (
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/web"
	"log"
	"micro-kit/apps/web/admin/internal/server"
	"micro-kit/global/reg"
	"path/filepath"
)

func main() {

	meta := reg.NewWebDemo()
	// create new web service
	service := web.NewService(
		web.Name(meta.GetName()),
		web.Version(meta.GerVersion()),
		web.Address(":9996"),
	)

	// initialise service
	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	// load config
	confPath, _ := filepath.Abs("configs/app.yaml")
	err := config.LoadFile(confPath)
	if err != nil {
		log.Fatal(err)
	}

	// register gin handler
	service.Handle("/", server.NewServer())

	// run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
