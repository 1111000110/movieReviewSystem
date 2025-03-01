package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	DB struct {
		Group struct {
			Url        string
			Db         string
			Collection string
		}
	}
}
