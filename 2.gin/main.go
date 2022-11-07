package main

import "github.com/gin-gonic/gin"

func main() { 
	r := gin.Default()  //携带基础中间件启动 logger
	r.GET("/ping/:id", func(c *gin.Context) {
		id := c.Param("id ")
		user := c.DefaultQuery("user", "zcccc")
		pwd := c.Query("pwd")
		c.JSON(200, gin.H{
			"message": "pong",
			"id": id,
			"user": user,
			"pwd":pwd,
		})
	})
	r.POST("/ping", func(c *gin.Context) {
		user := c.DefaultPostForm("user","zcccccccccccc")
		pwd := c.PostForm("pwd")
		c.JSON(200, gin.H{
			"message": "pong",
			"user": user,
			"pwd":pwd,
		})
	})
	r.DELETE("/ping/:id", func(c *gin.Context) {
		id := c.Param("id ")
		c.JSON(200, gin.H{
			"message": "pong",
			"id": id,
		})
	})
	r.PUT("/ping", func(c *gin.Context) {
		user := c.DefaultPostForm("user","zcccccccccccc")
		pwd := c.PostForm("pwd")
		c.JSON(200, gin.H{
			"message": "pong",
			"user": user,
			"pwd":pwd,
		})
	})
	r.Run(":1010") // listen and serve on 0.0.0.0:8080
}