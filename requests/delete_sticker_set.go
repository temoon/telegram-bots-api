package requests

import (
	"context"
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

func (r *DeleteStickerSet) IsMultipart() (multipart bool) {
	return false
}

func (r *DeleteStickerSet) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["name"] = r.Name

	return
}
