package cmd

import (
	"context"
	"os"

	"github.com/joho/godotenv"
	"github.com/rezaAmiri123/kingscomp/steps/06_web/internal/config"
	"github.com/rezaAmiri123/kingscomp/steps/06_web/internal/matchmaking"
	"github.com/rezaAmiri123/kingscomp/steps/06_web/internal/repository"
	"github.com/rezaAmiri123/kingscomp/steps/06_web/internal/repository/redis"
	"github.com/rezaAmiri123/kingscomp/steps/06_web/internal/service"
	"github.com/rezaAmiri123/kingscomp/steps/06_web/internal/telegram"
	"github.com/rezaAmiri123/kingscomp/steps/06_web/internal/webapp"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"golang.ngrok.com/ngrok"
	ngrokconfig "golang.ngrok.com/ngrok/config"
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
	if err != nil {
		logrus.WithError(err).Fatalln("couldn't connect to te redis server")
	}
	accountRepository := repository.NewAccountRedisRepository(redisClient)
	lobbyRepository := repository.NewLobbyRedisRepository(redisClient)
	// set up app
	app := service.NewApp(
		service.NewAccountService(accountRepository),
		service.NewLobbyService(lobbyRepository),
	)

	mm := matchmaking.NewRedisMatchmaking(redisClient, lobbyRepository)
	tg, err := telegram.NewTelegram(app, mm, os.Getenv("BOT_API"))
	if err != nil {
		logrus.WithError(err).Fatalln("couldn't connect to the telegram server")
	}
	go tg.Start()

	wa := webapp.NewWebApp(app, ":8080")

	// use ngrok if its local
	if os.Getenv("ENV") == "local" {
		listener, err := ngrok.Listen(context.Background(),
			ngrokconfig.HTTPEndpoint(),
			ngrok.WithAuthtokenFromEnv(),
		)
		if err != nil {
			logrus.WithError(err).Fatalln("couldn't set up ngrok")
		}
		config.Default.WebAppAddr = "https://" + listener.Addr().String()
		logrus.WithField("ngrok_addr", config.Default.WebAppAddr).Info("local server is now online")
		logrus.WithError(wa.StartDev(listener)).Errorln("http server error")
		return
	}

	wa.Start()
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
