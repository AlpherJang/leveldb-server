package main

import (
	"leveldb-server/configs"
	"leveldb-server/server"
	_ "net/http"
)

//@todo need to add glog and define some level for log
//level like this list :
//level 7 use to show some debug information
//level 10 use to show all error info and warning info some times can show some import info text
//level 20 use to show other info like license or help info

// @title Swagger Example API22222
// @version 1.0
// @description This is a sample server Petstore server
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache2.0
// @license.url http://www.apache.org/license/LICENSE-2.0.html
// @host petstore.swagger.io
// @BasePath /v1
func main() {
	err := configs.InitConfig("config.yaml")
	if err != nil {
		panic(err)
	}
	err = server.StartServer(configs.ServerConf.Model)
	if err != nil {
		panic(err)
	}
}
