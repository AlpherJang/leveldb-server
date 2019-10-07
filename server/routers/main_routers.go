package routers

import (
	"github.com/gin-gonic/gin"
)

type RouterInterface interface {
	GetGroupName() string
	GetRouterGroup(r *gin.RouterGroup)
}

type RouterItem struct {
	method   string
	function func(c *gin.Context)
}

const (
	MethodGet    = "get"
	MethodPost   = "post"
	MethodPut    = "put"
	MethodDelete = "delete"
	MethodPatch  = "patch"
)

func BuildRouter(engine *gin.RouterGroup, item RouterItem, path string) {
	switch item.method {
	case MethodGet:
		engine.GET(path, item.function)
	case MethodPost:
		engine.POST(path, item.function)
	case MethodPut:
		engine.PUT(path, item.function)
	case MethodDelete:
		engine.DELETE(path, item.function)
	case MethodPatch:
		engine.PATCH(path, item.function)
	default:
		panic("router method error, method value:" + item.method + " route path:" + path)
	}
}

func getRoutersSlice() []RouterInterface {
	return []RouterInterface{NewWriterRouters(), NewReaderRouters()}
}

func GetRouterGroup(routers RouterInterface, r *gin.Engine) {
	routerGroup := r.Group(routers.GetGroupName())
	{
		routers.GetRouterGroup(routerGroup)
	}
}

func InitRouters() *gin.Engine {
	r := gin.Default()
	routerSlice := getRoutersSlice()
	for _, item := range routerSlice {
		GetRouterGroup(item, r)
	}
	return r
}
