package cmd

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/rezaAmiri123/kingscomp/steps/05_matchmaking/internal/repository"
	"github.com/rezaAmiri123/kingscomp/steps/05_matchmaking/internal/repository/redis"
	"github.com/rezaAmiri123/kingscomp/steps/05_matchmaking/internal/service"
	"github.com/rezaAmiri123/kingscomp/steps/05_matchmaking/internal/telegram"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "serve the appliation",
	Run: func(cmd *cobra.Command, args []string) {
		serve()
	},
}

func serve() {
	_ = godotenv.Load()
	// set up repositories
	redisClient, err := redis.NewRedisClient(os.Getenv("REDIS_URL"))
	if err!= nil{
		logrus.WithError(err).Fatalln("couldn't connect to te redis server")
	}
	accountRepository := repository.NewAccountRedisRepository(redisClient)

	// set up app
	app := service.NewApp(
		service.NewAccountService(accountRepository),
	)

	tg,err:= telegram.NewTelegram(app,os.Getenv("BOT_API"))
	if err!= nil{
		logrus.WithError(err).Fatalln("couldn't connect to the telegram server")
	}
	tg.Start()
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
