package main

import (
	"go-gin-template/dao"
	"go-gin-template/router"
)

func main() {
	//初始化数据库链接
	dao.InitMysql()
	//关闭DB链接
	defer dao.DBClose()

	//路由
	r := router.InitRouter()

	r.Run(":50000")
}
