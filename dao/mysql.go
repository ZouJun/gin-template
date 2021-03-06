package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go-gin-template/env"
	"time"
)

var (
	MysqlDB *gorm.DB
)

func InitMysql() (err error) {
	MysqlDB, err = gorm.Open("mysql", env.DatabaseUrl)

	if err != nil {
		return err
	}
	fmt.Println("初始化 mysql成功~~~")

	//设置空闲连接池中连接的最大数量
	MysqlDB.DB().SetMaxIdleConns(100)
	//设置打开数据库连接的最大数量
	MysqlDB.DB().SetMaxOpenConns(200)
	//设置了连接可复用的最大时间
	MysqlDB.DB().SetConnMaxLifetime(time.Hour)

	return MysqlDB.DB().Ping()
}

func DBClose() {
	MysqlDB.Close()
}
