package config

import (
	"github.com/i-Things/things/shared/conf"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Event conf.EventConf //和things内部交互的设置
	Mysql struct {
		DataSource string
	}
	CacheRedis cache.ClusterConf
	DmRpc      conf.RpcClientConf `json:",optional"`
	DiRpc      conf.RpcClientConf `json:",optional"`
}
