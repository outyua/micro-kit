package service

import (
	"context"
	"log"

	demo "micro-kit/apps/srv/demo/api/demo"
)

type Demo struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Demo) Call(ctx context.Context, req *demo.Request, rsp *demo.Response) error {
	log.Println("Received Demo.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Demo) Stream(ctx context.Context, req *demo.StreamingRequest, stream demo.Demo_StreamStream) error {
	log.Printf("Received Demo.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Printf("Responding: %d", i)
		if err := stream.Send(&demo.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Demo) PingPong(ctx context.Context, stream demo.Demo_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Printf("Got ping %v", req.Stroke)
		if err := stream.Send(&demo.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
