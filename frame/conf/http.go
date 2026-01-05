package conf

import (
	"fmt"
	"github.com/panjiawan/go-lib/pkg/pcfg"
)

type HttpConf struct {
	EnableStdout      bool   `yaml:"enableStdout"`
	EnableDebug       bool   `yaml:"enableDebug"`
	Https             bool   `yaml:"https"`
	HttpsCertFile     string `yaml:"httpsCertFile"`
	HttpsKeyFile      string `yaml:"httpsKeyFile"`
	HttpPort          int    `yaml:"httpPort"`
	RateLimitPerSec   int    `yaml:"rateLimitPerSec"`
	RateLimitCapacity int    `yaml:"rateLimitCapacity"`
}

func (s *Handle) LoadHttp() {
	path := fmt.Sprintf("%s/%s", s.path, "http.yaml")
	err := pcfg.Load(pcfg.CfgTypeYaml, "http", path, &HttpConf{})
	if err != nil {
		panic(err)
	}
}
