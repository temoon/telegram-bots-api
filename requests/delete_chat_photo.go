package requests

import (
	"context"
	"io"

	"github.com/temoon/telegram-bots-api"
)

type DeleteChatPhoto struct {
	ChatId telegram.ChatId
}

func (r *DeleteChatPhoto) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "deleteChatPhoto", r, response)
	return
}

func (r *DeleteChatPhoto) GetValues() (values map[string]string, err error) {
	values = make(map[string]string)

	values["chat_id"] = r.ChatId.String()

	return
}

func (r *DeleteChatPhoto) GetFiles() (files map[string]io.Reader) {
	return
}
