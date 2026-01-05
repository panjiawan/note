package conf

import "github.com/panjiawan/go-lib/pkg/pcfg"

type Handle struct {
	path string
}

var handler *Handle = nil

func New(etcPath string) *Handle {
	if handler == nil {
		handler = &Handle{
			path: etcPath,
		}
	}

	return handler
}

func GetHandle() *Handle {
	return handler
}

func (s *Handle) Run() {
	s.LoadHttp()
	s.LoadMysql()
	s.LoadRedis()
}

func (s *Handle) GetHttpConf() *HttpConf {
	return pcfg.Get("http").(*HttpConf)
}

func (s *Handle) GetMysqlConf() *MysqlConf {
	return pcfg.Get("mysql").(*MysqlConf)
}

func (s *Handle) GetRedisConf() *RedisConf {
	return pcfg.Get("redis").(*RedisConf)
}

func (s *Handle) Close() {
}
