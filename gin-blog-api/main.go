package main

import (
	"gin-blog/model"
	"gin-blog/routes"
)

func main() {
	//引用数据库
	model.InitDb()
	//初始化路由
	routes.InitRoute()
}
