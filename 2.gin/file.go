package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func main() {
	r := gin.Default()
	r.POST("/testUpload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		//name:=c.PostForm("name")

		//c.SaveUploadedFile(file,"./"+file.Filename) //保存前端传过来的文件
		in, _ := file.Open()
		defer in.Close()
		out, _ := os.Create("./" + file.Filename)
		defer out.Close()
		io.Copy(out, in)

		//c.JSON(200,gin.H{ //返回json
		//	"msg":file,
		//	"name":name,
		//})
		c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment;filename=%s", file.Filename)) //给前端反东西
		c.File("./" + file.Filename)
	})
	r.Run(":080")
}
