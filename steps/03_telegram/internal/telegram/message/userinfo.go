package message

import "github.com/rezaAmiri123/kingscomp/steps/03_telegram/internal/entity"

func MainMenuText(account entity.Account) string {
	return `🏯 خوش آمدید %s

چه کاری میتونم براتون انجام بدم؟`
}
