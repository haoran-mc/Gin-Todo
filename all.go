package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

var (
	DB *gorm.DB
)

func initMySql() (err error) {
	dsn := "root:haoran232@(127.0.0.1:3306)/GoConnectMySql?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return
	}
	return DB.DB().Ping()
}

func main() {
	/*
		创建数据库
		sql: CREATE DATABASE bubble;
	*/
	// 连接数据库
	err := initMySql()
	if err != nil {
		panic(err)
	}
	defer DB.Close() // 程序退出，关闭数据库

	// 绑定模型
	DB.AutoMigrate(&Todo{})

	r := gin.Default()
	// 告诉gin框架模板文件引用的静态文件去哪里找，当gin框架找/static时，就会去找当前目录下的static文件夹
	r.Static("/static", "static")
	// 告诉gin框架去哪里找模板文件
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		// 响应，渲染文件，返回空
		c.HTML(http.StatusOK, "index.html", nil)
	})

	v1Group := r.Group("v1")
	{
		// 返回页面之后，有哪些操作
		// 添加
		v1Group.POST("/todo", func(c *gin.Context) {
			// 前端页面填写待办事项，点击提交，会发送请求到这里
			// 1. 从请求中把数据拿出来
			var todo Todo
			c.BindJSON(&todo)
			// 2. 存入数据库
			err := DB.Create(todo).Error
			// 3. 返回响应
			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"error": err.Error(),
				})
			} else {
				c.JSON(http.StatusOK, todo)
			}
		})
		// 查看
		v1Group.GET("/todo", func(c *gin.Context) {
			// 查询todo这个表里的所有数据
			// 1. 从数据库中拿出数据
			var todoList []Todo
			err := DB.Find(&todoList).Error
			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"error": err.Error(),
				})
			} else { // 如果拿到数据，就把todos这个表中的所有数据呈现出来
				c.JSON(http.StatusOK, todoList)
			}
			// 2.
		})
		// 修改
		v1Group.PUT("/todo/:id", func(c *gin.Context) {
			id, ok := c.Params.Get("id")
			if !ok {
				c.JSON(http.StatusOK, gin.H{
					"error": "无效的id",
				})
				return
			}
			var todo Todo
			err := DB.Where("id=?", id).First(&todo).Error
			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"error": err.Error(),
				})
				return
			}
			c.BindJSON(&todo)
			err = DB.Save(&todo).Error
			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"error": err.Error(),
				})
			} else {
				c.JSON(http.StatusOK, todo)
			}
		})
		// 删除
		v1Group.DELETE("/todo/:id", func(c *gin.Context) {
			id, ok := c.Params.Get("id")
			if !ok {
				c.JSON(http.StatusOK, gin.H{
					"error": "无效的id",
				})
				return
			}
			err := DB.Where("id=?", id).Delete(Todo{}).Error
			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"error": err.Error(),
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					id: "deleted",
				})
			}
		})
	}
	r.Run()
}
