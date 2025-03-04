package main

import (
	"flag"
	"fmt"

	"movieReviewSystem/movieReviewSystemApi/thumbup/thumbupRpc/generateFileV1/internal/config"
	"movieReviewSystem/movieReviewSystemApi/thumbup/thumbupRpc/generateFileV1/internal/server"
	"movieReviewSystem/movieReviewSystemApi/thumbup/thumbupRpc/generateFileV1/internal/svc"
	"movieReviewSystem/movieReviewSystemApi/thumbup/thumbupRpc/pb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/thumbup.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		__.RegisterThumbUpServiceServer(grpcServer, server.NewThumbUpServiceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
