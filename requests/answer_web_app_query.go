package requests

import (
	"context"
	"encoding/json"
	"github.com/temoon/telegram-bots-api"
	"io"
)

type AnswerWebAppQuery struct {
	Result        interface{}
	WebAppQueryId string
}

func (r *AnswerWebAppQuery) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.SentWebAppMessage)
	err = b.CallMethod(ctx, "answerWebAppQuery", r, response)
	return
}

func (r *AnswerWebAppQuery) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	var dataResult []byte
	if dataResult, err = json.Marshal(r.Result); err != nil {
		return
	}

	values["result"] = string(dataResult)

	values["web_app_query_id"] = r.WebAppQueryId

	return
}

func (r *AnswerWebAppQuery) GetFiles() (files map[string]io.Reader) {
	return
}
