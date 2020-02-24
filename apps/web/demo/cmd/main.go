package main

import (
	"github.com/micro/go-micro/v2/web"
	"log"
	handler "micro-kit/apps/web/demo/internal/server"
	"micro-kit/global/reg"
	"net/http"
)

func main() {

	meta := reg.NewWebDemo()
	// create new web service
	service := web.NewService(
		web.Name(meta.GetName()),
		web.Version(meta.GerVersion()),
		web.Address(":9991"),
	)

	// initialise service
	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	// register html handler
	service.Handle("/", http.FileServer(http.Dir("web")))

	// register call handler
	service.HandleFunc("/demo/call", handler.DemoCall)

	// run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
