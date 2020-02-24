package subscriber

import (
	"context"
	"log"

	demo "micro-kit/apps/srv/demo/api/demo"
)

type Demo struct{}

func (e *Demo) Handle(ctx context.Context, msg *demo.Message) error {
	log.Println("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *demo.Message) error {
	log.Println("Function Received message: ", msg.Say)
	return nil
}
