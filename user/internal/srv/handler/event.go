package handler

import (
	"fmt"
	"github.com/Jinof/go-micro-demo/user/pkg/pubsub"
	"github.com/micro/go-micro/v2/broker"
	log "github.com/micro/go-micro/v2/logger"
)

type Event struct {
}

func (e *User) Sub() {
	_, err := broker.Subscribe(pubsub.Topic, func(event broker.Event) error {
		log.Infof("[sub] received body %s, header: %s\n", string(event.Message().Body), event.Message().Header)
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
}
