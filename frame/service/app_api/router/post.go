package router

import "FRAME/service/app_api/control"

var postHandleList = map[string]*routerMethod{
	"/": {
		Handle: control.Index,
		Filter: false,
	},
}
