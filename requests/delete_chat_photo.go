package requests

import (
	"context"
	"github.com/temoon/telegram-bots-api"
	"io"
)

type DeleteChatPhoto struct {
	ChatId telegram.ChatId
}

func (r *DeleteChatPhoto) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "deleteChatPhoto", r, response)
	return
}

func (r *DeleteChatPhoto) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["chat_id"] = r.ChatId.String()

	return
}

func (r *DeleteChatPhoto) GetFiles() (files map[string]io.Reader) {
	return
}
