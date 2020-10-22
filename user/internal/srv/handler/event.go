package handler

import (
	"context"
	"github.com/Jinof/go-micro-demo/user/genproto/event"
	log "github.com/micro/go-micro/v2/logger"
)

// Event def
type Event struct {
}

func (e *Event) Sub(ctx context.Context, event *event.Event) error {
	log.Infof("[sub] id: %d, username: %s, data: %s\n", event.Id, event.Username, event.Data)
	return nil
}
