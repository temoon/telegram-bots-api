package requests

import (
"context"
	"github.com/temoon/telegram-bots-api"
)

type GetMyName struct {
LanguageCode *string
}

func (r *GetMyName) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.BotName)
	err = b.CallMethod(ctx, "getMyName", r, response)
	return
}



func (r *GetMyName) IsMultipart() (multipart bool) {
	return false
	}

func (r *GetMyName) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	
			if r.LanguageCode != nil {
			values["language_code"] = *r.LanguageCode
			}
			

	return
}
