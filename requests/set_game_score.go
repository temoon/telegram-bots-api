package requests

import (
	"context"
	"github.com/temoon/telegram-bots-api"
	"strconv"
)

type SetGameScore struct {
	ChatId             interface{}
	DisableEditMessage *bool
	Force              *bool
	InlineMessageId    *string
	MessageId          *int32
	Score              int32
	UserId             int64
}

func (r *SetGameScore) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(interface{})
	err = b.CallMethod(ctx, "setGameScore", r, response)
	return
}

func (r *SetGameScore) IsMultipart() (multipart bool) {
	return false
}

func (r *SetGameScore) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	switch value := r.ChatId.(type) {
	case int64:
		values["chat_id"] = strconv.FormatInt(value, 10)
	case string:
		values["chat_id"] = value
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
		values["message_id"] = strconv.FormatInt(int64(*r.MessageId), 10)
	}

	values["score"] = strconv.FormatInt(int64(r.Score), 10)

	values["user_id"] = strconv.FormatInt(r.UserId, 10)

	return
}
