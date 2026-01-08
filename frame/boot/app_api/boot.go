package app_api

import (
	"FRAME/conf"
	"FRAME/service/app_api/router"
	"FRAME/service/dao"
	"github.com/panjiawan/go-lib/pkg/plog"
)

type BootArgs struct {
	EtcPath string
	LogPath string
}

type BootHandler interface {
	Run()
	Close()
}

func Start(etcPath string, logPath string) {
	// load conf
	confHandle := conf.New(etcPath)
	confHandle.Run()

	// start log
	plog.Start(logPath, "app_api_log", confHandle.GetHttpConf().EnableDebug, confHandle.GetHttpConf().EnableStdout)

	// start signal
	go closeSignalListen()

	plog.Info("conf started")
	dao.Run()
	plog.Info("model started")
	route := router.New(confHandle.GetHttpConf())
	route.Run()
}

func Stay() {
	select {}
}

// 优雅关闭调用点
func close() {
}
