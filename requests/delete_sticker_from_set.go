package requests

import (
	"context"
	"github.com/temoon/telegram-bots-api"
	"io"
)

type DeleteStickerFromSet struct {
	Sticker string
}

func (r *DeleteStickerFromSet) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "deleteStickerFromSet", r, response)
	return
}

func (r *DeleteStickerFromSet) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["sticker"] = r.Sticker

	return
}

func (r *DeleteStickerFromSet) GetFiles() (files map[string]io.Reader) {
	return
}
