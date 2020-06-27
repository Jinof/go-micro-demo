package subscriber

import (
	"context"
	"github.com/Jinof/go-micro-demo/user-srv/proto/user"
	log "github.com/micro/go-micro/v2/logger"
)

type User struct{}

func (e *User) Handle(ctx context.Context, msg *user.Message) error {
	log.Info("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *user.Message) error {
	log.Info("Function Received message: ", msg.Say)
	return nil
}
