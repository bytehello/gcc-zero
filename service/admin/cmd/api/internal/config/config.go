package config

import "github.com/tal-tech/go-zero/rest"

type Config struct {
	rest.RestConf
	Mysql struct {
		DataSource string
	}
	Jwt struct {
		AccessSecret string
		AccessExpire int64
	}
}
