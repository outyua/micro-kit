package main

import (
	"github.com/micro/go-micro/v2/config"
	"log"
	auth "micro-kit/apps/srv/admin/api/auth"
	srv "micro-kit/apps/srv/admin/internal/service"
	"micro-kit/global/reg"
	"micro-kit/pkg/database"
	"path/filepath"

	"github.com/micro/go-micro/v2"
)

func main() {
	// New Service
	meta := reg.NewSrvAdmin()
	service := micro.NewService(
		micro.Name(meta.GetName()),
		micro.Version(meta.GerVersion()),
	)

	// load config
	confPath, _ := filepath.Abs("configs/app.yaml")
	err := config.LoadFile(confPath)
	if err != nil {
		log.Fatal(err)
	}

	var dbDefault database.Config
	err = config.Get("db", "default").Scan(&dbDefault)
	if err != nil {
		panic("db default config is error")
	}

	// Initialise service
	service.Init(
		micro.BeforeStart(func() error {
			return database.Init("default", dbDefault)
		}),
		micro.AfterStop(func() error {
			return database.Close("default")
		}),
	)
	// Register Handler
	auth.RegisterAuthHandler(service.Server(), new(srv.Auth))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
