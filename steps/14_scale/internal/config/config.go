package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type Config struct {
	WebAppAddr string
	BotToken   string

	AppURL string

	LobbyMaxPlayer     int
	LobbyQuestionCount int
}

var Default Config

func init() {
	_ = godotenv.Load()
	Default = Config{
		WebAppAddr: os.Getenv("WEBAPP_URL"),
		BotToken:   os.Getenv("BOT_TOKEN"),

		AppURL: os.Getenv("APP_URL"),

		LobbyMaxPlayer:     getInt("LOBBY_MAX_PLAYER"),
		LobbyQuestionCount: getInt("LOBBY_QUESTION_COUNT"),
	}
}

func getInt(key string)int{
	num,err := strconv.Atoi(os.Getenv(key))
	if err!= nil{
		logrus.WithError(err).WithField("key", key).Fatal("couldn't get env value")
	}
	return num
}
