package main

import (
    "github.com/Jinof/go-micro-demo/user-srv/handler"
    "github.com/Jinof/go-micro-demo/user-srv/proto/user"
    "github.com/Jinof/go-micro-demo/user-srv/subscriber"
    "github.com/micro/go-micro/v2"
    log "github.com/micro/go-micro/v2/logger"
)

func main() {
    // New Service
    service := micro.NewService(
        micro.Name("go.micro.service.user"),
        micro.Version("latest"),
    )

    // Initialise service
    service.Init()

    // Register Handler
    user.RegisterUserHandler(service.Server(), new(handler.User))

    // Register Struct as Subscriber
    micro.RegisterSubscriber("go.micro.service.user", service.Server(), new(subscriber.User))

    // Run service
    if err := service.Run(); err != nil {
        log.Fatal(err)
    }
}
