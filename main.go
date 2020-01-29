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
