package routers

import (
	"leveldb-server/server/handlers"

	"github.com/gin-gonic/gin"
)

type WriterRouters struct {
	name       string
	routerList map[string]RouterItem
}

func NewWriterRouters() *WriterRouters {
	writerHandler := handlers.WriterHandler{}
	routerList := map[string]RouterItem{
		"putOne": {MethodPost, writerHandler.PutOne},
	}
	return &WriterRouters{"writer", routerList}
}

func (n WriterRouters) GetGroupName() string {
	return "/" + n.name
}

func (n WriterRouters) GetRouterGroup(r *gin.RouterGroup) {
	for key, route := range n.routerList {
		BuildRouter(r, route, key)
	}
}
