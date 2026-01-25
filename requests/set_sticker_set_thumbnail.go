package requests

import (
	"context"
	"io"
	"strconv"

	"github.com/temoon/telegram-bots-api"
)

type SetStickerSetThumbnail struct {
	Format    string
	Name      string
	UserId    int64
	Thumbnail *telegram.InputFile
}

func (r *SetStickerSetThumbnail) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "setStickerSetThumbnail", r, response)
	return
}

func (r *SetStickerSetThumbnail) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["format"] = r.Format

	values["name"] = r.Name

	values["user_id"] = strconv.FormatInt(r.UserId, 10)

	if r.Thumbnail != nil {
		values["thumbnail"] = r.Thumbnail.GetValue()
	}

	return
}

func (r *SetStickerSetThumbnail) GetFiles() (files map[string]io.Reader) {
	return
}
