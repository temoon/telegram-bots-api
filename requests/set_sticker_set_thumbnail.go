package requests

import (
	"context"
	"io"
	"strconv"

	"github.com/temoon/telegram-bots-api"
)

type SetStickerSetThumbnail struct {
	Name      string
	Thumbnail interface{}
	UserId    int64
}

func (r *SetStickerSetThumbnail) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "setStickerSetThumbnail", r, response)
	return
}

func (r *SetStickerSetThumbnail) IsMultipart() (multipart bool) {
	return true
}

func (r *SetStickerSetThumbnail) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["name"] = r.Name

	switch value := r.Thumbnail.(type) {
	case io.Reader:
		values["thumbnail"] = value
	case string:
		values["thumbnail"] = value
	}

	values["user_id"] = strconv.FormatInt(r.UserId, 10)

	return
}
