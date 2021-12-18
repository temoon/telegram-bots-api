package requests

import (
	"context"
	"strconv"

	"github.com/temoon/go-telegram-bots-api"
)

type SetGameScore struct {
	ChatId             int64
	DisableEditMessage bool
	Force              bool
	InlineMessageId    string
	MessageId          int32
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

	if r.ChatId != 0 {
		values["chat_id"] = strconv.FormatInt(r.ChatId, 10)
	}

	if r.DisableEditMessage {
		values["disable_edit_message"] = "1"
	}

	if r.Force {
		values["force"] = "1"
	}

	if r.InlineMessageId != "" {
		values["inline_message_id"] = r.InlineMessageId
	}

	if r.MessageId != 0 {
		values["message_id"] = strconv.FormatInt(int64(r.MessageId), 10)
	}

	values["score"] = strconv.FormatInt(int64(r.Score), 10)

	values["user_id"] = strconv.FormatInt(r.UserId, 10)

	return
}
