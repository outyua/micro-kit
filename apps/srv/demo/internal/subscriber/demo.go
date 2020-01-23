package subscriber

import (
	"context"
	"github.com/micro/go-micro/util/log"

	demo "micro-kit/apps/srv/demo/api/demo"
)

type Demo struct{}

func (e *Demo) Handle(ctx context.Context, msg *demo.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *demo.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
