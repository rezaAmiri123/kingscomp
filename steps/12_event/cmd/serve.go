package cmd

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/joho/godotenv"
	"github.com/rezaAmiri123/kingscomp/steps/12_event/internal/gameserver"
	"github.com/rezaAmiri123/kingscomp/steps/12_event/internal/matchmaking"
	"github.com/rezaAmiri123/kingscomp/steps/12_event/internal/repository"
	"github.com/rezaAmiri123/kingscomp/steps/12_event/internal/repository/redis"
	"github.com/rezaAmiri123/kingscomp/steps/12_event/internal/service"
	"github.com/rezaAmiri123/kingscomp/steps/12_event/internal/telegram"
	"github.com/rezaAmiri123/kingscomp/steps/12_event/internal/webapp"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	_ "go.uber.org/automaxprocs"
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
	if os.Getenv("ENV") != "local" {
		logrus.SetLevel(logrus.ErrorLevel)
	}
	// set up repositories
	redisClient, err := redis.NewRedisClient(os.Getenv("REDIS_URL"))
	if err != nil {
		logrus.WithError(err).Fatalln("couldn't connect to the redis server")
	}
	accountRepository := repository.NewAccountRedisRepository(redisClient)
	lobbyRepository := repository.NewLobbyRedisRepository(redisClient)
	questionRepository := repository.NewQuestionRedisRepository(redisClient)

	// set up app
	app := service.NewApp(
		service.NewAccountService(accountRepository),
		service.NewLobbyService(lobbyRepository),
	)

	mm := matchmaking.NewRedisMatchmaking(redisClient, lobbyRepository, questionRepository, accountRepository)
	gs := gameserver.NewGameServer(app, gameserver.DefaultGameServerConfig())

	tg, err := telegram.NewTelegram(app, mm, gs, os.Getenv("BOT_API"))
	if err != nil {
		logrus.WithError(err).Fatalln("couldn't connect to the telegram server")
	}
	go tg.Start()

	wa := webapp.NewWebApp(app, gs, ":8080")

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	if os.Getenv("ENV") == "local" {
		go func() {
			logrus.WithError(wa.StartDev()).Errorln("http server error")
		}()
	} else {
		go func() {
			logrus.WithError(wa.Start()).Errorln("http server error")
		}()
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	defer wa.Shutdown(shutdownCtx)
	defer tg.Shutdown()

	logrus.Info("server is up and running")
	<-ctx.Done()
	logrus.Info("shutting down the server ... please wait ...")
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
