package requests

import (
	"context"
	"github.com/temoon/telegram-bots-api"
	"strconv"
)

type GetGameHighScores struct {
	ChatId          interface{}
	InlineMessageId *string
	MessageId       *int32
	UserId          int64
}

func (r *GetGameHighScores) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new([]telegram.GameHighScore)
	err = b.CallMethod(ctx, "getGameHighScores", r, response)
	return
}

func (r *GetGameHighScores) IsMultipart() (multipart bool) {
	return false
}

func (r *GetGameHighScores) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	switch value := r.ChatId.(type) {
	case int64:
		values["chat_id"] = strconv.FormatInt(value, 10)
	case string:
		values["chat_id"] = value
	}

	if r.InlineMessageId != nil {
		values["inline_message_id"] = *r.InlineMessageId
	}

	if r.MessageId != nil {
		values["message_id"] = strconv.FormatInt(int64(*r.MessageId), 10)
	}

	values["user_id"] = strconv.FormatInt(r.UserId, 10)

	return
}
