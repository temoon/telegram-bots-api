package requests

import (
	"context"
	"io"
	"strconv"

	"github.com/temoon/telegram-bots-api"
)

type DeclineSuggestedPost struct {
	ChatId    int64
	MessageId int64
	Comment   *string
}

func (r *DeclineSuggestedPost) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "declineSuggestedPost", r, response)
	return
}

func (r *DeclineSuggestedPost) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["chat_id"] = strconv.FormatInt(r.ChatId, 10)

	values["message_id"] = strconv.FormatInt(r.MessageId, 10)

	if r.Comment != nil {
		values["comment"] = *r.Comment
	}

	return
}

func (r *DeclineSuggestedPost) GetFiles() (files map[string]io.Reader) {
	return
}
