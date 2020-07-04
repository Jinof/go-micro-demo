package main

import (
	"context"
	"fmt"
	user "github.com/Jinof/go-micro-demo/user-srv/api/genproto/api"
	srv "github.com/Jinof/go-micro-demo/user-srv/api/genproto/srv"
	"github.com/micro/go-micro/v2"
	mApi "github.com/micro/go-micro/v2/api"
	hApi "github.com/micro/go-micro/v2/api/handler/api"
	api "github.com/micro/go-micro/v2/api/proto"
	"github.com/micro/go-micro/v2/client"
	merr "github.com/micro/go-micro/v2/errors"
	"log"
)

type User struct {
	Client client.Client
}

// Example.User is a method will be served by http request /example/srv
func (g *User) Call(ctx context.Context, req *api.Request, res *api.Response) error {
	usernamePair, ok := req.Header["Username"]
	if !ok {
		log.Println("err: cannot get username in header")
	}
	fmt.Println(usernamePair.Values)

	userClient := srv.NewUserService("go.micro.service.srv", g.Client)
	rsp, err := userClient.Call(ctx, &srv.Request{Name: fmt.Sprint(usernamePair.Values)})
	if err != nil {
		return merr.InternalServerError("api.greeter.call", err.Error())
	}

	fmt.Println(res)
	fmt.Println(rsp)

	res.Body = rsp.Msg

	return nil
}

func main() {

	service := micro.NewService(
		micro.Name("go.micro.api.example"),
	)

	service.Init()

	user.RegisterUserHandler(service.Server(), &User{Client: service.Client()}, mApi.WithEndpoint(&mApi.Endpoint{
		// The Rpc Method
		Name: "User.Call",
		// The HTTP paths. This can be a POSIX regex
		// Please check the url below before you assign the Path.
		// https://github.com/micro-in-cn/tutorials/tree/master/examples/micro-api
		Path: []string{"/user/call"},
		// The HTTP Methods for this endpoint.
		Method: []string{"POST", "GET"},
		// The API handler to use
		Handler: hApi.Handler,
	}))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
