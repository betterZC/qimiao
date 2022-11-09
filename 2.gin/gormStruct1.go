package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)
import _ "github.com/go-sql-driver/mysql"

type Class struct {
	gorm.Model
	ClassName string
	Students  []Student
}

type Student struct {
	gorm.Model
	StudentName string
	ClassID     uint
	IDCard      IDCard
	//多对多
	Teacher   []Teacher `gorm:"many2many:Student_Teacher;"` //关系表
	TeacherID uint
}

type IDCard struct {
	gorm.Model
	StudentID uint
	Num       int
}

type Teacher struct {
	gorm.Model
	StudentID   uint
	TeacherName string
	Students    []Student `gorm:"many2many:Student_Teacher;"`
}

func main() {
	db, err := gorm.Open("mysql", "root:1234@/gormclass?charset=utf8mb4&parseTime=True&loc=Local")
	db.AutoMigrate(&Teacher{}, &Class{}, &Student{}, &IDCard{})
	if err != nil {
		panic(err)
	}
	//i:=IDCard{
	//	Num: 123456,
	//}
	//s:=Student{
	//	StudentName: "zccc",
	//	IDCard: i,
	//}
	//
	//t:= Teacher{
	//	TeacherName: "qmm",
	//	Students: []Student{s},
	//}
	//c:=Class{
	//	ClassName: "classsssssssss",
	//	Students: []Student{s},
	//}
	//_ = db.Create(&c).Error
	//_ = db.Create(&t).Error

	defer db.Close()

	r := gin.Default()
	r.POST("/student", func(c *gin.Context) {
		var student Student
		_ = c.BindJSON(&student)
		db.Create(&student)
	})
	r.GET("/student/:ID", func(c *gin.Context) {
		id := c.Param("ID")
		var student Student
		_ = c.BindJSON(&student)
		db.Preload("Teachers").Preload("IDCard").First(&student, "id=?", id) //查不出来？预加载
		//db.Create(&student)
		c.JSON(200, gin.H{
			"s": student,
		})
	})
	r.GET("/class/:ID", func(c *gin.Context) { //多层关联 嵌套预加载
		id := c.Param("ID")
		var class Class
		db.Preload("Student").Preload("Students.Teachers").Preload("Students.IDCard").First(&class, "id=?", id)
		c.JSON(200, gin.H{
			"s": class,
		})
	})

	r.Run(":8888")
}
