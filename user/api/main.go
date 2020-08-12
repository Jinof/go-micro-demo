package main

import (
	"github.com/Jinof/go-micro-demo/user/api/handler"
	"github.com/micro/go-micro/v2"
	"log"
)

func main() {

	service := micro.NewService(
		micro.Name("go.micro.api.example"),
	)

	service.Init()

	handler.RegisterHandler(service)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
