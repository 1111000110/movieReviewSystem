package main

import (
	"flag"
	"fmt"

	"movieReviewSystem/movieReviewSystemApi/group/groupRpc/generateFileV1/internal/config"
	"movieReviewSystem/movieReviewSystemApi/group/groupRpc/generateFileV1/internal/server"
	"movieReviewSystem/movieReviewSystemApi/group/groupRpc/generateFileV1/internal/svc"
	"movieReviewSystem/movieReviewSystemApi/group/groupRpc/pb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/group.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		__.RegisterGroupServiceServer(grpcServer, server.NewGroupServiceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
