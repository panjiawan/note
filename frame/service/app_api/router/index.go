package router

import (
	"FRAME/service/app_api/control"
	"github.com/panjiawan/go-lib/pkg/phttp"
)

func init() {
	groupRoute := NewGroup("index")
	groupRoute.Add("get", &routerMethod{control.Index, phttp.MethodGet, false})
	groupRoute.Add("post", &routerMethod{control.Index, phttp.MethodPost, false})
}
