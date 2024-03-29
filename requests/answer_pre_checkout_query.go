package requests

import (
	"context"
	"github.com/temoon/telegram-bots-api"
	"io"
)

type AnswerPreCheckoutQuery struct {
	Ok                 bool
	PreCheckoutQueryId string
	ErrorMessage       *string
}

func (r *AnswerPreCheckoutQuery) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "answerPreCheckoutQuery", r, response)
	return
}

func (r *AnswerPreCheckoutQuery) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	if r.Ok {
		values["ok"] = "1"
	} else {
		values["ok"] = "0"
	}

	values["pre_checkout_query_id"] = r.PreCheckoutQueryId

	if r.ErrorMessage != nil {
		values["error_message"] = *r.ErrorMessage
	}

	return
}

func (r *AnswerPreCheckoutQuery) GetFiles() (files map[string]io.Reader) {
	return
}
