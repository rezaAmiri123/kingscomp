package telegram

import (
	"github.com/rezaAmiri123/kingscomp/steps/15_loadbalancer/internal/config"
	"github.com/rezaAmiri123/kingscomp/steps/15_loadbalancer/internal/gameserver"
	"github.com/rezaAmiri123/kingscomp/steps/15_loadbalancer/internal/matchmaking"
	"github.com/rezaAmiri123/kingscomp/steps/15_loadbalancer/internal/service"
	"github.com/rezaAmiri123/kingscomp/steps/15_loadbalancer/internal/telegram/teleprompt"
	"github.com/sirupsen/logrus"
	"gopkg.in/telebot.v3"
)

type Telegram struct {
	App *service.App
	Bot *telebot.Bot

	TelePrompt *teleprompt.TelePrompt
	mm         matchmaking.Matchmaking
	gs         *gameserver.GameServer
}

func NewTelegram(app *service.App,
	mm matchmaking.Matchmaking,
	gs *gameserver.GameServer,
	tp *teleprompt.TelePrompt,
	apiKey string,
) (*Telegram, error) {
	t := &Telegram{
		App:        app,
		TelePrompt: tp,
		mm:         mm,
		gs:         gs,
	}

	pref := telebot.Settings{
		Token: apiKey,
		Poller: &telebot.Webhook{
			MaxConnections: 100,
			Endpoint: &telebot.WebhookEndpoint{
				PublicURL: config.Default.AppURL + "/webhook/" + config.Default.BotToken,
			},
		},
		OnError: t.onError,
	}

	bot, err := telebot.NewBot(pref)
	if err != nil {
		logrus.WithError(err).Error("couldn't connect to telegram servers")
		return nil, err
	}

	t.Bot = bot

	t.setupHandlers()

	return t, nil
}

func (t *Telegram) Start() {
	t.Bot.Start()
}

func (t *Telegram) Shutdown() {
	t.Bot.Stop()
}
