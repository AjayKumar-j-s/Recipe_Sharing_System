package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
  dsn := "root:abc123@tcp(127.0.0.1:3306)/recipe?charset=utf8mb4&parseTime=True&loc=Local"
  var err error
  DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
  if(err!=nil){
	panic("cannot establish connection with db")
  }

}


