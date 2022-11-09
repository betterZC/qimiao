package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
)
import _ "github.com/go-sql-driver/mysql"

type HelloWorld struct {
	gorm.Model //自动创建点基础的东西
	Name       string
	Sex        bool
	Age        int
}

func main() {
	db, err := gorm.Open("mysql", "root:1234@/gormclass?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&HelloWorld{})

	//db.Create(&HelloWorld{
	//	Name: "zcccc",
	//	Sex: true,
	//	Age: 19,
	//})
	var hello HelloWorld
	var hello1 []HelloWorld

	//db.First(&hello, "name=?","zcccc")  //Find 接受切片地址  var hello []HelloWorld
	db.Where("name=?", "zcccc").Find(&hello1)
	db.Where("id= ?", 1).First(&HelloWorld{}).Update("name", "qimiap") //Update可放结构体或map
	//db.Where("id in (?)",[]int{1,2}).Find(&[]HelloWorld{}).Update("name", "qimiap") //批量
	fmt.Println(hello)

	db.Delete(&HelloWorld{}, "id = ?", 1) //软删除 硬：.Unscoped

	defer db.Close()
}
