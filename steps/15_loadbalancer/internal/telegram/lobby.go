package telegram

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/rezaAmiri123/kingscomp/steps/15_loadbalancer/internal/entity"
	"github.com/rezaAmiri123/kingscomp/steps/15_loadbalancer/internal/events"
	"github.com/rezaAmiri123/kingscomp/steps/15_loadbalancer/internal/matchmaking"
	"github.com/rezaAmiri123/kingscomp/steps/15_loadbalancer/internal/repository"
	"github.com/samber/lo"
	"gopkg.in/telebot.v3"
)

func (t *Telegram) joinMatchMaking(c telebot.Context) error {
	c.Respond()
	myAccount := GetAccount(c)

	if myAccount.CurrentLobby != "" { //todo: show the current game's status
		return c.Reply("در حال حاضر در حال انجام یک بازی هستید")
	}

	msg, err := t.Input(c, InputConfig{
		Prompt:         "⏰ هر بازی بین 2 تا 4 دقیقه طول میکشد و در صورت ورود باید اینترنت پایداری داشته باشید.\n\nجستجوی بازی جدید رو شروع کنیم؟",
		PromptKeyboard: [][]string{{TxtDecline, TxtConfirm}},
		Validator:      choiceValidator(TxtDecline, TxtConfirm),
	})
	if err != nil {
		return err
	}

	if msg.Text == TxtDecline {
		return t.myInfo(c)
	}

	ch := make(chan struct{}, 1)
	var lobby entity.Lobby
	var isHost bool
	go func() {
		lobby, isHost, err = t.mm.Join(context.Background(), c.Sender().ID, DefaultMatchmakingTimeout)
		ch <- struct{}{}
	}()

	ticker := time.NewTicker(DefaultMatchmakingLoadingInterval)
	loadingMessage, err := c.Bot().Send(c.Sender(), `🎮 درحال پیدا کردن حریف ... منتظر بمانید`)
	if err != nil {
		return err
	}
	defer func() {
		c.Bot().Delete(loadingMessage)
	}()
	s := time.Now()
Loading:
	for {
		select {
		case <-ticker.C:
			took := int(time.Since(s).Seconds())
			c.Bot().Edit(loadingMessage, fmt.Sprintf(`🎮 درحال پیدا کردن حریف ... منتظر بمانید

🕕 %d ثانیه از %d`, took, int(DefaultMatchmakingTimeout.Seconds())))
			continue
		case <-ch:
			break Loading
		}
	}

	if err != nil {
		if errors.Is(err, matchmaking.ErrTimeout) {
			c.Send(`🕕 به مدت 2 دقیقه دنبال بازی جدیدی گشتیم اما متاسفانه پیدا نشد. میتونید چند دقیقه دیگه دوباره تلاش کنید`)
			return t.myInfo(c)
		}
		return err
	}

	// setup reminder with goroutines
	if isHost {

		game, err := t.gs.Register(lobby.ID)
		if err != nil {
			return err
		}

		game.Events.Register("lobby."+lobby.ID, events.EventJoinReminder, func(info events.EventInfo) {
			c.Bot().Send(&telebot.User{ID: info.AccountID},
				`⚠️ بازی جدید برای شما ساخته شده اما هنوز بازی را باز نکرده اید! تا چند ثانیه دیگر اگر بازی را باز نکنید تسلیم شده در نظر گرفته میشوید.`,
				NewLobbyInlineKeyboards(lobby.ID))
		})

		game.Events.Register("lobby."+lobby.ID, events.EventLateResign, func(info events.EventInfo) {
			c.Bot().Send(&telebot.User{ID: info.AccountID},
				`😔 متاسفانه چون وارد بازی جدید نشدید مجبور شدیم وضعیتتون رو به «تسلیم شده» تغییر بدیم.`)
		})

		game.Events.Register("lobby."+lobby.ID, events.EventGameClosed, func(info events.EventInfo) {
			c.Bot().Send(&telebot.User{ID: info.AccountID}, `🎲 بازی با موفقیت به اتمام رسید. خسته نباشید.

اگه میخواید ربات رو استارت کنید یا بازی جدیدی شروع کنید روی /home کلیک کنید.`)
		})
	}

	myAccount.CurrentLobby = lobby.ID
	c.Set("account", myAccount)

	return t.currentLobby(c)
}

func (t *Telegram) currentLobby(c telebot.Context) error {
	myAccount := GetAccount(c)

	lobby, accounts, err := t.App.LobbyParticipants(context.Background(), myAccount.CurrentLobby)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			c.Respond(&telebot.CallbackResponse{
				Text: `این بازی به اتمام رسیده است`,
			})
			c.Bot().Delete(c.Message())
			myAccount.CurrentLobby = ""
			t.App.Account.Save(context.Background(), myAccount)
			return t.myInfo(c)
		}
		return err
	}

	return c.Send(fmt.Sprintf(`🏁 بازی درحال اجرای شما

بازیکنان شما:
%s

شناسه بازی: %s
`,
		strings.Join(lo.Map(accounts, func(item entity.Account, _ int) string {
			isMeTxt := ""
			if item.ID == myAccount.ID {
				isMeTxt = "(شما)"
			}
			return fmt.Sprintf(`🎴 %s %s`, item.DisplayName, isMeTxt)
		}), "\n"),
		lobby.ID,
	), NewLobbyInlineKeyboards(lobby.ID))
}

func NewLobbyInlineKeyboards(lobbyId string) *telebot.ReplyMarkup {
	selector := &telebot.ReplyMarkup{}
	selector.Inline(selector.Row(btnResignLobby, NewStartWebAppGame(lobbyId)))
	return selector
}

func (t *Telegram) resignLobby(c telebot.Context) error {
	defer c.Bot().Delete(c.Message())
	myAccount := GetAccount(c)
	myLobby := myAccount.CurrentLobby
	if myLobby == "" {
		c.Respond(&telebot.CallbackResponse{
			Text: `شما قبلا از این بازی انصراف داده بودید`,
		})
		return t.myInfo(c)
	}
	c.Respond(&telebot.CallbackResponse{
		Text: `✅ با موفقیت از بازی فعلی انصراف دادید`,
	})
	myAccount.CurrentLobby = ""
	if err := t.App.Account.Save(context.Background(), myAccount); err != nil {
		return err
	}

	t.App.Lobby.UpdateUserState(context.Background(),
		myLobby, myAccount.ID, "isResigned", true)

	t.gs.PubSub.Dispatch(
		context.Background(),
		"lobby."+myLobby,
		events.EventUserResigned,
		events.EventInfo{
			AccountID: myAccount.ID,
		})

	c.Set("account", myAccount)
	return t.myInfo(c)
}
