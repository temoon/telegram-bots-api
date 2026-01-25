package requests

import (
	"context"
	"io"

	"github.com/temoon/telegram-bots-api"
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

func (r *SetChatPhoto) GetValues() (values map[string]string, err error) {
	values = make(map[string]string)

	values["chat_id"] = r.ChatId.String()

	values["photo"] = r.Photo.String()

	return
}

func (r *SetChatPhoto) GetFiles() (files map[string]io.Reader) {
	files = make(map[string]io.Reader)

	if r.Photo.HasFile() {
		files[r.Photo.GetFormFieldName()] = r.Photo.GetFile()
	}

	return
}
