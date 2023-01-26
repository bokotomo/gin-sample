package driver

// classにする
import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// DB is
var DB *gorm.DB = nil

// TestDB is
var TestDB *gorm.DB = nil
var formmat string = "%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local"

// Init is
func Init() {
	if DB != nil {
		return
	}

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

// TestInit is
func TestInit() {
	if TestDB != nil {
		return
	}

	dataSource := fmt.Sprintf(formmat,
		os.Getenv("TEST_DB_USER"),
		os.Getenv("TEST_DB_PASSWORD"),
		os.Getenv("TEST_DB_HOST"),
		os.Getenv("TEST_DB_DATABASE"),
	)
	var err error
	TestDB, err = gorm.Open("mysql", dataSource)
	if err != nil {
		panic("failed to connect database")
	}
}
