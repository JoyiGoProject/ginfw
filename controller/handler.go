package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type JsonHolder struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

//中间件
func Middleware(c *gin.Context) {
	fmt.Println("this is a middleware")
}

func GetHandler(c *gin.Context) {
	value, exist := c.GetQuery("key")
	if !exist {
		value = "the key is not exist"
	}
	c.Data(http.StatusOK, "text/plain", []byte(fmt.Sprintf("get success! %s\n", value)))

	return
}

func PostHandler(c *gin.Context) {
	holder := JsonHolder{Id: 1, Name: "my name"}
	//若返回json数据，可以直接使用gin封装好的Json方法
	c.JSON(http.StatusOK, holder)
	return
}

func PutHandler(c *gin.Context) {
	c.Data(http.StatusOK, "text/plain", []byte("put success!\n"))
	return
}

func DeleteHandler(c *gin.Context) {
	c.Data(http.StatusOK, "text/plain", []byte("delete success!\n"))
	return
}
