package conf

import (
	"fmt"
	"github.com/panjiawan/go-lib/pkg/pcfg"
)

type RedisConf struct {
	Hosts []*RedisItem `yaml:"hosts"`
}

type RedisItem struct {
	Name    string `yaml:"name"`
	Host    string `yaml:"host"`
	Port    int    `yaml:"port"`
	Auth    string `yaml:"auth"`
	DB      int    `json:"db"`
	MinIdle int    `yaml:"minIdle"`
	MaxIdle int    `yaml:"maxIdle"`
	Timeout int    `yaml:"timeout"`
	Prefix  string `yaml:"prefix"`
}

func (s *Handle) LoadRedis() {
	path := fmt.Sprintf("%s/%s", s.path, "redis.yaml")
	err := pcfg.Load(pcfg.CfgTypeYaml, "redis", path, &RedisConf{})
	if err != nil {
		panic(err)
	}
}
