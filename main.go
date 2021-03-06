package main

import (
	"Gin-Todo/dao"
	"Gin-Todo/models"
	"Gin-Todo/routers"
)

func main() {
	// 创建数据库 CREATE DATABASE Gin_Todo;
	// 连接数据库
	if err := dao.InitMySQL(); err != nil {
		panic(err)
	}
	defer dao.Close()
	// 模型绑定
	dao.DB.AutoMigrate(&models.Todo{})

	r := routers.SetupRouter()
	r.Run(":8002")
}
