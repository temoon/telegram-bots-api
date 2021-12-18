package requests

import (
	"context"
	"strconv"

	"github.com/temoon/go-telegram-bots-api"
)

type GetGameHighScores struct {
	ChatId          int64
	InlineMessageId string
	MessageId       int32
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

	if r.ChatId != 0 {
		values["chat_id"] = strconv.FormatInt(r.ChatId, 10)
	}

	if r.InlineMessageId != "" {
		values["inline_message_id"] = r.InlineMessageId
	}

	if r.MessageId != 0 {
		values["message_id"] = strconv.FormatInt(int64(r.MessageId), 10)
	}

	values["user_id"] = strconv.FormatInt(r.UserId, 10)

	return
}
