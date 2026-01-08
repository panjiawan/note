package middleware

import (
	"github.com/valyala/fasthttp"
	"strings"
)

func OutputMiddleware(ctx *fasthttp.RequestCtx) {
	SetCORSHeader(ctx)
}

func SetCORSHeader(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.Add(fasthttp.HeaderAccessControlAllowOrigin, "*")
	ctx.Response.Header.Add(fasthttp.HeaderAccessControlAllowMethods, "GET,POST,PUT,DELETE,OPTIONS")
	ctx.Response.Header.Add(fasthttp.HeaderAccessControlAllowCredentials, "true")
	ctx.Response.Header.Add(fasthttp.HeaderAccessControlAllowHeaders, "Origin,X-Requested-With,Content-Type,Authorization")
}

func ClientIP(ctx *fasthttp.RequestCtx) string {
	clientIP := string(ctx.Request.Header.Peek("X-Forwarded-For"))
	if index := strings.IndexByte(clientIP, ','); index >= 0 {
		clientIP = clientIP[0:index]
	}
	clientIP = strings.TrimSpace(clientIP)
	if len(clientIP) > 0 {
		return clientIP
	}
	clientIP = strings.TrimSpace(string(ctx.Request.Header.Peek("X-Real-Ip")))
	if len(clientIP) > 0 {
		return clientIP
	}
	return ctx.RemoteIP().String()
}
