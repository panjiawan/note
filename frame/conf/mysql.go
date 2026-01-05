package conf

import (
	"fmt"
	"github.com/panjiawan/go-lib/pkg/pcfg"
)

type MysqlConf struct {
	Hosts []*MysqlItem `yaml:"hosts"`
}

type MysqlItem struct {
	Name     string `yaml:"name"`
	Debug    bool   `yaml:"debug"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DB       string `yaml:"db"`
	Prefix   string `yaml:"prefix"`
	MaxIdle  int    `yaml:"maxIdle"`
	MaxOpen  int    `yaml:"maxOpen"`
}

func (s *Handle) LoadMysql() {
	path := fmt.Sprintf("%s/%s", s.path, "mysql.yaml")
	err := pcfg.Load(pcfg.CfgTypeYaml, "mysql", path, &MysqlConf{})
	if err != nil {
		panic(err)
	}
}
