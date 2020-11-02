package main

import (
	"github.com/Jinof/go-micro-demo/pkg/plugins/auth"
	"github.com/micro/micro/v2/client/api"
	"github.com/micro/micro/v2/cmd"
)

func main() {
	cmd.Init()
}

func init() {
	if err := api.Register(auth.NewPlugin()); err != nil {
		panic(err)
	}
}

// Test trigger github actions
