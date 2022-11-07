package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type PostParams struct {
	//Name string `json:"name" uri:"name" form:"name"` //用uri绑定加uri，用query绑定加 form
	//Age  int `json:"age" uri:"age" form:"age"`
	//Sex bool `json:"sex" uri:"sex" form:"sex"`
	Name string `json:"name"` //用uri绑定加uri，用query绑定加 form
	Age  int `json:"age" binding:"required,mustBig"`
	Sex bool `json:"sex"`
}

func mustBig(fl validator.FieldLevel) bool {
	if(fl.Field().Interface().(int)<=18) {
		return false
	}//  .(int)类型断言
	return true
}


func main() {
	r:=gin.Default()
	//registervalidation验证规则略去;;;;不应略，加上
	//https://www.bilibili.com/video/BV1gt4y1173C/?p=24&spm_id_from=pageDriver
	//15min开始，思维过程 
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("mustBig",mustBig)
	}
	
	r.POST("/testBind", func(c *gin.Context) {
	//r.POST("/testBind/:name/:age/:sex", func(c *gin.Context) {
	//r.POST("/testBind", func(c *gin.Context) {
		var p PostParams
		err:=c.ShouldBindJSON(&p)
		//err:=c.ShouldBindUri(&p)
		//err:=c.ShouldBindQuery(&p)
		if err!= nil {
			c.JSON(400, gin.H{
				"msg":"errorrrrrr",
				"data":gin.H{},
			})
		}else {
			c.JSON(200,gin.H{
				"msg":"okk",
				"data":p,
			})
		}
		
	})
	r.Run(":8080")
}
