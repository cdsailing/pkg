package repository

import (
	"github.com/cdsailing/pkg/config"
	"github.com/cdsailing/pkg/repository/mysql"
	"github.com/cdsailing/pkg/repository/pgsql"
	"github.com/cdsailing/pkg/repository/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"strings"
	"time"
)

var (
	DbContext *gorm.DB
)

func init() {
	conf := &config.Config{}
	config.Init(conf)
	if strings.ToLower(conf.Db.Type) == "pgsql" {
		DbContext = pgsql.Init(conf)
	} else if strings.ToLower(conf.Db.Type) == "mysql" {
		DbContext = mysql.Init(conf)
	} else if strings.ToLower(conf.Db.Type) == "sqlite" {
		DbContext = sqlite.Init(conf)
	}
	if conf.Db.Debug {
		DbContext.Debug()
	}
if DbContext!=nil{
	DbContext.Logger = logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,   // 慢 SQL 阈值
			LogLevel:      logger.Silent, // Log level
			Colorful:      false,         // 禁用彩色打印
		},
	)
}
}
