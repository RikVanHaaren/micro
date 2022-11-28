package main

import (
	"github.com/RikVanHaaren/micro/services/example/handler"
	pb "github.com/RikVanHaaren/micro/services/example/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("example"),
	)

	// Register handler
	pb.RegisterExampleHandler(srv.Server(), handler.New())

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
