package router

import "FRAME/service/app_api/control"

var getHandleList = map[string]*routerMethod{
	"/": {
		Handle: control.Index,
		Filter: false,
	},
}
