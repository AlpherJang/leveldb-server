package server

import (
	"fmt"
	"github.com/gin-contrib/gzip"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"leveldb-server/configs"
	"leveldb-server/server/routers"

	"github.com/gin-gonic/gin"
	_ "leveldb-server/docs"
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
	r.Use(gzip.Gzip(gzip.BestSpeed))
	host := fmt.Sprintf("%s:%s", configs.ServerConf.Host, configs.ServerConf.Port)
	fmt.Printf("server run on  %s.......", host)
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	err = r.Run(host)
	return err
}
