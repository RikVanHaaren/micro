package handler

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"

	pb "github.com/RikVanHaaren/micro/services/example/proto"
)

type Example struct{}

// Return a new handler
func New() *Example {
	return &Example{}
}

// Call is a single request handler called via client.Call or the generated client code
func (e *Example) Call(ctx context.Context, req *pb.Request, rsp *pb.Response) error {
	log.Info("Received Example.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Example) Stream(ctx context.Context, req *pb.StreamingRequest, stream pb.Example_StreamStream) error {
	log.Infof("Received Example.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&pb.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Example) PingPong(ctx context.Context, stream pb.Example_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&pb.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
