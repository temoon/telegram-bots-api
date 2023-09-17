package requests

import (
	"context"
	"github.com/temoon/telegram-bots-api"
)

type AnswerPreCheckoutQuery struct {
	ErrorMessage       *string
	Ok                 bool
	PreCheckoutQueryId string
}

func (r *AnswerPreCheckoutQuery) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "answerPreCheckoutQuery", r, response)
	return
}

func (r *AnswerPreCheckoutQuery) IsMultipart() (multipart bool) {
	return false
}

func (r *AnswerPreCheckoutQuery) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	if r.ErrorMessage != nil {
		values["error_message"] = *r.ErrorMessage
	}

	if r.Ok {
		values["ok"] = "1"
	} else {
		values["ok"] = "0"
	}

	values["pre_checkout_query_id"] = r.PreCheckoutQueryId

	return
}
