package router

import (
	"FRAME/service/app_api/code"
	"github.com/valyala/fasthttp"
)

func (h *HttpRouter) Filter(ctx *fasthttp.RequestCtx) code.OutputCode {
	//if internal.VerifyAuth(ctx) {
	//	return code.Success
	//}

	//return code.ErrorAuth

	return code.Success
}
