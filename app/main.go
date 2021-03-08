package main

import (
	"context"
	"flag"
	"log"
	"net/http"

	"github.com/gin-gonic/gin" // 导入 Gin 依赖
	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

var (
	grpcOptions        []grpc.DialOption
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:9090", "gRpc server endpoint")
)

func httpRun() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	//err := sayHello.RegisterGreeterServer(ctx, mux, *grpcServerEndpoint, opts)
	err := registerRPCServer(ctx, mux, opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(":8080", mux)
}

func httpRun() {
	router := gin.Default() // 获取 engine

	err := router.Run(":9090") // 指定端口，运行 Gin
	if err != nil {
		log.Panicln("服务器启动失败：", err.Error())
	}

	println("server starting ...")
}

func main() {
	flag.Parse()
	defer glog.Flush()

	go httpRun()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
