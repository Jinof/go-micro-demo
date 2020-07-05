package main

import (
	"github.com/Jinof/go-micro-demo/pkg/plugins/login"
	"github.com/micro/micro/v2/client/api"
	"github.com/micro/micro/v2/cmd"
)

func main() {
	cmd.Init()
}

func init() {
	if err := api.Register(login.NewPlugin()); err != nil {
		panic(err)
	}
}
