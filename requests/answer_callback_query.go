package requests

import (
	"context"
	"strconv"

	"github.com/temoon/go-telegram-bots-api"
)

type AnswerCallbackQuery struct {
	CacheTime       int32
	CallbackQueryId string
	ShowAlert       bool
	Text            string
	Url             string
}

func (r *AnswerCallbackQuery) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "answerCallbackQuery", r, response)
	return
}

func (r *AnswerCallbackQuery) IsMultipart() (multipart bool) {
	return false
}

func (r *AnswerCallbackQuery) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	if r.CacheTime != 0 {
		values["cache_time"] = strconv.FormatInt(int64(r.CacheTime), 10)
	}

	values["callback_query_id"] = r.CallbackQueryId

	if r.ShowAlert {
		values["show_alert"] = "1"
	}

	if r.Text != "" {
		values["text"] = r.Text
	}

	if r.Url != "" {
		values["url"] = r.Url
	}

	return
}
