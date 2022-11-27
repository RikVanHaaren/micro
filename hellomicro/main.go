package main

import (
	"hellomicro/handler"
	pb "hellomicro/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("hellomicro"),
	)

	// Register handler
	pb.RegisterHellomicroHandler(srv.Server(), handler.New())

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
