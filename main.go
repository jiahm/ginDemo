package main

import (
	"github.com/gin-gonic/gin"
	"ginDemo/router"
	"ginDemo/config"
)

func main() {
	//fmt.Println("hello,go app!")
	//r := gin.Default()
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//})
	//r.Run() // listen and serve on 0.0.0.0:8080

	gin.SetMode(gin.DebugMode) // 默认为 debug 模式，设置为发布模式
	engine := gin.Default()
	router.InitRouter(engine) // 设置路由
	engine.Run(config.PORT)
}