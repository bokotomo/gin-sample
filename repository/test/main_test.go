package repository

import (
	. "gin-sample/driver"
	"github.com/joho/godotenv"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	if err := godotenv.Load("../../.env"); err != nil {
		panic(".env file cannot be load.")
	}
	Init()
	defer DB.Close()

	code := m.Run()

	os.Exit(code)
}
