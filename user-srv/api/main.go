package main

import (
	"context"
	"fmt"
	"github.com/Jinof/go-micro-demo/user-srv/api/proto/greeter"
	"github.com/Jinof/go-micro-demo/user-srv/proto/user"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/api"
	"github.com/micro/go-micro/v2/api/handler/rpc"
	"github.com/micro/go-micro/v2/client"
	merr "github.com/micro/go-micro/v2/errors"
	"log"
)

type User struct {
	Client client.Client
}

// Example.User is a method will be served by http request /example/user
func (g *User) Call(ctx context.Context, req *greeter.CallRequest, res *greeter.CallResponse) error {
	log.Println("Received " + req.Name)

	userClient := user.NewUserService("go.micro.service.user", g.Client)
	rsp, err := userClient.Call(ctx, &user.Request{Name: req.Name})
	if err != nil {
		return merr.InternalServerError("api.greeter.call", err.Error())
	}

	fmt.Println(res)
	fmt.Println(rsp)

	res.Message = rsp.Msg

	return nil
}

func main() {

	service := micro.NewService(
		micro.Name("go.micro.api.example"),
	)

	service.Init()

	greeter.RegisterUserHandler(service.Server(), &User{Client: service.Client()}, api.WithEndpoint(&api.Endpoint{
		// The Rpc Method
		Name: "User.Call",
		// The HTTP paths. This can be a POSIX regex
		// Please check the url below before you assign the Path.
		// https://github.com/micro-in-cn/tutorials/tree/master/examples/micro-api
		Path: []string{"/user/call"},
		// The HTTP Methods for this endpoint.
		Method: []string{"POST", "GET"},
		// The API handler to use
		Handler: rpc.Handler,
	}))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
