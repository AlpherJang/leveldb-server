package routers

import (
	"github.com/gin-gonic/gin"
	"leveldb-server/server/handlers"
)

type ReaderRouters struct {
	name       string
	routerList map[string]RouterItem
}

func NewReaderRouters() *ReaderRouters {
	readerHandler := handlers.ReaderHandler{}
	routerList := map[string]RouterItem{
		"readOne/:key": {MethodGet, readerHandler.ReadOne},
	}
	return &ReaderRouters{name: "reader", routerList: routerList}
}

func (n ReaderRouters) GetGroupName() string {
	return "/" + n.name
}

func (n ReaderRouters) GetRouterGroup(r *gin.RouterGroup) {
	for key, route := range n.routerList {
		BuildRouter(r, route, key)
	}
}
