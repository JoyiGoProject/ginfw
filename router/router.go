package router

import (
	c "ginfw/controller"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	router := router()
	return router
}

func router() *gin.Engine {
	router := gin.Default()

	//添加中间件
	router.Use(c.Middleware)

	//注册接口
	router.GET("/simple/server/get", c.GetHandler)
	router.POST("/simple/server/post", c.PostHandler)
	router.PUT("/simple/server/put", c.PutHandler)
	router.DELETE("/simple/server/delete", c.DeleteHandler)
	return router
}
