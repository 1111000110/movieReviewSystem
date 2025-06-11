package main

import (
	"flag"
	"fmt"
	configurator "github.com/zeromicro/go-zero/core/configcenter"
	"github.com/zeromicro/go-zero/core/configcenter/subscriber"
	"movieReviewSystem/movieReviewSystemApi/user/user-api/generate-file/internal/middleware"

	"movieReviewSystem/movieReviewSystemApi/user/user-api/generate-file/internal/config"
	"movieReviewSystem/movieReviewSystemApi/user/user-api/generate-file/internal/handler"
	"movieReviewSystem/movieReviewSystemApi/user/user-api/generate-file/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

type TestSt struct {
	Name string `json:"name"`
}

func main() {
	flag.Parse()
	var c config.Config
	conf.MustLoad(*configFile, &c)
	ss := subscriber.MustNewEtcdSubscriber(subscriber.EtcdConf{
		Hosts: []string{"localhost:2379"}, // etcd 地址
		Key:   "test",                     // 配置key
	})
	cc := configurator.MustNewConfigCenter[config.Config](configurator.Config{
		Type: "json", // 配置值类型：json,yaml,toml
	}, ss)
	fmt.Println(c.Test)
	c, err := cc.GetConfig()
	if err != nil {
		panic(err)
	}
	println(c.Test)
	cc.AddListener(func() {
		c, err := cc.GetConfig()
		if err != nil {
			panic(err)
		}
		println(c.Test)
	})
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()
	server.Use(middleware.NewMiddleware().Handle) //使用全局中间件
	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)
	server.Start()
}
