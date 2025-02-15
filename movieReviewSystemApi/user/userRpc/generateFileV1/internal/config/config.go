package config

import (
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	DB struct {
		User struct {
			Url        string
			Db         string
			Collection string
		}
		UserRelations struct {
			Url        string
			Db         string
			Collection string
		}
	}
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
}
