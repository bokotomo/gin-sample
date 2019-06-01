package driver

import (
	"os"
  "github.com/jinzhu/gorm"
  _ "github.com/go-sql-driver/mysql"
)

var DB *gorm.DB = nil

func Init() {
	if DB != nil {
		return
	}
  host := os.Getenv("DB_HOST")
  database := os.Getenv("DB_DATABASE")
  user := os.Getenv("DB_USER")
  password := os.Getenv("DB_PASSWORD")
  var err error
  DB, err = gorm.Open(
    "mysql",
    user + ":" + password + "@tcp(" + host + ")/" + database + "?charset=utf8&parseTime=True&loc=Local",
  )
  if err != nil {
    panic("failed to connect database")
  }
}
