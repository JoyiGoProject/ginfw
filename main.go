package main

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
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

/* func main() {
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
} */

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

/* func main() {
	//获取一个路由handler
	router := gin.Default()

	router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		fmt.Println("获取的name是：", name)
		action := c.Param("action")
		fmt.Println("获取的action是：", action)
		// c.String(http.StatusOK, "Hello World")
		c.String(http.StatusOK, "Hello %s", name+"\n")
		c.String(http.StatusOK, name+" is "+action)
	})
	router.Run(":8000")
}
*/

//极简实例
/* func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
} */

//构建restful api
func main() {
	//获取路由实例
	router := gin.Default()
	gin.SetMode(gin.DebugMode)
	// gin.SetMode(gin.ReleaseMode)
	v1 := router.Group("/api/v1/userinfo")
	{
		v1.POST("/", CreateUser)
		v1.GET("/", FetchAllUsers)
		v1.GET("/:id", FetchSingleUser)
		v1.PUT("/:id", UpdateUser)
		v1.DELETE("/:id", DeleteUser)
	}

	router.Run()
}

type Person struct {
	Id   int64
	Name string
	Age  int64
}

//查询单个用户
func FetchSingleUser(c *gin.Context) {
	id := c.Param("id")
	db, err := sql.Open("mysql", "root:admin@/test?charset=utf8")
	checkErr(err)

	defer db.Close()

	err = db.Ping()
	checkErr(err)

	var (
		person Person
		result gin.H
	)

	row := db.QueryRow("select * from user where id = ?;", id)
	row.Scan(&person.Id, &person.Name, &person.Age)
	if err != nil {
		checkErr(err)
		//If no results send null
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": person,
			"count":  1,
		}
	}
	c.JSON(http.StatusOK, result)
}

//创建用户
func CreateUser(c *gin.Context) {

}

//查询所有用户
func FetchAllUsers(c *gin.Context) {

}

//更新用户
func UpdateUser(c *gin.Context) {

}

//删除用户
func DeleteUser(c *gin.Context) {

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
