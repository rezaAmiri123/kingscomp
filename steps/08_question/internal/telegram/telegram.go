package telegram

import (
	"time"

	"github.com/rezaAmiri123/kingscomp/steps/08_question/internal/matchmaking"
	"github.com/rezaAmiri123/kingscomp/steps/08_question/internal/service"
	"github.com/rezaAmiri123/kingscomp/steps/08_question/internal/telegram/teleprompt"
	"github.com/sirupsen/logrus"
	"gopkg.in/telebot.v3"
)

type Telegram struct {
	App *service.App
	bot *telebot.Bot

	TelePrompt *teleprompt.TelePrompt
	mm         matchmaking.Matchmaking
}

func NewTelegram(app *service.App, mm matchmaking.Matchmaking, apiKey string) (*Telegram, error) {
	t := &Telegram{
		App:        app,
		TelePrompt: teleprompt.NewTelePrompt(),
		mm: mm,
	}

	pref := telebot.Settings{
		Token:   apiKey,
		Poller:  &telebot.LongPoller{Timeout: time.Minute},
		OnError: t.onError,
	}

	bot, err := telebot.NewBot(pref)
	if err != nil {
		logrus.WithError(err).Error("couldn't connect to telegram servers")
		return nil, err
	}

	t.bot = bot

	t.setupHandlers()

	return t, nil
}

func (t *Telegram) Start() {
	t.bot.Start()
}

func(t *Telegram)Shutdown(){
	t.bot.Stop()
}