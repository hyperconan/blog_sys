package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	Database DatabaseConfig
}

type DatabaseConfig struct {
	DSN string
}
