package main

import (
	"flag"
	"movieReviewSystem/movieReviewSystemApi/user/userApi/generateFileV1/internal/middleware"

	"movieReviewSystem/movieReviewSystemApi/user/userApi/generateFileV1/internal/config"
	"movieReviewSystem/movieReviewSystemApi/user/userApi/generateFileV1/internal/handler"
	"movieReviewSystem/movieReviewSystemApi/user/userApi/generateFileV1/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()
	server.Use(middleware.NewMiddleware().Handle) //使用全局中间件
	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)
	server.Start()
}
