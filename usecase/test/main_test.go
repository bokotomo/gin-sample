package usecase

import (
	"os"
	"testing"
	"github.com/joho/godotenv"
	. "gin-sample/driver"
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
