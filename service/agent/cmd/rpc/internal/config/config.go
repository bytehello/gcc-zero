package config

import "github.com/tal-tech/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	Cc struct {
		ReleasePath string
	}
}
