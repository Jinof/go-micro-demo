package main

import (
	"github.com/Jinof/go-micro-demo/user/internal/api/handler"
	"github.com/micro/go-micro/v2"
	ocPlugin "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	oc "github.com/opentracing/opentracing-go"
	"log"
)

func main() {

	service := micro.NewService(
		micro.Name("go.micro.api.example"),
		micro.Version("latest"),
		micro.WrapHandler(ocPlugin.NewHandlerWrapper(oc.GlobalTracer())),
	)

	service.Init()

	handler.RegisterHandler(service)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
