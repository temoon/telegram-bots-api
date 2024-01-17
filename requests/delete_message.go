package requests

import (
	"context"
	"github.com/temoon/telegram-bots-api"
	"io"
	"strconv"
)

type DeleteMessage struct {
	ChatId    telegram.ChatId
	MessageId int64
}

func (r *DeleteMessage) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "deleteMessage", r, response)
	return
}

func (r *DeleteMessage) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["chat_id"] = r.ChatId.String()

	values["message_id"] = strconv.FormatInt(r.MessageId, 10)

	return
}

func (r *DeleteMessage) GetFiles() (files map[string]io.Reader) {
	return
}
