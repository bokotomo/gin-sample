package repository

import (
	"os"
	"testing"

	"github.com/joho/godotenv"

	. "gin-sample/driver"
)

func TestMain(m *testing.M) {
	// Load ENV
	if err := godotenv.Load(); err != nil {
		panic(".env file cannot be load.")
	}
	Init()

	code := m.Run()

	os.Exit(code)
  defer DB.Close()
}
