package requests

import (
	"context"

	"github.com/temoon/go-telegram-bots-api"
)

type DeleteStickerFromSet struct {
	Sticker string
}

func (r *DeleteStickerFromSet) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "deleteStickerFromSet", r, response)
	return
}

func (r *DeleteStickerFromSet) IsMultipart() (multipart bool) {
	return false
}

func (r *DeleteStickerFromSet) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["sticker"] = r.Sticker

	return
}
