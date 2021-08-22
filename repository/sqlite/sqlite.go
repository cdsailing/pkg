package sqlite

//
//import (
//	"github.com/cdsailing/pkg/config"
//	"gorm.io/driver/sqlite"
//	"gorm.io/gorm"
//	"os"
//)
//
//func Init(conf *config.Config) *gorm.DB {
//	dir, _ := os.Getwd()
//	path := dir
//	app := os.Getenv("app")
//	if len(app) > 0 {
//		path = path + "/" + app
//	}
//	path += conf.Db.Connection
//	_, err := os.Stat(path)
//	if err != nil {
//		f, _ := os.Create(path)
//		defer f.Close()
//	}
//
//	context, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
//	if err != nil {
//		os.Stderr.WriteString("数据库打开失败 " + err.Error())
//		panic(0)
//	}
//	return context
//}
