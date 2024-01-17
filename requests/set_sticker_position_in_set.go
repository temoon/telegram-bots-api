package requests

import (
	"context"
	"github.com/temoon/telegram-bots-api"
	"io"
	"strconv"
)

type SetStickerPositionInSet struct {
	Sticker  string
	Position int64
}

func (r *SetStickerPositionInSet) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "setStickerPositionInSet", r, response)
	return
}

func (r *SetStickerPositionInSet) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["sticker"] = r.Sticker

	values["position"] = strconv.FormatInt(r.Position, 10)

	return
}

func (r *SetStickerPositionInSet) GetFiles() (files map[string]io.Reader) {
	return
}
