package router

import (
	"FRAME/conf"
	"FRAME/service/app_api/code"
	"FRAME/service/app_api/internal"
	"FRAME/service/app_api/middleware"
	"fmt"
	"github.com/panjiawan/go-lib/pkg/phttp"
	"github.com/panjiawan/go-lib/pkg/plog"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
	"runtime/debug"
)

type HttpRouter struct {
	httpConf *conf.HttpConf
	handle   *phttp.Service
}

type routerMethod struct {
	Handle func(ctx *fasthttp.RequestCtx)
	Filter bool
}

func New(cfg *conf.HttpConf) *HttpRouter {
	return &HttpRouter{
		httpConf: cfg,
	}
}

// Run 启动函数
func (h *HttpRouter) Run() {
	h.handle = phttp.New(
		phttp.WithAddress("", h.httpConf.HttpPort),
		phttp.WithCertificate(h.httpConf.HttpsCertFile, h.httpConf.HttpsKeyFile),
		phttp.WithRate(h.httpConf.RateLimitPerSec, h.httpConf.RateLimitCapacity),
	)

	h.Register()

	if err := h.handle.Run(); err != nil {
		plog.Error("http server start error", zap.Error(err))
	}
}

func (h *HttpRouter) Register() {
	for path := range getHandleList {
		h.handle.Register(path, phttp.MethodGet, h.PrepareCall)
		plog.Info("register", zap.String("path", path), zap.String("method", phttp.MethodGet))

	}

	for path := range postHandleList {
		h.handle.Register(path, phttp.MethodPost, h.PrepareCall)
		plog.Info("register", zap.String("path", path), zap.String("method", phttp.MethodPost))
	}
}

func (h *HttpRouter) PrepareCall(ctx *fasthttp.RequestCtx) {
	defer func() {
		if e := recover(); e != nil {
			plog.Error("panic prepareCall", zap.Error(fmt.Errorf("%v", e)), zap.String("trace", string(debug.Stack())))
		}
	}()

	h.options(ctx)
	path := string(ctx.URI().Path())
	method := string(ctx.Request.Header.Method())

	if method == "GET" {
		if _, ok := getHandleList[path]; ok {
			if getHandleList[path].Filter {
				if outputCode := h.Filter(ctx); outputCode != code.Success {
					internal.OutputError(ctx, outputCode)
					return
				}
			}
			getHandleList[path].Handle(ctx)
		}
	} else if method == "POST" {
		if _, ok := postHandleList[path]; ok {
			if postHandleList[path].Filter {
				if outputCode := h.Filter(ctx); outputCode != code.Success {
					internal.OutputError(ctx, outputCode)
					return
				}
			}
			postHandleList[path].Handle(ctx)
		}
	}
}

func (h *HttpRouter) options(ctx *fasthttp.RequestCtx) {
	// 处理OPTIONS
	middleware.SetCORSHeader(ctx)
	ctx.SetStatusCode(fasthttp.StatusAccepted)
}

func (h *HttpRouter) Close() {
}
