package handler

import (
	user "github.com/Jinof/go-micro-demo/user_srv/api/genproto/api"
	"github.com/micro/go-micro/v2"
	mApi "github.com/micro/go-micro/v2/api"
	hApi "github.com/micro/go-micro/v2/api/handler/api"
	"github.com/micro/go-micro/v2/util/log"
)

func RegisterHandler(service micro.Service) {
	if err := registerUser(service); err != nil {
		log.Error(err)
	}
}

func registerUser(service micro.Service) error {
	return user.RegisterUserHandler(service.Server(), &User{Client: service.Client()}, mApi.WithEndpoint(&mApi.Endpoint{
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
}
