package model

import (
	"jhgocms/config"
	"jhgocms/utils"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)
var DB *gorm.DB
func initMysqlEngine() {
	db,err := gorm.Open("mysql",config.SysConfig.MySql.Connstr)
	if err != nil{
		log.Print(err.Error())
	}
	// 设置到数据库的最大打开连接数
	db.DB().SetMaxOpenConns(config.SysConfig.MySql.MaxOpenConns)
	// 设置可重用连接的最大时间量
	db.DB().SetConnMaxLifetime(time.Duration(config.SysConfig.MySql.MaxLifetime)*time.Second)
	// 设置空闲连接池中的最大连接数
	db.DB().SetMaxIdleConns(config.SysConfig.MySql.MaxIdleConns)
	//设置表前缀
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return config.SysConfig.MySql.Prefix+defaultTableName
	}
	//是否开启sql语句调试
	db.LogMode(config.SysConfig.MySql.Debug)

	//整合第三方logrus日志，这里只是最简单的使用，更多复杂的可以查看logrus相关配置
	db.SetLogger(utils.NewLogrusLogger())

	//赋值给外层定义的变量
	DB = db
}


func init() {
	initMysqlEngine()
}