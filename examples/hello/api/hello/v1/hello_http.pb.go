// Code generated by protoc-gen-go-gin. DO NOT EDIT.

package v1

import (
	context "context"
	gin "github.com/gin-gonic/gin"
	binding "github.com/gittduse/protoc-gen-go-gin/pkg/transport/http/binding"
	metadata "google.golang.org/grpc/metadata"
)

// context.
// metadata.
// gin.
// binding.

type HelloServiceHttpServer interface {
	GetHello(context.Context, *GetHelloRequest) (*GetHelloReply, error)

	PostHello(context.Context, *PostHelloRequest) (*PostHelloReply, error)
}

type HelloService struct {
	server HelloServiceHttpServer
	router *gin.Engine
}

func RegisterHelloServiceHttpServer(r *gin.Engine, srv HelloServiceHttpServer) {
	s := HelloService{
		server: srv,
		router: r,
	}
	s.RegisterService()
}

func (s *HelloService) GetHello_0(ctx *gin.Context) {
	var in GetHelloRequest

	if err := binding.MapProto(&in, binding.RestParams(ctx)); err != nil {
		binding.Response(ctx, nil, err)
		return
	}

	if err := binding.MapProto(&in, binding.QueryParams(ctx)); err != nil {
		binding.Response(ctx, nil, err)
		return
	}

	md := metadata.New(nil)
	for k, v := range ctx.Request.Header {
		md.Set(k, v...)
	}
	newCtx := metadata.NewIncomingContext(ctx, md)
	out, err := s.server.GetHello(newCtx, &in)
	binding.Response(ctx, out, err)
}

func (s *HelloService) GetHello_1(ctx *gin.Context) {
	var in GetHelloRequest

	if err := binding.MapProto(&in, binding.RestParams(ctx)); err != nil {
		binding.Response(ctx, nil, err)
		return
	}

	if err := binding.MapProto(&in, binding.QueryParams(ctx)); err != nil {
		binding.Response(ctx, nil, err)
		return
	}

	md := metadata.New(nil)
	for k, v := range ctx.Request.Header {
		md.Set(k, v...)
	}
	newCtx := metadata.NewIncomingContext(ctx, md)
	out, err := s.server.GetHello(newCtx, &in)
	binding.Response(ctx, out, err)
}

func (s *HelloService) PostHello_0(ctx *gin.Context) {
	var in PostHelloRequest

	if err := ctx.ShouldBindJSON(&in); err != nil {
		binding.Response(ctx, nil, err)
		return
	}

	md := metadata.New(nil)
	for k, v := range ctx.Request.Header {
		md.Set(k, v...)
	}
	newCtx := metadata.NewIncomingContext(ctx, md)
	out, err := s.server.PostHello(newCtx, &in)
	binding.Response(ctx, out, err)
}

func (s *HelloService) RegisterService() {

	s.router.Handle("GET", "/v1/hello_world/:name", s.GetHello_0)

	s.router.Handle("GET", "/v1/hello/:name", s.GetHello_1)

	s.router.Handle("POST", "/v1/hello", s.PostHello_0)

}
