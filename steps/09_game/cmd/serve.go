package cmd

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/joho/godotenv"
	"github.com/rezaAmiri123/kingscomp/steps/09_game/internal/config"
	"github.com/rezaAmiri123/kingscomp/steps/09_game/internal/gameserver"
	"github.com/rezaAmiri123/kingscomp/steps/09_game/internal/matchmaking"
	"github.com/rezaAmiri123/kingscomp/steps/09_game/internal/repository"
	"github.com/rezaAmiri123/kingscomp/steps/09_game/internal/repository/redis"
	"github.com/rezaAmiri123/kingscomp/steps/09_game/internal/service"
	"github.com/rezaAmiri123/kingscomp/steps/09_game/internal/telegram"
	"github.com/rezaAmiri123/kingscomp/steps/09_game/internal/webapp"
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

	mm := matchmaking.NewRedisMatchmaking(redisClient, lobbyRepository, questionRepository)
	gs := gameserver.NewGameServer(app)

	tg, err := telegram.NewTelegram(app, mm, os.Getenv("BOT_API"))
	if err != nil {
		logrus.WithError(err).Fatalln("couldn't connect to the telegram server")
	}
	go tg.Start()

	wa := webapp.NewWebApp(app, gs, ":8080")

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	// use ngrok if its local
	if os.Getenv("ENV") == "local" {
		listener, err := ngrok.Listen(ctx,
			ngrokconfig.HTTPEndpoint(ngrokconfig.WithDomain(os.Getenv("NGROK_DOMAIN"))),
			ngrok.WithAuthtokenFromEnv(),
		)
		if err != nil {
			logrus.WithError(err).Fatalln("couldn't set up ngrok")
		}
		defer listener.Close()
		config.Default.WebAppAddr = "https://" + listener.Addr().String()
		logrus.WithField("ngrok_addr", config.Default.WebAppAddr).Info("local server is now online")
		logrus.WithError(wa.StartDev(listener)).Errorln("http server error")
	} else {
		wa.Start()
	}

	defer wa.Shutdown(context.Background())
	defer tg.Shutdown()

	<-ctx.Done()
	logrus.Info("shutting down the server ... please wait ...")
	<-time.After(time.Second)
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
