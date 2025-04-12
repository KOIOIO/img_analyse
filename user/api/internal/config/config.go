package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	UserRpc zrpc.RpcClientConf
	Redis   struct {
		Addr     string
		Password string
		DB       int
	}
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
}
