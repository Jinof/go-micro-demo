package main

import (
	srv "github.com/Jinof/go-micro-demo/user/genproto/srv"
	"github.com/Jinof/go-micro-demo/user/internal/srv/handler"
	"github.com/Jinof/go-micro-demo/user/pkg/pubsub"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
)

func main() {

	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.srv"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	srv.RegisterUserHandler(service.Server(), new(handler.User))

	micro.RegisterSubscriber(pubsub.Topic, service.Server(), new(handler.Event))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
