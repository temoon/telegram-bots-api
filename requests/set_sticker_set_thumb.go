package requests

import (
	"context"
	"io"
	"strconv"

	"github.com/temoon/telegram-bots-api"
)

type SetStickerSetThumb struct {
	Name   string
	Thumb  interface{}
	UserId int64
}

func (r *SetStickerSetThumb) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "setStickerSetThumb", r, response)
	return
}

func (r *SetStickerSetThumb) IsMultipart() (multipart bool) {
	return true
}

func (r *SetStickerSetThumb) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["name"] = r.Name

	switch value := r.Thumb.(type) {
	case io.Reader:
		values["thumb"] = value
	case string:
		values["thumb"] = value
	}

	values["user_id"] = strconv.FormatInt(r.UserId, 10)

	return
}
