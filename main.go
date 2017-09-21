package main

import (
	"ginfw/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

/* func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})
	router.Run(":8000")
} */

/* func main() {
	router := gin.Default()
	router.POST("/upload", func(c *gin.Context) {
		name := c.PostForm("name")
		fmt.Println(name)
		file, header, err := c.Request.FormFile("upload")
		if err != nil {
			c.String(http.StatusBadRequest, "Bad request")
			return
		}

		filename := header.Filename

		fmt.Println(file, err, filename)

		out, err := os.Create(filename)
		if err != nil {
			log.Fatal(err)
			return
		}

		defer out.Close()

		_, err = io.Copy(out, file)
		if err != nil {
			log.Fatal(err)
		}
		c.String(http.StatusCreated, "upload successful")
	})
	router.Run(":8080")
}
*/

func main() {
	gin.SetMode(gin.DebugMode) //设置全局环境，当前为开发环境
	router := gin.Default()

	router.Use(controller.Middleware)

	//注册接口
	router.GET("/simple/server/get", controller.GetHandler)
	router.POST("/simple/server/post", controller.PostHandler)
	router.PUT("/simple/server/put", controller.PutHandler)
	router.DELETE("/simple/server/delete", controller.DeleteHandler)

	//监听端口
	http.ListenAndServe(":8000", router)
}

/* func Middleware(c *gin.Context) {
	fmt.Println("this is a middleware")
}

func GetHandler(c *gin.Context) {
	value, exist := c.GetQuery("key")
	if !exist {
		value = "the key is not exist!"
	}
	c.Data(http.StatusOK, "text/plain", []byte(fmt.Sprintf("get success! %s\n", value)))

	return
}

func PostHandler(c *gin.Context) {
	type JsonHolder struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}
	holder := JsonHolder{Id: 1, Name: "my name"}
	//若返回json数据，可以直接使用gin封装好的JSON方法
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
} */

/* func main() {
	gin.SetMode(gin.DebugMode)

	//监听端口
	http.ListenAndServe(":8000", r.Init())
}
*/
