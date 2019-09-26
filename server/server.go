package server

import (
	"fmt"
	"leveldb-server/configs"
	"leveldb-server/server/routers"

	"github.com/gin-gonic/gin"
)

func StartServer(model string) (err error) {
	serverModel := gin.DebugMode
	if model == "pro" {
		serverModel = gin.ReleaseMode
	} else if model == "test" {
		serverModel = gin.TestMode
	}
	gin.SetMode(serverModel)
	r := routers.InitRouters()
	host := fmt.Sprintf("%s:%s", configs.ServerConf.Host, configs.ServerConf.Port)
	fmt.Printf("server run on  %s......", host)
	err = r.Run(host)
	return err
}
