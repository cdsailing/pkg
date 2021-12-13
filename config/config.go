package config

import (
	"github.com/cdsailing/pkg/log"
	"github.com/cdsailing/pkg/utils"
	"github.com/jinzhu/configor"
	"os"
)

var (
	Conf = &Config{}
)

type Config struct {
	App struct {
		Name    string
		Version string
	}
	Server struct {
		ApiPrefix string
		Port      int `default:8080`
		Version   string
		PageSize  int `default:10`
	}
	Db struct {
		Connection string
		Type       string
		Debug      bool
	}
	Log struct {
		Level int
	}
}

func init() {
	var err error
	BaseDir := utils.GetCurrentDirectory()
	path := BaseDir
	app := os.Getenv("app")
	if len(app) > 0 {
		path = path + "/" + app
	}
	err = configor.Load(Conf, path+"/conf.d/app.toml")
	if err != nil {
		log.Errorf("无法读取配置文件 %v", err)
	}
}

func Init(instance interface{}) {
	var err error
	BaseDir := utils.GetCurrentDirectory()
	path := BaseDir
	app := os.Getenv("app")
	if len(app) > 0 {
		path = path + "/" + app
	}
	err = configor.Load(instance, path+"/conf.d/app.toml")
	if err != nil {
		log.Errorf("无法读取配置文件 %v", err)
	}
}
