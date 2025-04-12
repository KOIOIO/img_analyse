package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"img_analyse/user/api/internal/config"
	"img_analyse/user/commn/init_redis"
	"img_analyse/user/rpc/userclient"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc userclient.User
	Redis   *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	Redis := init_redis.Init_redis()
	if Redis == nil {
		panic("redis init fail")
	}
	defer Redis.Close()
	return &ServiceContext{
		Config:  c,
		UserRpc: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
