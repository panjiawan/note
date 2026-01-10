package router

import (
	"fmt"
)

var group = map[string]map[string]*routerMethod{}
var routesList = map[string]*routerMethod{}

type GroupRoute struct {
	Name string
}

func NewGroup(name string) *GroupRoute {
	res := &GroupRoute{
		Name: name,
	}
	if _, ok := group[name]; !ok {
		group[name] = make(map[string]*routerMethod)
	}

	return res
}

func (g *GroupRoute) Add(path string, function *routerMethod) error {
	group[g.Name][path] = function
	return nil
}

func parseGroup() {
	for gp, route := range group {
		for path, v := range route {
			allPath := fmt.Sprintf("/%s/%s", gp, path)
			routesList[allPath] = v
		}
	}
}
