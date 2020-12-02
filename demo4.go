package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

// 自定义的中间件 返回值类型为：gin.HandlerFunc
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// 设置 example 变量
		c.Set("example", "111")

		// 请求之前...

		c.Next()

		// 请求之后...

		latency := time.Since(t)
		// 打印请求处理时间
		log.Print(latency)

		// 访问即将发送的响应状态码
		status := c.Writer.Status()
		log.Println(status)
	}
}

func main() {
	// gin.Default() 默认启动方式，包含 Logger、Recovery 中间件
	r := gin.New() // 不使用中间件
	// 使用自定义的 Logger 中间件
	r.Use(Logger())

	// 定义路由
	r.GET("/test", func(c *gin.Context) {
		example := c.MustGet("example").(string)
		// 打印 example 值 111
		log.Println(example)
	})

	// 监听 0.0.0.0:8080
	r.Run(":8080")
}
