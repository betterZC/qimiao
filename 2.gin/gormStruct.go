package main

import "github.com/jinzhu/gorm"
import _ "github.com/go-sql-driver/mysql"

type User struct {
	gorm.Model
	Name string `gorm:"primary_key;column:user_name;type:varchar(100);"`
}

func (u User) TableName() string { //自定义表名
	return "zc_users"
}

func main() {
	db, err := gorm.Open("mysql", "root:1234@/gormclass?charset=utf8mb4&parseTime=True&loc=Local")
	db.AutoMigrate(&User{})
	if err != nil {
		panic(err)
	}
	defer db.Close()
}
