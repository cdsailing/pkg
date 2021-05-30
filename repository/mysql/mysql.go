package mysql

import (
	"github.com/cdsailing/pkg/config"
	"github.com/cdsailing/pkg/log"
	_mysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Init(conf *config.Config) *gorm.DB {
	if len(conf.Db.Connection) == 0 {
		panic("数据库连接字符未配置")
	}
	context, err := gorm.Open(_mysql.Open(conf.Db.Connection), &gorm.Config{})
	if err != nil {
		log.Errorf("数据库打开失败 %v", err.Error())
	}
	return context
}
