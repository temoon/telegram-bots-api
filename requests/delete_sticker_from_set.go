package requests

import (
	"context"
	"io"

	"github.com/temoon/telegram-bots-api"
)

type DeleteStickerFromSet struct {
	Sticker string
}

func (r *DeleteStickerFromSet) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "deleteStickerFromSet", r, response)
	return
}

func (r *DeleteStickerFromSet) GetValues() (values map[string]string, err error) {
	values = make(map[string]string)

	values["sticker"] = r.Sticker

	return
}

func (r *DeleteStickerFromSet) GetFiles() (files map[string]io.Reader) {
	return
}
