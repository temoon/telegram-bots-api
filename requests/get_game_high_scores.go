package requests

import (
	"context"
	"github.com/temoon/telegram-bots-api"
	"io"
	"strconv"
)

type GetGameHighScores struct {
	ChatId          *int64
	MessageId       *int64
	InlineMessageId *string
	UserId          int64
}

func (r *GetGameHighScores) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new([]telegram.GameHighScore)
	err = b.CallMethod(ctx, "getGameHighScores", r, response)
	return
}

func (r *GetGameHighScores) GetValues() (values map[string]interface{}, err error) {
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

	return
}

func (r *GetGameHighScores) GetFiles() (files map[string]io.Reader) {
	return
}
