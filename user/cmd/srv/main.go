package main

import (
	"context"
	srv "github.com/Jinof/go-micro-demo/user/genproto/srv"
	"github.com/Jinof/go-micro-demo/user/internal/srv/handler"
	"github.com/Jinof/go-micro-demo/user/pkg/pubsub"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/broker"
	log "github.com/micro/go-micro/v2/logger"
	ocPlugin "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	oc "github.com/opentracing/opentracing-go"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.srv"),
		micro.Version("latest"),
		micro.WrapHandler(ocPlugin.NewHandlerWrapper(oc.GlobalTracer())),
	)

	// Initialise service
	service.Init()

	// Register Handler
	err := srv.RegisterUserHandler(service.Server(), new(handler.User))
	if err != nil {
		log.Fatal(err)
	}

	err = micro.RegisterSubscriber(pubsub.Topic, service.Server(), new(handler.Event))
	if err != nil {
		log.Fatal(err)
	}

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

//func sub() {
//	_, err := broker.Subscribe(pubsub.Topic, func(p broker.Event) error {
//		log.Infof("[sub] Received Body: %s, Header: %s\n", string(p.Message().Body), p.Message().Header)
//		return nil
//	})
//	if err != nil {
//		fmt.Println(err)
//	}
//}

func sub(ctx context.Context, p broker.Event) error {
	log.Infof("[sub] Received Body: %s, Header: %s\n", string(p.Message().Body), p.Message().Header)
	return nil
}
