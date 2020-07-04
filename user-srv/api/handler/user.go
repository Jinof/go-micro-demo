package handler

import (
	"context"
	"encoding/json"
	"fmt"
	srv "github.com/Jinof/go-micro-demo/user-srv/user/genproto/srv"
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
	// 请注意这里，req.Header内部默认key为大写
	usernamePair, ok := req.Header["Username"]
	if !ok {
		log.Println("err: cannot get username in header")
	}
	fmt.Println(usernamePair.Values)

	data := new(struct {
		Name string
	})
	err := json.Unmarshal([]byte(req.Body), &data)
	userClient := srv.NewUserService("go.micro.service.srv", g.Client)
	rsp, err := userClient.Call(ctx, &srv.Request{Name: data.Name})
	if err != nil {
		return merr.InternalServerError("api.greeter.call", err.Error())
	}

	fmt.Println(res)
	fmt.Println(rsp)

	b, err := ResponseBody(0, "成功调用User.Call", rsp.Msg)
	if err != nil {
		return merr.InternalServerError("api.greeter.call", err.Error())
	}

	res.Body = b

	return nil
}
