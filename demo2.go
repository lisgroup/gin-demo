package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// ##  路由(Route)方法
	// ### 支持方法
	r.Any("/ping", anything)
	// r.POST("/post", posting)
	// r.PUT("/put", putting)
	// r.DELETE("/delete", deleting)
	// r.PATCH("/patch", patching)
	// r.HEAD("/head", head)
	// r.OPTIONS("/options", options)

	// ### 解析动态路径参数
	r.GET("/user/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(200, gin.H{
			"id": id,
		})
	})

	// ### 1.获取Query参数
	// 匹配 /search?keyword=xxx&weight=xxx ，weight可选
	r.GET("/search", func(c *gin.Context) {
		keyword := c.Query("keyword")
		weight := c.DefaultQuery("weight", "")
		c.JSON(200, gin.H{
			"keyword": keyword,
			"weight":  weight,
		})
	})
	// ### 2.获取POST参数
	// POST application/x-www-form-urlencoded
	r.POST("/login", func(c *gin.Context) {
		username := c.PostForm("username")
		pwd := c.PostForm("pwd")
		c.JSON(200, gin.H{
			"username": username,
			"pwd":      pwd,
		})
	})

	// ### 3.Query和POST混合参数
	r.POST("/any", func(c *gin.Context) {
		id := c.Query("id")
		username := c.PostForm("username")
		pwd := c.DefaultPostForm("pwd", "") // 默认空

		c.JSON(200, gin.H{
			"id":       id,
			"username": username,
			"password": pwd,
		})
	})

	// 路由重定向
	r.GET("/goto", func(c *gin.Context) {
		c.Redirect(301, "/ping") // 301 永久重定向
	})

	// 获取路由内容
	r.GET("/index", func(c *gin.Context) {
		c.Request.URL.Path = "/ping"
		r.HandleContext(c)
	})

	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}

// 函数形式
func anything(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
