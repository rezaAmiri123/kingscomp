package telegram

import (
	"time"

	"github.com/rezaAmiri123/kingscomp/steps/15_loadbalancer/internal/config"
	"github.com/rezaAmiri123/kingscomp/steps/15_loadbalancer/internal/entity"
	"gopkg.in/telebot.v3"
)

var (
	DefaultMatchmakingTimeout         = time.Second * 120
	DefaultMatchmakingLoadingInterval = DefaultMatchmakingTimeout / 13 //todo: increase in the production
	DefaultInputTimeout               = time.Minute * 5
	DefaultTimeoutText                = `🕗 منتظر پیامت بودیم چیزی ارسال نکردی. لطفا هر وقت برگشتی دوباره پیام بده.`

	TxtConfirm = `✅ بله`
	TxtDecline = `✖ خیر`
)

func GetAccount(c telebot.Context)entity.Account{
	return c.Get("account").(entity.Account)
}

var (
	selector           = &telebot.ReplyMarkup{}
	btnEditDisplayName = selector.Data("📝 ویرایش نام‌نمایشی", "btnEditDisplayName")
	btnJoinMatchmaking = selector.Data("🎮 شروع بازی جدید", "btnJoinMatchmaking")
	btnCurrentMatch    = selector.Data("🎲 بازی در حال اجرای من", "btnCurrentMatch")
	btnResignLobby     = selector.Data("🏳 تسلیم شدن", "btnResignLobby")
	btnStartGameWebApp = selector.Data("🎮 باز کردن بازی", "btnStartGameWebApp")
)

func NewStartWebAppGame(lobbyId string)telebot.Btn{
	return selector.WebApp("🎮 باز کردن بازی",&telebot.WebApp{
		URL: config.Default.AppURL + "/lobby/" + lobbyId,
	})
}
