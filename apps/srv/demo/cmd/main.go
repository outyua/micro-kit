package main

import (
	"log"
	demo "micro-kit/apps/srv/demo/api/demo"
	srv "micro-kit/apps/srv/demo/internal/service"
	"micro-kit/apps/srv/demo/internal/subscriber"
	"micro-kit/global/reg"

	"github.com/micro/go-micro/v2"
)

func main() {
	// New Service
	meta := reg.NewSrvDemo()
	service := micro.NewService(
		micro.Name(meta.GetName()),
		micro.Version(meta.GerVersion()),
	)

	// Initialise service
	service.Init()
	// Register Handler
	demo.RegisterDemoHandler(service.Server(), new(srv.Demo))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.demo", service.Server(), new(subscriber.Demo))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.demo", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
