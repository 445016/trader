package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	AppKey string
	AppSecret string
	SpotRestHost string
}
