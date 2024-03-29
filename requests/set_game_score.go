package requests

import (
	"context"
	"github.com/temoon/telegram-bots-api"
	"io"
	"strconv"
)

type SetGameScore struct {
	Score              int64
	UserId             int64
	ChatId             *int64
	DisableEditMessage *bool
	Force              *bool
	InlineMessageId    *string
	MessageId          *int64
}

func (r *SetGameScore) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.Message)
	err = b.CallMethod(ctx, "setGameScore", r, response)
	return
}

func (r *SetGameScore) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["score"] = strconv.FormatInt(r.Score, 10)

	values["user_id"] = strconv.FormatInt(r.UserId, 10)

	if r.ChatId != nil {
		values["chat_id"] = strconv.FormatInt(*r.ChatId, 10)
	}

	if r.DisableEditMessage != nil {
		if *r.DisableEditMessage {
			values["disable_edit_message"] = "1"
		} else {
			values["disable_edit_message"] = "0"
		}
	}

	if r.Force != nil {
		if *r.Force {
			values["force"] = "1"
		} else {
			values["force"] = "0"
		}
	}

	if r.InlineMessageId != nil {
		values["inline_message_id"] = *r.InlineMessageId
	}

	if r.MessageId != nil {
		values["message_id"] = strconv.FormatInt(*r.MessageId, 10)
	}

	return
}

func (r *SetGameScore) GetFiles() (files map[string]io.Reader) {
	return
}
