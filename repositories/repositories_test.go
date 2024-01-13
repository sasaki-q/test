package repositories

import (
	"fmt"
	"testing"

	"github.com/joho/godotenv"
	"github.com/sasaki-q/test/config"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	err := godotenv.Load("../envs/.env.test")
	if err != nil {
		panic("sad .env.test file found")
	}
}

func TestMain(m *testing.M) {
	fmt.Print("Start: test\n")
	DB = config.New()
	config.Migrate(DB)
	m.Run()
	fmt.Print("End: test\n")
}
