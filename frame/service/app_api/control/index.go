package control

import "github.com/valyala/fasthttp"

func Index(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("Hello World!")
}
