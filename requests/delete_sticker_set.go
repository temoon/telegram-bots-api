package requests

import (
	"context"
	"github.com/temoon/telegram-bots-api"
	"io"
)

type DeleteStickerSet struct {
	Name string
}

func (r *DeleteStickerSet) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "deleteStickerSet", r, response)
	return
}

func (r *DeleteStickerSet) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["name"] = r.Name

	return
}

func (r *DeleteStickerSet) GetFiles() (files map[string]io.Reader) {
	return
}
