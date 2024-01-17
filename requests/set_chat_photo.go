package requests

import (
	"context"
	"github.com/temoon/telegram-bots-api"
	"io"
)

type SetChatPhoto struct {
	ChatId telegram.ChatId
	Photo  telegram.InputFile
}

func (r *SetChatPhoto) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "setChatPhoto", r, response)
	return
}

func (r *SetChatPhoto) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["chat_id"] = r.ChatId.String()

	values["photo"] = r.Photo.GetValue()

	return
}

func (r *SetChatPhoto) GetFiles() (files map[string]io.Reader) {
	return
}
