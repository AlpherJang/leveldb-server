package routers

import (
	"leveldb-server/server/handlers"

	"github.com/gin-gonic/gin"
)

type JobRouters struct {
	name       string
	routerList map[string]RouterItem
}

func NewJobRouters() *JobRouters {
	jobHandler := handlers.JobHandler{}
	routerList := map[string]RouterItem{
		"": {MethodGet, jobHandler.List},
	}
	return &JobRouters{"job", routerList}
}

func (n JobRouters) GetGroupName() string {
	return "/" + n.name
}

func (n JobRouters) GetRouterGroup(r *gin.RouterGroup) {
	for key, route := range n.routerList {
		BuildRouter(r, route, key)
	}
}
