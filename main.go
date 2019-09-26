package main

import (
	"leveldb-server/configs"
	"leveldb-server/server"
	_ "net/http"
)

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
