package driver

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"os"
)

var DB *gorm.DB = nil

func Init() {
	if DB != nil {
		return
	}

	formmat := "%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local"
	dataSource := fmt.Sprintf(formmat,
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_DATABASE"),
	)
	var err error
	DB, err = gorm.Open("mysql", dataSource)
	if err != nil {
		panic("failed to connect database")
	}
}
