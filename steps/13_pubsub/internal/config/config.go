package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct{
	WebAppAddr string
}

var Default Config

func init(){
	_ = godotenv.Load()
	Default = Config{
		WebAppAddr: os.Getenv("WEBAPP_URL"),
	}
}
