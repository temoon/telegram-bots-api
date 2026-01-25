package requests

import (
	"context"
	"io"

	"github.com/temoon/telegram-bots-api"
)

type DeleteStickerSet struct {
	Name string
}

func (r *DeleteStickerSet) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "deleteStickerSet", r, response)
	return
}

func (r *DeleteStickerSet) GetValues() (values map[string]string, err error) {
	values = make(map[string]string)

	values["name"] = r.Name

	return
}

func (r *DeleteStickerSet) GetFiles() (files map[string]io.Reader) {
	return
}
