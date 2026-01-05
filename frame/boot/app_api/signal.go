package app_api

import (
	"github.com/panjiawan/go-lib/pkg/plog"
	"github.com/panjiawan/go-lib/pkg/psignal"
	"go.uber.org/zap"
	"os"
)

func closeSignalListen() {
	defer func() {
		if err := recover(); err != nil {
			plog.Error("signal listen error", zap.Any("err", err))
		}
	}()

	signal := psignal.New()
	signal.Register(os.Interrupt, func(signal os.Signal, args interface{}) {
		close()
		os.Exit(0)
	})
	signal.Listen()
}
