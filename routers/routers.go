package routers

import (
	"Gin-Todo/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	// 静态文件
	r.Static("./static", "static")
	// 模板文件
	r.LoadHTMLGlob("./templates/*")
	r.GET("/", controller.IndexHandler)

	// v1
	v1Group := r.Group("vi")
	{
		// 待办事项
		// 添加
		v1Group.POST("/todo", controller.CreateTodo)
		// 查看所有的待办事项
		v1Group.GET("/todo", controller.GetTodoList)
		// 修改一个待办事项
		v1Group.PUT("/todo:id", controller.UpdateTodo)
		// 删除一个待办事项
		v1Group.DELETE("/tode:id", controller.DeleteTodo)
	}
	return r
}