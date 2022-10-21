package config

import (
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	RpcConfig
}

type RpcConfig struct {
	HellWorldEndpoint string `json:"HellWorldEndpoint"`
}
