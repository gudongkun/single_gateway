package global

import (
	"github.com/micro/go-micro/v2/config"
)

type Config struct {
	App     string  `json:"app"`
	AppAddr     string  `json:"appAddr"`
	Consul  string  `json:"consul"`
	Dsn     string  `json:"dsn"`
	Service Service `json:"service"`
	Jaeger  string  `json:"jaeger"`
}

type Service struct {
	UCenter string `json:"uCenter"`
}

var Conf Config

func init() {
	config.LoadFile("./config.json")
	config.Scan(&Conf)
}

// 加载 json 配置文件
