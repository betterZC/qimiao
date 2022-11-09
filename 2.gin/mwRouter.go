package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func middel() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("我在方法前1111")
		c.Next()
		fmt.Println("我在方法后1111")
	}
}

func middel2() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("我在方法前2222")
		c.Next()
		fmt.Println("我在方法后2222")
	}
}

func main() {
	r := gin.Default()
	v1 := r.Group("v1").Use(middel(), middel2()) //创建分组 .Use 中间件
	v1.GET("test", func(c *gin.Context) {
		fmt.Println("分组方法内部")
		c.JSON(200, gin.H{
			"success": true,
		})
	})
	r.Run(":8080")
}
