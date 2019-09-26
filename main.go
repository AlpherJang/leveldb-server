package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/configor"
	"github.com/syndtr/goleveldb/leveldb"
	"leveldb-server/config"
	_ "net/http"
)

func main() {
	var err error
	config.ServerConf = config.Config{}
	err = configor.Load(&config.ServerConf, "config.yaml")
	if err != nil {
		panic(err)
	}
	config.Db, err = leveldb.OpenFile(config.ServerConf.DbPath, nil)
	if err != nil {
		fmt.Println("there is some thing wrong")
		panic(err)
	}
	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	fmt.Println("server run on 9650 pod")
	err = r.Run()
	if err != nil {
		panic(err)
	}
	//gin.SetMode(gin.DebugMode)
	//r := gin.Default()
}
