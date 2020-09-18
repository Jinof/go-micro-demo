package handler

import (
	"context"
	"fmt"
	user "github.com/Jinof/go-micro-demo/user/genproto/srv"

	log "github.com/micro/go-micro/v2/logger"
)

type User struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *User) Call(ctx context.Context, req *user.Request, rsp *user.Response) error {
	log.Info("Received User.Call request")
	fmt.Printf("received data: %s from internal: %s", req.Data, req.Name)
	rsp.Msg = "Hello " + req.Name + " your data has been received"
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *User) Stream(ctx context.Context, req *user.StreamingRequest, stream user.User_StreamStream) error {
	log.Infof("Received User.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&user.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *User) PingPong(ctx context.Context, stream user.User_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&user.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}

// Hello is a bidirectional stream handler called via client.Stream or the generated client code
func (e *User) Hello(ctx context.Context, req *user.HelloReq, rsp *user.HelloRes) error {
	log.Info("Received User.Hello request")
	rsp.Msg = "hello from service hello"
	return nil
}
