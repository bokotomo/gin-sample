package repository

import (
	"gin-sample/driver"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestMain(m *testing.M) {
	if err := godotenv.Load("../../.env"); err != nil {
		panic(".env file cannot be load.")
	}
	driver.Init()
	defer driver.DB.Close()

	code := m.Run()

	os.Exit(code)
}
