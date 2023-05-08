package main

import (
	"context"
	"github.com/gin-gonic/gin"
	v1 "github.com/gittduse/protoc-gen-go-gin/examples/hello/api/hello/v1"
	"github.com/gittduse/protoc-gen-go-gin/pkg/errors"
	"log"
)

type helloService struct {
}

func (s *helloService) GetHello(ctx context.Context, request *v1.GetHelloRequest) (response *v1.GetHelloReply, err error) {
	if request.GetName() != "aaron" {
		err = errors.New(400, "PARAMS ERROR", "he name is invalid")
	} else {
		response = &v1.GetHelloReply{
			Message: "hello " + request.GetName(),
		}
	}
	return
}
func (s *helloService) PostHello(ctx context.Context, request *v1.PostHelloRequest) (response *v1.PostHelloReply, err error) {
	log.Println("request name", request.GetName())
	return
}
func main() {
	e := gin.New()
	v1.RegisterHelloServiceHttpServer(e, &helloService{})
	e.Run()
}
