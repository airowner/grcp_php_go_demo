package main

import (
	"context"

	"app/api/sayHello"
	"github.com/golang/glog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// Implements of GreeterServer
type HelloServer struct{}

func newHelloServer() sayHello.GreeterServer {
	return new(HelloServer)
}

func (s *HelloServer) SayHello(ctx context.Context, request *sayHello.HelloRequest) (reply *sayHello.HelloReply, err error) {
	header := new(ReplyHeader)
	header.setCode(0)
	header.setMessage("")
	reply := new(sayHello.HelloReply)
	reply.setHeader(header)
	reply.setMessage("OK")
	return
}

func (s *HelloServer) SayHelloAgain(ctx context.Context, request *sayHello.HelloRequest) (reply *SayHello.HelloReply, err error) {
	header := new(ReplyHeader)
	header.setCode(0)
	header.setMessage("")
	reply := new(HelloReply)
	reply.setHeader(header)
	reply.setMessage("OK")
	return
}

// func (s *HelloServer) Echo(ctx context.Context, msg *examples.SimpleMessage) (*examples.SimpleMessage, error) {
// 	glog.Info(msg)
// 	return msg, nil
// }
//
// func (s *HelloServer) EchoBody(ctx context.Context, msg *examples.SimpleMessage) (*examples.SimpleMessage, error) {
// 	glog.Info(msg)
// 	grpc.SendHeader(ctx, metadata.New(map[string]string{
// 		"foo": "foo1",
// 		"bar": "bar1",
// 	}))
// 	grpc.SetTrailer(ctx, metadata.New(map[string]string{
// 		"foo": "foo2",
// 		"bar": "bar2",
// 	}))
// 	return msg, nil
// }
