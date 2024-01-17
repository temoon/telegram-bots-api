package requests

import (
	"context"
	"github.com/temoon/telegram-bots-api"
	"io"
	"strconv"
)

type SetGameScore struct {
	ChatId             *int64
	MessageId          *int64
	InlineMessageId    *string
	UserId             int64
	Score              int64
	Force              *bool
	DisableEditMessage *bool
}

func (r *SetGameScore) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.Message)
	err = b.CallMethod(ctx, "setGameScore", r, response)
	return
}

func (r *SetGameScore) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	if r.ChatId != nil {
		values["chat_id"] = strconv.FormatInt(*r.ChatId, 10)
	}

	if r.MessageId != nil {
		values["message_id"] = strconv.FormatInt(*r.MessageId, 10)
	}

	if r.InlineMessageId != nil {
		values["inline_message_id"] = *r.InlineMessageId
	}

	values["user_id"] = strconv.FormatInt(r.UserId, 10)

	values["score"] = strconv.FormatInt(r.Score, 10)

	if r.Force != nil {
		if *r.Force {
			values["force"] = "1"
		} else {
			values["force"] = "0"
		}
	}

	if r.DisableEditMessage != nil {
		if *r.DisableEditMessage {
			values["disable_edit_message"] = "1"
		} else {
			values["disable_edit_message"] = "0"
		}
	}

	return
}

func (r *SetGameScore) GetFiles() (files map[string]io.Reader) {
	return
}
