package requests

import (
	"context"
	"github.com/temoon/telegram-bots-api"
	"io"
)

type SetStickerSetTitle struct {
	Name  string
	Title string
}

func (r *SetStickerSetTitle) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "setStickerSetTitle", r, response)
	return
}

func (r *SetStickerSetTitle) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["name"] = r.Name

	values["title"] = r.Title

	return
}

func (r *SetStickerSetTitle) GetFiles() (files map[string]io.Reader) {
	return
}
