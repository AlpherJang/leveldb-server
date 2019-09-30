package configs

import (
	"fmt"
	"leveldb-server/database"

	"github.com/jinzhu/configor"
)

type Config struct {
	Model   string `env:"model" default:"debug"`
	AppName string `default:"leveldb-server"`
	DbPath  string `env:"DBPath"`
	Port    string `env:"Port" default:"9650"`
	Host    string `env:"Host" default:127.0.0.1`
}

var ServerConf Config
var Db *database.LevelDB

func InitConfig(configPath string) (err error) {
	ServerConf = Config{}
	err = configor.Load(&ServerConf, configPath)
	if err != nil {
		fmt.Println("there is some thing wrong when init server config")
		return err
	}
	Db = database.NewLevelDB(ServerConf.DbPath)
	err = Db.OpenDataBase(nil)
	if err != nil {
		fmt.Println("there is some thing wrong when open db file")
		return err
	}
	return nil
}
