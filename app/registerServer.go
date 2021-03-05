
package main

import (
    "reflect"
    "context"
    "runtime"
    "strings"
    "io/ioutil"

    "google.golang.org/grpc"
)

var (
    PathSep := string(os.PathSeparator)
)

type GRPCServerFilePath string

type RPCServer struct {

}

func (s *RPCServer) register(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {

}

// 反射获取RPC服务的信息
func scanServerFile(serverfiles []GRPCServerFilePath, path string) error {
    dir, err := ioutil.ReadDir(path)
    if err != nil {
        return err
    }

    for _, f := range dir {
        if f.isDir() {
            err := scanRPCServer(servers, path+PathSep+f)
            if err != nil {
                return err
            }
        } else {
            if ok := strings.HasSuffix(f.Name(), ".gw.go"), ok {
                serverfiles = append(serverfiles, path+PathSep+f.Name())
            }
        }
    }
}

func registerGRPCServer(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {

    path := "../go/service"
    serverfiles := make([]GRPCServerFilePath)
    scanServerFile(serverfiles, path)

    for _, file = range serverfiles {
        err := server.register(ctx, mux, *endpoint, opts)
        if err != nil {
            return err
        }
    }
}
